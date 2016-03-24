package isin

import (
	"reflect"
	"testing"
)

func TestNewIsin(t *testing.T) {
	isin, err := NewIsin("US037833100")

	if err == nil {
		if isin.CountryCode != "US" {
			t.Fatalf("Expected %s, got %s", "US", isin.CountryCode)
		}

		if isin.Nsin != "037833100" {
			t.Fatalf("Expected %s, got %s", "037833100", isin.Nsin)
		}
	}

	_, err2 := NewIsin("")

	if err2 == nil {
		t.Fatalf("Expected %s, got %s", err2.Msg, err2)
	}

	_, err3 := NewIsin("US0378331")

	if err3 == nil {
		t.Fatalf("Expected %s, got %s", err3.Msg, err3)
	}

	_, err4 := NewIsin("11037833100")

	if err4 == nil {
		t.Fatalf("Expected %s, got %s", err4.Msg, err4)
	}
}

func TestFormat(t *testing.T) {
	expected := "US0378331005"
	isin, _ := NewIsin(expected)
	result := isin.Format()

	if result != expected {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}

func TestValid(t *testing.T) {
	str := "US0378331005"

	isin, _ := NewIsin(str)
	result := isin.Valid()

	if result != true {
		t.Fatalf("Expected true, got %t", result)
	}

	str2 := "US0378331004"
	isin2, _ := NewIsin(str2)
	result2 := isin2.Valid()

	if result2 != false {
		t.Fatalf("Expected false, got %t", result2)
	}
}

func TestChecksum(t *testing.T) {
	str := "US0378331005"
	isin, _ := NewIsin(str)
	result := isin.Checksum()

	if result != 5 {
		t.Fatalf("Expected %s, got %s", 5, result)
	}

	str2 := "AU0000XVGZA3"
	isin2, _ := NewIsin(str2)
	result2 := isin2.Checksum()

	if result2 != 3 {
		t.Fatalf("Expected %s, got %s", 3, result2)
	}

	str3 := "AU0000VXGZA3"
	isin3, _ := NewIsin(str3)
	result3 := isin3.Checksum()

	if result3 != 3 {
		t.Fatalf("Expected %s, got %s", 3, result3)
	}

	str4 := "IE00B3RBWM25"
	isin4, _ := NewIsin(str4)
	result4 := isin4.Checksum()

	if result4 != 5 {
		t.Fatalf("Expected %s, got %s", 5, result4)
	}
}

func TestDigits(t *testing.T) {
	isin, _ := NewIsin("US0378331005")
	result := isin.digits()
	expected := []int{3, 0, 2, 8, 0, 3, 7, 8, 3, 3, 1, 0, 0}

	if reflect.DeepEqual(result, expected) != true {
		t.Fatalf("Expected %s, got %s", expected, result)
	}
}
