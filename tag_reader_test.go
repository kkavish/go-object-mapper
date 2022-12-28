package go_property_mapper

import (
	"fmt"
	"testing"
)

func TestReadTags(t *testing.T) {
	m := GetTagsReflect()
	println(fmt.Sprintf("%v", m))
}

func TestReadTagsAndVal1(t *testing.T) {
	u := User{
		UserId:        GetStrPtr("kk"),
		FirstName:     GetStrPtr("kumar kavish"),
		LastName:      nil,
		Address:       nil,
		Phone:         GetInt64Ptr(9011911100),
		Country:       GetStrPtr("IN"),
		Pin:           nil,
		Age:           nil,
		Nominee:       nil,
		Qualification: nil,
		BirthDate:     nil,
		Salary:        nil,
		Father:        nil,
		Mother:        nil,
		SpouseName:    nil,
		NumChildren:   nil,
		VehiclesOwned: nil,
		Company:       nil,
		Designation:   nil,
		Department:    nil,
		IsVeteran:     nil,
	}
	m := GetColsAndVal(u)
	for k, v := range m {
		println(fmt.Sprintf("'%s' -> '%v'", k, v))
	}
}

func TestReadTagsAndVal2(t *testing.T) {
	u := User{
		UserId:        GetStrPtr("kk"),
		FirstName:     GetStrPtr("kumar kavish"),
		LastName:      nil,
		Address:       nil,
		Phone:         GetInt64Ptr(9011911100),
		Country:       GetStrPtr("IN"),
		Pin:           nil,
		Age:           nil,
		Nominee:       nil,
		Qualification: nil,
		BirthDate:     nil,
		Salary:        GetFloat64Ptr(6400000.768),
		Father:        nil,
		Mother:        nil,
		SpouseName:    nil,
		NumChildren:   nil,
		VehiclesOwned: nil,
		Company:       nil,
		Designation:   nil,
		Department:    nil,
		IsVeteran:     GetBoolPtr(true),
	}
	m := GetColsAndVal(u)
	for k, v := range m {
		println(fmt.Sprintf("'%s' -> '%v'", k, v))
	}
}

func TestReadTagsAndVal3(t *testing.T) {
	type Tiger struct {
		Type      *string  `col:"type"`
		IsExtinct *bool    `col:"is-extinct"`
		TagId     *int64   `col:"tag-id"`
		Length    *float64 `col:"length"`
	}
	u := Tiger{
		Type:      GetStrPtr("royal bengal"),
		IsExtinct: nil,
		TagId:     GetInt64Ptr(10001),
		Length:    nil,
	}
	m := GetColsAndVal(u)
	for k, v := range m {
		println(fmt.Sprintf("'%s' -> '%v'", k, v))
	}
}
