package go_property_mapper

type B struct {
	Id      int
	Salary  float64
	Name    string
	Phone   int64
	Address string
}

type BSlice struct {
	Slice []string
}

type BMap struct {
	Map map[int]string
}

type BCollection struct {
	Map   map[string]float64
	Slice []int
}

type BComplex struct {
	Info     string
	A        *B
	SlicePtr *BSlice
	Row      int
	MapPtr   *BMap
}
