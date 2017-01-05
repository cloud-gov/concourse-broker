package broker

import (
	"code.cloudfoundry.org/lager"
	"context"
	"errors"
	"github.com/18F/concourse-broker/cf"
	"github.com/18F/concourse-broker/concourse"
	"github.com/18F/concourse-broker/config"
	"github.com/pivotal-cf/brokerapi"
)

// New returns a new concourse service broker instance.
func New(services []brokerapi.Service, logger lager.Logger, env config.Env) brokerapi.ServiceBroker {
	return &concourseBroker{services: services, logger: logger, env: env}
}

type concourseBroker struct {
	services []brokerapi.Service
	logger   lager.Logger
	env      config.Env
}

func (c *concourseBroker) Services(context context.Context) []brokerapi.Service {
	return c.services
}

func (c *concourseBroker) Provision(context context.Context, instanceID string,
	details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	cfClient, err := cf.NewClient(c.env)
	if err != nil {
		return brokerapi.ProvisionedServiceSpec{}, err
	}
	cfDetails, err := cfClient.GetProvisionDetails(details.SpaceGUID)
	cfDetails.SpaceGUID = details.SpaceGUID
	if err != nil {
		return brokerapi.ProvisionedServiceSpec{}, err
	}
	concourseClient := concourse.NewClient(c.env)
	err = concourseClient.CreateTeam(cfDetails)
	if err != nil {
		return brokerapi.ProvisionedServiceSpec{}, err
	}
	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (c *concourseBroker) Deprovision(context context.Context, instanceID string,
	details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
	cfClient, err := cf.NewClient(c.env)
	if err != nil {
		return brokerapi.DeprovisionServiceSpec{}, err
	}
	cfDetails, err := cfClient.GetDeprovisionDetails(instanceID)
	if err != nil {
		return brokerapi.DeprovisionServiceSpec{}, err
	}
	concourseClient := concourse.NewClient(c.env)
	err = concourseClient.DeleteTeam(cfDetails)
	if err != nil {
		return brokerapi.DeprovisionServiceSpec{}, err
	}
	return brokerapi.DeprovisionServiceSpec{}, nil
}

func (c *concourseBroker) Bind(context context.Context, instanceID,
	bindingID string, details brokerapi.BindDetails) (brokerapi.Binding, error) {
	return brokerapi.Binding{}, errors.New("service does not support bind")
}
func (c *concourseBroker) Unbind(context context.Context, instanceID, bindingID string,
	details brokerapi.UnbindDetails) error {
	return errors.New("service does not support bind")
}

func (c *concourseBroker) Update(context context.Context, instanceID string,
	details brokerapi.UpdateDetails, asyncAllowed bool) (brokerapi.UpdateServiceSpec, error) {
	return brokerapi.UpdateServiceSpec{}, nil
}

func (c *concourseBroker) LastOperation(context context.Context, instanceID,
	operationData string) (brokerapi.LastOperation, error) {
	return brokerapi.LastOperation{}, nil
}
