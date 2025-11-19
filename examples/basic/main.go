package main

import (
	"fmt"
	"log"

	"github.com/proencaj/gorthanc"
)

func main() {
	// Create a new Orthanc client
	// Replace with your Orthanc server URL and credentials
	client, err := gorthanc.NewClient(
		"http://localhost:8243",
		gorthanc.WithBasicAuth("orthanc", "orthanc"),
		// gorthanc.WithTimeout(30*time.Second),
	)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	// Get system information
	fmt.Println("Fetching Orthanc system information...")
	info, err := client.GetSystem()
	if err != nil {
		// Check for specific error types
		if gorthanc.IsUnauthorized(err) {
			log.Fatal("Authentication failed - check your credentials")
		}
		if gorthanc.IsNotFound(err) {
			log.Fatal("Endpoint not found - check your Orthanc server URL")
		}
		log.Fatalf("Failed to get system info: %v", err)
	}

	// Display system information
	fmt.Println("\n=== Orthanc System Information ===")
	fmt.Printf("Server Name:           %s\n", info.Name)
	fmt.Printf("Orthanc Version:       %s\n", info.Version)
	fmt.Printf("API Version:           %v\n", info.ApiVersion)
	fmt.Printf("Database Version:      %d\n", info.DatabaseVersion)
	fmt.Println()
	fmt.Printf("DICOM AET:             %s\n", info.DicomAet)
	fmt.Printf("DICOM Port:            %d\n", info.DicomPort)
	fmt.Printf("HTTP Port:             %d\n", info.HttpPort)
	fmt.Println()
	fmt.Printf("Plugins Enabled:       %t\n", info.PluginsEnabled)
	fmt.Printf("Check Revisions:       %t\n", info.CheckRevisions)

	if info.DatabaseBackendPlugin != "" {
		fmt.Printf("Database Plugin:       %s\n", info.DatabaseBackendPlugin)
	}

	if info.StorageAreaPlugin != "" {
		fmt.Printf("Storage Plugin:        %s\n", info.StorageAreaPlugin)
	}

	if info.MaximumStorageSize > 0 {
		fmt.Printf("Max Storage Size:      %d bytes (%.2f GB)\n", 
			info.MaximumStorageSize, 
			float64(info.MaximumStorageSize)/(1024*1024*1024))
	}

	if info.MaximumPatients > 0 {
		fmt.Printf("Max Patients:          %d\n", info.MaximumPatients)
	}

	fmt.Println("\n✓ Successfully connected to Orthanc server!")
}