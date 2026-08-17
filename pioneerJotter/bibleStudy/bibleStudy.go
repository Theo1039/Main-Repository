package bibleStudy

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type BibleStudy struct {
	Name               string
	PhoneNo            string
	StreetName         string
	HouseNo            string
	BookOrBrochureName string
	ChapterToStart     string
	Paragraph          string
	QuestionAsked      string
}

func CapitalizeTheFirstLetter(s string) string {
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
func (bs *BibleStudy) FilledFromUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("enter name: ")
	bs.Name, _ = reader.ReadString('\n')
	bs.Name = CapitalizeTheFirstLetter(bs.Name)

	fmt.Print("enter street name: ")
	bs.StreetName, _ = reader.ReadString('\n')
	bs.StreetName = CapitalizeTheFirstLetter(bs.StreetName)

	fmt.Print("enter phone number: ")
	bs.PhoneNo, _ = reader.ReadString('\n')
	bs.PhoneNo = strings.TrimSpace(bs.PhoneNo)

	fmt.Print("enter house number: ")
	bs.HouseNo, _ = reader.ReadString('\n')
	bs.HouseNo = strings.TrimSpace(bs.HouseNo)

	fmt.Println("enter book or brochure name: ")
	bs.BookOrBrochureName, _ = reader.ReadString('\n')
	bs.BookOrBrochureName = strings.TrimSpace(bs.BookOrBrochureName)

	fmt.Print("enter chapter to start from: ")
	bs.ChapterToStart, _ = reader.ReadString('\n')
	bs.ChapterToStart = strings.TrimSpace(bs.ChapterToStart)

	fmt.Println("enter paragraph number: ")
	bs.Paragraph, _ = reader.ReadString('\n')
	bs.Paragraph = strings.TrimSpace(bs.Paragraph)

	fmt.Print("enter question asked: ")
	bs.QuestionAsked, _ = reader.ReadString('\n')
	bs.QuestionAsked = strings.TrimSpace(bs.QuestionAsked)
}

func PrintFilledField(bs BibleStudy) {
	fmt.Println("====BIBLE STUDY DETAIL===")
	if bs.Name != "" {

	}
	fmt.Println("Name:", bs.Name)

	if bs.PhoneNo != "" {

	}
	fmt.Println("Phone Number: ", bs.PhoneNo)

	if bs.StreetName != "" {

	}
	fmt.Println("Street Name: ", bs.StreetName)

	if bs.HouseNo != "" {

	}
	fmt.Println("House Number: ", bs.HouseNo)

	if bs.BookOrBrochureName != "" {

	}
	fmt.Println("Book/Brochure name: ", bs.BookOrBrochureName)

	if bs.ChapterToStart != "" {

	}
	fmt.Println("Chapter to start from: ", bs.ChapterToStart)

	if bs.Paragraph != "" {

	}
	fmt.Println("Paragraph to start: ", bs.Paragraph)

	if bs.QuestionAsked != "" {

	}
	fmt.Println("Question asked: ", bs.QuestionAsked)

}
