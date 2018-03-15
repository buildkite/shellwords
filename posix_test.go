package shellwords_test

import (
	"reflect"
	"testing"

	"github.com/buildkite/shellwords"
)

func TestSplitPosix(t *testing.T) {
	var testCases = []struct {
		String   string
		Expected []string
	}{
		{
			`simple --string "quoted"`, []string{
				`simple`,
				`--string`,
				`quoted`,
			},
		},
		{
			`\\\""quoted" llamas 'test\''`, []string{
				`\"quoted`,
				`llamas`,
				`test'`,
			},
		},
		{
			`/usr/bin/bash -e -c "llamas are the \"best\" && echo 'alpacas'"`, []string{
				`/usr/bin/bash`,
				`-e`,
				`-c`,
				`llamas are the "best" && echo 'alpacas'`,
			},
		},
		{
			`"/bin"/ba'sh' -c echo\ \\\\"fo real"`, []string{
				`/bin/bash`,
				`-c`,
				`echo \\fo real`,
			},
		},
		{
			`echo 'abc'\''abc'`, []string{
				`echo`,
				`abc'abc`,
			},
		},
		{
			`echo "abc"\""abc"`, []string{
				`echo`,
				`abc"abc`,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run("", func(t *testing.T) {
			actual, err := shellwords.SplitPosix(tc.String)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.Expected, actual) {
				t.Fatalf("Expected vs Actual: \n%#v\n\n%#v", tc.Expected, actual)
			}
		})
	}
}