package concourse

import (
	"errors"
	"fmt"
	"github.com/18F/concourse-broker/cf"
	"github.com/18F/concourse-broker/config"
	"github.com/concourse/atc"
	"github.com/concourse/go-concourse/concourse"
	"log"
)

const adminTeam = "main"

type Client interface {
	CreateTeam(details cf.Details) error
	DeleteTeam(details cf.Details) error
}

func NewClient(env config.Env) Client {
	httpClient := newBasicAuthClient(env.AdminUsername, env.AdminPassword)
	return &concourseClient{client: concourse.NewClient(env.ConcourseURL, httpClient), env: env}
}

type concourseClient struct {
	client concourse.Client
	env    config.Env
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

func (c *concourseClient) CreateTeam(details cf.Details) error {
	teamName := details.OrgName
	team := atc.Team{
		UAAAuth: &atc.UAAAuth{
			ClientID:     c.env.ClientID,
			ClientSecret: c.env.ClientSecret,
			AuthURL:      c.env.AuthURL,
			TokenURL:     c.env.TokenURL,
			CFSpaces:     []string{details.SpaceGUID},
			CFCACert:     "",
			CFURL:        c.env.CFURL,
		},
	}
	client, err := c.getAuthClient(c.env.ConcourseURL)
	if err != nil {
		log.Println("can't get auth client")
		return err
	}
	authMethods, err := client.Team(teamName).ListAuthMethods()
	if err == nil || len(authMethods) > 0 {
		return fmt.Errorf("Team %s already exists", teamName)
	}
	_, created, updated, err := client.Team(teamName).CreateOrUpdate(team)
	if err != nil {
		log.Printf("Unable to create team %s\n", team.Name)
		return err
	}
	if !created || updated {
		return errors.New("Unable to provision instance")
	}
	return nil
}

func (c *concourseClient) DeleteTeam(details cf.Details) error {
	client, err := c.getAuthClient(c.env.ConcourseURL)
	if err != nil {
		log.Println("can't get auth client")
		return err
	}
	err = client.Team(details.OrgName).DestroyTeam(details.OrgName)
	if err != nil {
		log.Println("couldn't destroy team.")
		return err
	}
	return nil
}
