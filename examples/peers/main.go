package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/proencaj/gorthanc"
	"github.com/proencaj/gorthanc/types"
)

func main() {
	client, err := gorthanc.NewClient(
		"http://localhost:8042",
		gorthanc.WithBasicAuth("orthanc", "orthanc"),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: GetPeers
	peers, err := client.GetPeers()
	if err != nil {
		log.Fatalf("Failed to get peers: %v", err)
	}
	fmt.Println(peers)

	// Example: GetPeerDetails
	if len(peers) > 0 {
		peer, err := client.GetPeerDetails(peers[0])
		if err != nil {
			log.Fatalf("Failed to get peer details: %v", err)
		}

		jsonData, _ := json.MarshalIndent(peer, "", "  ")
		fmt.Println(string(jsonData))
	}

	// Example: GetPeerSystem (test connectivity)
	if len(peers) > 0 {
		system, err := client.GetPeerSystem(peers[0])
		if err != nil {
			fmt.Printf("Peer not reachable: %v\n", err)
		} else {
			jsonData, _ := json.MarshalIndent(system, "", "  ")
			fmt.Println(string(jsonData))
		}
	}

	// Example: CreateOrUpdatePeer
	request := &types.PeerCreateRequest{
		URL:      "http://host.docker.internal:8043",
		Username: "orthanc",
		Password: "orthanc",
	}
	err = client.CreateOrUpdatePeer("TEST_PEER", request)
	if err != nil {
		log.Fatalf("Failed to create peer: %v", err)
	}

	// Example: StoreToPeer (single resource)
	err = client.StoreToPeer("TEST_PEER", "0f74e061-061039a7-aa467254-1f777b3f-480dede3")
	if err != nil {
		log.Fatalf("Failed to store to peer: %v", err)
	}

	// Example: StoreToPeerWithOptions
	storeRequest := &types.PeerStoreRequest{
		Resources:   []string{"0f74e061-061039a7-aa467254-1f777b3f-480dede3", "182b9f4f-87bbfed0-62269def-a8301708-135c2420"},
		Compress:    gorthanc.BoolPtr(true),
		Synchronous: gorthanc.BoolPtr(true),
	}
	storeResult, err := client.StoreToPeerWithOptions("TEST_PEER", storeRequest)
	if err != nil {
		log.Fatalf("Failed to store to peer: %v", err)
	}
	jsonData, _ := json.MarshalIndent(storeResult, "", "  ")
	fmt.Println(string(jsonData))

	// Example: DeletePeer
	err = client.DeletePeer("TEST_PEER")
	if err != nil {
		log.Fatalf("Failed to delete peer: %v", err)
	}
}
