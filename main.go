package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/vanderscycle/eaglesovertherhine/engine"
	"github.com/vanderscycle/eaglesovertherhine/lib/nations"
	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
	"github.com/vanderscycle/eaglesovertherhine/lib/units"
)

func main() {

	// Batalion
	// firstBattation := units.NewBattalion("1st Battalion", nations.FrenchLightInfantryBattalion, nations.FrenchCompanyDefault, pops.France)
	// firstBattation.EarnBloodName("Unbowed")
	// firstBattation.Status()
	// comp := firstBattation.GetCompany("1 Voltigeurs")
	// comp.EarnBloodName("Unbroken")
	// fmt.Printf("%+v/n", comp)
	//firstBattation.ListCompanies()

	firstCompany := units.NewCompany("1st Company", nations.FrenchCompanyDefault, pops.France)
	firstCompany.Status()
	firstCompany.ListSoldiers()

	firstSoldat := pops.NewSoldier(pops.AdjudantMajor, pops.France, pops.Nco)
	firstSoldat.Status()

	firstCiv := pops.NewCivilian(pops.France, pops.Farmer)
	firstCiv.Status()

	g := engine.NewGame()

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
