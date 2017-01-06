# Concourse Broker

[![Code Climate](https://codeclimate.com/github/18F/concourse-broker/badges/gpa.svg)](https://codeclimate.com/github/18F/concourse-broker)

[![Build Status](https://travis-ci.org/18F/concourse-broker.svg?branch=master)](https://travis-ci.org/18F/concourse-broker)

This is an **experimental** Cloud Foundry Service Broker for
provisioning teams on a deployed [Concourse CI](https://concourse.ci/)
instance.


## Setup

#### Create a client in UAA for this app

This application uses oauth to perform actions on your behalf in UAA.  To add a new oauth client in UAA, run the following command:

	uaac client add [your-client-id] --name "Concourse CI Broker" --scope "cloud_controller.read" --authorized_grant_types "authorization_code,client_credentials,refresh_token" --authorities "cloud_controller.admin" --autoapprove "true" --redirect_uri [url-for-concourse-ci]/auth/uaa/callback -s [your-client-secret]

**Remember the client-id and client-secret, you'll need them for Deployment**

## Deployment

### Automated

The easiest/recommended way to deploy the broker is via the [Concourse](http://concourse.ci/) pipeline.

1. Create a `ci/credentials.yml` file from the `ci/credentials.example.yml` (i.e. `cp ci/credentials.example.yml ci/credentials.yml`), and fill in the templated values from [the pipeline](ci/pipeline.yml).
1. Deploy the pipeline.

    ```bash
    fly -t lite set-pipeline -n -c ci/pipeline.yml -p deploy-concourse-broker -l ci/credentials.yml
    ```

### Manual

1. Clone this repository, and `cd` into it.
1. Target the space you want to deploy the broker to.

    ```bash
    $ cf target -o <org> -s <space>
    ```

1. The configuration is entirely read from environment variables. Edit the manifest.yml files and update your settings as necessary.
1. Deploy the broker as an application.

    ```bash
    $ cf push
    ```

1. [Register the broker](http://docs.cloudfoundry.org/services/managing-service-brokers.html#register-broker).

    ```bash
    $ cf create-service-broker concourse-broker [username] [password] [app-url] --space-scoped
    ```

### Explanation of Environment Variables

* `BROKER_USERNAME`
  * The username for providing [HTTP Basic Auth](https://docs.cloudfoundry.org/services/api.html#authentication) for the broker.
* `BROKER_PASSWORD`
  * The password for providing [HTTP Basic Auth](https://docs.cloudfoundry.org/services/api.html#authentication) for the broker.
* `ADMIN_USERNAME`
  * The username for the user that has access to the main team of the Concourse deployment.
* `ADMIN_PASSWORD`
  * The password for the user that has access to the main team of the Concourse deployment.
* `CONCOURSE_URL`
	* The base URL for the Concourse instance.
* `CF_URL`
	* The CF API URL for the Cloud Foundry deployment. (e.g. `https://api.bosh-lite.com`)
* `AUTH_URL`
	* The authorization url for UAA. (e.g. `https://login.bosh-lite.com/oauth/authorize`)
* `TOKEN_URL`
	* The token url for UAA. (e.g. `https://uaa.bosh-lite.com/oauth/token`)
* `CLIENT_ID`
	* The Client ID from [Setup](#setup)
* `CLIENT_SECRET`
	* The Client Setup from [Setup](#setup)

## Developing

In order to contribute to the broker, you will need:
* [Go 1.7](https://golang.org/dl/)
* [Glide](https://glide.sh/)
* [Ginkgo & Gomega](https://github.com/onsi/ginkgo#set-me-up)

### Adding new Dependencies

In order to add new dependencies, use Glide from the root of the project:

```
glide get github.com/org/projectname
```

*Please remember to add the new dependencies in a separate commit from the rest of the commits in the PR.*

### Running tests

In order to run the tests for the project, in the root of the project run:

```
ginkgo -r .
```
