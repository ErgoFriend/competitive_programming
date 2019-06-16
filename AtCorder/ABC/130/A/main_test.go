package main

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

type testValue struct {
	arg []string
	ans string
}

func getTests() []testValue {
	testValues := []testValue{
		testValue{
			[]string{
				"1",
				"2 3",
				"test",
			},
			"6 test"},
		testValue{
			[]string{
				"72",
				"128 256",
				"myonmyon",
			},
			"456 myonmyon"},
	}

	return testValues
}

func Test_main(t *testing.T) {
	tests := getTests()
	for i, tt := range tests {
		si := strconv.Itoa(i)
		t.Run("Case "+si, func(t *testing.T) {
			ret, err := stubio(strings.Join(tt.arg, " "), main)
			if err != "" {
				t.Errorf("error func: %s ", err)
			}
			ans := fmt.Sprintf("%s\n", tt.ans)
			if ret != ans {
				t.Errorf("Unexpected output: '%s' Need: '%s'", ret, ans)
			}
		})
	}
}
