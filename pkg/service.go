package pkg

import (
	"encoding/json"
	"fmt"
	"fruites_crud/connect/db"

	"log"
	"time"
)

func FruitInsertService(reqbody Fruites) bool {

	res := InsertFruite(reqbody)
	//fmt.Println("aaaa", res)

	if res == true {
		data := GetFruite()
		//fmt.Println("mmm", data)
		jsonData, _ := json.Marshal(data)
		r := db.Client.Set("all", jsonData, time.Minute*20)

		fmt.Println("test", r)
	}

	return res
}

func GetFruitsService() (f []Fruites) {

	val, err := db.Client.Get("all").Bytes()
	json.Unmarshal(val, &f)
	if err != nil {
		log.Println("error while unmarshalling data: ", err)
	}
	if len(f) > 0 {
		return f
	}

	data := GetFruite()
	return data
}

func UpdateFruitsService(reqbody Fruites) bool {

	res := FruitUpdate(reqbody)

	return res
}

func DeleteFruitsService(id int) bool {

	res := DelFruites(id)

	return res
}
