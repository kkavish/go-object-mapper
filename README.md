# golang-object-mapper
map one golang object to other, uses reflection

1. Transform takes pointers to both source and destination objects, see tests.
2. Works with complex structs with nested structs as properties.
3. Feel free to update the code/tests.


Benchmark Results
   debugserver-@(#)PROGRAM:LLDB  PROJECT:lldb-1316.0.9.46
   for arm64.
   Got a connection, launched process /private/var/folders/9w/pnljdlc955nb947mcdtlnrf00000gn/T/GoLand/___1gobench_github_com_angel_one_go_property_mapper.test (pid = 66445).
   goos: darwin
   goarch: arm64
   pkg: github.com/angel-one/go-property-mapper
   BenchmarkTransform
   BenchmarkTransform-10               	  660039	      1881 ns/op
   BenchmarkTransform-10               	  662365	      1867 ns/op
   BenchmarkTransform-10               	  667057	      1838 ns/op
   BenchmarkTransform-10               	  677646	      1847 ns/op
   BenchmarkTransform-10               	  681721	      1845 ns/op
   BenchmarkTransform-10               	  673772	      1854 ns/op
   BenchmarkTransform-10               	  650306	      1859 ns/op
   BenchmarkTransform-10               	  665235	      1864 ns/op
   BenchmarkTransform-10               	  671467	      1862 ns/op
   BenchmarkTransform-10               	  667620	      1869 ns/op
   BenchmarkTransformSingleLevel
   BenchmarkTransformSingleLevel-10    	  330792	      3771 ns/op
   BenchmarkTransformSingleLevel-10    	  328476	      3786 ns/op
   BenchmarkTransformSingleLevel-10    	  329373	      3838 ns/op
   BenchmarkTransformSingleLevel-10    	  327301	      3748 ns/op
   BenchmarkTransformSingleLevel-10    	  307011	      3953 ns/op
   BenchmarkTransformSingleLevel-10    	  302007	      4258 ns/op
   BenchmarkTransformSingleLevel-10    	  321460	      3808 ns/op
   BenchmarkTransformSingleLevel-10    	  330157	      3721 ns/op
   BenchmarkTransformSingleLevel-10    	  329463	      3732 ns/op
   BenchmarkTransformSingleLevel-10    	  330657	      3727 ns/op
   BenchmarkTransformCollection1
   BenchmarkTransformCollection1-10    	  886543	      1354 ns/op
   BenchmarkTransformCollection1-10    	  908239	      1362 ns/op
   BenchmarkTransformCollection1-10    	  907286	      1363 ns/op
   BenchmarkTransformCollection1-10    	  921127	      1353 ns/op
   BenchmarkTransformCollection1-10    	  902152	      1350 ns/op
   BenchmarkTransformCollection1-10    	  895257	      1357 ns/op
   BenchmarkTransformCollection1-10    	  897492	      1362 ns/op
   BenchmarkTransformCollection1-10    	  870526	      1350 ns/op
   BenchmarkTransformCollection1-10    	  904780	      1349 ns/op
   BenchmarkTransformCollection1-10    	  901560	      1362 ns/op
   BenchmarkTransformCollection2
   BenchmarkTransformCollection2-10    	 1334337	       892.4 ns/op
   BenchmarkTransformCollection2-10    	 1340718	       895.7 ns/op
   BenchmarkTransformCollection2-10    	 1344886	       897.2 ns/op
   BenchmarkTransformCollection2-10    	 1341655	       909.2 ns/op
   BenchmarkTransformCollection2-10    	 1331512	       887.5 ns/op
   BenchmarkTransformCollection2-10    	 1327039	       905.8 ns/op
   BenchmarkTransformCollection2-10    	 1331307	       919.0 ns/op
   BenchmarkTransformCollection2-10    	 1343074	       889.9 ns/op
   BenchmarkTransformCollection2-10    	 1345050	       894.4 ns/op
   BenchmarkTransformCollection2-10    	 1346365	       893.8 ns/op
   BenchmarkTransformCollections
   BenchmarkTransformCollections-10    	  589357	      2107 ns/op
   BenchmarkTransformCollections-10    	  600440	      2131 ns/op
   BenchmarkTransformCollections-10    	  580111	      2069 ns/op
   BenchmarkTransformCollections-10    	  602272	      2073 ns/op
   BenchmarkTransformCollections-10    	  607626	      2080 ns/op
   BenchmarkTransformCollections-10    	  591717	      2079 ns/op
   BenchmarkTransformCollections-10    	  600542	      2074 ns/op
   BenchmarkTransformCollections-10    	  589729	      2051 ns/op
   BenchmarkTransformCollections-10    	  607506	      2081 ns/op
   BenchmarkTransformCollections-10    	  597606	      2082 ns/op
   BenchmarkTransformComplex
   BenchmarkTransformComplex-10        	  116577	     10444 ns/op
   BenchmarkTransformComplex-10        	  116157	     10398 ns/op
   BenchmarkTransformComplex-10        	  115240	     10354 ns/op
   BenchmarkTransformComplex-10        	  116320	     10325 ns/op
   BenchmarkTransformComplex-10        	  117358	     10337 ns/op
   BenchmarkTransformComplex-10        	  116035	     10343 ns/op
   BenchmarkTransformComplex-10        	  115976	     10395 ns/op
   BenchmarkTransformComplex-10        	  115610	     10445 ns/op
   BenchmarkTransformComplex-10        	  115321	     10328 ns/op
   BenchmarkTransformComplex-10        	  116313	     10420 ns/op
   PASS
   Exiting.