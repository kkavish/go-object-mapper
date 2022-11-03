package go_property_mapper

import "time"

type ADAO struct {
	Id          int
	Name        string
	Address     string
	Phone       int64
	Salary      float64
	DBId        int64
	CreatedTime time.Time
}
