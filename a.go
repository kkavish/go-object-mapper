package go_property_mapper

type A struct {
	Id      int
	Name    string
	Address string
	Phone   int64
	Salary  float64
}

type ASlice struct {
	Slice []string
}

type AMap struct {
	Map map[int]string
}

type ACollection struct {
	Slice []int
	Map   map[string]float64
}

type AComplex struct {
	A        *A
	SlicePtr *ASlice
	MapPtr   *AMap
	Info     string
	Row      int
}
