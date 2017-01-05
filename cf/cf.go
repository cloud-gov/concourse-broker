package cf

import (
	"encoding/json"
	"fmt"
	"github.com/18F/concourse-broker/config"
	"github.com/cloudfoundry-community/go-cfclient"
	"io/ioutil"
	"log"
)

type Details struct {
	OrgGUID   string
	OrgName   string
	SpaceGUID string
	SpaceName string
}

type Client interface {
	GetProvisionDetails(spaceGUID string) (Details, error)
	GetDeprovisionDetails(serviceGUID string) (Details, error)
}

func NewClient(env config.Env) (Client, error) {
	config := &cfclient.Config{
		ClientID:     env.ClientID,
		ClientSecret: env.ClientSecret,
		ApiAddress:   env.CFURL,
	}
	log.Printf("API ADDRESS %s", config.ApiAddress)
	client, err := cfclient.NewClient(config)
	if err != nil {
		return nil, err
	}
	services, err := client.ListServices()
	log.Printf("NUMBER OF SERVICES %d", len(services))
	return &cfClient{client: client}, nil
}

type cfClient struct {
	client *cfclient.Client
}

func (c *cfClient) GetProvisionDetails(spaceGUID string) (Details, error) {
	requestURI := fmt.Sprintf("/v2/spaces/%s", spaceGUID)
	orgName, err := c.getOrgName(requestURI)
	if err != nil {
		return Details{}, err
	}
	return Details{OrgName: orgName}, nil
}

func (c *cfClient) GetDeprovisionDetails(serviceGUID string) (Details, error) {
	serviceInstance, err := c.client.ServiceInstanceByGuid(serviceGUID)
	if err != nil {
		return Details{}, err
	}
	services, err := c.client.ListServices()
	log.Printf("NUMBER OF SERVICES ALSO %d", len(services))
	orgName, err := c.getOrgName(serviceInstance.SpaceUrl)
	if err != nil {
		return Details{}, err
	}
	return Details{OrgName: orgName}, nil
}

func (c *cfClient) getOrgName(requestUrl string) (string, error) {
	var spaceResp cfclient.SpaceResource
	r := c.client.NewRequest("GET", requestUrl)
	resp, err := c.client.DoRequest(r)
	if err != nil {
		return "", fmt.Errorf("Error requesting spaces %v", err)
	}
	resBody, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("Error reading space request %v", err)
	}
	err = json.Unmarshal(resBody, &spaceResp)
	if err != nil {
		return "", fmt.Errorf("Error unmarshalling space %v", err)
	}
	var orgResp cfclient.OrgResource
	r = c.client.NewRequest("GET", spaceResp.Entity.OrgURL)
	resp, err = c.client.DoRequest(r)
	if err != nil {
		return "", fmt.Errorf("Error requesting orgs %v", err)
	}
	resBody, err = ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return "", fmt.Errorf("Error reading org request %v", err)
	}
	err = json.Unmarshal(resBody, &orgResp)
	if err != nil {
		return "", fmt.Errorf("Error unmarshalling org %v", err)
	}
	return orgResp.Entity.Name, nil
}
