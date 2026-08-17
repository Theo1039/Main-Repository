package bibleStudy

import (
    "fieldConnector/stdInput"
    "database/sql"
    "log"
    "bufio"
    "fmt"
    "os"
)

type BibleStudy struct{
    Id int
    Name string
    PhoneNumber string
    HouseNumber string
    StreetName string
    Book_brochure_name string
    Chapter string
    Paragraph string
    Next_appointment string
    QuestionAsked string
}

func CreateBibleStudy() BibleStudy{
    reader := bufio.NewReader(os.Stdin)
    b := BibleStudy{
        Name: stdInput.GetInput(reader, "Enter Name: "),
        PhoneNumber: stdInput.GetInput(reader, "Enter Phone number: "),
        HouseNumber: stdInput.GetInput(reader, "Enter House number: "),
        StreetName: stdInput.GetInput(reader, "Enter Street Name: "),
        Book_brochure_name: stdInput.GetInput(reader, "Enter Book Or Brochure Name: "),
        Chapter: stdInput.GetInput(reader, "Enter Chapter: "),
        Paragraph: stdInput.GetInput(reader, "Enter Paragraph: "),
        Next_appointment: stdInput.GetInput(reader, "Enter Next Appointment Date: "),
        QuestionAsked: stdInput.GetInput(reader, "Enter Question Asked: "),
    }
	return b
}
func SaveBibleStudy(db *sql.DB, b BibleStudy){
    query := `CREATE TABLE IF NOT EXISTS bible_study(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT,
    phoneNumber TEXT, 
    houseNumber TEXT,
    streetName TEXT,
    book_brochure_name TEXT,
    Chapter TEXT,
    paragraph TEXT,
    next_appointment TEXT,
    questionAsked TEXT
    )`

    _, err := db.Exec(query)
    if err != nil{
        log.Fatalf("Creating Database Table Failed %v", err)
    }
    query = `INSERT INTO bible_study(name,phoneNumber,
    houseNumber,streetName,book_brochure_name,chapter,paragraph,
    next_appointment,questionAsked) VALUES(?,?,?,?,?,?,?,?,?)`

    _,err = db.Exec(query,b.Name,b.PhoneNumber,b.HouseNumber,b.StreetName,b.Book_brochure_name,b.Chapter,b.Paragraph,b.Next_appointment,b.QuestionAsked)
    if err != nil {
       log.Fatalf("Database Insertion Failed %v", err)
    }
    fmt.Println("Blble Study saved successfully!")
}

func DisplayBibleStudy(db *sql.DB){

    query := `SELECT
    id,
    COALESCE(name, ""),
    COALESCE(phoneNumber, ""),
    COALESCE(houseNumber,""),
    COALESCE(streetName,""),
    COALESCE(book_brochure_name,""),
    COALESCE(chapter,""),
    COALESCE(paragraph,""),
    COALESCE(next_appointment,""),
    COALESCE(questionAsked,"")
    FROM bible_study WHERE 1=1`

    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()
    printRows(rows)

    }
    


func printRows(rows *sql.Rows) {
	count := 0

	for rows.Next() {
        var b BibleStudy
		
        err := rows.Scan(&b.Id, &b.Name, &b.PhoneNumber, &b.HouseNumber, &b.StreetName, &b.Book_brochure_name, &b.Chapter, &b.Paragraph, &b.Next_appointment, &b.QuestionAsked)
		
		if err != nil {
			fmt.Println("Error scanning row data:", err)
			return
		}
		count++
        fmt.Printf("ID: %d | Name: %s | Phone No: %s | House No: %s | Street No: %s | Book Or Brochure: %s | Chapter: %s | Paragraph: %s | Next Appointment: %s | Question: %s\n",
        b.Id,b.Name,b.PhoneNumber,b.HouseNumber,b.StreetName,b.Book_brochure_name,b.Chapter,b.Paragraph,b.Next_appointment,b.QuestionAsked)
		
	}

	if count == 0 {
		fmt.Println("No matching return visits found.")
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during row processing:", err)
	}
}


func UpdateBibleStudy(db *sql.DB){
    var id int
	fmt.Print("Enter ID to Update: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid ID format")
		return
	}

	var oldName, oldPhoneNo, oldHouseNo, oldStreetName, oldBook_brochure_name string

	err = db.QueryRow("SELECT name, phoneNumber, houseNumber, streetName, book_brochure_name FROM bible_study WHERE id = ?", id).
		Scan(&oldName, &oldPhoneNo, &oldHouseNo, &oldStreetName, &oldBook_brochure_name)
	if err != nil {
		fmt.Println("Bible Study  Not Found")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	
	newName := stdInput.GetInput(reader, "New Name (Leave Empty to Keep Current): ")
	newPhoneNo := stdInput.GetInput(reader, "New Phone Number (Leave Empty to Keep Current): ")
	newHouseNo := stdInput.GetInput(reader, "New House Number (Leave Empty to Keep Current): ")
	newStreetName := stdInput.GetInput(reader, "New Street Name (Leave Empty to Keep Current): ")
	newBook_brochure_name := stdInput.GetInput(reader, "New Land Mark (Leave Empty to Keep Current): ")

	if newName == "" {
		newName = oldName
	}
	if newPhoneNo == "" {
		newPhoneNo = oldPhoneNo
	}
	if newHouseNo == "" {
		newHouseNo = oldHouseNo
	}
	if newStreetName == "" {
		newStreetName = oldStreetName
	}
	if newBook_brochure_name == "" {
		newBook_brochure_name = oldBook_brochure_name
	}

	_, err = db.Exec(
		"UPDATE bible_study SET name = ?, phoneNumber = ?, houseNumber = ?, streetName = ?, book_brochure_name = ? WHERE id = ?",
		newName,
		newPhoneNo,
		newHouseNo,
		newStreetName,
		newBook_brochure_name,
		id,
	)
	if err != nil {
		fmt.Println("Update failed:", err)
		return
	}
	fmt.Println("Bible Study Updated Successfully!")
    
}


func DeleteBibleStudy(db *sql.DB, id string) {
	_, err := db.Exec("DELETE FROM bible_study WHERE id = ?", id)
	if err != nil {
		fmt.Println("Delete failed:", err)
		return
	}
	fmt.Println("Deleted successfully")
}


func FilterBibleStudyByFields(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	nameFilter := stdInput.GetInput(reader, "Filter by Name (Leave empty to skip): ")
	streetFilter := stdInput.GetInput(reader, "Filter by Street Name (Leave empty to skip): ")

	query := `SELECT 
		id, 
		COALESCE(name, ''), 
		COALESCE(phoneNumber, ''), 
		COALESCE(houseNumber, ''),  
		COALESCE(streetName, ''), 
        COALESCE(book_brochure_name, ""),
        COALESCE(chapter,""),
        COALESCE(paragraph,""),
        COALESCE(next_appointment, ""),
		COALESCE(questionAsked, '')
	FROM bible_study WHERE 1=1`
   

	var args []interface{}

	if nameFilter != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+nameFilter+"%")
	}
	if streetFilter != "" {
		query += " AND streetName LIKE ?"
		args = append(args, "%"+streetFilter+"%")
	}

	rows, err := db.Query(query, args...)
	if err != nil {
		fmt.Println("Filter query failed:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n--- FILTERED RESULTS ---")
	printRows(rows)
    if err := rows.Err(); err != nil {
        fmt.Println("Rows Error:", err)
    }
}