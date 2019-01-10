package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	main()
}

func TestMain_Sum(t *testing.T) {
	result := sum(1, 1);
	assert.Equal(t, 2, result)
}

func TestMain_Subtract(t *testing.T) {
	result := subtract(1, 1);
	assert.Equal(t, 0, result)
}

func TestMain_Divide(t *testing.T) {
	result := divide(1, 1);
	assert.Equal(t, 1, result)
}

func TestMain_Print(t *testing.T) {
	print("oi");
}

func TestMain_Print2(t *testing.T) {
	print2("oi");
}