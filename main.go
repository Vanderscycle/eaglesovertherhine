package main

import (
	"github.com/vanderscycle/eaglesovertherhine/lib/nations"
	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
	"github.com/vanderscycle/eaglesovertherhine/lib/units"
)

func main() {
	firstCompany := units.NewCompany("1st Company", nations.FrenchCompanyDefault, pops.France)
	firstCompany.Status()

	// firstSoldat := pops.NewSoldier()
	// firstSoldat.Status()
	firstCompany.ListSoldiers()
}
