package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/weaveworks/tcptracer-bpf/pkg/tracer"
)

type callbacks struct {
	// the connections are store from source to destination stored in the form of ip:port
	connections map[string]string
}

func (c *callbacks) TCPEventV4(e tracer.TcpV4) {
	if e.SPort != 2380 && e.DPort != 2380 {
		return
	}
	//fmt.Printf("Got event %s from %s:%d to %s:%d\n", e.Type, e.SAddr, e.SPort, e.DAddr, e.DPort)
	source := fmt.Sprintf("%s:%d", e.SAddr, e.SPort)
	destination := fmt.Sprintf("%s:%d", e.DAddr, e.DPort)
	switch {
	case e.Type == tracer.EventConnect && e.DPort == 2380:
		// add if this is a connection request from client
		c.connections[source] = destination
	case e.Type == tracer.EventAccept && e.SPort == 2380:
		// add only if the server accepted connection the connection
		c.connections[source] = destination
	case e.Type == tracer.EventClose:
		_, ok := c.connections[source]
		if !ok {
			fmt.Printf("unsuccessful connection attempt from %s to %s with process id %d\n", source, destination, e.Pid)
			return
		}
		//fmt.Printf("connection closed from %s to %s\n", source, destination)
		delete(c.connections, source)
	}
}

func (c *callbacks) TCPEventV6(e tracer.TcpV6) {
	fmt.Printf("Got V6 event %s from %s:%d to %s:%d\n", e.Type, e.SAddr, e.SPort, e.DAddr, e.DPort)
}

func (c *callbacks) LostV4(connId uint64) {
	fmt.Printf("Lost V4 connection %d\n", connId)
}

func (c *callbacks) LostV6(connId uint64) {
	fmt.Printf("Lost V6 connection %d\n", connId)
}

var pid int

func main() {
	flag.IntVar(&pid, "pid", 1, "The pid of process to watch network connections for")
	flag.Parse()
	t, err := tracer.NewTracer(&callbacks{
		connections: map[string]string{},
	})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(t)
	t.Start()
	defer t.Stop()
	err = t.AddFdInstallWatcher(uint32(pid))
	if err != nil {
		log.Fatalf("error installing fd watcher: %v\n", err)
	}
	defer t.RemoveFdInstallWatcher(uint32(pid))
	fmt.Scanf("%d", pid)
}
