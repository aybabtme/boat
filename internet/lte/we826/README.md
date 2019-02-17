# we826

Router is a cheap thing bought on AliExpress, Amazon or wherever. Buy a Sierra Wireless MC7455 LTE modem (or a newer one if you want - not covered here).

## instructions

> Currently following these: https://ltehacks.com/viewtopic.php?f=21&t=233 

Connect to the device via an ethernet cable.

### Flash the firmware. 

Verify that your router's address matches the one in this script:

```bash
./script/flash.sh
```

If it matches, execute the script. It will take a few minutes to run, and then the device will reboot and take another few minutes to come back online. It will come back as `192.168.1.1`.

### Configuring the device.

> Note 2: the TLS certificate presented by the router is not valid, so your browser might get in the way of opening the UI. Ignore the browser, or coerce it into letting you in.

> Note 2: The `configure.sh` script does not work just yet. I need to figure out a way to make the firmware be configurable over SSH from firstboot time.

Open your browser to `http://192.168.1.1` and login with the `admin` username and a blank password. Go to the [admin page](http://192.168.1.1/cgi-bin/luci/admin/system/admin) and add your SSH key to the box asking for it.

Then visit the [modem config](http://192.168.1.1/cgi-bin/luci/admin/modem/prof) page. Add the following APN depending on your SIM card's provider:

- T-Mobile: `fast.t-mobile.com`

Go to the [Firewall TTL](http://192.168.1.1/cgi-bin/luci/admin/network/firewall/ttl) rule page and enable a TTL of 65.

Save and apply. The router should have internet through your LTE card!