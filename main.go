package main

import (
	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
	"github.com/vanderscycle/eaglesovertherhine/lib/units"
)

func main() {
	firstCompany := units.NewCompany("1st Company")
	firstCompany.Status()

	firstSoldat := pops.NewSoldier()
	firstSoldat.Status()
	firstCompany.ListPrivates()
}
