package shellwords_test

import (
	"reflect"
	"testing"

	"github.com/buildkite/shellwords"
)

func TestSplitBatch(t *testing.T) {
	var testCases = []struct {
		String   string
		Expected []string
	}{
		{
			`"\\vmware-host\Shared Folders\src\github.com\buildkite\agent\llamas @% test\buildkite-agent.exe" start`, []string{
				`\\vmware-host\Shared Folders\src\github.com\buildkite\agent\llamas @% test\buildkite-agent.exe`,
				`start`,
			},
		},
		{
			`simple 🙌🏻 --string "quo""ted"`, []string{
				`simple`,
				`🙌🏻`,
				`--string`,
				`quo"ted`,
			},
		},

		{
			`simple --string "quo""ted"`, []string{
				`simple`,
				`--string`,
				`quo"ted`,
			},
		},
		{
			`mkdir "My favorite "^%OS^%`, []string{
				`mkdir`,
				`My favorite %OS%`,
			},
		},
		{
			`runme.exe /password:"~!@#$^%^^^&*()_+^|-=\][{}'^;:""/.>?,<"`, []string{
				`runme.exe`,
				`/password:~!@#$%^&*()_+|-=\][{}';:"/.>?,<`,
			},
		},
		{
			`echo ^^^^^&`, []string{
				`echo`,
				`^^&`,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run("", func(t *testing.T) {
			actual, err := shellwords.SplitBatch(tc.String)
			if err != nil {
				t.Fatal(err)
			}

			if !reflect.DeepEqual(tc.Expected, actual) {
				t.Fatalf("Expected vs Actual: \n%#v\n\n%#v", tc.Expected, actual)
			}
		})
	}
}
