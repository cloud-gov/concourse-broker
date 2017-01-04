package concourse

import (
	"github.com/18F/concourse-broker/cf"
	"github.com/concourse/go-concourse/concourse"
	"github.com/concourse/atc"
	"errors"
	"github.com/18F/concourse-broker/config"
	"fmt"
)

const adminTeam = "main"

type Client interface {
	CreateTeam(details cf.Details, env config.Env) error
	DeleteTeam(details cf.Details, env config.Env) error
}

func NewClient() Client {
	return &concourseClient{}
}

type concourseClient struct {
	client concourse.Client
}

func (c *concourseClient) CreateTeam(details cf.Details, env config.Env) error {
	team := atc.Team{
		Name: fmt.Sprintf("%s-%s", details.OrgName, details.SpaceName),
		UAAAuth: &atc.UAAAuth{
			ClientID: env.ClientID,
			ClientSecret: env.ClientSecret,
			AuthURL: env.AuthURL,
			TokenURL: "",
			CFSpaces: []string{details.SpaceGUID},
			CFCACert: "",
			CFURL: env.CFURL,
		},
	}
	_, created, updated, err := c.client.Team(adminTeam).CreateOrUpdate(team)
	if err != nil {
		return err
	}
	if !created || updated {
		return errors.New("Unable to provision instance")
	}
	return nil
}

func (c *concourseClient) DeleteTeam(details cf.Details, env config.Env) error {
	err := c.client.Team(adminTeam).DestroyTeam(fmt.Sprintf("%s-%s", details.OrgName, details.SpaceName))
	if err != nil {
		return err
	}
	return nil
}