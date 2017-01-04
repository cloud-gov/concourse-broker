#!/bin/bash

set -e
set -u

cf login -a $CF_API_URL -u $CF_USERNAME -p $CF_PASSWORD -o $CF_ORGANIZATION -s $CF_SPACE

set -x

space_guid=$(cf space "${CF_SPACE}" --guid)

# Test concourse ci plan

# Create service instance
cf create-service concourse-ci "${PLAN_NAME}" "${SERVICE_INSTANCE_NAME}"
instance_guid=$(cf service "${SERVICE_INSTANCE_NAME}" --guid)

# TODO Run actual tests

# Delete service instance
cf delete-service -f "${SERVICE_INSTANCE_NAME}"

# TODO Check if service was really deleted.

####

# Ensure service instance is deleted
teardown() {
  cf delete-service -f "${SERVICE_INSTANCE_NAME}"
}
trap teardown EXIT
