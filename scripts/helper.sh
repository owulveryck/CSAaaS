#!/bin/sh

# CSA helper lib

CSAFQDN=${CSAFQDN:-aws-hpe.owulveryck.info}
CSAPORT=${CSAPORT:-18444}
CSASCHEME=${CSASCHEME:-https}
CSABASEDIR=${CSABASEDIR:-/csa}
CSATOKEN=${CSATOKEN:-NULL}

if [ _$CSATOKEN == _NULL ] 
then
	echo "No token defined run ./getToken.sh"
fi

export CSAFQDN
export CSAPORT
export CSABASEDIR
export CSATOKEN
export CSASCHEME


