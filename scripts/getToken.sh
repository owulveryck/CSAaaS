#!/bin/sh
URL=$1
TOKEN=$(curl \
    -s \
    -H "Accept: application/json" \
    -d@credentials.json \
    -k \
    -H "Content-Type: application/json" \
    --user idmTransportUser:idmTransportUser \
    -XPOST $URL/idm-service/v2.0/tokens | egrep -o ".{36}\..{779}\..{43}")

echo "TOKEN=$TOKEN;export TOKEN"


