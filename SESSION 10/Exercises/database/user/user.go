package user

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Doc_id      int    `json: "DocID"`
	Doc_license string `json: "DocLicense"`
	Doc_name    string `json: "DocName"`
}

func check_error(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func connection() sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/hospital_management")
	check_error(err)
	defer db.Close()
	return *db
}

func SelectAll() []User {
	db := connection()

	result, err := db.Query("SELECT * FROM doctor")
	check_error(err)
	var docArray []User
	var usr = User{}

	for result.Next() {
		err := result.Scan(&usr.Doc_id, &usr.Doc_license, &usr.Doc_name)
		check_error(err)
		docArray = append(docArray, User{Doc_id: usr.Doc_id, Doc_license: usr.Doc_license, Doc_name: usr.Doc_name})
	}
	fmt.Print(len(docArray))
	for i := 0; i < len(docArray); i++ {
		fmt.Println(docArray[i])
	}
	return docArray
}

func main() {
	SelectAll()
}
