package pkg

import (
	"fruites_crud/connect/db"
	"log"
)

//insert fruites here...

func InsertFruite(reqBody Fruites) bool {

	var result = true

	sqlStatement := `INSERT INTO fruite(fruite_name)VALUES ($1)`

	_, err2 := db.DB.Exec(sqlStatement, reqBody.FruitesName)

	if err2 != nil {
		//log.Fatal("ERror in insert: ", err2)
		return false
	}
	//fmt.Println(user)
	return result
}

//get fruites details ....

func GetFruite() []Fruites {
	Data := []Fruites{}

	//SQL := `SELECT "email", "first_name", "last_name", "date_of_birth", "sex" FROM "persons"`

	SQL := `SELECT id,fruite_name  from fruite`

	rows, err := db.DB.Query(SQL)

	if err != nil {
		log.Println("Failed to execute query: ", err)
		return Data
	}
	defer rows.Close()
	detailFruits := Fruites{}

	for rows.Next() {
		rows.Scan(&detailFruits.ID, &detailFruits.FruitesName)

		Data = append(Data, detailFruits)

	}
	return Data
}

//delete function ....
func DelFruites(fruitId int) int {
	a := 1
	//msg := ""
	sqlStatement := `DELETE FROM fruite where id = $1`

	_, err := db.DB.Exec(sqlStatement, fruitId)

	if err != nil {
		log.Println("ERror in deleting1: ", err)
		a = 0
		return a
	}
	return a
}

//update call function...

func FruitUpdate(reqbody Fruites) bool {

	SQL := `UPDATE fruite SET  fruite_name=$1 WHERE id=$2`

	_, err2 := db.DB.Exec(SQL, reqbody.FruitesName, reqbody.ID)

	if err2 != nil {
		//log.Fatal("ERror in update: ", err2)
		return false
	}

	return true
}
