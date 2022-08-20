package main

import "fmt"

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("init db failed,err:%v\n", err)
		return
	}
	// CRUD
	//insertRowDemo()
	//fmt.Println("========================")
	//queryRowDemo()
	//fmt.Println("========================")
	//queryMultiRowDemo()
	//fmt.Println("========================")
	//updateRowDemo()
	//fmt.Println("========================")
	//deleteRowDemo()
	//fmt.Println("========================")

	// prepare
	//prepareQuery()
	//prepareInsertQuery()

	// sql inject
	//sqlInject("xxx' or 1=1#")
	//sqlInject("xxx' union select * from user #")
	//sqlInject("xxx' and (select count(*) from user) < 10 #")

	// transaction
	transaction()
}
