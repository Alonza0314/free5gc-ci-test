#!/bin/bash

##########################
#
# usage:
# ./test-ulcl-mp.sh <test-name>
#
# e.g. ./test-ulcl-mp.sh <TestULCLMultiPathCi1|TestULCLMultiPathCi2>
#
##########################

echo "test $1"

target_webconsole_subscription_data_file=""

case "$1" in
    "TestULCLMultiPathCi1")
        target_webconsole_subscription_data_file="json/webconsole-subscription-data-mp-1.json"
    ;;
    "TestULCLMultiPathCi2")
        target_webconsole_subscription_data_file="json/webconsole-subscription-data-mp-2.json"
    ;;
esac

echo "target_webconsole_subscription_data_file: $target_webconsole_subscription_data_file"

./api-webconsole-subscribtion-data-action.sh post $target_webconsole_subscription_data_file
if [ $? -ne 0 ]; then
    echo "Failed to post subscription data"
    exit 1
fi

./api-webconsole-subscribtion-data-action.sh delete $target_webconsole_subscription_data_file
if [ $? -ne 0 ]; then
    echo "Failed to delete subscription data"
    exit 1
fi

# exit $go_test_exit_code