package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

var fname []string
var lname []string
var email []string
var phone []string

func main() {
	f, err := excelize.OpenFile("sample.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Get all the rows in the sheet.
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, row := range rows[1:] {
		// append column values to slice
		fname = append(fname, row[1])
		lname = append(lname, row[2])
		email = append(email, row[3])
		phone = append(phone, row[4])
	}

	fmt.Printf("First Name:\n%v \n\n", fname)
	fmt.Printf("Last Name:\n%v \n\n", lname)
	fmt.Printf("Email:\n%v \n\n", email)
	fmt.Printf("Phone:\n%v \n\n", phone)
}
