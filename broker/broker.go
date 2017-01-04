package broker

import (
	"code.cloudfoundry.org/lager"
	"context"
	"errors"
	"github.com/pivotal-cf/brokerapi"
	"github.com/18F/concourse-broker/config"
)

// New returns a new concourse service broker instance.
func New(services []brokerapi.Service, logger lager.Logger, env config.Env) brokerapi.ServiceBroker {
	return &concourseBroker{services: services, logger: logger, env: env}
}

type concourseBroker struct {
	services []brokerapi.Service
	logger   lager.Logger
	env config.Env
}

func (c *concourseBroker) Services(context context.Context) []brokerapi.Service {
	return c.services
}

func (c *concourseBroker) Provision(context context.Context, instanceID string,
	details brokerapi.ProvisionDetails, asyncAllowed bool) (brokerapi.ProvisionedServiceSpec, error) {
	return brokerapi.ProvisionedServiceSpec{}, nil
}

func (c *concourseBroker) Deprovision(context context.Context, instanceID string,
	details brokerapi.DeprovisionDetails, asyncAllowed bool) (brokerapi.DeprovisionServiceSpec, error) {
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
