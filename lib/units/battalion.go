package units

import (
	"fmt"
	"reflect"

	"github.com/vanderscycle/eaglesovertherhine/lib/pops"
)

type battalion struct {
	name                 string
	nickname             string
	adjudantSousOfficier []pops.MilitaryPop
	adjudantMajor        []pops.MilitaryPop
	chefDeBattalion      []pops.MilitaryPop
	companies            []company
	companyTemplate CompanyTemplate
	template             BattalionCompanyTemplate
	nation               pops.Nation
	history              string
}

func NewBattalion(name string, template BattalionCompanyTemplate, companyTemplate CompanyTemplate, nation pops.Nation)battalion {
	b := battalion{name: name, nickname: "unearned", template: template, history: "TODO"}


	// battalion staff generation
	b.adjudantSousOfficier = append(b.adjudantSousOfficier, pops.NewSoldier(pops.AdjudantSousOfficier,nation, pops.Nco))
	b.adjudantMajor = append(b.adjudantMajor, pops.NewSoldier(pops.AdjudantMajor,nation, pops.Nco))
	b.chefDeBattalion = append(b.chefDeBattalion, pops.NewSoldier(pops.ChefDeBattalion,nation, pops.Officier))

	// company generation
	tmplVal := reflect.ValueOf(b.template)
	tmplType := reflect.TypeOf(b.template)

	for i := 0; i < tmplVal.NumField(); i++ {
		fieldName := tmplType.Field(i).Name
		count := int(tmplVal.Field(i).Int())
		//fmt.Printf("%s\n", fieldName)
		for j := 0; j < count; j++ {
			compInt := j+1
			fmt.Printf("\n%d %s\n", j, fieldName)
			switch fieldName {
			case "Chasseurs":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Chasseurs", compInt), companyTemplate, nation))
			case "Voltigeurs":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Voltigeurs", compInt), companyTemplate, nation))
			case "Carabiniers":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Carabiniers", compInt), companyTemplate, nation))
			case "Fusiliers":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Fusiliers", compInt), companyTemplate, nation))
			case "Grenadiers":
				b.companies = append(b.companies, NewCompany(fmt.Sprintf("%d Chasseur", compInt), companyTemplate, nation))
			}
		}
	}
	return b
}

func (b battalion) Status() {
	fmt.Printf("The battalion's name: %s and country %s\nTemplate: %+v\n", b.name, b.nation, b.template)
}

func (b battalion) ListCompanies() {
	for _, s := range b.companies {
		fmt.Printf("%+v\n", s)
	}
	}
