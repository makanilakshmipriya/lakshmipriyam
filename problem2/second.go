package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Read input JSON file
	inputJSON, err := os.ReadFile("input.json")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}

	var input map[string]interface{}
	err = json.Unmarshal(inputJSON, &input)
	if err != nil {
		fmt.Println("Error parsing input JSON:", err)
		return
	}

	output := convert(input)
	// output to JSON format
	outputJSON, err := json.MarshalIndent([]interface{}{output}, "", "  ")
	if err != nil {
		fmt.Println("Error marshalling output JSON:", err)
		return
	}
	fmt.Println(string(outputJSON))
}
func convert(input map[string]interface{}) map[string]interface{} {
	var output = make(map[string]interface{}, 0)
	for dataType, value := range input {
		// Remove leading and trailing whitespace from dataTypes
		dataType = strings.TrimSpace(dataType)
		if dataType != "" {
			switch valuePresented := value.(type) {
			case map[string]interface{}:
				if val, ok := valuePresented["N"].(string); ok {
					// If value is of Number data type
					if num, err := strconv.ParseFloat(strings.TrimSpace(val), 64); err == nil {
						output[dataType] = num
					}
				} else if val, ok := valuePresented["S"].(string); ok {
					// If value is of String data type
					value := strings.TrimSpace(val)
					if value != "" {
						if timestamp, err := time.Parse(time.RFC3339, value); err == nil {
							// convert RFC3339 formatted strings to Unix Epoch
							output[dataType] = timestamp.Unix()
						} else {
							output[dataType] = value
						}
					}
				} else if val, ok := valuePresented["BOOL"].(string); ok {
					// If value is of Boolean data type
					output[dataType] = handlEBool(val)
				} else if val, ok := valuePresented["NULL"].(string); ok {
					// If value is of Null data type
					nullVal := strings.TrimSpace(strings.ToLower(val))
					if nullVal == "1" || nullVal == "t" || nullVal == "true" {
						output[dataType] = nil
					}
				} else if val, ok := valuePresented["L"].([]interface{}); ok {
					// If value is of List data type
					if len(val) > 0 {
						convertedList := []interface{}{}
						for _, item := range val {
							itemMap, ok := item.(map[string]interface{})
							if !ok {
								continue
							}
							for itemType, itemValue := range itemMap {
								switch itemType {
								case "BOOL":
									val, _ := itemValue.(string)
									convertedList = append(convertedList, handlEBool(val))
								case "NULL":
									nullVal := strings.TrimSpace(strings.ToLower(fmt.Sprintf("%v", itemValue)))
									if nullVal == "1" || nullVal == "t" || nullVal == "true" {
										convertedList = append(convertedList, nil)
									}
								case "N":
									if num, err := strconv.ParseFloat(strings.TrimSpace(fmt.Sprintf("%v", itemValue)), 64); err == nil {
										convertedList = append(convertedList, num)
									}
								case "S":
									strVal := strings.TrimSpace(fmt.Sprintf("%v", itemValue))
									if strVal != "" {
										convertedList = append(convertedList, strVal)
									}

								}
							}
						}
						if len(convertedList) > 0 {
							output[dataType] = convertedList
						}
					}
				} else if val, ok := valuePresented["M"].(map[string]interface{}); ok {
					// If value is of Map data type
					convertedMap := convert(val)
					if len(convertedMap) > 0 {
						output[dataType] = convertedMap
					}
				}
			}
		}
	}
	return output
}

func handlEBool(val string) bool {
	boolValue := strings.TrimSpace(strings.ToLower(val))
	if boolValue == "1" || boolValue == "t" || boolValue == "true" {
		return true
	}
	return false
}
