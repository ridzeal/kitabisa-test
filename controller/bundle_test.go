package controller

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestBundlePositiveSmall(t *testing.T) {
	total, apple, cake := BundleIt(1, 1)

	assert.Equal(t, 1, total, "Total should be 1")
	assert.Equal(t, 1, apple, "Apple should be 1")
	assert.Equal(t, 1, cake, "Cake should be 1")
}

func TestBundlePositiveBigDiff(t *testing.T) {
	total, apple, cake := BundleIt(25, 23)

	assert.Equal(t, 1, total, "Total should be 1")
	assert.Equal(t, 25, apple, "Apple should be 25")
	assert.Equal(t, 23, cake, "Cake should be 23")
}

func TestBundleNegativeZero(t *testing.T) {
	total, apple, cake := BundleIt(0, 0)

	assert.Equal(t, 0, total, "Total should be 0")
	assert.Equal(t, 0, apple, "Apple should be 0")
	assert.Equal(t, 0, cake, "Cake should be 0")
}

func TestGCD(t *testing.T) {
	res1 := GCD(20,25)
	res2 := GCD(23,25)

	assert.Equal(t, 5, res1, "Res1 should be 5")
	assert.Equal(t, 1, res2, "Res2 should be 1")
}