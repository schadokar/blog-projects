package main

import "testing"

func TestEnvVariable(t *testing.T) {

	value := envVariable("name")

	expected := "gopher"

	if value != expected {
		t.Errorf("Expected value %s, got %s", expected, value)
	}
}

func TestGoDotEnvVariable(t *testing.T) {

	key := "STRONGEST_AVENGER"

	value := goDotEnvVariable(key)

	expected := "Thor"

	if value != expected {
		t.Errorf("Expected value %s, got %s", expected, value)
	}
}

// func TestViperEnvVariable(t *testing.T) {
// 	key := "STRONGEST_AVENGER"

// 	value := viperEnvVariable(key)

// 	expected := "Thor"

// 	if value != expected {
// 		t.Errorf("Expected value %s, got %s", expected, value)
// 	}
// }

// func TestViperConfigVariable(t *testing.T) {

// 	key := "I_AM_INEVITABLE"

// 	value := viperConfigVariable(key)

// 	expected := "I am Iron Man"

// 	if value != expected {
// 		t.Errorf("Expected value %s, got %s", expected, value)
// 	}
// }
