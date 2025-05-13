package main

import (
	"testing"
)

func TestCalcAdd(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, 1, 0},
		{100, 200, 300},
	}

	for _, test := range tests {
		result := calcAdd(test.a, test.b)
		if result != test.expected {
			t.Errorf("calcAdd(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestCalcSub(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{5, 3, 2},
		{0, 0, 0},
		{-1, -1, 0},
		{100, 200, -100},
	}

	for _, test := range tests {
		result := calcSub(test.a, test.b)
		if result != test.expected {
			t.Errorf("calcSub(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestCalcMultiply(t *testing.T) {
	tests := []struct {
		a, b, expected int
	}{
		{2, 3, 6},
		{-2, -3, 6},
		{0, 5, 0},
		{-1, 5, -5},
		{10, 10, 100},
	}

	for _, test := range tests {
		result := calcMultiply(test.a, test.b)
		if result != test.expected {
			t.Errorf("calcMultiply(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
	}
}

func TestCalcDivide(t *testing.T) {
	tests := []struct {
		a, b        int
		expected    int
		success     bool
		expectedErr string
	}{
		{6, 3, 2, true, ""},
		{10, 2, 5, true, ""},
		{5, 0, 0, false, "Can't divide by Zero !"},
		{-10, 2, -5, true, ""},
	}

	for _, test := range tests {
		result, success, errMsg := calcDivide(test.a, test.b)
		if success != test.success {
			t.Errorf("calcDivide(%d, %d) success = %v; want %v", test.a, test.b, success, test.success)
		}
		if success && result != test.expected {
			t.Errorf("calcDivide(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
		}
		if !success && errMsg != test.expectedErr {
			t.Errorf("calcDivide(%d, %d) error = %s; want %s", test.a, test.b, errMsg, test.expectedErr)
		}
	}
}
