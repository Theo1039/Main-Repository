package hour

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type HourRecorded struct {
	Year      string
	Day       string
	Date      string
	Hour      int
	TotalHour int
}

func (h *HourRecorded) RecordHour() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Year: ")
	h.Year, _ = reader.ReadString('\n')
	h.Year = strings.TrimSpace(h.Year)

	fmt.Print("Enter Day: ")
	h.Day, _ = reader.ReadString('\n')
	h.Day = strings.TrimSpace(h.Day)

	fmt.Print("Enter Date: ")
	h.Date, _ = reader.ReadString('\n')
	h.Date = strings.TrimSpace(h.Date)

	fmt.Print("Enter Hour: ")
	hourInput, _ := reader.ReadString('\n')
	hourInput = strings.TrimSpace(hourInput)

	hour, err := strconv.Atoi(hourInput)
	if err != nil {
		fmt.Println("Invalid hour. Setting hour to 0")
		hour = 0
	}

	h.Hour = hour
	h.TotalHour += hour
}

func PrintFilledHour(h *HourRecorded) {
	if h.Year != "" {
		fmt.Println("Year:", h.Year)
	}
	if h.Day != "" {
		fmt.Println("Day:", h.Day)
	}
	if h.Date != "" {
		fmt.Println("Date:", h.Date)
	}
	fmt.Println("Hour Entered:", h.Hour)
	fmt.Println("Total Hours:", h.TotalHour)
}
