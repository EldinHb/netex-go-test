package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"

	"gotest/netex"
)

func main() {
	fmt.Println("Loading and unmarshalling NeTEx_HTM_test_20210301.xml...")

	// Open the XML file
	file, err := os.Open("./NeTEx_HTM_test_20210301.xml")
	// file, err := os.Open("./netexret.xml")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()

	xmlData, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	var publicationDelivery netex.PublicationDelivery

	err = xml.Unmarshal(xmlData, &publicationDelivery)
	if err != nil {
		fmt.Printf("Error unmarshalling XML: %v\n", err)
		return
	}

	// Print the struct to stdout
	fmt.Printf("Successfully unmarshalled XML into PublicationDelivery struct:\n")
	fmt.Printf("Version: %s\n", *publicationDelivery.Version)
	fmt.Printf("Publication Timestamp: %s\n", publicationDelivery.PublicationTimestamp)
	fmt.Printf("Participant Ref: %s\n", publicationDelivery.ParticipantRef.Value)

	if publicationDelivery.Description != nil {
		fmt.Printf("Description: %+v\n", publicationDelivery.Description)
	}

	if publicationDelivery.DataObjects != nil {
		fmt.Printf("Data Objects present: %+v\n", publicationDelivery.DataObjects)
	}

	marshaled, err := xml.MarshalIndent(publicationDelivery, "", "  ")
	if err != nil {
		log.Fatalf("Error marshalling back to XML: %v", err)
	}

	// save in file
	err = os.WriteFile("output.xml", marshaled, 0644)
	if err != nil {
		log.Fatalf("Error writing to file: %v", err)
	}

	// Print the marshalled XML to stdout

	fmt.Println("XML unmarshalling completed successfully!")
}
