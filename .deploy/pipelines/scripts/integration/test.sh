#!/bin/bash

# If running on Alpine, this script requires the following packages:
# * curl
# * ncurses
# * bash

color_red=$(tput setaf 1)
color_green=$(tput setaf 2)
color_cyan=$(tput setaf 6)
color_white=$(tput setaf 7)
separator="------------------------------"

if [ $# -lt 2 ]; then
    cat <<EOF
${color_red}
ERROR: Incorrect number of arguments supplied
Usage: ./test.sh [HOSTNAME] [PORT]
EOF
    exit 1
fi

service_name="$1"
service_port="$2"

function test_endpoint() {
  endpoint=$1
  expected_status_code=$2
  status_code="$(curl -o /dev/null -sw '%{http_code}' ${service_name}:${service_port}/${endpoint})"
  result=""
  if [[ "${status_code}" == "${expected_status_code}" ]]; then
    result="${color_green}Success.${color_white}"
  else
    result="${color_red}Failure.${color_white}"
  fi

  cat <<EOF
${separator}
${color_cyan}Testing endpoint:${color_white} ${service_name}:${service_port}/${endpoint}
${color_cyan}Expected status code:${color_white} ${expected_status_code}
${color_cyan}Status code:${color_white} ${status_code}
${color_cyan}Result:${color_white} ${result}
EOF
}

echo "${color_cyan}Running integration tests:${color_white}"

# main endpoints
test_endpoint "health-check" "200"
test_endpoint "" "200" # index
# misc assets
test_endpoint "favicon.ico" "200"
test_endpoint "style.css" "200"
test_endpoint "index.html" "200"
# 404
test_endpoint "asdkasldjal" "404"

echo ${separator}