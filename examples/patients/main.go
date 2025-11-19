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
			Force: false,
			Permissive: false,
			KeepSource: true,
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
	

	// 	// Example 5: Get patient statistics
	// 	fmt.Printf("5. Fetching statistics for patient: %s...\n", patientIDs[0])
	// 	stats, err := client.GetPatientStatistics(patientIDs[0])
	// 	if err != nil {
	// 		log.Fatalf("Failed to get patient statistics: %v", err)
	// 	}


	// if len(series) > 1 {
	// 	err := client.DeleteSeries(expandedSeries[1].ID)
	// 	if err != nil {
	// 		log.Fatalf("Failed to delete serie: %v", err)
	// 	}
	// 	fmt.Printf("Serie %s deleted with success! \n", expandedSeries[1].ID)
	// }




	// fmt.Println("=== Orthanc Patients Examples ===\n")

	// // Example 1: Get all patient IDs
	// fmt.Println("1. Fetching all patient IDs...")
	// patientIDs, err := client.GetPatients(nil)
	// if err != nil {
	// 	log.Fatalf("Failed to get patients: %v", err)
	// }
	// fmt.Printf("   Found %d patients\n", len(patientIDs))
	// if len(patientIDs) > 0 {
	// 	fmt.Printf("   First patient ID: %s\n", patientIDs[0])
	// }
	// fmt.Println()

	// // Example 2: Get paginated patient IDs
	// fmt.Println("2. Fetching first 5 patients (paginated)...")
	// params := &types.PatientQueryParams{
	// 	Limit: 5,
	// 	Since: 0,
	// }
	// paginatedPatients, err := client.GetPatients(params)
	// if err != nil {
	// 	log.Fatalf("Failed to get paginated patients: %v", err)
	// }
	// fmt.Printf("   Retrieved %d patients\n", len(paginatedPatients))
	// for i, patientID := range paginatedPatients {
	// 	fmt.Printf("   [%d] %s\n", i+1, patientID)
	// }
	// fmt.Println()

	// // Example 4: Get specific patient details
	// if len(patientIDs) > 0 {
	// 	fmt.Printf("4. Fetching details for specific patient: %s...\n", patientIDs[0])
	// 	patient, err := client.GetPatient(patientIDs[0])
	// 	if err != nil {
	// 		log.Fatalf("Failed to get patient details: %v", err)
	// 	}

	// 	fmt.Printf("\n   Patient Details:\n")
	// 	fmt.Printf("   ├─ Patient Name:    %s\n", patient.MainDicomTags.PatientName)
	// 	fmt.Printf("   ├─ Patient ID:      %s\n", patient.MainDicomTags.PatientID)
	// 	fmt.Printf("   ├─ Birth Date:      %s\n", patient.MainDicomTags.PatientBirthDate)
	// 	fmt.Printf("   ├─ Sex:             %s\n", patient.MainDicomTags.PatientSex)
	// 	fmt.Printf("   └─ Number of Studies: %d\n", len(patient.Studies))

	// 	if len(patient.Studies) > 0 {
	// 		fmt.Printf("\n   Study IDs:\n")
	// 		for i, studyID := range patient.Studies {
	// 			fmt.Printf("   [%d] %s\n", i+1, studyID)
	// 		}
	// 	}
	// 	fmt.Println()

	// 	// Example 5: Get patient statistics
	// 	fmt.Printf("5. Fetching statistics for patient: %s...\n", patientIDs[0])
	// 	stats, err := client.GetPatientStatistics(patientIDs[0])
	// 	if err != nil {
	// 		log.Fatalf("Failed to get patient statistics: %v", err)
	// 	}

	// 	fmt.Printf("\n   Patient Statistics:\n")
	// 	fmt.Printf("   ├─ Disk Size:       %s (%d MB)\n", stats.DiskSize, stats.DiskSizeMB)
	// 	fmt.Printf("   ├─ Studies Count:   %d\n", stats.CountStudies)
	// 	fmt.Printf("   ├─ Series Count:    %d\n", stats.CountSeries)
	// 	fmt.Printf("   └─ Instance Count:  %d\n", stats.CountInstances)
	// 	if stats.UncompressedSize != "" {
	// 		fmt.Printf("   └─ Uncompressed:    %s (%d MB)\n",
	// 			stats.UncompressedSize, stats.UncompressedSizeMB)
	// 	}
	// 	fmt.Println()
	// }

	// err = client.DeletePatient(patientIDs[0])

	

	// Example 9: Advanced - Delete patient (commented out for safety)
// 	fmt.Println("9. Example: Delete patient (commented out)")
// 	fmt.Println("   Uncomment the code below to delete a patient:")
// 	fmt.Println(`
//    // WARNING: This will permanently delete the patient and all associated data!
//    // remainingResources, err := client.DeletePatient(patientIDs[0])
//    // if err != nil {
//    //     log.Fatalf("Failed to delete patient: %v", err)
//    // }
//    // fmt.Printf("Patient deleted. Remaining resources: %+v\n", remainingResources)
// 	`)
// 	fmt.Println()

// 	fmt.Println("✓ All examples completed successfully!")
}
