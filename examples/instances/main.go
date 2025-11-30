package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"os"
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


	// Example: GetAllInstances

	params := &types.InstancesQueryParams{
		Limit: 100,
	}
	instances, err := client.GetAllInstances(params)
	if err != nil {
		log.Fatalf("Failed to get instances: %v", err)
	}
	fmt.Printf("Found %d instances \n", len(instances))
	fmt.Println(instances)


	// Example: GetInstanceDetails

	fmt.Println("Logging information about the first 5 instances")
	for _, instanceId := range instances[:min(5, len(instances))] {
		instanceDetails, err := client.GetInstanceDetails(instanceId)
		if err != nil { 
			log.Fatalf("Failed to get instance details: %v ", err)
		}
		jsonData, err := json.MarshalIndent(instanceDetails, "", "  ")
		if err != nil { 
			log.Fatalf("Failed to marshal instance: %v ", err)
		}
		fmt.Println(string(jsonData))
	}


	// Example: DownloadInstance

	filepath := "./tmp/" + instances[1] + ".dcm"

	output, err := os.Create(filepath)
	if err != nil {
		log.Fatalf("Failed to write file on disk")
	}
	defer output.Close()

	resp, err := client.DownloadDicomFile(instances[1])
	if err != nil {
		log.Fatalf("Failed to download DICOM file: %v", err)
	}
	defer resp.Body.Close()

	writer := bufio.NewWriterSize(output, 256 * 1024)
	defer writer.Flush()

	buf := make([]byte, 128 * 1024)

	_, err = io.CopyBuffer(writer, resp.Body, buf)
	if err != nil {
		log.Fatalf("Failed to write the file: %v", err)
	}

	if err := writer.Flush(); err != nil {
		log.Fatalf("Failed to flush: %v", err)
	}

	fmt.Printf("Instance %s downloaded with success. \n", instances[1])

	
	// Example: DeleteInstance (Commented to not delete nothing)

	if len(instances) > 1 {
		err := client.DeleteInstance(instances[1])
		if err != nil {
			log.Fatalf("Failed to delete instance: %v", err)
		}
		fmt.Printf("Instance %s deleted with success! \n", instances[1])
	}


	// Example< UploadDicomFile

	file, err := os.Open(filepath)

	if err != nil {
		log.Fatalf("Failed to read dicom file: %v", err)
	}

	defer file.Close()

	reader := bufio.NewReaderSize(file, 256*1024)

	result, err := client.UploadDicomFile(reader)

	if err != nil {
		log.Fatalf("Failed to upload the dicom file: %v", err)
	}

	fmt.Printf("Created instance with ID : %s and status: %s \n", result.ID, result.Status)


	// Example: AnonymizeInstance

	if len(instances) > 0 {
		insId := "0c3c027d-784ff6e6-788354b4-2b90ac5f-e6307490"
		filepath := "./tmp/" + "anonymized_" + insId + ".dcm"

		output, err := os.Create(filepath)
		if err != nil {
			log.Fatalf("Failed to write file on disk")
		}
		defer output.Close()

		anonymizeRequest := &types.InstancesAnonymizeRequest{
			Force:      gorthanc.BoolPtr(false),
			KeepSource: gorthanc.BoolPtr(true),
		}

		resp, err := client.AnonymizeInstance(insId, anonymizeRequest)
		if err != nil {
			log.Fatalf("Failed to download anonymized DICOM file: %v", err)
		}
		defer resp.Body.Close()

		writer := bufio.NewWriterSize(output, 256 * 1024)
		defer writer.Flush()

		buf := make([]byte, 128 * 1024)

		_, err = io.CopyBuffer(writer, resp.Body, buf)
		if err != nil {
			log.Fatalf("Failed to write the anonymized dicom file: %v", err)
		}

		if err := writer.Flush(); err != nil {
			log.Fatalf("Failed to flush: %v", err)
		}

		fmt.Printf("Anonymized instance Instance %s downloaded with success. \n", instances[1])
	}

}
