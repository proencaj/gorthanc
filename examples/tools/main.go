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

	// Example 1: Find all patients
	fmt.Println("=== Example 1: Find all patients ===")
	findPatientsRequest := &types.ToolsFindRequest{
		Level: types.ResourceLevelPatient,
		Query: map[string]string{},
	}
	patientIDs, err := client.Find(findPatientsRequest)
	if err != nil {
		log.Fatalf("Failed to find patients: %v", err)
	}
	fmt.Printf("Found %d patients\n", len(patientIDs))
	jsonData, _ := json.MarshalIndent(patientIDs, "", "  ")
	fmt.Println(string(jsonData))
	fmt.Println()

	// Example 2: Find patients by name pattern
	fmt.Println("=== Example 2: Find patients by name pattern ===")
	findByNameRequest := &types.ToolsFindRequest{
		Level: types.ResourceLevelPatient,
		Query: map[string]string{
			"PatientName": "*",
		},
	}
	results, err := client.Find(findByNameRequest)
	if err != nil {
		log.Fatalf("Failed to find patients by name: %v", err)
	}
	fmt.Printf("Found %d patients matching pattern\n", len(results))
	jsonData, _ = json.MarshalIndent(results, "", "  ")
	fmt.Println(string(jsonData))
	fmt.Println()

	// Example 3: Find studies with limit
	fmt.Println("=== Example 3: Find studies with limit ===")
	limit := 5
	findStudiesRequest := &types.ToolsFindRequest{
		Level: types.ResourceLevelStudy,
		Query: map[string]string{},
		Limit: &limit,
	}
	studyIDs, err := client.Find(findStudiesRequest)
	if err != nil {
		log.Fatalf("Failed to find studies: %v", err)
	}
	fmt.Printf("Found %d studies (limited to %d)\n", len(studyIDs), limit)
	jsonData, _ = json.MarshalIndent(studyIDs, "", "  ")
	fmt.Println(string(jsonData))
	fmt.Println()

	// Example 4: Find studies by date range
	fmt.Println("=== Example 4: Find studies by date range ===")
	findByDateRequest := &types.ToolsFindRequest{
		Level: types.ResourceLevelStudy,
		Query: map[string]string{
			"StudyDate": "20240101-20241231",
		},
	}
	dateResults, err := client.Find(findByDateRequest)
	if err != nil {
		log.Fatalf("Failed to find studies by date: %v", err)
	}
	fmt.Printf("Found %d studies in date range\n", len(dateResults))
	jsonData, _ = json.MarshalIndent(dateResults, "", "  ")
	fmt.Println(string(jsonData))
	fmt.Println()

	// Example 5: Find series with specific modality
	fmt.Println("=== Example 5: Find series with specific modality ===")
	findByModalityRequest := &types.ToolsFindRequest{
		Level: types.ResourceLevelSeries,
		Query: map[string]string{
			"Modality": "CT",
		},
	}
	modalityResults, err := client.Find(findByModalityRequest)
	if err != nil {
		log.Fatalf("Failed to find series by modality: %v", err)
	}
	fmt.Printf("Found %d CT series\n", len(modalityResults))
	jsonData, _ = json.MarshalIndent(modalityResults, "", "  ")
	fmt.Println(string(jsonData))
	fmt.Println()

	// Example 6: Find instances
	fmt.Println("=== Example 6: Find instances ===")
	instanceLimit := 3
	findInstancesRequest := &types.ToolsFindRequest{
		Level: types.ResourceLevelInstance,
		Query: map[string]string{},
		Limit: &instanceLimit,
	}
	instanceIDs, err := client.Find(findInstancesRequest)
	if err != nil {
		log.Fatalf("Failed to find instances: %v", err)
	}
	fmt.Printf("Found %d instances (limited to %d)\n", len(instanceIDs), instanceLimit)
	jsonData, _ = json.MarshalIndent(instanceIDs, "", "  ")
	fmt.Println(string(jsonData))
	fmt.Println()

	// Example 7: Find patients with expanded details
	fmt.Println("=== Example 7: Find patients with expanded details ===")
	expandLimit := 2
	findExpandedRequest := &types.ToolsFindRequest{
		Level: types.ResourceLevelPatient,
		Query: map[string]string{},
		Limit: &expandLimit,
	}
	expandedPatients, err := client.FindExpanded(findExpandedRequest)
	if err != nil {
		log.Fatalf("Failed to find expanded patients: %v", err)
	}
	fmt.Printf("Found %d expanded patients (limited to %d)\n", len(expandedPatients), expandLimit)
	for i, patient := range expandedPatients {
		fmt.Printf("\n--- Patient %d ---\n", i+1)
		fmt.Printf("ID: %s\n", patient.ID)
		fmt.Printf("Type: %s\n", patient.Type)
		fmt.Printf("IsStable: %t\n", patient.IsStable)
		fmt.Printf("LastUpdate: %s\n", patient.LastUpdate)
		fmt.Printf("Studies: %v\n", patient.Studies)
		fmt.Printf("MainDicomTags:\n")
		tagData, _ := json.MarshalIndent(patient.MainDicomTags, "  ", "  ")
		fmt.Printf("  %s\n", string(tagData))
	}
	fmt.Println()

	// Example 8: Find studies with expanded details
	fmt.Println("=== Example 8: Find studies with expanded details ===")
	studyExpandLimit := 2
	findStudiesExpandedRequest := &types.ToolsFindRequest{
		Level: types.ResourceLevelStudy,
		Query: map[string]string{},
		Limit: &studyExpandLimit,
	}
	expandedStudies, err := client.FindExpanded(findStudiesExpandedRequest)
	if err != nil {
		log.Fatalf("Failed to find expanded studies: %v", err)
	}
	fmt.Printf("Found %d expanded studies (limited to %d)\n", len(expandedStudies), studyExpandLimit)
	for i, study := range expandedStudies {
		fmt.Printf("\n--- Study %d ---\n", i+1)
		fmt.Printf("ID: %s\n", study.ID)
		fmt.Printf("Type: %s\n", study.Type)
		fmt.Printf("ParentPatient: %s\n", study.ParentPatient)
		fmt.Printf("Series: %v\n", study.Series)
		fmt.Printf("MainDicomTags:\n")
		tagData, _ := json.MarshalIndent(study.MainDicomTags, "  ", "  ")
		fmt.Printf("  %s\n", string(tagData))
	}
	fmt.Println()

	// Example 9: Get and Set Log Level
	fmt.Println("=== Example 9: Get and Set Log Level ===")
	currentLevel, err := client.GetLogLevel()
	if err != nil {
		log.Fatalf("Failed to get log level: %v", err)
	}
	fmt.Printf("Current log level: %s\n", currentLevel)

	// Set to verbose
	err = client.SetLogLevel(types.LogLevelVerbose)
	if err != nil {
		log.Fatalf("Failed to set log level to verbose: %v", err)
	}
	fmt.Println("Log level set to: verbose")

	// Verify the change
	newLevel, err := client.GetLogLevel()
	if err != nil {
		log.Fatalf("Failed to get log level: %v", err)
	}
	fmt.Printf("New log level: %s\n", newLevel)

	// Restore to original level
	err = client.SetLogLevel(currentLevel)
	if err != nil {
		log.Fatalf("Failed to restore log level: %v", err)
	}
	fmt.Printf("Log level restored to: %s\n", currentLevel)
	fmt.Println()

	// Example 10: Reset Orthanc (commented out for safety)
	// WARNING: This will restart the Orthanc server!
	// Uncomment only if you want to test the reset functionality
	/*
	fmt.Println("=== Example 10: Reset Orthanc ===")
	err = client.Reset()
	if err != nil {
		log.Fatalf("Failed to reset Orthanc: %v", err)
	}
	fmt.Println("Orthanc has been restarted successfully")
	fmt.Println()
	*/

	// Example 11: Shutdown Orthanc (commented out for safety)
	// WARNING: This will shutdown the Orthanc server!
	// Uncomment only if you want to test the shutdown functionality
	/*
	fmt.Println("=== Example 11: Shutdown Orthanc ===")
	err = client.Shutdown()
	if err != nil {
		log.Fatalf("Failed to shutdown Orthanc: %v", err)
	}
	fmt.Println("Orthanc has been shutdown successfully")
	*/
}
