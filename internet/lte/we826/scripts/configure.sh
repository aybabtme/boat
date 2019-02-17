#!/usr/bin/env bash

set -eu

echo "this script is not yet functional"
exit 1

router_user=root
router_addr=192.168.1.1
ssh_pub_key=$HOME/.ssh/id_rsa.pub

root=$(git rev-parse --show-toplevel)
cd $root/internet/lte/we826/

http_addr="http://$router_addr"

base_url="$http_addr/cgi-bin/luci"
login_url=$base_url
admin_url="$base_url/admin/system/admin"

base_username="root"
base_password=

curl -X GET  $login_url --cookie-jar cookie.jar --data "luci_username=$base_username&luci_password="

curl -X POST $admin_url --cookie-jar cookie.jar -F cbid.dropbear._keys._data=@$ssh_pub_key