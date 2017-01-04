package concourse

import (
	"github.com/18F/concourse-broker/cf"
	"github.com/concourse/go-concourse/concourse"
	"github.com/concourse/atc"
	"errors"
	"github.com/18F/concourse-broker/config"
	"fmt"
	"net/http"
	"crypto/x509"
	//"crypto/tls"
	"net"
	"time"
	"log"
)

const adminTeam = "main"

type Client interface {
	CreateTeam(details cf.Details, env config.Env) error
	DeleteTeam(details cf.Details, env config.Env) error
}


type basicAuthTransport struct {
	username string
	password string

	base http.RoundTripper
}

// https://github.com/concourse/fly/blob/6fb036ef31f6e6f3e74f0089f2d59d2722f0580c/rc/target.go#L378
func (t basicAuthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	log.Printf("IS REQUEST EQUAL TO NIL %v", r == nil)
	log.Printf("username %s password length %d\n", t.username, len(t.password))
	r.SetBasicAuth(t.username, t.password)
	log.Printf("AUTH HEADER %s", r.Header.Get("Authorization"))
	return t.base.RoundTrip(r)
}

func transport(insecure bool, caCertPool *x509.CertPool) http.RoundTripper {
	var transport http.RoundTripper

	transport = &http.Transport{
		/*
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: insecure,
			RootCAs:            caCertPool,
		},
		*/
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		//Proxy: http.ProxyFromEnvironment,
	}

	return transport
}

func NewClient(env config.Env) Client {
	httpClient := &http.Client{
		Transport: basicAuthTransport{
			username: env.AdminUsername,
			password: env.AdminPassword,
			base: transport(false, nil),
		},
	}
	return &concourseClient{client: concourse.NewClient(env.ConcourseURL, httpClient)}
}

type concourseClient struct {
	client concourse.Client
}

func (c *concourseClient) CreateTeam(details cf.Details, env config.Env) error {
	fmt.Println("made it")
	teamName := details.OrgName
	team := atc.Team{
		//Name: teamName,
		UAAAuth: &atc.UAAAuth{
			ClientID: env.ClientID,
			ClientSecret: env.ClientSecret,
			AuthURL: env.AuthURL,
			TokenURL: env.TokenURL,
			CFSpaces: []string{details.SpaceGUID},
			CFCACert: "",
			CFURL: env.CFURL,
		},
	}
	_, created, updated, err := c.client.Team(teamName).CreateOrUpdate(team)
	if err != nil {
		log.Printf("Unable to create team %s\n", team.Name)
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