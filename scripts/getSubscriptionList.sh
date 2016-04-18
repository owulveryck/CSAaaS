#!/bin/sh

curl -v \
    -H "Accept: application/json" \
    -H "X-Auth-Token: $CSATOKEN" \
    -d'{"name": null, "approval": "ALL", "category": null}' \
    -k \
    -H "Content-Type: application/json" \
    --user idmTransportUser:cloud \
    -XPOST \
    $CSASCHEME://$CSAFQDN:$CSAPORT$CSABASEDIR/api/mpp/mpp-offering/filter 

