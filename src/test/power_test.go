package main

import "testing"

/*
 * TestPower will test whether the individual program power.go
 * works fine and gives correct results
 * naming conventions are to be followed strictly
 * unit test function name should start with TestXxx
 * Xxx is the function we are testing should start with capital letter
 * also filename should be xxx_test.go
 */
func TestPower(t *testing.T) {
	// base = 0, exp = n (>0)
	test1 := pow(0, 23)
	// test1 should have  output as 0 as 0^n = 0
	// if any other result if found then it has failed the test, we now throw an error

	if test1 != 0 {
		t.Errorf("pow(0, 23) = %f; expected: 0", test1)
	}

	// base = 45, exp = 0 should return 1, as n^0 = 1
	test2 := pow(45, 0)

	if test2 != 1 {
		t.Errorf("pow(45,0) = %f; expected: 1", test2)
	}
}
