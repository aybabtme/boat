#!/usr/bin/env bash

set -eu

router_user=root
router_addr=192.168.1.1
router_firmware=bin/openwrt-we826-GO2019-01-13-upgrade.bin

root=$(git rev-parse --show-toplevel)
cd $root/internet/lte/we826/

scp $router_firmware $router_user@$router_addr:/tmp/firmware.bin
ssh $router_user@$router_addr 'sysupgrade -n -v /tmp/firmware.bin'