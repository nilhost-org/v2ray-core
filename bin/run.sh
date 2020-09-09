#!/usr/bin/env bash
case "$1" in

client)
    ./v2ray --config="./config/config-client.json"
;;

client-prod)
    ./v2ray --config="./config/config-client-prod.json"
;;

server)
    ./v2ray --config="./config/config-server.json"
;;

build-mac)
    cd ../core && make macos && cd -
;;
*)
echo "usage: ./run.sh client/server"

esac
