package go_property_mapper

import (
	"fmt"
	"testing"
)

func BenchmarkTransform(t *testing.B) {
	for i := 0; i < t.N; i++ {
		a := &A{
			Id:      i,
			Name:    fmt.Sprintf("Kavish-%d", i),
			Address: fmt.Sprintf("PNBE-%d", i),
			Phone:   9019010 + int64(i),
			Salary:  504.13 + float64(i),
		}
		aSlice := &ASlice{
			Slice: []string{fmt.Sprintf("%d-Apple", i), fmt.Sprintf("%d-Ball", i), fmt.Sprintf("%d-Cat", i)},
		}
		aMap := &AMap{
			Map: map[int]string{1: fmt.Sprintf("one-%d", i), 100: fmt.Sprintf("hundred-%d", i), 36: fmt.Sprintf("thirty-six-%d", i)},
		}
		aComplex := &AComplex{
			A:        a,
			SlicePtr: aSlice,
			MapPtr:   aMap,
			Info:     fmt.Sprintf("AComplex-%d", i),
			Row:      3112022 + i,
		}
		transformACToBC(aComplex)
	}
}

func transformACToBC(aComplex *AComplex) *BComplex {
	b := transformAToB(aComplex.A)
	spr := &BSlice{Slice: aComplex.SlicePtr.Slice}
	mpr := &BMap{Map: aComplex.MapPtr.Map}
	return &BComplex{
		Info:     aComplex.Info,
		A:        b,
		SlicePtr: spr,
		Row:      aComplex.Row,
		MapPtr:   mpr,
	}
}

func transformAToB(a *A) *B {
	return &B{
		Id:      a.Id,
		Salary:  a.Salary,
		Name:    a.Name,
		Phone:   a.Phone,
		Address: a.Address,
	}
}

func BenchmarkTransformSingleLevel(t *testing.B) {
	for i := 0; i < t.N; i++ {
		a := &A{
			Id:      i,
			Name:    fmt.Sprintf("KK-%d", i),
			Address: fmt.Sprintf("Patna-%d", i),
			Phone:   9019000000 + int64(i),
			Salary:  55 + float64(i/10),
		}
		dest := &ADAO{}
		_, err := Transform(a, dest)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTransformCollection1(t *testing.B) {
	for i := 0; i < t.N; i++ {
		aSlice := &ASlice{
			Slice: []string{fmt.Sprintf("Apple-%d", i), fmt.Sprintf("Ball-%d", i), fmt.Sprintf("Cat-%d", i)},
		}
		bSlice := &BSlice{}
		_, err := Transform(aSlice, bSlice)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTransformCollection2(t *testing.B) {
	for i := 0; i < t.N; i++ {
		aMap := &AMap{
			Map: map[int]string{1 + i: "one", 100 + i: "hundred", 36 + i: "thirty-six"},
		}
		bMap := &BMap{}
		_, err := Transform(aMap, bMap)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTransformCollections(t *testing.B) {
	for i := 0; i < t.N; i++ {
		aCollection := &ACollection{
			Slice: []int{9 + i, 11 + i, 2001 + i, 30 + i, 9 + i},
			Map:   map[string]float64{fmt.Sprintf("house-%d", i): 3.0, fmt.Sprintf("car-%d", i): 2.2, fmt.Sprintf("weight-%d", i): 74.72},
		}
		bCollection := &BCollection{}
		_, err := Transform(aCollection, bCollection)
		if err != nil {
			panic(err)
		}
	}
}

func BenchmarkTransformComplex(t *testing.B) {
	for i := 0; i < t.N; i++ {
		a := &A{
			Id:      i,
			Name:    fmt.Sprintf("Kavish-%d", i),
			Address: fmt.Sprintf("PNBE-%d", i),
			Phone:   9019010 + int64(i),
			Salary:  504.13 + float64(i),
		}
		aSlice := &ASlice{
			Slice: []string{fmt.Sprintf("%d-Apple", i), fmt.Sprintf("%d-Ball", i), fmt.Sprintf("%d-Cat", i)},
		}
		aMap := &AMap{
			Map: map[int]string{1: fmt.Sprintf("one-%d", i), 100: fmt.Sprintf("hundred-%d", i), 36: fmt.Sprintf("thirty-six-%d", i)},
		}
		aComplex := &AComplex{
			A:        a,
			SlicePtr: aSlice,
			MapPtr:   aMap,
			Info:     fmt.Sprintf("AComplex-%d", i),
			Row:      3112022 + i,
		}
		bComplex := &BComplex{}
		_, err := Transform(aComplex, bComplex)
		if err != nil {
			panic(err)
		}
	}
}
