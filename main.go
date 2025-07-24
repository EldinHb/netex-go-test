package main

import (
	"encoding/xml"
	"fmt"
	"io"
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

	for _, choice := range publicationDelivery.DataObjects.Choice {
		for _, frame := range choice.CompositeFrame.Frames.CommonFrame {
			if frame.ServiceFrame != nil {
				fmt.Printf("Destination %+v\n", frame.ServiceFrame.DestinationDisplays.DestinationDisplay[0].Name.Value)
			}
		}
	}

	fmt.Println("XML unmarshalling completed successfully!")
}
