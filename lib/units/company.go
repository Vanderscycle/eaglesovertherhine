package units

import (
	"fmt"
	"reflect"

	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
)

// TODO: create person class w/ rank, name,  + history (battles/medals/promotions)
type CompanyTemplate struct {
	Soldats         int16
	Caporals        int16
	CaporalFourrier int16
	Drummers        int16
	Sergents        int16
	SergentMajor    int16
	SousLieutenants int16
	Lieutenants     int16
	Capitaine       int16
}

type company struct {
	name            string
	nickname        string
	soldats         []pops.MilitaryPop
	caporals        []pops.MilitaryPop
	caporalFourrier []pops.MilitaryPop
	drummers        []pops.MilitaryPop
	sergents        []pops.MilitaryPop
	sergentMajor    []pops.MilitaryPop
	sousLieutenants []pops.MilitaryPop
	lieutenants     []pops.MilitaryPop
	capitaine       []pops.MilitaryPop
	template        CompanyTemplate
	nation          pops.Nation
	history string
}

func NewCompany(name string, template CompanyTemplate, nation pops.Nation) company {
	c := company{name: name, nickname: "unearned", template: template, history: "TODO"}

	tmplVal := reflect.ValueOf(c.template)
	tmplType := reflect.TypeOf(c.template)

	for i := 0; i < tmplVal.NumField(); i++ {
		fieldName := tmplType.Field(i).Name
		count := int(tmplVal.Field(i).Int())
		//fmt.Printf("%s\n", fieldName)
		for j := 0; j <= count; j++ {
			//fmt.Printf("%d\n", j)
			switch fieldName {
			case "soldats":
				c.soldats = append(c.soldats, pops.NewSoldier(pops.Soldat, nation, pops.Soldier))
			case "caporals":
				c.caporals = append(c.caporals, pops.NewSoldier(pops.Caporal,nation,pops.Soldier))
			case "caporalFourrier":
				c.caporalFourrier = append(c.caporalFourrier, pops.NewSoldier(pops.CaporalFourrier,nation,pops.Support))
			case "drummers":
				c.drummers = append(c.drummers, pops.NewSoldier(pops.Soldat,nation,pops.Support))
			case "sergents":
				c.sergents = append(c.sergents, pops.NewSoldier(pops.Sergent,nation, pops.Nco))
			case "sergentMajor":
				c.sergentMajor = append(c.sergentMajor, pops.NewSoldier(pops.SergentMajor,nation, pops.Nco))
			case "sousLieutenant":
				c.sousLieutenants = append(c.sousLieutenants, pops.NewSoldier(pops.SousLieutenant,nation, pops.Officier))
			case "lieutenant":
				c.lieutenants = append(c.lieutenants, pops.NewSoldier(pops.Lieutenant,nation, pops.Officier))
			case "captain":
				c.capitaine = append(c.capitaine, pops.NewSoldier(pops.Capitaine,nation, pops.Officier))
			}
		}
	}
	return c
}

func (c company) Status() {
	fmt.Printf("The company's name: %s and country %s\nTemplate: %+v\n", c.name, c.nation, c.template)
}

func (c company) ListSoldiers() {
	for _, s := range c.soldats {
		s.Status()
	}
	for _, s := range c.caporals {
		s.Status()
	}
	for _, s := range c.sergents {
		s.Status()
	}
}
