package go_property_mapper

import "reflect"

type User struct {
	UserId        string  `col:"user-id"`
	FirstName     string  `col:"first-name"`
	LastName      string  `col:"last-name"`
	Address       string  `col:"address"`
	Phone         int64   `col:"phone"`
	Country       string  `col:"country"`
	Pin           int     `col:"pin"`
	Age           int     `col:"age"`
	Nominee       string  `col:"nominee"`
	Qualification string  `col:"qualification"`
	BirthDate     string  `col:"birth-date"`
	Salary        float64 `col:"salary"`
	Father        string  `col:"father"`
	Mother        string  `col:"mother"`
	SpouseName    string  `col:"spouse-name"`
	NumChildren   int     `col:"num-children"`
	VehiclesOwned int     `col:"vehicles-owned"`
	Company       string  `col:"company"`
	Designation   string  `col:"designation"`
	Department    string  `col:"department"`
	IsVeteran     bool    `col:"is-veteran"`
}

func GetTagsReflect() map[string]bool {
	e := User{}
	userProps := map[string]bool{}
	r := reflect.TypeOf(e)
	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)
		n := f.Name
		//print(n + " : ")
		if s, ex := r.FieldByName(n); ex {
			userProps[s.Tag.Get("col")] = true
			//println(s.Tag.Get("col"))
		}
	}
	return userProps
}

func GetTagName(i int) string {
	switch i {
	case 1:
		return "user-id"
	case 2:
		return "first-name"
	case 3:
		return "last-name"
	case 4:
		return "address"
	case 5:
		return "phone"
	case 6:
		return "country"
	case 7:
		return "pin"
	case 8:
		return "age"
	case 9:
		return "nominee"
	case 10:
		return "qualification"
	case 11:
		return "birth-date"
	case 12:
		return "salary"
	case 13:
		return "father"
	case 14:
		return "mother"
	case 15:
		return "spouse-name"
	case 16:
		return "num-children"
	case 17:
		return "vehicles-owned"
	case 18:
		return "company"
	case 19:
		return "designation"
	case 20:
		return "department"
	case 21:
		return "is-veteran"
	default:
		return ""
	}
}
