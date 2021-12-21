package main

import "testing"

func test_add_1_1(t *testing.T) {
	if add(1, 1) != 2 {
		t.Error("wrong result")
	}
}

func test_add_2_3(t *testing.T) {
	if add(2, 3) != 5 {
		t.Error("wrong result")
	}
}
