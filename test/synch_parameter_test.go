package test

import (
	"geometris-go/parser"
	"reflect"
	"strings"
	"testing"
)

func TestSynchParameter(t *testing.T) {
	synchParam := "28.65.9.36.3.4.7.8.11.12.14.17.24.50.56.51.55.70.71.72.73.74.75.76.77.80.81.82"
	synchParameter := parser.NewSynchParameter(synchParam)
	if !reflect.DeepEqual(synchParameter.ColumnsID(), strings.Split(synchParam, ".")) {
		t.Error("Unexpected value")
	}
}
