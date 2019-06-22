package reflect

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Mo"},
			[]string{"Mo"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Mo", "Riyadh"},
			[]string{"Mo", "Riyadh"},
		},
		{
			"Struct with non string fields",
			struct {
				Name string
				Age  int
			}{"Mo", 30},
			[]string{"Mo"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}

}
