package cluster

import (
	"log"

	"github.com/hashicorp/memberlist"
)

var clusterConf *memberlist.Config
var clusterNodes *memberlist.Memberlist

// Connect will form a cluster and optionally, connect to a seed node.
func Connect(nodeName, advertiseAddr string, advertisePort int, seed string) {
	clusterConf = memberlist.DefaultLocalConfig()
	clusterConf.Name = nodeName
	clusterConf.AdvertiseAddr = advertiseAddr
	clusterConf.AdvertisePort = advertisePort
	clusterConf.BindAddr = advertiseAddr
	clusterConf.BindPort = advertisePort

	// Create a memberlist
	var err error
	clusterNodes, err = memberlist.Create(clusterConf)
	if err != nil {
		log.Fatal("Failed to create memberlist: " + err.Error())
	}

	if seed != "" {
		// Join an existing cluster by specifying at least one known member.
		_, err = clusterNodes.Join([]string{seed})
		if err != nil {
			log.Fatal("Failed to join cluster: " + err.Error())
		}
	}

	// Ask for members of the cluster
	log.Println("Connected to cluster. Members:")
	for _, member := range clusterNodes.Members() {
		log.Printf("Node: %s %s\n", member.Name, member.Addr)
	}
}
