package logging

import (
	"github.com/valyala/fasthttp"
	"reflect"
	"testing"
)

func Test_parseHeader(t *testing.T) {
	want := getTestData()

	requestHeader := &fasthttp.RequestHeader{}
	for sIndex := range want {
		for i := range want[sIndex] {
			requestHeader.Add(sIndex, want[sIndex][i])
		}
	}

	if got := parseHeader(requestHeader); !reflect.DeepEqual(got, want) {
		t.Errorf("parseHeader() = %v, want %v", got, want)
	}
}
