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

	// Example: GetDicomWebServers
	servers, err := client.GetDicomWebServers()
	if err != nil {
		log.Fatalf("Failed to get DICOMweb servers: %v", err)
	}
	fmt.Println("DICOMweb Servers:")
	fmt.Println(servers)

	// Example: GetDicomWebServersExpanded
	expandedServers, err := client.GetDicomWebServersExpanded()
	if err != nil {
		log.Fatalf("Failed to get expanded DICOMweb servers: %v", err)
	}

	jsonData, err := json.MarshalIndent(expandedServers, "", "  ")
	if err != nil { 
		log.Fatalf("Failed to marshal instance: %v ", err)
	}
	fmt.Println(string(jsonData))


	// Example: CreateOrUpdateDicomWebServer
	createRequest := &types.DicomWebServerCreateRequest{
		Url:      "http://example-dicomweb-server:8042/dicom-web",
		Username: "username",
		Password: "password",
		HasDelete: gorthanc.BoolPtr(true),
		ChunkedTransfers: gorthanc.BoolPtr(true),
		HasWadoRsUniversalTransferSyntax: gorthanc.BoolPtr(true),
	}
	err = client.CreateOrUpdateDicomWebServer("sample", createRequest)
	if err != nil {
		log.Fatalf("Failed to create DICOMweb server: %v", err)
	}
	fmt.Println("DICOMweb server 'sample' created successfully")

	// Example: CreateOrUpdateDicomWebServer (minimal configuration)
	minimalRequest := &types.DicomWebServerCreateRequest{
		Url: "http://another-server:8042/dicom-web",
	}
	err = client.CreateOrUpdateDicomWebServer("minimal-server", minimalRequest)
	if err != nil {
		log.Fatalf("Failed to create minimal DICOMweb server: %v", err)
	}
	fmt.Println("DICOMweb server 'minimal-server' created successfully")

	// Example: CreateOrUpdateDicomWebServer (for older Orthanc versions)
	legacyRequest := &types.DicomWebServerCreateRequest{
		Url:      "http://legacy-orthanc:8042/dicom-web",
		Username: "orthanc",
		Password: "orthanc",
		ChunkedTransfers: gorthanc.BoolPtr(false), // Orthanc <= 1.5.6
		HasWadoRsUniversalTransferSyntax: gorthanc.BoolPtr(false), // DICOMweb plugin <= 1.0
	}
	err = client.CreateOrUpdateDicomWebServer("legacy-server", legacyRequest)
	if err != nil {
		log.Fatalf("Failed to create legacy DICOMweb server: %v", err)
	}
	fmt.Println("DICOMweb server 'legacy-server' created successfully")


	// Example: Update existing DICOMweb server
	updateRequest := &types.DicomWebServerCreateRequest{
		Url:      "http://updated-server:8042/dicom-web",
		Username: "new-username",
		Password: "new-password",
		HasDelete: gorthanc.BoolPtr(false),
	}
	err = client.CreateOrUpdateDicomWebServer("sample", updateRequest)
	if err != nil {
		log.Fatalf("Failed to update DICOMweb server: %v", err)
	}
	fmt.Println("\nDICOMweb server 'sample' updated successfully")

	// Example: GetDicomWebServers (after creating new servers)
	updatedServers, err := client.GetDicomWebServers()
	if err != nil {
		log.Fatalf("Failed to get updated DICOMweb servers: %v", err)
	}
	fmt.Println("\nUpdated DICOMweb Servers list:")
	fmt.Println(updatedServers)

	// Example: DeleteDicomWebServer
	err = client.DeleteDicomWebServer("minimal-server")
	if err != nil {
		log.Fatalf("Failed to delete DICOMweb server: %v", err)
	}
	fmt.Println("\nDICOMweb server 'minimal-server' deleted successfully")

	err = client.DeleteDicomWebServer("legacy-server")
	if err != nil {
		log.Fatalf("Failed to delete DICOMweb server: %v", err)
	}
	fmt.Println("DICOMweb server 'legacy-server' deleted successfully")

	err = client.DeleteDicomWebServer("sample")
	if err != nil {
		log.Fatalf("Failed to delete DICOMweb server: %v", err)
	}
	fmt.Println("DICOMweb server 'sample' deleted successfully")
}
