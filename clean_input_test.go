package main

import "testing"


func TestNormalInput(t *testing.T) {
	input := "hello"
	expected := []string{"hello"}
	actual := cleanInput(input)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestWhiteSpacesInput(t *testing.T) {
	input := "  hello     		"
	expected := []string{"hello"}
	actual := cleanInput(input)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}


func TestEmptyInput(t *testing.T) {
	input := ""
	expected := []string{}
	actual := cleanInput(input)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestMultipleWords(t *testing.T){
	input := "    hello    from the      other side "
	expected := []string{"hello", "from", "the", "other", "side"}
	actual := cleanInput(input)
	if len(actual) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
