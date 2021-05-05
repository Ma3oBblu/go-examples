package internal

import "testing"
import "github.com/ma3obblu/masker"

//BenchmarkBool
//BenchmarkBool/check_text1_is_exist_in_bool_var
//BenchmarkBool/check_text1_is_exist_in_bool_var-8         	181123724	         6.56 ns/op	       0 B/op	       0 allocs/op
//BenchmarkBool/check_text3_is_not_exist_in_bool_var
//BenchmarkBool/check_text3_is_not_exist_in_bool_var-8     	75207015	        15.4 ns/op	       0 B/op	       0 allocs/op
func BenchmarkBool(b *testing.B) {
	boolVar := map[string]bool{
		"text1": true,
		"text2": true,
	}
	testsForBool := []struct {
		name         string
		checkedValue string
		boolVar      map[string]bool
	}{
		{
			"check text1 is exist in bool var",
			"text1",
			boolVar,
		},
		{
			"check text3 is not exist in bool var",
			"text3",
			boolVar,
		},
	}
	for _, tt := range testsForBool {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, ok := tt.boolVar[tt.checkedValue]; ok {
					continue
				}
			}
		})
	}
}

//BenchmarkStruct
//BenchmarkStruct/check_text1_is_exist_in_struct_var
//BenchmarkStruct/check_text1_is_exist_in_struct_var-8         	183780481	         6.50 ns/op	       0 B/op	       0 allocs/op
//BenchmarkStruct/check_text3_is_not_exist_in_struct_var
//BenchmarkStruct/check_text3_is_not_exist_in_struct_var-8     	74461951	        15.3 ns/op	       0 B/op	       0 allocs/op
func BenchmarkStruct(b *testing.B) {
	structVar := map[string]struct{}{
		"text1": {},
		"text2": {},
	}
	testsForStruct := []struct {
		name         string
		checkedValue string
		structVar    map[string]struct{}
	}{
		{
			"check text1 is exist in struct var",
			"text1",
			structVar,
		},
		{
			"check text3 is not exist in struct var",
			"text3",
			structVar,
		},
	}
	for _, tt := range testsForStruct {
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				if _, ok := tt.structVar[tt.checkedValue]; ok {
					continue
				}
			}
		})
	}
}

// BenchmarkStringMethods/check_ConvertSymbolsWithRune
//BenchmarkStringMethods/check_ConvertSymbolsWithRune-8         	 4831677	       227.7 ns/op	      24 B/op	       1 allocs/op
//BenchmarkStringMethods/check_ConvertSymbolsWithBytes
//BenchmarkStringMethods/check_ConvertSymbolsWithBytes-8        	13781157	        82.14 ns/op	      48 B/op	       2 allocs/op
//BenchmarkStringMethods/check_ConvertSymbolsWithSlicing
//BenchmarkStringMethods/check_ConvertSymbolsWithSlicing-8      	 9721963	       119.0 ns/op	      16 B/op	       1 allocs/op
//BenchmarkStringMethods/sensitive
//BenchmarkStringMethods/sensitive-8                            	 1458075	       847.4 ns/op	     290 B/op	       4 allocs/op
//BenchmarkStringMethods/masker
//BenchmarkStringMethods/masker-8                               	 4024092	       292.0 ns/op	      56 B/op	       2 allocs/op
//BenchmarkStringMethods/masker#01
//BenchmarkStringMethods/masker#01-8                            	 9733102	       119.4 ns/op	      32 B/op	       2 allocs/op
func BenchmarkStringMethods(b *testing.B) {
	testsForStruct := []struct {
		name string
		make func()
	}{
		{
			"check ConvertSymbolsWithRune",
			func() { ConvertSymbolsWithRune("Ma3oBblu@gmail.com") },
		},
		{
			"check ConvertSymbolsWithBytes",
			func() { ConvertSymbolsWithBytes("Ma3oBblu@gmail.com") },
		},
		{
			"check ConvertSymbolsWithSlicing",
			func() { ConvertSymbolsWithSlicing("Ma3oBblu@gmail.com") },
		},
		{
			"sensitive",
			func() { MaskSensitiveDataRunes("123456") },
		},
		{
			"masker",
			func() { masker.Email("Ma3oBblu@gmail.com") },
		},
		{
			"masker",
			func() { masker.LastFourDigits("79191232323") },
		},
	}
	for _, tt := range testsForStruct {
		tt := tt
		b.Run(tt.name, func(b *testing.B) {
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				tt.make()
			}
		})
	}
}
