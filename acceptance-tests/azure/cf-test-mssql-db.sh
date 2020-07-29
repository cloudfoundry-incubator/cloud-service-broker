#!/usr/bin/env bash

set -o nounset

SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

. "${SCRIPT_DIR}/../functions.sh"

SERVER_NAME=mssql-server-$$
USERNAME=anadminuser
PASSWORD=This_S0uld-3eC0mplex~

SERVER_RG=rg-test-service-$$
RESULT=1
if create_service csb-azure-resource-group standard "${SERVER_RG}" "{\"instance_name\":\"${SERVER_RG}\"}"; then
  if "${SCRIPT_DIR}/cf-create-mssql-server.sh" "${SERVER_NAME}" "${USERNAME}" "${PASSWORD}" "${SERVER_RG}" centralus; then
      CONFIG="{ 
        \"server\": \"test_server\", \
        \"server_credentials\": { \
          \"test_server\": { \
            \"server_name\":\"${SERVER_NAME}\", \
            \"admin_username\":\"${USERNAME}\", \
            \"admin_password\":\"${PASSWORD}\", \
            \"server_resource_group\":\"${SERVER_RG}\" \
          }, \
          \"fail_server\": { \
            \"server_name\":\"missing\", \
            \"admin_username\":\"bogus\", \
            \"admin_password\":\"bad-password\", \
            \"server_resource_group\":\"rg\" \
          } \
        } \
      }"

      echo $CONFIG

      ${SCRIPT_DIR}/../cf-test-spring-music.sh csb-azure-mssql-db small -u large "${CONFIG}"
      RESULT=$?
  fi

  "${SCRIPT_DIR}/cf-delete-mssql-server.sh" "${SERVER_NAME}"

  delete_service "${SERVER_RG}"
else
  echo "Failed creating resource group ${SERVER_RG} for test services"
fi

exit ${RESULT}