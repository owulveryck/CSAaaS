#!/bin/sh

URL=$1
eval $(echo $URL | sed 's|\(http.://.*/csa\).*catalog/\(.*\)/category/\(.*\)/service/\(.*\)|BASEURL=\1;CATALOGID=\2;CATEGORY=\2;ID=\4|')


#ID=$1
#CATALOGID=$2
#CATEGORY=$3

unset https_proxy
curl -s \
    -H "Accept: application/json" \
    -H "X-Auth-Token: $CSATOKEN" \
    -k \
    -H "Content-Type: application/json" \
    --user idmTransportUser:cloud \
    -XGET \
    "$BASEURL/api/mpp/mpp-offering/$ID?catalogId=$CATALOGID&category=$CATEGORY"


