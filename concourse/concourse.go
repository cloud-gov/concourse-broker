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
	"golang.org/x/oauth2"
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

func (c *concourseClient) getAuthClient(concourseURL string) (concourse.Client, error) {
	team := c.client.Team(adminTeam)
	token, err := team.AuthToken()
	if err != nil {
		return nil, err
	}
	var oAuthToken *oauth2.Token
	oAuthToken = &oauth2.Token{
		TokenType:   token.Type,
		AccessToken: token.Value,
	}

	transport := transport(false, nil)

	transport = &oauth2.Transport{
		Source: oauth2.StaticTokenSource(oAuthToken),
		Base:   transport,
	}

	httpClient := &http.Client{Transport: transport}
	return concourse.NewClient(concourseURL, httpClient), nil
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
	client, err := c.getAuthClient(env.ConcourseURL)
	if err != nil {
		log.Println("can't get auth client")
		return err
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

func (c *concourseClient) DeleteTeam(details cf.Details, env config.Env) error {
	client, err := c.getAuthClient(env.ConcourseURL)
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