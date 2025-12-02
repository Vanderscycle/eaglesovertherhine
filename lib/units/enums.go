package units

type CompanySpecialization int

const (
	Chasseur CompanySpecialization = iota
	Voltigeurs
	Carabiniers
	Fusilers
	Grenadiers
)

var AssignedCompanySpecialization = map[CompanySpecialization]string{
	Chasseur:    "Chasseur",
	Voltigeurs:  "Voltigeurs",
	Carabiniers: "Carabiniers",
	Fusilers:    "Fusilers",
	Grenadiers:  "Grenadiers",
}

func (s CompanySpecialization) String() string {
	return AssignedCompanySpecialization[s]
}

type CompanyTemplate struct {
	Soldats         int16
	Caporals        int16
	CaporalFourrier int16
	Drummers        int16
	Sergents        int16
	SergentMajor    int16
	SousLieutenants int16
	Lieutenants     int16
	Capitaines      int16
}

type BattalionCompanyTemplate struct {
	Chasseurs   int16
	Voltigeurs  int16
	Carabiniers int16
	Fusiliers   int16
	Grenadiers  int16
}
