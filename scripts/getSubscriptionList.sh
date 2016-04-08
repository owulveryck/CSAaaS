#!/bin/ksh

URL=$1

if [ "_$TOKEN" == "_" ]
then
    echo "No token provided"
    exit
fi

unset https_proxy
curl -s \
    -H "Accept: application/json" \
    -H "X-Auth-Token: $TOKEN" \
    -d'{"name": null, "approval": "ALL", "category": null}' \
    -k \
    -H "Content-Type: application/json" \
    --user idmTransportUser:idmTransportUser \
    -XPOST \
    $URL/csa/api/mpp/mpp-offering/filter | jsonformat

