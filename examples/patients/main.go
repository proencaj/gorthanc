package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/proencaj/gorthanc"
	"github.com/proencaj/gorthanc/types"
)

func main() {
	// Create a new Orthanc client
	// Replace with your Orthanc server URL and credentials
	client, err := gorthanc.NewClient(
		"http://localhost:8243",
		gorthanc.WithBasicAuth("orthanc", "orthanc"),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}


	// Example: GetPatients

	params := &types.PatientQueryParams{
		Limit: 5,		
	}
	patients, err := client.GetPatients(params)
	if err != nil {
		log.Fatalf("Failed to get patients: %v", err)
	}
	fmt.Printf("Found %d patients\n", len(patients))
	fmt.Println(patients)

	
	// Example: GetPatientDetails

	for _, patientID := range patients {
		patientDetails, err := client.GetPatientDetails(patientID)		
		if err != nil {
			log.Fatalf("Failed to get patient details: %v ", err)
		}
		jsonData, err := json.MarshalIndent(patientDetails, "", "  ")
		if err != nil { 
			log.Fatalf("Failed to marshal patient: %v ", err)
		}
		fmt.Println(string(jsonData))
	}


	// Example: AnonymizePatient

	var anonymizedPatientID string

	if len(patients) > 0 {
		anonymizeRequest := &types.PatientAnonymizeRequest{
			Force:      gorthanc.BoolPtr(false),
			Permissive: gorthanc.BoolPtr(false),
			KeepSource: gorthanc.BoolPtr(true),
		}
		anonymizedPatient, err := client.AnonymizePatient(patients[0], anonymizeRequest)
		if err != nil {
			log.Fatalf("Failed to anonymize series: %v", err)
		}

		anonymizedPatientID = anonymizedPatient.ID

		jsonData, _ := json.MarshalIndent(anonymizedPatient, "", "  ")
		fmt.Println(string(jsonData))
	}


	// Example: DeletePatient

	if len(patients) > 0 {
		err := client.DeletePatient(anonymizedPatientID)
		if err != nil { 
			log.Fatalf("Failed to delete patient: %v", err)
		}
		fmt.Printf("Patient %s deleted with success! \n", anonymizedPatientID)

	}

	
	// Example GetPatientStatistics

	if len(patients) > 0 {
		stats, err := client.GetPatientStatistics(patients[0])
		if err != nil {
			log.Fatalf("Failed get series stats: %v", err)
		}

		jsonData, err := json.MarshalIndent(stats, "", "  ")
		if err != nil { 
			log.Fatalf("Failed to marshal series: %v ", err)
		}
		
		fmt.Println(string(jsonData))
	}
}
