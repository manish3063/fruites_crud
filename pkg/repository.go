package pkg

import "fruites_crud/connect/db"

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
