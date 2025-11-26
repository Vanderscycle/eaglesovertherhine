package main

import (
	"fmt"

	"github.com/vanderscycle/eaglesovertherhine/lib"
	"github.com/vanderscycle/eaglesovertherhine/lib/company"
)

func main() {
	firstCompany := company.New("1st Company ")
	firstCompany.Status

	result := lib.Hello("bob")
	fmt.Printf("%s",result)
}
