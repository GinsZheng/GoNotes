package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  int
	ManagerID int
}

func structPrint() {
	var gins = Employee{ID: 1, Name: "Gins", Address: "深圳", Position: 3, ManagerID: 123}
	fmt.Println(gins)
}
