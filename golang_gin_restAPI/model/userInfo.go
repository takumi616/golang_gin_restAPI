package model

import (
	"log"
    "github.com/takumi616/golang_gin_restAPI/db"
)

type UserInfo struct {
	Id      int      `json:"id"`	
	Name    string   `json:"name"`
	Age     int      `json:age`
	Country string   `json:"country"`
	Work    string   `json:"work"`
}

func GetAllUserInfo() ([]UserInfo) {
	Db := db.Init()
    rows, err := Db.Query("SELECT * FROM userInfo")
	
	if err != nil {
		log.Fatalf("failed to get records: %v", err)
	}

	var userInfoList []UserInfo 
	for rows.Next() {
		userInfo := UserInfo{}
		err = rows.Scan(&userInfo.Id, &userInfo.Name, &userInfo.Age, &userInfo.Country, &userInfo.Work)

		if err != nil {
			log.Fatalf("failed to copy a record into UserInfo struct: %v", err)
		}

		userInfoList = append(userInfoList, userInfo)
	}

	return userInfoList
}

func GetSingleUserInfo(id int) UserInfo {
	Db := db.Init()
	row := Db.QueryRow("SELECT * FROM userInfo WHERE id = $1", id)
	userInfo := UserInfo{}
	err := row.Scan(&userInfo.Id, &userInfo.Name, &userInfo.Age, &userInfo.Country, &userInfo.Work)

	if err != nil {
		log.Fatalf("failed to copy a record into UserInfo struct: %v", err)
	}

	return userInfo
}

func CreateUserInfo(userInfo UserInfo) int {
	Db := db.Init()
	var createdRecordId int
	err := Db.QueryRow("INSERT INTO userInfo(name, age, country, work) VALUES($1, $2, $3, $4) RETURNING id", userInfo.Name, userInfo.Age, userInfo.Country, userInfo.Work).Scan(&createdRecordId)

	if err != nil {
		log.Fatalf("failed to copy returnedId into variable: %v", err)
	}

	return createdRecordId
}

func UpdateUserInfo(id int, userInfo UserInfo) int {
	Db := db.Init()
	var updatedRecordId int
	err := Db.QueryRow("UPDATE userInfo SET name = $2, age = $3, country = $4, work = $5 WHERE id = $1 RETURNING id", id, userInfo.Name, userInfo.Age, userInfo.Country, userInfo.Work).Scan(&updatedRecordId)

	if err != nil {
		log.Fatalf("failed to copy returnedId into variable: %v", err)
	}

	return updatedRecordId
}

func DeleteUserInfo(id int) int {
	Db := db.Init()
	var deletedRecordId int
	err := Db.QueryRow("DELETE FROM userInfo WHERE id = $1 RETURNING id", id).Scan(&deletedRecordId)

	if err != nil {
		log.Fatalf("failed to copy returnedId into variable: %v", err)
	}

	return deletedRecordId
}
