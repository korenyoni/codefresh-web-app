#!/bin/bash

service_name="codefresh-web-app"
service_port="8080"

echo "Testing ${service_name}:${service_port}/health-check ..."

[[ "$(curl ${service_name}:${service_port}/health-check | jq -r .status)" == "alive" ]]
[[ "$(curl -o /dev/null -sw '%{http_code}' ${service_name}:${service_port}/health-check)" == "200" ]]

echo "Testing ${service_name}:${service_port}/ ..."

[[ "$(curl -o /dev/null -sw '%{http_code}' ${service_name}:${service_port}/)" == "200" ]]

echo "Testing ${service_name}:${service_port}/asdkasldjal ..."

[[ "$(curl -o /dev/null -sw '%{http_code}' ${service_name}:${service_port}/asdkasldjal)" == "404" ]]