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

	// Fetch system-wide statistics
	fmt.Println("Fetching Orthanc system statistics...")
	stats, err := client.GetSystemStatistics() // Note: In gorthanc, this is usually GetStatistics()
	if err != nil {
		if gorthanc.IsUnauthorized(err) {
			log.Fatal("Authentication failed - check your credentials")
		}
		log.Fatalf("Failed to get statistics: %v", err)
	}

	// Display Statistics Information
	fmt.Println("\n=== Orthanc System Statistics ===")
	
	// Entity Counts
	fmt.Printf("Total Patients:        %d\n", stats.CountPatients)
	fmt.Printf("Total Studies:         %d\n", stats.CountStudies)
	fmt.Printf("Total Series:          %d\n", stats.CountSeries)
	fmt.Printf("Total Instances:       %d\n", stats.CountInstances)
	
	fmt.Println("\n--- Storage Information ---")
	
	// Disk Usage
	fmt.Printf("Disk Size (MB):        %d MB\n", stats.TotalDiskSizeMB)
	fmt.Printf("Disk Size (Bytes):     %s\n", stats.TotalDiskSize)
	
	// Compression Info
	fmt.Printf("Uncompressed (MB):     %d MB\n", stats.TotalUncompressedSizeMB)
	fmt.Printf("Uncompressed (Bytes):  %s\n", stats.TotalUncompressedSize)

	// Calculate Compression Ratio if applicable
	if stats.TotalUncompressedSizeMB > 0 {
		ratio := float64(stats.TotalDiskSizeMB) / float64(stats.TotalUncompressedSizeMB) * 100
		fmt.Printf("Compression Ratio:     %.2f%%\n", ratio)
	}

	fmt.Println("\n✓ Statistics retrieved successfully!")
}