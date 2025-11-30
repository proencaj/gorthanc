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
		"http://localhost:8243",
		gorthanc.WithBasicAuth("orthanc", "orthanc"),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Example: GetModalities
	modalities, err := client.GetModalities()
	if err != nil {
		log.Fatalf("Failed to get modalities: %v", err)
	}
	fmt.Println(modalities)

	// Example: GetModalityDetails
	if len(modalities) > 0 {
		modality, err := client.GetModalityDetails(modalities[0])
		if err != nil {
			log.Fatalf("Failed to get modality details: %v", err)
		}

		jsonData, _ := json.MarshalIndent(modality, "", "  ")
		fmt.Println(string(jsonData))
	}

	// Example: EchoModality
	if len(modalities) > 0 {
		err = client.EchoModality(modalities[0])
		if err != nil {
			fmt.Printf("Modality not reachable: %v\n", err)
		}
	}

	// Example: CreateOrUpdateModality
	request := &types.ModalityCreateRequest{
		AET:  "TEST_MODALITY",
		Host: "localhost",
		Port: 4242,
	}
	err = client.CreateOrUpdateModality("TEST_MODALITY", request)
	if err != nil {
		log.Fatalf("Failed to create modality: %v", err)
	}

	// Example: FindInModality
	findRequest := &types.ModalityFindRequest{
		Level:     "Study",
		Normalize: gorthanc.BoolPtr(true),
		Query: map[string]string{
			"PatientName": "*",
		},
	}
	findResults, err := client.FindInModality("PACS", findRequest)
	if err != nil {
		log.Fatalf("Failed to query modality: %v", err)
	}
	jsonData, _ := json.MarshalIndent(findResults, "", "  ")
	fmt.Println(string(jsonData))

	// Example: MoveFromModality
	moveRequest := &types.ModalityMoveRequest{
		Level:        "Study",
		TargetAet:    "ORTHANC",
		Asynchronous: gorthanc.BoolPtr(false),
		Permissive:   gorthanc.BoolPtr(false),
		Timeout:      30,
		Resources: []map[string]interface{}{
			{"StudyInstanceUID": "1.2.840.113619.2.55.3.123456789"},
		},
	}
	moveResults, err := client.MoveFromModality("PACS", moveRequest)
	if err != nil {
		log.Fatalf("Failed to move study: %v", err)
	}
	jsonData, _ = json.MarshalIndent(moveResults, "", "  ")
	fmt.Println(string(jsonData))

	// Example: GetFromModality
	getRequest := &types.ModalityGetRequest{
		Level:        "Study",
		Asynchronous: gorthanc.BoolPtr(false),
		Permissive:   gorthanc.BoolPtr(false),
		Timeout:      30,
		Resources: []map[string]interface{}{
			{"StudyInstanceUID": "1.2.840.113619.2.55.3.123456789"},
		},
	}
	err = client.GetFromModality("PACS", getRequest)
	if err != nil {
		log.Fatalf("Failed to get study: %v", err)
	}

	// Example: StoreToModality
	err = client.StoreToModality("PACS", "study-orthanc-id")
	if err != nil {
		log.Fatalf("Failed to store study: %v", err)
	}

	// Example: DeleteModality
	err = client.DeleteModality("TEST_MODALITY")
	if err != nil {
		log.Fatalf("Failed to delete modality: %v", err)
	}
}
