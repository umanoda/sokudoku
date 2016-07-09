package sokudoku

import (
	"reflect"
	"strings"
	"testing"
)

func TestShowWord(t *testing.T) {
	type testCase struct {
		sample string
		expect []string
	}

	space := func(n int) string { return strings.Repeat(" ", 22+n) }

	for _, test_case := range []testCase{
		{"あ", []string{space(0), "", "あ", ""}},
		{"あい", []string{space(0), "", "あ", "い"}},
		{"あいう", []string{space(-2), "あ", "い", "う"}},
		{"あいうえ", []string{space(-2), "あ", "い", "うえ"}},
		{"あいうえお", []string{space(-4), "あい", "う", "えお"}},
	} {
		got := _showWord(test_case.sample)
		expect := test_case.expect
		if !reflect.DeepEqual(got, expect) {
			t.Error(
				"\n   got:", got,
				"\nexpect:", expect,
			)
		}
	}
}
