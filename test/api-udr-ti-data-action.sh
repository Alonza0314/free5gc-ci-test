#!/bin/bash

##########################
#
# usage:
# ./ci-test-ti-data-insert.sh put
# ./ci-test-ti-data-insert.sh get
# ./ci-test-ti-data-insert.sh delete
#
##########################

set -e

# ti data request
case "$1" in
    "put")
        curl -X PUT -H "Content-Type: application/json" --data @json/ti-data.json \
            http://udr.free5gc.org:8000/nudr-dr/v1/application-data/influenceData/1
        ;;
    "get")
        curl -X GET -H "Content-Type: application/json" \
            http://udr.free5gc.org:8000/nudr-dr/v1/application-data/influenceData/
        ;;
    "delete")
        curl -X DELETE -H "Content-Type: application/json" \
            http://udr.free5gc.org:8000/nudr-dr/v1/application-data/influenceData/1
        ;;
    *)
        echo "error: invalid parameter"
        echo "usage: $0 [put|get|delete]"
        exit 1
        ;;
esac
