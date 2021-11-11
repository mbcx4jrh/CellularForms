package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var delta float64 = 0.00001

func TestZero(t *testing.T) {
	v1 := Zero()

	assert.Equal(t, Vector3{0, 0, 0}, v1)
}

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	v1 := Vector3{1.0, 3.2, 1.2}
	v2 := Vector3{1.1, 1.3, 5.0}

	v1.Add(&v2)
	assert.Equal(v1.x, 2.1)
	assert.Equal(v1.y, 4.5)
	assert.Equal(v1.z, 6.2)

	//check v2 hasnt changed
	assert.Equal(v2.x, 1.1)
}

func TestSubtract(t *testing.T) {
	assert := assert.New(t)

	v1 := Vector3{1.0, 3.2, 1.2}
	v2 := Vector3{1.1, 1.3, 5.0}

	v1.Subtract(&v2)
	assert.InDelta(v1.x, -0.1, delta)
	assert.InDelta(v1.y, 1.9, delta)
	assert.InDelta(v1.z, -3.8, delta)

	//check v2 hasnt changed
	assert.Equal(v2.x, 1.1)

}

func TestMultiply(t *testing.T) {
	assert := assert.New(t)

	v1 := Vector3{1.0, 1.0, 1.0}

	v1.Multiply(5.0)

	assert.Equal(Vector3{5.0, 5.0, 5.0}, v1)
}

func TestSqrMagnitude(t *testing.T) {
	v1 := Vector3{3.0, 3.0, 4.0}
	expected := 3.0*3.0 + 3.0*3.0 + 4.0*4.0
	sqrMag := SqrMagnitude(&v1)
	assert.Equal(t, expected, sqrMag)
}
