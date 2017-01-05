#!/bin/bash

set -e
set -u

cf login -a $CF_API_URL -u $CF_USERNAME -p $CF_PASSWORD -o $CF_ORGANIZATION -s $CF_SPACE

set -x

space_guid=$(cf space "${CF_SPACE}" --guid)

# Test concourse ci plan

# Try to find team before creation and should find nothing

RESULT=$(curl -L "${CONCOURSE_URL}/api/v1/teams" -s | jq -c ".[] | select(.name | contains (\"${CF_ORGANIZATION}\"))")
if [ ! -z "$RESULT" ]; then
    echo "Found concourse team before creation. Failing for now."
    exit 1
fi

# Create service instance
cf create-service concourse-ci "${PLAN_NAME}" "${SERVICE_INSTANCE_NAME}"
instance_guid=$(cf service "${SERVICE_INSTANCE_NAME}" --guid)


# Try to find team after service creation and should find the list.

RESULT=$(curl -L "${CONCOURSE_URL}/api/v1/teams" -s | jq -c ".[] | select(.name | contains (\"${CF_ORGANIZATION}\"))")
if [ -z "$RESULT" ]; then
    echo "Could not find concourse team after service creation. Failing..."
    exit 1
fi

# Delete service instance
cf delete-service -f "${SERVICE_INSTANCE_NAME}"


RESULT=$(curl -L "${CONCOURSE_URL}/api/v1/teams" -s | jq -c ".[] | select(.name | contains (\"${CF_ORGANIZATION}\"))")
if [ ! -z "$RESULT" ]; then
    echo "Found concourse team after deletion. Failing..."
    exit 1
fi

####

# Ensure service instance is deleted
teardown() {
  cf delete-service -f "${SERVICE_INSTANCE_NAME}"
}
trap teardown EXIT
