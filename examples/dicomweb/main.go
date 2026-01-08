package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

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

	// =========================================================================
	// QIDO-RS Examples (Query)
	// =========================================================================

	fmt.Println("=== QIDO-RS Examples ===")
	fmt.Println()

	// Example: Search all studies
	fmt.Println("--- Search All Studies ---")
	studies, err := client.QidoSearchStudies(nil)
	if err != nil {
		log.Fatalf("Failed to search studies: %v", err)
	}
	fmt.Printf("Found %d studies\n", len(studies))
	printJSON(studies)

	// Example: Search studies with filters
	fmt.Println("\n--- Search Studies with Filters ---")
	studyParams := &types.QidoStudyQueryParams{
		QidoQueryParams: types.QidoQueryParams{
			Limit: 5,
		},
	}
	filteredStudies, err := client.QidoSearchStudies(studyParams)
	if err != nil {
		log.Fatalf("Failed to search filtered studies: %v", err)
	}
	fmt.Printf("Found %d studies (limited to 5)\n", len(filteredStudies))

	if len(studies) > 0 {
		// Get the first study's UID from the response
		studyUID := getStudyInstanceUID(studies[0])
		if studyUID != "" {
			fmt.Printf("\nUsing Study UID: %s\n", studyUID)

			// Example: Search series in a study
			fmt.Println("\n--- Search Series in Study ---")
			series, err := client.QidoSearchSeries(studyUID, nil)
			if err != nil {
				log.Fatalf("Failed to search series: %v", err)
			}
			fmt.Printf("Found %d series in study\n", len(series))
			printJSON(series)

			// If we have series, search for instances
			if len(series) > 0 {
				seriesUID := getSeriesInstanceUID(series[0])
				if seriesUID != "" {
					fmt.Printf("\nUsing Series UID: %s\n", seriesUID)

					// Example: Search instances in a series
					fmt.Println("\n--- Search Instances in Series ---")
					instances, err := client.QidoSearchInstances(studyUID, seriesUID, nil)
					if err != nil {
						log.Fatalf("Failed to search instances: %v", err)
					}
					fmt.Printf("Found %d instances\n", len(instances))

					// =========================================================
					// WADO-RS Examples (Retrieve)
					// =========================================================

					fmt.Println("\n=== WADO-RS Examples ===")
					fmt.Println()

					// Example: Retrieve study metadata
					fmt.Println("--- Retrieve Study Metadata ---")
					studyMetadata, err := client.WadoRsRetrieveStudyMetadata(studyUID)
					if err != nil {
						log.Printf("Failed to retrieve study metadata: %v", err)
					} else {
						fmt.Printf("Retrieved metadata for %d instances in study\n", len(studyMetadata))
					}
					printJSON(studyMetadata)

					// Example: Retrieve series metadata
					fmt.Println("\n--- Retrieve Series Metadata ---")
					seriesMetadata, err := client.WadoRsRetrieveSeriesMetadata(studyUID, seriesUID)
					if err != nil {
						log.Printf("Failed to retrieve series metadata: %v", err)
					} else {
						fmt.Printf("Retrieved metadata for %d instances in series\n", len(seriesMetadata))
					}

					// Example: Retrieve rendered instance (if we have instances)
					if len(instances) > 0 {
						instanceUID := getSOPInstanceUID(instances[0])
						if instanceUID != "" {
							fmt.Printf("\nUsing Instance UID: %s\n", instanceUID)

							// Example: Retrieve instance metadata
							fmt.Println("\n--- Retrieve Instance Metadata ---")
							instanceMetadata, err := client.WadoRsRetrieveInstanceMetadata(studyUID, seriesUID, instanceUID)
							if err != nil {
								log.Printf("Failed to retrieve instance metadata: %v", err)
							} else {
								fmt.Printf("Retrieved metadata for instance\n")
								printJSON(instanceMetadata)
							}

							// Example: Retrieve rendered instance as JPEG
							fmt.Println("\n--- Retrieve Rendered Instance ---")
							renderedParams := &types.WadoRsRenderedParams{
								Quality: 90,
							}
							resp, err := client.WadoRsRetrieveRenderedInstance(studyUID, seriesUID, instanceUID, renderedParams)
							if err != nil {
								log.Printf("Failed to retrieve rendered instance: %v", err)
							} else {
								defer resp.Body.Close()
								body, _ := io.ReadAll(resp.Body)
								fmt.Printf("Retrieved rendered instance: %d bytes, Content-Type: %s\n",
									len(body), resp.Header.Get("Content-Type"))

								// Save the image to tmp folder
								imagePath := "./tmp/dicomweb_rendered.jpg"
								if err := os.MkdirAll(filepath.Dir(imagePath), 0755); err != nil {
									log.Printf("Failed to create directory: %v", err)
								} else if err := os.WriteFile(imagePath, body, 0644); err != nil {
									log.Printf("Failed to save image: %v", err)
								} else {
									fmt.Printf("Image saved to: %s\n", imagePath)
								}
							}

							// Example: Retrieve single DICOM instance
							fmt.Println("\n--- Retrieve Single DICOM Instance ---")
							dicomResp, err := client.WadoRsRetrieveInstance(studyUID, seriesUID, instanceUID)
							if err != nil {
								log.Printf("Failed to retrieve DICOM instance: %v", err)
							} else {
								defer dicomResp.Body.Close()

								// Save DICOM instance to tmp folder
								instanceDir := "./tmp/instance"
								if err := os.MkdirAll(instanceDir, 0755); err != nil {
									log.Printf("Failed to create directory: %v", err)
								} else {
									count, err := saveMultipartDicomFiles(dicomResp.Body, dicomResp.Header.Get("Content-Type"), instanceDir)
									if err != nil {
										log.Printf("Failed to save DICOM instance: %v", err)
									} else {
										fmt.Printf("Saved %d DICOM file(s) to: %s\n", count, instanceDir)
									}
								}
							}

							// Example: Retrieve entire series as DICOM files
							fmt.Println("\n--- Retrieve Entire Series ---")
							seriesResp, err := client.WadoRsRetrieveSeries(studyUID, seriesUID)
							if err != nil {
								log.Printf("Failed to retrieve series: %v", err)
							} else {
								defer seriesResp.Body.Close()

								// Save all DICOM files from series to tmp folder
								seriesDir := fmt.Sprintf("./tmp/series/%s", sanitizeUID(seriesUID))
								if err := os.MkdirAll(seriesDir, 0755); err != nil {
									log.Printf("Failed to create directory: %v", err)
								} else {
									count, err := saveMultipartDicomFiles(seriesResp.Body, seriesResp.Header.Get("Content-Type"), seriesDir)
									if err != nil {
										log.Printf("Failed to save series DICOM files: %v", err)
									} else {
										fmt.Printf("Saved %d DICOM file(s) to: %s\n", count, seriesDir)
									}
								}
							}

							// Example: Retrieve entire study as DICOM files
							fmt.Println("\n--- Retrieve Entire Study ---")
							studyResp, err := client.WadoRsRetrieveStudy(studyUID)
							if err != nil {
								log.Printf("Failed to retrieve study: %v", err)
							} else {
								defer studyResp.Body.Close()

								// Save all DICOM files from study to tmp folder
								studyDir := fmt.Sprintf("./tmp/study/%s", sanitizeUID(studyUID))
								if err := os.MkdirAll(studyDir, 0755); err != nil {
									log.Printf("Failed to create directory: %v", err)
								} else {
									count, err := saveMultipartDicomFiles(studyResp.Body, studyResp.Header.Get("Content-Type"), studyDir)
									if err != nil {
										log.Printf("Failed to save study DICOM files: %v", err)
									} else {
										fmt.Printf("Saved %d DICOM file(s) to: %s\n", count, studyDir)
									}
								}
							}
						}
					}
				}
			}
		}

		fmt.Println("\nDICOMweb examples completed.")
	}
}

