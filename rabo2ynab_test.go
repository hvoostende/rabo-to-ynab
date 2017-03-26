package main

import (
	"errors"
	"testing"
)

func TestReadCSV(t *testing.T) {
	v, _ := readCSV("test.csv")
	_, err := readCSV("error")
	c := [][]string{{"test1", "test2"}, {"test3", "test4"}}
	if v[0][0] != c[0][0] {
		t.Error("Expected test1 got:", v[0][0])
	}

	errMessage := errors.New("open error: no such file or directory")
	if err.Error() != errMessage.Error() {
		t.Error("Expected open error: no such file or directory got:", err)
	}

}

func TestGetYNABDate(t *testing.T) {
	v := GetYNABDate("20161028")
	if v != "28/10/2016" {
		t.Error("Expected [[test1 test2] [test3 test4]] got:", v)
	}
}

func TestGetYNABDPayee(t *testing.T) {
	v1 := GetYNABPayee("", "second")
	v2 := GetYNABPayee("first", "second")
	if v1 != "second" {
		t.Error("Expected second got:", v1)
	}
	if v2 != "first" {
		t.Error("Expected first got:", v2)
	}
}

func TestGetYNABDInfow(t *testing.T) {
	v1 := GetYNABInflow("C", "100.1")
	v2 := GetYNABInflow("D", "100.1")
	if v1 != "100.1" {
		t.Error("Expected 100.1 got:", v1)
	}
	if v2 != "" {
		t.Error("Expected nothing got:", v2)
	}
}

func TestGetYNABDOutfow(t *testing.T) {
	v1 := GetYNABOutflow("C", "100.1")
	v2 := GetYNABOutflow("D", "100.1")
	if v1 != "" {
		t.Error("Expected nothing got:", v1)
	}
	if v2 != "100.1" {
		t.Error("Expected 100.1 got:", v2)
	}
}

func TestMain(t *testing.T) {
	input = "test.txt"
	main()

}
