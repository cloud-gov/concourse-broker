package concourse

import (
	"encoding/json"
	"errors"
	"fmt"

	"code.cloudfoundry.org/lager"
	"github.com/18F/concourse-broker/cf"
	"github.com/18F/concourse-broker/config"
	"github.com/concourse/atc"
	"github.com/concourse/atc/auth/uaa"
	"github.com/concourse/go-concourse/concourse"
	"github.com/hashicorp/go-multierror"
)

const adminTeam = "main"

// Client defines the capabilities that any concourse client should be able to do.
type Client interface {
	CreateTeam(details cf.Details) error
	DeleteTeam(details cf.Details) error
	UpdateTeams() error
}

// NewClient returns a client that can be used to interface with a deployed Concourse CI instance.
func NewClient(env config.Env, logger lager.Logger) Client {
	httpClient := newBasicAuthClient(env.AdminUsername, env.AdminPassword)

	return &concourseClient{
		client: concourse.NewClient(env.ConcourseURL, httpClient),
		env:    env,
		logger: logger.Session("concourse-client"),
	}
}

type concourseClient struct {
	client concourse.Client
	env    config.Env
	logger lager.Logger
}

func (c *concourseClient) getAuthClient(concourseURL string) (concourse.Client, error) {
	team := c.client.Team(adminTeam)
	token, err := team.AuthToken()
	if err != nil {
		return nil, err
	}
	httpClient := newOAuthClient(token.Type, token.Value)
	return concourse.NewClient(concourseURL, httpClient), nil
}

func (c *concourseClient) getTeamName(details cf.Details) string {
	return details.OrgName
}

func (c *concourseClient) CreateTeam(details cf.Details) error {
	teamName := c.getTeamName(details)
	auth := uaa.UAAAuthConfig{
		ClientID:     c.env.ClientID,
		ClientSecret: c.env.ClientSecret,
		CFURL:        c.env.CFURL,
		AuthURL:      c.env.AuthURL,
		TokenURL:     c.env.TokenURL,
		CFSpaces:     []string{details.SpaceGUID},
		CFCACert:     "",
	}
	authData, err := json.Marshal(auth)
	if err != nil {
		c.logger.Error("create-team.marshal-error", err)
		return err
	}
	team := atc.Team{
		Auth: map[string]*json.RawMessage{
			uaa.ProviderName: (*json.RawMessage)(&authData),
		},
	}
	client, err := c.getAuthClient(c.env.ConcourseURL)
	if err != nil {
		c.logger.Error("create-team.auth-client-error", err)
		return err
	}
	authMethods, err := client.Team(teamName).ListAuthMethods()
	if err == nil || len(authMethods) > 0 {
		err := fmt.Errorf("Team %s already exists", teamName)
		c.logger.Error("create-team.existing-team-error", err,
			lager.Data{
				"team-name":         teamName,
				"auth-methods-size": len(authMethods),
			})
		return err
	}
	_, created, updated, err := client.Team(teamName).CreateOrUpdate(team)
	if err != nil {
		c.logger.Error("create-team.unknown-create-error", err,
			lager.Data{
				"team-name": teamName,
			})
		return err
	}
	if !created || updated {
		err := errors.New("Unable to provision instance")
		c.logger.Error("create-team.unknown-create-error", err,
			lager.Data{
				"team-name": teamName,
			})
		return err
	}
	return nil
}

func (c *concourseClient) DeleteTeam(details cf.Details) error {
	teamName := c.getTeamName(details)
	client, err := c.getAuthClient(c.env.ConcourseURL)
	if err != nil {
		c.logger.Error("delete-team.auth-client-error", err)
		return err
	}
	err = client.Team(details.OrgName).DestroyTeam(teamName)
	if err != nil {
		c.logger.Error("delete-team.unknown-delete-error", err,
			lager.Data{
				"team-name": teamName,
			})
		return err
	}
	return nil
}

func (c *concourseClient) UpdateTeams() error {
	client, err := c.getAuthClient(c.env.ConcourseURL)
	if err != nil {
		c.logger.Error("update-teams.auth-client-error", err)
		return err
	}

	teams, err := client.ListTeams()
	if err != nil {
		c.logger.Error("update-teams.list-error", err)
		return err
	}

	errc := make(chan error)
	for _, team := range teams {
		go func(team atc.Team) {
			auth := &uaa.UAAAuthConfig{}
			if err := json.Unmarshal(*team.Auth[uaa.ProviderName], &auth); err != nil {
				errc <- err
				return
			}

			if auth.ClientID == c.env.ClientID && auth.ClientSecret == c.env.ClientSecret {
				errc <- nil
				return
			}

			auth.ClientID = c.env.ClientID
			auth.ClientSecret = c.env.ClientSecret

			authData, err := json.Marshal(auth)
			if err != nil {
				c.logger.Error("update-teams.marshal-error", err)
				errc <- err
				return
			}
			team.Auth[uaa.ProviderName] = (*json.RawMessage)(&authData)

			_, _, _, err = client.Team(team.Name).CreateOrUpdate(team)
			errc <- err
		}(team)
	}

	var result error
	for _ = range teams {
		if err := <-errc; err != nil {
			multierror.Append(result, err)
		}
	}

	return result
}
