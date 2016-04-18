#!/bin/sh

TOKEN=$(curl \
    -s \
    -H "Accept: application/json" \
    -d@credentials.json \
    -k \
    -H "Content-Type: application/json" \
    --user idmTransportUser:cloud \
    -XPOST ${CSASCHEME}://${CSAFQDN}:${CSAPORT}/idm-service/v2.0/tokens | egrep -o ".{36}\.[a-zA-Z0-9]+\..{43}")

echo "CSATOKEN=$TOKEN;export CSATOKEN"


