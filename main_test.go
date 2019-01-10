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