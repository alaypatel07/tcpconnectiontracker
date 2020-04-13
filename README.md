### TCP Connection Tracker

This is a POC to notify when connection attempts on etcd peer port fails. This will raise an event if etcd pod is running but facing underlying network problems. It can also be extended for KAS, OAS and etcd client port.

# Run the binary
 1. docker build . -t quay.io/<USER>/tcpconnectiontracker
 2. If you are testing on OpenShift `oc debug node/<node>` followed by `chroot /host`
 3. ```podman run -it --rm   --privileged   -v /sys/kernel/debug:/sys/kernel/debug:rw   -v /lib/modules:/lib/mdules:ro  --pid=host  quay.io/user/tcpconnectiontracker ./tcpconnectiontracker```
 4. Sample output when I block network connection using iptables.
 ```$xslt
sh-4.4# podman run -it --rm   --privileged   -v /sys/kernel/debug:/sys/kernel/debug:rw   -v /lib/modules:/lib/modules:ro  --pid=host  quay.io/alaypatel07/tcpconnection
 tracker:trial ./tcpconnectiontracker
 &{0xc000188000 0xc00006c040 0xc00006c080 0xc00006f8c0}
 Got V6 event connect from ::1:50248 to ::1:6443
 Got V6 event accept from ::1:6443 to ::1:50248
 unsuccessful connection attempt from 10.0.0.5:35210 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:53396 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:53394 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:2380 to 10.0.0.4:48094 with process id 19521
 Got V6 event close from ::1:54114 to ::1:6443
 Got V6 event close from ::1:6443 to ::1:54114
 unsuccessful connection attempt from 10.0.0.5:35232 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:35236 to 10.0.0.4:2380 with process id 19521
 Got V6 event connect from ::1:50306 to ::1:6443
 Got V6 event accept from ::1:6443 to ::1:50306
 unsuccessful connection attempt from 10.0.0.5:35244 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:53698 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:35246 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:35250 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:35284 to 10.0.0.4:2380 with process id 19521
 unsuccessful connection attempt from 10.0.0.5:35286 to 10.0.0.4:2380 with process id 19521
```
 