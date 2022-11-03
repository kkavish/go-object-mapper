package go_property_mapper

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTrasformSingleLevel(t *testing.T) {
	a := &A{
		Id:      90382,
		Name:    "KK",
		Address: "Patna",
		Phone:   9019010,
		Salary:  55,
	}
	dest := &ADAO{}
	obj, err := Transform(a, dest)
	if err != nil {
		panic(err)
	}
	dao := obj.(*ADAO)
	assert.Equal(t, dao.Id, 90382)
	assert.Equal(t, dao.Name, "KK")
	assert.Equal(t, dao.Address, "Patna")
	assert.Equal(t, dao.Phone, int64(9019010))
	assert.Equal(t, dao.Salary, float64(55))
}

func TestTransformCollection(t *testing.T) {
	aSlice := &ASlice{
		Slice: []string{"Apple", "Ball", "Cat"},
	}
	bSlice := &BSlice{}
	obj, err := Transform(aSlice, bSlice)
	if err != nil {
		panic(err)
	}
	bs := obj.(*BSlice)
	assert.Equal(t, len(bs.Slice), 3)
	assert.Equal(t, bs.Slice[0], "Apple")
	assert.Equal(t, bs.Slice[1], "Ball")
	assert.Equal(t, bs.Slice[2], "Cat")

	aMap := &AMap{
		Map: map[int]string{1: "one", 100: "hundred", 36: "thirty-six"},
	}
	bMap := &BMap{}
	obj, err = Transform(aMap, bMap)
	if err != nil {
		panic(err)
	}
	bm := obj.(*BMap)
	assert.Equal(t, len(bm.Map), 3)
	assert.Equal(t, bm.Map[1], "one")
	assert.Equal(t, bm.Map[100], "hundred")
	assert.Equal(t, bm.Map[36], "thirty-six")
}

func TestTransformCollections(t *testing.T) {
	aCollection := &ACollection{
		Slice: []int{9, 11, 2001, 30, 9},
		Map:   map[string]float64{"house": 3.0, "car": 2.2, "weight": 74.72},
	}
	bCollection := &BCollection{}
	obj, err := Transform(aCollection, bCollection)
	if err != nil {
		panic(err)
	}
	bm := obj.(*BCollection)
	assert.Equal(t, len(bm.Slice), 5)
	assert.Equal(t, len(bm.Map), 3)
	assert.Equal(t, bm.Slice[0], 9)
	assert.Equal(t, bm.Slice[1], 11)
	assert.Equal(t, bm.Slice[2], 2001)
	assert.Equal(t, bm.Slice[3], 30)
	assert.Equal(t, bm.Slice[4], 9)
	assert.Equal(t, bm.Map["house"], float64(3.0))
	assert.Equal(t, bm.Map["car"], float64(2.2))
	assert.Equal(t, bm.Map["weight"], float64(74.72))
}

func TestTransformComplex(t *testing.T) {
	a := &A{
		Id:      90382,
		Name:    "KK",
		Address: "Patna",
		Phone:   9019010,
		Salary:  55,
	}
	aSlice := &ASlice{
		Slice: []string{"Apple", "Ball", "Cat"},
	}
	aMap := &AMap{
		Map: map[int]string{1: "one", 100: "hundred", 36: "thirty-six"},
	}
	aComplex := &AComplex{
		A:        a,
		SlicePtr: aSlice,
		MapPtr:   aMap,
		Info:     "AComplex",
		Row:      3112022,
	}
	bComplex := &BComplex{}
	obj, err := Transform(aComplex, bComplex)
	if err != nil {
		panic(err)
	}
	bm := obj.(*BComplex)
	assert.Equal(t, bm.A.Id, 90382)
	assert.Equal(t, bm.A.Name, "KK")
	assert.Equal(t, bm.A.Address, "Patna")
	assert.Equal(t, bm.A.Phone, int64(9019010))
	assert.Equal(t, bm.A.Salary, float64(55))
	assert.Equal(t, len(bm.SlicePtr.Slice), 3)
	assert.Equal(t, bm.SlicePtr.Slice[0], "Apple")
	assert.Equal(t, bm.SlicePtr.Slice[1], "Ball")
	assert.Equal(t, bm.SlicePtr.Slice[2], "Cat")
	assert.Equal(t, len(bm.MapPtr.Map), 3)
	assert.Equal(t, bm.MapPtr.Map[1], "one")
	assert.Equal(t, bm.MapPtr.Map[100], "hundred")
	assert.Equal(t, bm.MapPtr.Map[36], "thirty-six")
	assert.Equal(t, bm.Info, "AComplex")
	assert.Equal(t, bm.Row, 3112022)
}
