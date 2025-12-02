package records

type Record int

const (
	Battle Record = iota
)

// TODO: add timestamp + information for each
var AssignedRecord = map[Record]string{
	Battle: "Battle",
}

func (n Record) String() string {
	return AssignedRecord[n]
}
