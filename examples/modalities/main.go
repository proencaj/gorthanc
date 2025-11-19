package main

import (
	"fmt"
	"log"
	"log/slog"

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

	fmt.Println("=== Orthanc Modalities Examples ===")

	// Example 5: Create/Update a modality (commented out for safety)
	// request := &types.ModalityCreateRequest{
	// 	AET:  "ROTEADOR",
	// 	Host: "10.193.16.30",
	// 	Port: 4243,
	// }

	// request2 := &types.ModalityCreateRequest{
	// 	AET:  "ROTEADOR2",
	// 	Host: "localhost",
	// 	Port: 4243,
	// }

	// request3 := &types.ModalityCreateRequest{
	// 	AET:  "ROTEADOR3",
	// 	Host: "localhost",
	// 	Port: 4244,
	// }

	// err = client.CreateModality("ROTEADOR", request)

	// if err != nil {
	// 	log.Fatalf("Failed to create modality: %v", err)
	// }

	// err = client.CreateModality("ROTEADOR2", request2)

	// if err != nil {
	// 	log.Fatalf("Failed to create modality: %v", err)
	// }

	// err = client.CreateModality("ROTEADOR3", request3)

	// if err != nil {
	// 	log.Fatalf("Failed to create modality: %v", err)
	// }

	// fmt.Println("Modality created successfully!")

	// // Example 1: Get all configured modalities
	// fmt.Println("1. Fetching all configured modalities...")
	// modalities, err := client.GetModalities()
	// if err != nil {
	// 	log.Fatalf("Failed to get modalities: %v", err)
	// }
	// fmt.Printf("   Found %d configured modalities\n", len(modalities))
	// if len(modalities) > 0 {
	// 	for i, modalityName := range modalities {
	// 		fmt.Printf("   [%d] %s\n", i+1, modalityName)
	// 	}
	// } else {
	// 	fmt.Println("   No modalities configured")
	// }
	// fmt.Println()

	// Example 2: Get details for a specific modality
	// if len(modalities) > 0 {
	// 	fmt.Printf("2. Fetching details for modality: %s...\n", modalities[0])
	// 	modality, err := client.GetModalityDetails(modalities[0])
	// 	if err != nil {
	// 		log.Fatalf("Failed to get modality details: %v", err)
	// 	}

	// 	fmt.Printf("\n   Modality Configuration:\n")
	// 	fmt.Printf("   ├─ AET:          %s\n", modality.AET)
	// 	fmt.Printf("   ├─ Host:         %s\n", modality.Host)
	// 	fmt.Printf("   ├─ Port:         %d\n", modality.Port)
	// 	if modality.Manufacturer != "" {
	// 		fmt.Printf("   ├─ Manufacturer: %s\n", modality.Manufacturer)
	// 	}
	// 	if modality.Timeout > 0 {
	// 		fmt.Printf("   └─ Timeout:      %d seconds\n", modality.Timeout)
	// 	}
	// 	fmt.Println()

	// 	// Example 3: Test modality connection with C-ECHO
	// 	fmt.Printf("3. Testing connection to modality: %s...\n", modalities[0])
	// 	err = client.EchoModality(modalities[0])
	// 	if err != nil {
	// 		fmt.Printf("   ✗ Modality is not reachable: %v\n", err)
	// 	} else {
	// 		fmt.Printf("   ✓ Modality is reachable and responding\n")
	// 	}
	// 	fmt.Println()
	// }

	find_request := &types.ModalityFindRequest{
		Level: "Study",
		Normalize: true,
		Query: map[string]string{
			// "0010,0010": "ANTONIO DOS SANTOS PEREIRA",
			"PatientName": "ANTONIO DOS SANTOS PEREIRA",
		},
	}

	find_results, err := client.FindInModality("ORTHANC_02", find_request)

	if err != nil {
		log.Fatalf("Failed to query modality: %v", err)
	}

	// for i, result := range results {
	// 	fmt.Printf("Study %d: %+v\n", i+1, result)
	// }
	fmt.Println(find_results)

		// Example 12: C-MOVE operation
// 	fmt.Println("12. Example: Retrieve study with C-MOVE (commented out)")
// 	fmt.Println("    Uncomment the code below to retrieve a study:")
// 	fmt.Println(`
//    // request := &types.ModalityMoveRequest{
//    //     Level:     "Study",
//    //     TargetAet: "ORTHANC",
//    //     Query: map[string]string{
//    //         "StudyInstanceUID": "1.2.840.113619.2.55.3.123456789",
//    //     },
//    // }
//    // result, err := client.MoveFromModality("PACS", request)
//    // if err != nil {
//    //     log.Fatalf("Failed to move study: %v", err)
//    // }
//    // fmt.Printf("Move operation completed: %s\n", result.Description)
// 	`)
// 	fmt.Println()

move_request := &types.ModalityMoveRequest{
	Level: "Study",
	TargetAet: "ORTHANC_01",
	Asynchronous: false,
	Limit: 0,
	Timeout: 30,
	Priority: 0,
	Permissive: false,
	Resources: []map[string]interface{}{ 
		{ "StudyInstanceUID": "1.3.840.20240710.1503.20031000007811920" },
	},
}


move_results, err := client.MoveFromModality("ORTHANC_02", move_request)

if err != nil { 
	log.Fatalf("Failed to move study: %v", err)	
}

fmt.Println(move_results)

err = client.StoreToModality("ORTHANC_02", "86d42ea0-2d1628f9-e28942c0-a6067b20-8f245335")

if err != nil {
	log.Fatal("Study sent failure!")
}

slog.Info("Study successfuly sent")

get_request := &types.ModalityGetRequest{
	Level: "Study",
	Resources: []map[string]interface{}{
		{ "StudyInstanceUID": "1.3.840.20240710.1503.20031000007811920" },
	},
	Timeout: 30,
	Permissive: false,
	Asynchronous: false,
}

err = client.GetFromModality("ORTHANC_02", get_request)

if err != nil {
	log.Fatal("Study get failure!")
}

slog.Info("Study get successfuly")

err = client.DeleteModality("ORTHANC_02")

if err != nil {
	log.Fatal("Fail to delete modality!")
}

modalities, err := client.GetModalities()

if err != nil {
	log.Fatal("Fail to list modalities!")
}

slog.Info("Logging client modalities", "modalities", modalities)
}
