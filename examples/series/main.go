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


	// Example: GetSeries

	params := &types.SeriesQueryParams{
		Limit: 5,
	}
	series, err := client.GetSeries(params)
	if err != nil {
		log.Fatalf("Failed to get series: %v", err)
	}
	fmt.Printf("Found %d series \n", len(series))
	fmt.Println(series)


	// // Example: GetSeriesExpanded

	expandParams := &types.SeriesQueryParams{
		Limit: 5,
		Expand: true,
	}
	
	expandedSeries, err := client.GetSeriesExpanded(expandParams)
	if err != nil {
		log.Fatalf("Failed to get series: %v", err)
	}
	
	for _, study := range expandedSeries {
		jsonData, err := json.MarshalIndent(study, "", "  ")
		if err != nil { 
			log.Fatalf("Failed to marshal series: %v ", err)
		}
		
		fmt.Println(string(jsonData))	
	}


	// Example: GetSeriesDetail

	for _, serieId := range series[:min(5, len(series))] {
		serieDetails, err := client.GetSeriesDetail(serieId)
		if err != nil { 
			log.Fatalf("Failed to get serie details: %v ", err)
		}
		jsonData, err := json.MarshalIndent(serieDetails, "", "  ")
		if err != nil { 
			log.Fatalf("Failed to marshal serie: %v ", err)
		}
		fmt.Println(string(jsonData))
	}


	// Example: DeleteSerie (Commented to not delete nothing)

	// if len(series) > 1 {
	// 	err := client.DeleteSeries(expandedSeries[1].ID)
	// 	if err != nil {
	// 		log.Fatalf("Failed to delete serie: %v", err)
	// 	}
	// 	fmt.Printf("Serie %s deleted with success! \n", expandedSeries[1].ID)
	// }


	// Example: AnonymizeSeries (Only synchronous)

	if len(expandedSeries) > 0 {
		anonymizeRequest := &types.SeriesAnonymizeRequest{
			Force:      gorthanc.BoolPtr(false),
			Permissive: gorthanc.BoolPtr(false),
			KeepSource: gorthanc.BoolPtr(true),
		}
		anonymizedSeries, err := client.AnonymizeSeries(expandedSeries[0].ID, anonymizeRequest)
		if err != nil {
			log.Fatalf("Failed to anonymize series: %v", err)
		}

		jsonData, _ := json.MarshalIndent(anonymizedSeries, "", "  ")
		fmt.Println(string(jsonData))
	}


	// Example: DownloadSeriesArchive

	if len(expandedSeries) > 0 {
		archive, err := client.DownloadSeriesArchive(expandedSeries[0].ID)
		if err != nil {
			log.Fatalf("Failed to download series archive: %v", err)
		}
		defer archive.Body.Close()

		output, err := os.Create("tmp/" + "series_" + expandedSeries[0].ID + ".zip")
		if err != nil {
			log.Fatalf("Failed to write file: %v", err)
		}
		defer output.Close()

		_, err = io.Copy(output, archive.Body)
		if err != nil {
			log.Fatalf("Failed to create output file: %v", err)
		}
		fmt.Printf("Series archive downloaded to: series_%s.zip\n", expandedSeries[0].ID)
	}


	// Example: GetSeriesStatistics

	stats, err := client.GetSeriesStatistics(expandedSeries[0].ID)
	if err != nil {
		log.Fatalf("Failed get series stats: %v", err)
	}

	jsonData, err := json.MarshalIndent(stats, "", "  ")
	if err != nil { 
		log.Fatalf("Failed to marshal series: %v ", err)
	}
	
	fmt.Println(string(jsonData))	


	// Example: GetSeriesInstances

	if len(expandedSeries) > 0 {
		instanceIDs, err := client.GetSeriesInstances(expandedSeries[0].ID)
		if err != nil {
			log.Fatalf("Failed to get series instances: %v", err)
		}
		fmt.Printf("Found %d instances in series %s\n", len(instanceIDs), expandedSeries[0].ID)
		fmt.Println(instanceIDs)
	}

	// Example: GetSeriesInstancesExpanded

	if len(expandedSeries) > 0 {
		instances, err := client.GetSeriesInstancesExpanded(expandedSeries[0].ID)
		if err != nil {
			log.Fatalf("Failed to get expanded series instances: %v", err)
		}
		fmt.Printf("Found %d instances with full details in series %s\n", len(instances), expandedSeries[0].ID)

		for _, i := range instances {
			jsonData, err := json.MarshalIndent(i, "", "  ")
			if err != nil {
				log.Fatalf("Failed to marshal instance: %v", err)
			}
			fmt.Println(string(jsonData))
		}
	}
}
