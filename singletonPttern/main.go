package main

import (
	"fmt"
	"sync"
)

type DbConnection struct{}

var dbConn *DbConnection

var lock = &sync.Mutex{}

func getInstance() *DbConnection {
	if dbConn != nil {
		fmt.Println("Single instance already created")
	} else {
		lock.Lock()
		defer lock.Unlock()
		if dbConn == nil {
			fmt.Println("Creating single instance")
			dbConn = &DbConnection{}
		} else {
			fmt.Println("Single instance already created")
		}

	}

	return dbConn
}
func main() {
	for i := 0; i < 10; i++ {
		go getInstance()
	}

	fmt.Scanln()

}
