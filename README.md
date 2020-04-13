### TCP Connection Tracker

This is a POC to notify when connection attempts on etcd peer port fails. This will raise an event if etcd pod is running but facing underlying network problems. It can also be extended for KAS, OAS and etcd client port.

# Run the binary
 1. docker build . -t quay.io/<USER>/tcpconnectiontracker
 2. If you are testing on OpenShift `oc debug node/<node>` followed by `chroot /host`
 3. `podman run -it --rm   --privileged   -v /sys/kernel/debug:/sys/kernel/debug:rw   -v /lib/modules:/lib/mdules:ro  --pid=host  quay.io/user/tcpconnectiontracker ./tcpconnectiontracker`