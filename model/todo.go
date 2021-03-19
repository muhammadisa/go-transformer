package model

type Code struct {
	ID int64
}
type Todo struct {
	ID            string
	Name          string
	Completed     bool
	NumberCode    int32
	NumberProduct uint64
	Codes         []Code
}
