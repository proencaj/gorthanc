package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

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

	// Example: GetStudies
	params := &types.StudiesQueryParams{
		Limit: 5,
	}
	studies, err := client.GetStudies(params)
	if err != nil {
		log.Fatalf("Failed to get studies: %v", err)
	}
	fmt.Printf("Found %d studies \n", len(studies))
	fmt.Println(studies)

	// Example: GetStudiesExpanded
	expandParams := &types.StudiesQueryParams{
		Limit: 5,
		Expand: true,
	}
	
	expandedStudies, err := client.GetStudiesExpanded(expandParams)
	if err != nil {
		log.Fatalf("Failed to get studies: %v", err)
	}
	
	for _, study := range expandedStudies {
		jsonData, err := json.MarshalIndent(study, "", "  ")
		if err != nil { 
			log.Fatalf("Failed to marshal studies: %v ", err)
		}
		
		fmt.Println(string(jsonData))	
	}

	// Example: GetStudy

	for _, studyId := range studies {
		studyDetails, err := client.GetStudy(studyId)
		if err != nil { 
			log.Fatalf("Failed to get study details: %v ", err)
		}
		jsonData, err := json.MarshalIndent(studyDetails, "", "  ")
		if err != nil { 
			log.Fatalf("Failed to marshal study: %v ", err)
		}
		fmt.Println(string(jsonData))	
	}

	// Example: DeleteStudy

	// if len(studies) > 1 {
	// 	err := client.DeleteStudy(expandedStudies[1].ID)
	// 	if err != nil {
	// 		log.Fatalf("Failed to delet study: %v", err)
	// 	}
	// 	fmt.Printf("Study %s deleted with success! \n", expandedStudies[1].ID)
	// }

	// Example: AnonymizeStudy

	if len(expandedStudies) > 0 {
		anonymizeRequest := &types.StudyAnonymizeRequest{
			Force:      gorthanc.BoolPtr(false),
			Permissive: gorthanc.BoolPtr(false),
			KeepSource: gorthanc.BoolPtr(true),
		}
		anonymizedStudy, err := client.AnonymizeStudy(expandedStudies[0].ID, anonymizeRequest)
		if err != nil {
			log.Fatalf("Failed to anonymize study: %v", err)
		}

		jsonData, _ := json.MarshalIndent(anonymizedStudy, "", "  ")
		fmt.Println(string(jsonData))
	}

	// Example: DownloadStudyArchive

	if len(expandedStudies) > 0 {
		archive, err := client.DownloadStudyArchive(expandedStudies[0].ID)
		if err != nil {
			log.Fatalf("Failed to download study archive: %v", err)
		}
		defer archive.Body.Close()

		output, err := os.Create("tmp/" + "study_" + expandedStudies[0].ID + ".zip")
		if err != nil {
			log.Fatalf("Failed to write file: %v", err)
		}
		defer output.Close()

		_, err = io.Copy(output, archive.Body)
		if err != nil {
			log.Fatalf("Failed to create output file: %v", err)
		}
		fmt.Printf("Study archive downloaded to: study_%s.zip\n", expandedStudies[0].ID)
	}
}