// saveMultipartDicomFiles parses a multipart response and saves each DICOM part as a file
func saveMultipartDicomFiles(body io.Reader, contentType string, outputDir string) (int, error) {
	// Parse the Content-Type header to get the boundary
	mediaType, params, err := mime.ParseMediaType(contentType)
	if err != nil {
		return 0, fmt.Errorf("failed to parse content type: %w", err)
	}

	if !strings.HasPrefix(mediaType, "multipart/") {
		// Not multipart, save as single file
		data, err := io.ReadAll(body)
		if err != nil {
			return 0, fmt.Errorf("failed to read body: %w", err)
		}
		filePath := filepath.Join(outputDir, "instance_0.dcm")
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return 0, fmt.Errorf("failed to write file: %w", err)
		}
		return 1, nil
	}

	boundary := params["boundary"]
	if boundary == "" {
		return 0, fmt.Errorf("no boundary found in content type")
	}

	// Read entire body into buffer for multipart parsing
	bodyData, err := io.ReadAll(body)
	if err != nil {
		return 0, fmt.Errorf("failed to read body: %w", err)
	}

	reader := multipart.NewReader(bytes.NewReader(bodyData), boundary)
	count := 0

	for {
		part, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			// Try alternative parsing if standard multipart fails
			return parseManualMultipart(bodyData, boundary, outputDir)
		}

		data, err := io.ReadAll(part)
		if err != nil {
			return count, fmt.Errorf("failed to read part: %w", err)
		}

		filePath := filepath.Join(outputDir, fmt.Sprintf("instance_%d.dcm", count))
		if err := os.WriteFile(filePath, data, 0644); err != nil {
			return count, fmt.Errorf("failed to write file: %w", err)
		}

		count++
		part.Close()
	}

	return count, nil
}

