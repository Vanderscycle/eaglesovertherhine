package main

import (
	"fmt"

	"github.com/vanderscycle/eaglesovertherhine/lib/nations"
	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
	"github.com/vanderscycle/eaglesovertherhine/lib/units"
)

func main() {
	// firstCompany := units.NewCompany("1st Company", nations.FrenchCompanyDefault, pops.France)
	// firstCompany.Status()
	// firstCompany.ListSoldiers()

	firstBattation := units.NewBattalion("1st Battalion", nations.FrenchLightInfantryBattalion, nations.FrenchCompanyDefault, pops.France)
	firstBattation.EarnBloodName("Unbowed")
	firstBattation.Status()
	comp := firstBattation.GetCompany("1 Voltigeurs")
	comp.EarnBloodName("Unbroken")
	fmt.Printf("%+v/n", comp)
	//firstBattation.ListCompanies()

	// firstSoldat := pops.NewSoldier()
	// firstSoldat.Status()
}
