package returnVisit

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	"fieldConnector/stdInput"
)

type ReturnVisit struct {
	ID             int
	Name           string
	PhoneNumber    string
	NextDiscussion string
	StreetName     string
	HouseNumber    string
	LandMark       string
	Others         string
}

func UserInput()ReturnVisit {
	reader := bufio.NewReader(os.Stdin)
	name := stdInput.GetInput(reader, "Enter Return Visit Name: ")
	if name == "" {
		fmt.Println("\U0001F6D1")
		fmt.Println("Please! Enter Return Visit Name")
		return ReturnVisit{
            Name: name,
        }
	}
	user := ReturnVisit{
		Name:           name,
		PhoneNumber:    stdInput.GetInput(reader, "Enter phone number: "),
		NextDiscussion: stdInput.GetInput(reader, "Enter your next discussion: "),
		StreetName:     stdInput.GetInput(reader, "Enter street name: "),
		HouseNumber:    stdInput.GetInput(reader, "Enter house number: "),
		LandMark:       stdInput.GetInput(reader, "Enter land mark: "),
		Others:         stdInput.GetInput(reader, "Enter others: "),
	}
    return user
}
func Save(db *sql.DB, user ReturnVisit){
	createTableQuery := `CREATE TABLE IF NOT EXISTS return_visit (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		phoneNumber TEXT,
		nextDiscussion TEXT,
		streetName TEXT,
		houseNumber TEXT,
		landMark TEXT,
		others TEXT
	);`

	_, err := db.Exec(createTableQuery)
	if err != nil {
		fmt.Println("Failed to create table:", err)
		return
	}

	insertQuery := `INSERT INTO return_visit (
		name, phoneNumber, nextDiscussion, streetName, houseNumber, landMark, others
	) VALUES (?, ?, ?, ?, ?, ?, ?);`

	_, err = db.Exec(insertQuery,
		user.Name,
		user.PhoneNumber,
		user.NextDiscussion,
		user.StreetName,
		user.HouseNumber,
		user.LandMark,
		user.Others,
	)

	if err != nil {
		fmt.Println("Database insertion failed:", err)
		return
	}

	fmt.Println("Database insertion successful")
}

func DeleteUser(db *sql.DB, id string) {
	_, err := db.Exec("DELETE FROM return_visit WHERE id = ?", id)
	if err != nil {
		fmt.Println("Delete failed:", err)
		return
	}
	fmt.Println("Deleted successfully")
}

func Display(db *sql.DB) {
	query := `SELECT 
		id, 
		COALESCE(name, ''), 
		COALESCE(phoneNumber, ''), 
		COALESCE(nextDiscussion, ''), 
		COALESCE(streetName, ''), 
		COALESCE(houseNumber, ''), 
		COALESCE(landMark, ''), 
		COALESCE(others, '') 
	FROM return_visit`

	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("Error reading from the database:", err)
		return
	}
	defer rows.Close()

	fmt.Println("\n--- ALL RETURN VISITS ---")
	printRows(rows)
}


func FilterByFields(db *sql.DB) {
	reader := bufio.NewReader(os.Stdin)

	nameFilter := stdInput.GetInput(reader, "Filter by Name (Leave empty to skip): ")
	streetFilter := stdInput.GetInput(reader, "Filter by Street Name (Leave empty to skip): ")

	query := `SELECT 
		id, 
		COALESCE(name, ''), 
		COALESCE(phoneNumber, ''), 
		COALESCE(nextDiscussion, ''), 
		COALESCE(streetName, ''), 
		COALESCE(houseNumber, ''), 
		COALESCE(landMark, ''), 
		COALESCE(others, '') 
	FROM return_visit WHERE 1=1`

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
}


func printRows(rows *sql.Rows) {
	count := 0
	for rows.Next() {
		var u ReturnVisit
		err := rows.Scan(&u.ID, &u.Name, &u.PhoneNumber, &u.NextDiscussion, &u.StreetName, &u.HouseNumber, &u.LandMark, &u.Others)
		if err != nil {
			fmt.Println("Error scanning row data:", err)
			return
		}
		count++
		fmt.Printf("ID: %d | Name: %s | Phone: %s | Next Discussion: %s | Address: %s, %s (%s) | Others: %s\n",
			u.ID, u.Name, u.PhoneNumber, u.NextDiscussion, u.HouseNumber, u.StreetName, u.LandMark, u.Others)
	}

	if count == 0 {
		fmt.Println("No matching return visits found.")
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error during row processing:", err)
	}
}

func UpdateUser(db *sql.DB) {
	var id int
	fmt.Print("Enter ID to Update: ")
	_, err := fmt.Scanln(&id)
	if err != nil {
		fmt.Println("Invalid ID format")
		return
	}

	var oldName, oldPhoneNo, oldHouseNo, oldStreetName, oldLandMark string

	err = db.QueryRow("SELECT name, phoneNumber, houseNumber, streetName, landMark FROM return_visit WHERE id = ?", id).
		Scan(&oldName, &oldPhoneNo, &oldHouseNo, &oldStreetName, &oldLandMark)
	if err != nil {
		fmt.Println("Return Visit Not Found")
		return
	}

	reader := bufio.NewReader(os.Stdin)

	
	newName := stdInput.GetInput(reader, "New Name (Leave Empty to Keep Current): ")
	newPhoneNo := stdInput.GetInput(reader, "New Phone Number (Leave Empty to Keep Current): ")
	newHouseNo := stdInput.GetInput(reader, "New House Number (Leave Empty to Keep Current): ")
	newStreetName := stdInput.GetInput(reader, "New Street Name (Leave Empty to Keep Current): ")
	newLandMark := stdInput.GetInput(reader, "New Land Mark (Leave Empty to Keep Current): ")

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
	if newLandMark == "" {
		newLandMark = oldLandMark
	}

	_, err = db.Exec(
		"UPDATE return_visit SET name = ?, phoneNumber = ?, houseNumber = ?, streetName = ?, landMark = ? WHERE id = ?",
		newName,
		newPhoneNo,
		newHouseNo,
		newStreetName,
		newLandMark,
		id,
	)
	if err != nil {
		fmt.Println("Update failed:", err)
		return
	}
	fmt.Println("Return Visit Updated Successfully!")
}
	