package main

import "fmt"

func main() {
	// Uruguay is a country whose name is a reference to its location
	// Similarly, a pointer is a variable that references a memory location
	
	// Direct value (like saying "Uruguay" informally)
	countryName := "República Oriental del Uruguay"
	
	// Creating a pointer (a reference to where the actual name is stored)
	countryNamePtr := &countryName
	
	fmt.Println("The actual name:", countryName)
	fmt.Println("Memory address where name is stored:", countryNamePtr)
	fmt.Println("Accessing name through pointer:", *countryNamePtr)
	
	// Modifying value through pointer
	*countryNamePtr = "Eastern Republic of Uruguay"
	fmt.Println("Modified name:", countryName)
	
	// Function demonstrating pointers
	displayAndChange(&countryName)
	fmt.Println("Name after function call:", countryName)
	
	// Real-world example with a struct
	country := Country{
		OfficialName: "República Oriental del Uruguay",
		CommonName:   "Uruguay",
		Location:     "East of Uruguay River",
	}
	
	updateLocation(&country)
	fmt.Println("Updated country info:", country)
}

func displayAndChange(namePtr *string) {
	fmt.Println("In function - received pointer to:", *namePtr)
	*namePtr = "Uruguay (modified by function)"
}

type Country struct {
	OfficialName string
	CommonName   string
	Location     string
}

func updateLocation(c *Country) {
	c.Location = "Eastern Bank of Uruguay River"
}
