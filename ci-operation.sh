#! /bin/bash

##########################
#
# This script is used for quickly testing the function
#
##########################
#
# usage:
# ./ ci-operation.sh [action]
#
# e.g. ./ci-operation.sh test ulcl-ti
#
##########################

usage() {
    echo "usage: ./ci-operation.sh [action] [target]"
    echo "  - pull: remove the existed free5gc repo under base/ and clone a new free5gc with its NFs"
    echo "  - fetch: fetch the target NF's PR"
    echo "  - build: build the necessary images"
    echo "  - up <ulcl-ti>: bring up the compose"
    echo "  - down <ulcl-ti>: shut down the compose"
    echo "  - test <ulcl-ti>: run ULCL test"
    echo "  - exec <ci>: enter the ci container"
}

main() {
    if [ $# -ne 1 ] && [ $# -ne 2 ] && [ $# -ne 3 ]; then
        usage
    fi

    case "$1" in
        "pull")
            cd base
            rm -rf free5gc
            git clone -j `nproc` --recursive https://github.com/free5gc/free5gc
            cd ..
        ;;
        "fetch")
            cd base/free5gc/NFs/$2
            git fetch origin pull/$3/head:pr-$3
            git checkout pr-$3
            cd ../../../../
        ;;
        "build")
            make ulcl
        ;;
        "up")
            case "$2" in
                "ulcl-ti")
                    docker compose -f docker-compose-ulcl-ti.yaml up
                    ;;
                *)
                    usage
            esac
        ;;
        "down")
            case "$2" in
                "ulcl-ti")
                    docker compose -f docker-compose-ulcl-ti.yaml down
                    ;;
                *)
                    usage
            esac
        ;;
        "test")
            case "$2" in
                "ulcl-ti")
                    docker exec ci /bin/bash -c "cd /root/test && ./test-ulcl-ti.sh TestULCLTrafficInfluence"
                ;;
                *)
                    usage
            esac
        ;;
        "exec")
            case "$2" in
                "ci")
                    docker exec -it ci bash
                ;;
                *)
                    usage
            esac
        ;;
    esac
}

main "$@"