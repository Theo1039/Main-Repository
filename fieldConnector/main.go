package main

import (
	"database/sql"
	"fmt"
	"log" // Added for cleaner error handling exit
    "bufio"
	"os"
	_ "github.com/mattn/go-sqlite3"
	"fieldConnector/stdInput"
	"fieldConnector/returnVisit"
	"fieldConnector/bibleStudy"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\n WELCOME! TO FIELDCONNECTOR APP") 
	fmt.Println()
	fmt.Println("\tOPTIONS")
	fmt.Println()

	db, err := sql.Open("sqlite3", "mainDB.db")
	if err != nil {
		log.Fatalf("Database initialization failed: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

	for{
		fmt.Println("....MAIN OPTIONS...\n")
		fmt.Println("Option 1. Return Visit Section\nOption 2. Bible Study Section\nOption 3. Exit")
		option := stdInput.GetInput(reader, "Enter Main Option: ")
		switch option{
			case "1":
			returnVisitMenu(db)
			case "2":
			bibleStudyMenu(db)
			case "3":
			return
			default:
			fmt.Println("Invalid Input")
		}
	}
}
	
func returnVisitMenu(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n*=====__ RETURN VISITS SECTION __=====*")
		fmt.Println()
		fmt.Println("1. Add")
		fmt.Println("2. Display")
		fmt.Println("3. Search")
		fmt.Println("4. Update")
		fmt.Println("5. Delete")
		fmt.Println("6. Back")

		
		option := stdInput.GetInput(reader, "\nEnter Return Visit Option: ")

		switch option {

		case "1":
			inputData := returnVisit.UserInput()
			returnVisit.Save(db,inputData)

		case "2":
			returnVisit.Display(db)
		

		case "3":
			returnVisit.FilterByFields(db)

		case "4":
			returnVisit.Display(db)
			fmt.Println()
			returnVisit.UpdateUser(db)

		case "5":
			returnVisit.Display(db)
			fmt.Println()
		    id := stdInput.GetInput(reader, "Enter ID: ")
			returnVisit.DeleteUser(db, id)

		case "6":
			return

		default:
			fmt.Println("Invalid option")
		}
	}
}

func bibleStudyMenu(db *sql.DB){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\t*.......WELCOME! TO BIBLE STUDY SECTION.....*")
	fmt.Println()
	fmt.Println("\t*****____OPTIONS____*****")
	fmt.Println()
	fmt.Println("Choice 1. add Bible Study\nChoice 2. Display Bible Study\nchoice3. Update Bible Study\nchoice 4. Delete Bible Study\nChoice 5. Filter Bible Study By Name Or Street\nChoice 6. Exit")
	choice := stdInput.GetInput(reader, "\nEnter Bible Study Choice: ")
	fmt.Println()
	switch choice{
		case "1":
		bStudy := bibleStudy.CreateBibleStudy()
		bibleStudy.SaveBibleStudy(db, bStudy)
		case "2":
		bibleStudy.DisplayBibleStudy(db)
		case "3":
		bibleStudy.DisplayBibleStudy(db)
		fmt.Println()
		bibleStudy.UpdateBibleStudy(db)
		case "4":
		bibleStudy.DisplayBibleStudy(db)
	    id := stdInput.GetInput(reader, "\nEnter Delete ID: ")
		bibleStudy.DeleteBibleStudy(db,id)
		case "5":
		bibleStudy.FilterBibleStudyByFields(db)
        case "6":
		return
		default:
		fmt.Println("Invalid Input")
	}

}
