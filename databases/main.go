package main

import (
	"database/sql"
	"fmt"

	_ "github.com/godror/godror"
)

// 1. Create a small script that will create a table called Users. This table must have three
// columns: ID, Name, and Email.
// 2. Add the details of two users with their data into the table. They should have unique
// names, IDs, and email addresses.
// 3. Then you need to update the email of the first user to user@packt.com and remove
// the second user. Make sure that none of the fields are NULL, and the ID is the
// primary key, so it needs to be unique.
// 4. When you are inserting, updating, and deleting from the table, please use the
// Prepare() function to protect against SQL injection attacks.
// 5. You should use a struct to store the user information you would like to insert, and
// when you are inserting, iterate over the struct with a for loop.
// 6. Once the insert, update, and delete calls are complete, make sure you use Close()
// when appropriate and finally close the connection to the database.

func main() {

	db, err := sql.Open("godror", `user="SYSTEM" password="sdhcbv27" connectString="127.0.0.1:1521/pdb1"`)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	// createTableQuery := `
	// 					CREATE TABLE USERS (
	// 						ID INTEGER NOT NULL,
	// 						NAME VARCHAR2(200 CHAR),
	// 						EMAIL VARCHAR2(200 CHAR)
	// 					)
	// 					`
	createUserRecordsQuery, err := db.Prepare(`
							INSERT INTO USERS (ID, NAME, EMAIL) VALUES (:1, :2, :3)
						`)

	updateUserEmailQuery, err := db.Prepare(`
						UPDATE USERS 
						SET EMAIL = :1
						WHERE ID = :2
						`)

	deleteUserQuery, err := db.Prepare(`
						DELETE FROM USERS
						WHERE ID = :1
						`)

	if err != nil {
		panic(err)
	}

	// _, err = db.Exec(createTableQuery)
	// if err != nil {
	// 	fmt.Println("Error creating table")
	// 	println(err)
	// 	return
	// }

	_, err = createUserRecordsQuery.Exec(1, "volek", "volek@gmail.com")
	if err != nil {
		fmt.Println("Error inserting data 1")
		println(err.Error())
		return
	}

	_, err = createUserRecordsQuery.Exec(2, "orets", "orets_motobyker_sosi_hui_2000@gmail.com")
	if err != nil {
		fmt.Println("Error inserting data 2")
		println(err.Error())
		return
	}

	_, err = updateUserEmailQuery.Exec("user@packt.com", 1)
	if err != nil {
		fmt.Println("Error updating data")
		println(err.Error())
		return
	}

	_, err = deleteUserQuery.Exec(2)
	if err != nil {
		fmt.Println("Error deleting data")
		println(err.Error())
		return
	}

}

// rows, err := db.Query("select sysdate from dual")
// if err != nil {
// 	fmt.Println("Error running query")
// 	fmt.Println(err)
// 	return
// }
// defer rows.Close()

// var thedate string
// for rows.Next() {

// 	rows.Scan(&thedate)
// }
// fmt.Printf("The date is: %s\n", thedate)
