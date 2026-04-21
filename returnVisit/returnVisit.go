package returnVisit

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Struct definition
type ReturnVisit struct {
	Name                string
	PhoneNo             string
	StreetName          string
	Description         string
	HouseNo             string
	QuestionLeftOrAsked string
	LandMark            string
}

// Function to capitalize first letter
func CapitalizeFirst(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}
	s = strings.ToUpper(s[:1]) + (s[1:])
	return s
}

func (rv *ReturnVisit) FillFromUser() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Name: ")
	rv.Name, _ = reader.ReadString('\n')
	rv.Name = CapitalizeFirst(rv.Name)

	fmt.Print("Enter Street Name: ") // Add this block
	rv.StreetName, _ = reader.ReadString('\n')
	rv.StreetName = CapitalizeFirst(rv.StreetName)
	fmt.Print("Enter Phone Number: ")
	rv.PhoneNo, _ = reader.ReadString('\n')
	rv.PhoneNo = strings.TrimSpace(rv.PhoneNo)

	fmt.Print("Enter Description: ")
	rv.Description, _ = reader.ReadString('\n')
	rv.Description = CapitalizeFirst(rv.Description)

	fmt.Print("Enter House Number: ")
	rv.HouseNo, _ = reader.ReadString('\n')
	rv.HouseNo = strings.TrimSpace(rv.HouseNo)

	fmt.Print("Enter Question Left or Asked: ")
	rv.QuestionLeftOrAsked, _ = reader.ReadString('\n')
	rv.QuestionLeftOrAsked = CapitalizeFirst(rv.QuestionLeftOrAsked)

	fmt.Print("Enter Landmark: ")
	rv.LandMark, _ = reader.ReadString('\n')
	rv.LandMark = CapitalizeFirst(rv.LandMark)
}

// Method to print only filled fields
func (rv *ReturnVisit) PrintFilledFields() {

	fmt.Println("\n--- Return Visit Details ---")

	if rv.Name != "" {
		fmt.Println("Name:", rv.Name)
	}
	if rv.PhoneNo != "" {
		fmt.Println("Phone Number:", rv.PhoneNo)
	}
	if rv.Description != "" {
		fmt.Println("Description:", rv.Description)
	}
	if rv.HouseNo != "" {
		fmt.Println("House Number:", rv.HouseNo)
	}
	if rv.QuestionLeftOrAsked != "" {
		fmt.Println("Question Left/Asked:", rv.QuestionLeftOrAsked)
	}
	if rv.LandMark != "" {
		fmt.Println("Landmark:", rv.LandMark)
	}
}
