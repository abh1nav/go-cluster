package main

import (
	"flag"
	"log"
	"runtime"
	"syscall"

	"github.com/abh1nav/go-cluster-test/api"
	"github.com/abh1nav/go-cluster-test/cluster"
	"github.com/abh1nav/go-cluster-test/utils"
	"github.com/zenazn/goji/graceful"
)

var (
	bind          = flag.CommandLine.String("bind", "127.0.0.1:5000", "<address>:<port> to bind HTTP server")
	name          = flag.CommandLine.String("node", "", "Give this node a name")
	seed          = flag.CommandLine.String("seed", "", "Seed name, optionally to join an existing cluster")
	advertiseAddr = flag.CommandLine.String("advertise-addr", "127.0.0.1", "Cluster advertise address")
	advertisePort = flag.CommandLine.Int("advertise-port", 6000, "Cluster advertise port")
)

func main() {
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())

	// If no nodename was provided, generate a random one
	var nodeName = ""
	if *name == "" {
		nodeName = utils.RandSeq(10)
	} else {
		nodeName = *name
	}

	// Connect to the cluster
	cluster.Connect(nodeName, *advertiseAddr, *advertisePort, *seed)

	graceful.AddSignal(syscall.SIGINT, syscall.SIGTERM)
	app := api.New()
	log.Println("go-cluster server listening on " + *bind)
	err := graceful.ListenAndServe(*bind, app)
	if err != nil {
		log.Fatal(err)
	}

	graceful.Wait()
}
