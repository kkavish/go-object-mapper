package go_property_mapper

import (
	"fmt"
	"testing"
)

//goos: darwin
//goarch: arm64
//pkg: github.com/angel-one/go-property-mapper
//BenchmarkTagReflectReader-10              272232              4294 ns/op
//BenchmarkTagReflectReader-10              276799              4301 ns/op
//BenchmarkTagReflectReader-10              280483              4263 ns/op
//BenchmarkTagReflectReader-10              272572              4279 ns/op
//BenchmarkTagReflectReader-10              276536              4298 ns/op
//BenchmarkTagReflectReader-10              278106              4288 ns/op
//BenchmarkTagReflectReader-10              278223              4318 ns/op
//BenchmarkTagReflectReader-10              280474              4279 ns/op
//BenchmarkTagReflectReader-10              278982              4304 ns/op
//BenchmarkTagReflectReader-10              264958              4269 ns/op
//BenchmarkTagReflectReader-10              278872              4295 ns/op
//BenchmarkTagReflectReader-10              277105              4278 ns/op
//BenchmarkTagReflectReader-10              276721              4300 ns/op
//BenchmarkTagReflectReader-10              279140              4276 ns/op
//BenchmarkTagReflectReader-10              279638              4283 ns/op
//BenchmarkTagReflectReader-10              275811              4280 ns/op
//BenchmarkTagReflectReader-10              280012              4280 ns/op
//BenchmarkTagReflectReader-10              273303              4264 ns/op
//BenchmarkTagReflectReader-10              280342              4303 ns/op
//BenchmarkTagReflectReader-10              272752              4272 ns/op
//BenchmarkTagReflectReader-10              278098              4310 ns/op
//BenchmarkTagReflectReader-10              281055              4274 ns/op
//BenchmarkTagReflectReader-10              278434              4278 ns/op
//BenchmarkTagReflectReader-10              280435              4309 ns/op
//BenchmarkTagReflectReader-10              276349              4294 ns/op
//BenchmarkTagReflectReader-10              279258              4296 ns/op
//BenchmarkTagReflectReader-10              275833              4278 ns/op
//BenchmarkTagReflectReader-10              280849              4271 ns/op
//BenchmarkTagReflectReader-10              279808              4283 ns/op
//BenchmarkTagReflectReader-10              280036              4569 ns/op
//BenchmarkTagReflectReader-10              255748              4702 ns/op
//BenchmarkTagReflectReader-10              232551              5215 ns/op
//BenchmarkTagReflectReader-10              267530              4451 ns/op
//BenchmarkTagReflectReader-10              266203              4627 ns/op
//BenchmarkTagReflectReader-10              265858              4556 ns/op
//BenchmarkTagReflectReader-10              247594              4499 ns/op
//BenchmarkTagReflectReader-10              266880              4350 ns/op
//BenchmarkTagReflectReader-10              281048              4305 ns/op
//BenchmarkTagReflectReader-10              277050              4347 ns/op
//BenchmarkTagReflectReader-10              280206              4348 ns/op
//BenchmarkTagReflectReader-10              256393              4380 ns/op
//BenchmarkTagReflectReader-10              279542              4340 ns/op
//BenchmarkTagReflectReader-10              273573              4861 ns/op
//BenchmarkTagReflectReader-10              248766              4395 ns/op
//BenchmarkTagReflectReader-10              260176              4463 ns/op
//BenchmarkTagReflectReader-10              280059              4374 ns/op
//BenchmarkTagReflectReader-10              275054              4384 ns/op
//BenchmarkTagReflectReader-10              278782              4334 ns/op
//BenchmarkTagReflectReader-10              254760              4328 ns/op
//BenchmarkTagReflectReader-10              278898              4309 ns/op
//BenchmarkTagReflectReader-10              275514              4318 ns/op
//BenchmarkTagReflectReader-10              279585              4301 ns/op
//BenchmarkTagReflectReader-10              276847              4332 ns/op
//BenchmarkTagReflectReader-10              270124              4458 ns/op
//BenchmarkTagReflectReader-10              261871              4667 ns/op
//BenchmarkTagReflectReader-10              262738              4404 ns/op
//BenchmarkTagReflectReader-10              259376              4383 ns/op
//BenchmarkTagReflectReader-10              276800              4395 ns/op
//BenchmarkTagReflectReader-10              279196              4359 ns/op
//BenchmarkTagReflectReader-10              278269              4421 ns/op
//BenchmarkTagReflectReader-10              278037              4312 ns/op
//BenchmarkTagReflectReader-10              278121              4377 ns/op
//BenchmarkTagReflectReader-10              259596              4510 ns/op
//BenchmarkTagReflectReader-10              267967              4492 ns/op
//BenchmarkTagReflectReader-10              270468              4367 ns/op
//BenchmarkTagReflectReader-10              259072              4300 ns/op
//BenchmarkTagReflectReader-10              279121              4279 ns/op
//BenchmarkTagReflectReader-10              279237              4362 ns/op
//BenchmarkTagReflectReader-10              268478              4418 ns/op
//BenchmarkTagReflectReader-10              268014              4464 ns/op
//BenchmarkTagReflectReader-10              261819              4533 ns/op
//BenchmarkTagReflectReader-10              260166              4479 ns/op
//BenchmarkTagReflectReader-10              278318              4538 ns/op
//BenchmarkTagReflectReader-10              260420              4364 ns/op
//BenchmarkTagReflectReader-10              278216              4330 ns/op
//BenchmarkTagReflectReader-10              277549              4421 ns/op
//BenchmarkTagReflectReader-10              262629              4443 ns/op
//BenchmarkTagReflectReader-10              279684              4369 ns/op
//BenchmarkTagReflectReader-10              278092              4537 ns/op
//BenchmarkTagReflectReader-10              277021              4346 ns/op
//BenchmarkTagReflectReader-10              272673              4433 ns/op
//BenchmarkTagReflectReader-10              282728              4285 ns/op
//BenchmarkTagReflectReader-10              280837              4288 ns/op
//BenchmarkTagReflectReader-10              279283              4411 ns/op
//BenchmarkTagReflectReader-10              259321              4452 ns/op
//BenchmarkTagReflectReader-10              275100              4420 ns/op
//BenchmarkTagReflectReader-10              281548              4500 ns/op
//BenchmarkTagReflectReader-10              260710              4412 ns/op
//BenchmarkTagReflectReader-10              271472              4415 ns/op
//BenchmarkTagReflectReader-10              271699              4477 ns/op
//BenchmarkTagReflectReader-10              267234              4485 ns/op
//BenchmarkTagReflectReader-10              279028              4479 ns/op
//BenchmarkTagReflectReader-10              278979              4354 ns/op
//BenchmarkTagReflectReader-10              264603              4555 ns/op
//BenchmarkTagReflectReader-10              253518              4507 ns/op
//BenchmarkTagReflectReader-10              264253              4413 ns/op
//BenchmarkTagReflectReader-10              279529              4424 ns/op
//BenchmarkTagReflectReader-10              281443              4401 ns/op
//BenchmarkTagReflectReader-10              263876              4444 ns/op
//BenchmarkTagReflectReader-10              265516              4458 ns/op
//BenchmarkTagReader-10                   19554284                62.17 ns/op
//BenchmarkTagReader-10                   20006779                61.87 ns/op
//BenchmarkTagReader-10                   19319587                61.27 ns/op
//BenchmarkTagReader-10                   19777000                59.73 ns/op
//BenchmarkTagReader-10                   19373794                60.27 ns/op
//BenchmarkTagReader-10                   19941656                62.25 ns/op
//BenchmarkTagReader-10                   19083691                63.44 ns/op
//BenchmarkTagReader-10                   19727080                62.54 ns/op
//BenchmarkTagReader-10                   22444939                60.66 ns/op
//BenchmarkTagReader-10                   19873402                58.89 ns/op
//BenchmarkTagReader-10                   21164316                60.58 ns/op
//BenchmarkTagReader-10                   20569503                58.81 ns/op
//BenchmarkTagReader-10                   19993321                59.71 ns/op
//BenchmarkTagReader-10                   19865082                58.43 ns/op
//BenchmarkTagReader-10                   19944790                59.89 ns/op
//BenchmarkTagReader-10                   19961600                59.35 ns/op
//BenchmarkTagReader-10                   19356346                61.93 ns/op
//BenchmarkTagReader-10                   19756568                62.20 ns/op
//BenchmarkTagReader-10                   19920290                61.75 ns/op
//BenchmarkTagReader-10                   21035670                59.45 ns/op
//BenchmarkTagReader-10                   19982750                62.37 ns/op
//BenchmarkTagReader-10                   18928489                63.38 ns/op
//BenchmarkTagReader-10                   19774406                60.54 ns/op
//BenchmarkTagReader-10                   19904018                59.22 ns/op
//BenchmarkTagReader-10                   19884338                58.96 ns/op
//BenchmarkTagReader-10                   19877161                61.29 ns/op
//BenchmarkTagReader-10                   19637658                60.13 ns/op
//BenchmarkTagReader-10                   19031330                62.45 ns/op
//BenchmarkTagReader-10                   20551917                59.72 ns/op
//BenchmarkTagReader-10                   18596097                61.70 ns/op
//BenchmarkTagReader-10                   19903893                60.84 ns/op
//BenchmarkTagReader-10                   19743985                61.08 ns/op
//BenchmarkTagReader-10                   19535528                62.54 ns/op
//BenchmarkTagReader-10                   18710373                59.76 ns/op
//BenchmarkTagReader-10                   19774786                59.49 ns/op
//BenchmarkTagReader-10                   19920318                60.24 ns/op
//BenchmarkTagReader-10                   19706764                60.87 ns/op
//BenchmarkTagReader-10                   18756483                63.15 ns/op
//BenchmarkTagReader-10                   19509049                59.32 ns/op
//BenchmarkTagReader-10                   22346022                60.67 ns/op
//BenchmarkTagReader-10                   19189905                64.01 ns/op
//BenchmarkTagReader-10                   17842328                71.49 ns/op
//BenchmarkTagReader-10                   19429696                63.46 ns/op
//BenchmarkTagReader-10                   19401307                63.49 ns/op
//BenchmarkTagReader-10                   18892790                64.11 ns/op
//BenchmarkTagReader-10                   18801213                63.55 ns/op
//BenchmarkTagReader-10                   18620685                62.65 ns/op
//BenchmarkTagReader-10                   18898852                64.22 ns/op
//BenchmarkTagReader-10                   19748588                63.87 ns/op
//BenchmarkTagReader-10                   18977257                63.37 ns/op
//BenchmarkTagReader-10                   19574246                62.29 ns/op
//BenchmarkTagReader-10                   17794429                62.34 ns/op
//BenchmarkTagReader-10                   19198195                63.60 ns/op
//BenchmarkTagReader-10                   19502931                61.67 ns/op
//BenchmarkTagReader-10                   19468496                61.63 ns/op
//BenchmarkTagReader-10                   19051612                60.60 ns/op
//BenchmarkTagReader-10                   19598983                61.51 ns/op
//BenchmarkTagReader-10                   19438798                62.46 ns/op
//BenchmarkTagReader-10                   19481638                60.71 ns/op
//BenchmarkTagReader-10                   20586058                60.11 ns/op
//BenchmarkTagReader-10                   19662165                59.76 ns/op
//BenchmarkTagReader-10                   19950207                60.27 ns/op
//BenchmarkTagReader-10                   19965572                60.22 ns/op
//BenchmarkTagReader-10                   19921048                59.71 ns/op
//BenchmarkTagReader-10                   19373782                60.21 ns/op
//BenchmarkTagReader-10                   19901322                60.26 ns/op
//BenchmarkTagReader-10                   19868755                59.47 ns/op
//BenchmarkTagReader-10                   19700484                60.26 ns/op
//BenchmarkTagReader-10                   19575842                60.45 ns/op
//BenchmarkTagReader-10                   19926589                60.39 ns/op
//BenchmarkTagReader-10                   19710730                59.85 ns/op
//BenchmarkTagReader-10                   19477804                60.22 ns/op
//BenchmarkTagReader-10                   19353486                60.70 ns/op
//BenchmarkTagReader-10                   18697352                63.72 ns/op
//BenchmarkTagReader-10                   18916566                63.04 ns/op
//BenchmarkTagReader-10                   19165756                62.22 ns/op
//BenchmarkTagReader-10                   19752015                62.61 ns/op
//BenchmarkTagReader-10                   18925404                63.61 ns/op
//BenchmarkTagReader-10                   18618686                62.36 ns/op
//BenchmarkTagReader-10                   20671759                60.08 ns/op
//BenchmarkTagReader-10                   19873046                60.40 ns/op
//BenchmarkTagReader-10                   19610713                60.13 ns/op
//BenchmarkTagReader-10                   20826643                59.11 ns/op
//BenchmarkTagReader-10                   20609364                60.25 ns/op
//BenchmarkTagReader-10                   19942954                59.91 ns/op
//BenchmarkTagReader-10                   19939143                59.78 ns/op
//BenchmarkTagReader-10                   18333088                64.35 ns/op
//BenchmarkTagReader-10                   18970694                62.66 ns/op
//BenchmarkTagReader-10                   19465956                62.61 ns/op
//BenchmarkTagReader-10                   19804035                61.13 ns/op
//BenchmarkTagReader-10                   19722514                60.71 ns/op
//BenchmarkTagReader-10                   19473076                60.37 ns/op
//BenchmarkTagReader-10                   19928781                60.23 ns/op
//BenchmarkTagReader-10                   19248441                60.27 ns/op
//BenchmarkTagReader-10                   20589693                58.46 ns/op
//BenchmarkTagReader-10                   20845471                60.35 ns/op
//BenchmarkTagReader-10                   19550978                60.21 ns/op
//BenchmarkTagReader-10                   21815107                60.15 ns/op
//BenchmarkTagReader-10                   19624383                60.19 ns/op
//BenchmarkTagReader-10                   19607802                60.71 ns/op
//BenchmarkTransform-10                    1814976               659.9 ns/op
//BenchmarkTransform-10                    1922643               625.4 ns/op
//BenchmarkTransform-10                    1789682               669.8 ns/op
//BenchmarkTransform-10                    1787217               683.0 ns/op
//BenchmarkTransform-10                    1869127               639.1 ns/op
//BenchmarkTransform-10                    1754010               712.7 ns/op
//BenchmarkTransform-10                    1790620               660.5 ns/op
//BenchmarkTransform-10                    1845256               661.1 ns/op
//BenchmarkTransform-10                    1863577               626.5 ns/op
//BenchmarkTransform-10                    1909795               627.1 ns/op
//BenchmarkTransform-10                    1930539               625.3 ns/op
//BenchmarkTransform-10                    1883419               636.0 ns/op
//BenchmarkTransform-10                    1932385               632.0 ns/op
//BenchmarkTransform-10                    1930437               624.1 ns/op
//BenchmarkTransform-10                    1922174               647.3 ns/op
//BenchmarkTransform-10                    1782777               737.3 ns/op
//BenchmarkTransform-10                    1854876               641.7 ns/op
//BenchmarkTransform-10                    1877692               635.8 ns/op
//BenchmarkTransform-10                    1905655               630.4 ns/op
//BenchmarkTransform-10                    1814412               669.6 ns/op
//BenchmarkTransform-10                    1847413               633.7 ns/op
//BenchmarkTransform-10                    1907743               626.1 ns/op
//BenchmarkTransform-10                    1915254               627.1 ns/op
//BenchmarkTransform-10                    1917787               628.2 ns/op
//BenchmarkTransform-10                    1895367               634.0 ns/op
//BenchmarkTransform-10                    1916192               636.0 ns/op
//BenchmarkTransform-10                    1878591               645.9 ns/op
//BenchmarkTransform-10                    1900256               638.4 ns/op
//BenchmarkTransform-10                    1895178               651.1 ns/op
//BenchmarkTransform-10                    1901004               631.2 ns/op
//BenchmarkTransform-10                    1928004               634.2 ns/op
//BenchmarkTransform-10                    1846032               627.4 ns/op
//BenchmarkTransform-10                    1800304               647.7 ns/op
//BenchmarkTransform-10                    1900329               626.8 ns/op
//BenchmarkTransform-10                    1923495               626.3 ns/op
//BenchmarkTransform-10                    1916299               661.2 ns/op
//BenchmarkTransform-10                    1820174               655.2 ns/op
//BenchmarkTransform-10                    1817926               647.3 ns/op
//BenchmarkTransform-10                    1874398               636.8 ns/op
//BenchmarkTransform-10                    1810844               653.8 ns/op
//BenchmarkTransform-10                    1862674               641.4 ns/op
//BenchmarkTransform-10                    1894236               634.1 ns/op
//BenchmarkTransform-10                    1793372               657.9 ns/op
//BenchmarkTransform-10                    1858160               646.9 ns/op
//BenchmarkTransform-10                    1840293               638.5 ns/op
//BenchmarkTransform-10                    1859895               625.9 ns/op
//BenchmarkTransform-10                    1909354               659.7 ns/op
//BenchmarkTransform-10                    1848919               672.1 ns/op
//BenchmarkTransform-10                    1924538               633.7 ns/op
//BenchmarkTransform-10                    1875637               653.4 ns/op
//BenchmarkTransform-10                    1882368               649.4 ns/op
//BenchmarkTransform-10                    1930514               631.3 ns/op
//BenchmarkTransform-10                    1850397               646.5 ns/op
//BenchmarkTransform-10                    1871373               654.3 ns/op
//BenchmarkTransform-10                    1921468               628.4 ns/op
//BenchmarkTransform-10                    1899313               641.9 ns/op
//BenchmarkTransform-10                    1887372               645.3 ns/op
//BenchmarkTransform-10                    1895610               628.7 ns/op
//BenchmarkTransform-10                    1899655               626.6 ns/op
//BenchmarkTransform-10                    1933438               626.8 ns/op
//BenchmarkTransform-10                    1917310               648.4 ns/op
//BenchmarkTransform-10                    1932554               629.3 ns/op
//BenchmarkTransform-10                    1928004               626.5 ns/op
//BenchmarkTransform-10                    1900200               646.7 ns/op
//BenchmarkTransform-10                    1931527               645.0 ns/op
//BenchmarkTransform-10                    1904152               640.4 ns/op
//BenchmarkTransform-10                    1895478               640.6 ns/op
//BenchmarkTransform-10                    1909050               642.6 ns/op
//BenchmarkTransform-10                    1683219               633.8 ns/op
//BenchmarkTransform-10                    1934443               626.8 ns/op
//BenchmarkTransform-10                    1936051               640.3 ns/op
//BenchmarkTransform-10                    1895665               635.8 ns/op
//BenchmarkTransform-10                    1896244               635.6 ns/op
//BenchmarkTransform-10                    1912848               625.1 ns/op
//BenchmarkTransform-10                    1929219               633.3 ns/op
//BenchmarkTransform-10                    1931035               625.9 ns/op
//BenchmarkTransform-10                    1932465               627.3 ns/op
//BenchmarkTransform-10                    1934853               626.6 ns/op
//BenchmarkTransform-10                    1931722               625.2 ns/op
//BenchmarkTransform-10                    1933102               634.1 ns/op
//BenchmarkTransform-10                    1891965               626.4 ns/op
//BenchmarkTransform-10                    1911234               627.1 ns/op
//BenchmarkTransform-10                    1934227               636.6 ns/op
//BenchmarkTransform-10                    1933237               627.6 ns/op
//BenchmarkTransform-10                    1919881               631.7 ns/op
//BenchmarkTransform-10                    1920494               630.8 ns/op
//BenchmarkTransform-10                    1929314               625.8 ns/op
//BenchmarkTransform-10                    1931521               624.1 ns/op
//BenchmarkTransform-10                    1933783               653.2 ns/op
//BenchmarkTransform-10                    1803085               867.4 ns/op
//BenchmarkTransform-10                    1682533               696.4 ns/op
//BenchmarkTransform-10                    1758164               652.0 ns/op
//BenchmarkTransform-10                    1922739               650.4 ns/op
//BenchmarkTransform-10                    1847094               674.1 ns/op
//BenchmarkTransform-10                    1855372               635.2 ns/op
//BenchmarkTransform-10                    1854753               624.8 ns/op
//BenchmarkTransform-10                    1917486               626.9 ns/op
//BenchmarkTransform-10                    1847856               662.0 ns/op
//BenchmarkTransform-10                    1806765               637.0 ns/op
//BenchmarkTransform-10                    1870268               657.0 ns/op

// BenchmarkTagReflectReader-10    	   70278	     16710 ns/op
// BenchmarkTagReflectReader-10    	  281818	      4244 ns/op
func BenchmarkTagReflectReader(t *testing.B) {
	for i := 0; i < t.N; i++ {
		/*m := */ GetTagsReflect()
		//println(fmt.Sprintf("%v", m))
	}
}

// BenchmarkTagReader-10    	  108134	      9852 ns/op
// BenchmarkTagReader-10    	18109801	        59.87 ns/op
func BenchmarkTagReader(t *testing.B) {
	for i := 0; i < t.N; i++ {
		//p := make([]string, 21)
		for j := 1; j < 22; j++ {
			/*p[j-1] = */ GetTagName(j)
		}
		//println(fmt.Sprintf("%v", p))
	}
}

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
