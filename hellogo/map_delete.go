package main

import "fmt"

func main() {
	countryCaptitalMap := map[string]string{
		"France": "Paris",
		"Italy":  "Rome",
		"Japan":  "Tokyo",
		"India":  "New Delhi",
	}

	fmt.Println("原始Map")
	for country := range countryCaptitalMap {
		fmt.Println("Capital of", country, "is", countryCaptitalMap[country])
	}

	delete(countryCaptitalMap, "France")
	fmt.Println("Entry for France is deleted")

	fmt.Println("删除后Map")
	for country := range countryCaptitalMap {
		fmt.Println("Capital of", country, "is", countryCaptitalMap[country])
	}

}