// parseManualMultipart handles multipart parsing manually when standard parser fails
func parseManualMultipart(data []byte, boundary string, outputDir string) (int, error) {
	delimiter := []byte("--" + boundary)
	parts := bytes.Split(data, delimiter)
	count := 0

	for _, part := range parts {
		// Skip empty parts and closing boundary
		if len(part) == 0 || bytes.Equal(bytes.TrimSpace(part), []byte("--")) {
			continue
		}

		// Find the end of headers (double CRLF or double LF)
		headerEnd := bytes.Index(part, []byte("\r\n\r\n"))
		if headerEnd == -1 {
			headerEnd = bytes.Index(part, []byte("\n\n"))
			if headerEnd == -1 {
				continue
			}
			headerEnd += 2
		} else {
			headerEnd += 4
		}

		// Extract body (skip headers)
		body := part[headerEnd:]
		// Trim trailing CRLF
		body = bytes.TrimSuffix(body, []byte("\r\n"))
		body = bytes.TrimSuffix(body, []byte("\n"))

		if len(body) == 0 {
			continue
		}

		filePath := filepath.Join(outputDir, fmt.Sprintf("instance_%d.dcm", count))
		if err := os.WriteFile(filePath, body, 0644); err != nil {
			return count, fmt.Errorf("failed to write file: %w", err)
		}

		count++
	}

	return count, nil
}

func sanitizeUID(uid string) string {
	return strings.ReplaceAll(uid, "/", "_")
}

// Helper function to print JSON data
func printJSON(data interface{}) {
	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal JSON: %v", err)
		return
	}
	// Limit output to first 1000 characters
	output := string(jsonData)
	if len(output) > 1000 {
		output = output[:1000] + "\n... (truncated)"
	}
	fmt.Println(output)
}

// Helper function to extract Study Instance UID from QIDO response
func getStudyInstanceUID(study map[string]interface{}) string {
	// Study Instance UID is at tag 0020000D
	if val, ok := study["0020000D"]; ok {
		if tagVal, ok := val.(map[string]interface{}); ok {
			if value, ok := tagVal["Value"].([]interface{}); ok && len(value) > 0 {
				if str, ok := value[0].(string); ok {
					return str
				}
			}
		}
	}
	return ""
}

// Helper function to extract Series Instance UID from QIDO response
func getSeriesInstanceUID(series map[string]interface{}) string {
	// Series Instance UID is at tag 0020000E
	if val, ok := series["0020000E"]; ok {
		if tagVal, ok := val.(map[string]interface{}); ok {
			if value, ok := tagVal["Value"].([]interface{}); ok && len(value) > 0 {
				if str, ok := value[0].(string); ok {
					return str
				}
			}
		}
	}
	return ""
}

// Helper function to extract SOP Instance UID from QIDO response
func getSOPInstanceUID(instance map[string]interface{}) string {
	// SOP Instance UID is at tag 00080018
	if val, ok := instance["00080018"]; ok {
		if tagVal, ok := val.(map[string]interface{}); ok {
			if value, ok := tagVal["Value"].([]interface{}); ok && len(value) > 0 {
				if str, ok := value[0].(string); ok {
					return str
				}
			}
		}
	}
	return ""
}
