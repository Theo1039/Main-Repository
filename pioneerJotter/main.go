package main

import (
	"bufio"
	"fmt"
	"os"
	"pioneerJotter/bibleStudy"
	"pioneerJotter/hour"
	"pioneerJotter/returnVisit"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// 1. Lists to store multiple entries while the app is running
	var bsList []bibleStudy.BibleStudy
	var rvList []returnVisit.ReturnVisit
	var hourList []hour.HourRecorded

	for {
		fmt.Println("\n--- PIONEER JOTTER MENU ---")
		fmt.Println("1. Add Return Visit")
		fmt.Println("2. Add Bible Study")
		fmt.Println("3. Record Hours")
		fmt.Println("4. Filter by Street")
		fmt.Println("5. Show All Street Names")
		fmt.Println("6. Exit")
		fmt.Print("Select option: ")

		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			var rv returnVisit.ReturnVisit
			rv.FillFromUser()
			rvList = append(rvList, rv)

		case "2":
			var bs bibleStudy.BibleStudy
			bs.FilledFromUser()
			bsList = append(bsList, bs)

		case "3":
			var h hour.HourRecorded
			h.RecordHour()
			hourList = append(hourList, h)

		case "4":
			// FILTER LOGIC
			fmt.Print("Enter street name to filter: ")
			search, _ := reader.ReadString('\n')
			search = strings.TrimSpace(strings.ToLower(search))

			fmt.Printf("\n--- Results for: %s ---\n", search)
			// Filter Bible Studies
			for _, b := range bsList {
				if strings.Contains(strings.ToLower(b.StreetName), search) {
					bibleStudy.PrintFilledField(b)
				}
			}
			// Filter Return Visits
			for _, r := range rvList {
				if strings.Contains(strings.ToLower(r.StreetName), search) {
					r.PrintFilledFields()
				}
			}

		case "5":
			// LIST UNIQUE STREETS
			fmt.Println("\n--- List of Inputted Streets ---")
			streets := make(map[string]bool) // Use a map to avoid duplicates
			for _, b := range bsList {
				streets[b.StreetName] = true
			}
			for _, r := range rvList {
				streets[r.StreetName] = true
			}
			for s := range streets {
				if s != "" {
					fmt.Println("-", s)
				}
			}

		case "6":
			fmt.Println("Goodbye!")
			return
		}
	}
}
