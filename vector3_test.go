package main

import (
	"math"
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
	mag := Magnitude(&v1)
	assert.Equal(t, expected, sqrMag)
	assert.InDelta(t, math.Sqrt(expected), mag, delta)
}

func TestDot(t *testing.T) {
	v1 := Vector3{1.0, 1.0, 1.0}
	v2 := Vector3{1.0, 2.0, 3.0}
	dot := Dot(&v1, &v2)
	assert.InDelta(t, 6.0, dot, delta)
}

func TestNormalise(t *testing.T) {
	v1 := Vector3{3.0, 0.0, 0.0}

	v1.Normalise()
	assert := assert.New(t)
	assert.InDelta(v1.x, 1.0, delta)
	assert.InDelta(v1.y, 0.0, delta)
	assert.InDelta(v1.z, 0.0, delta)
}

func TestCross(t *testing.T) {

	i := Vector3{1, 0, 0}
	j := Vector3{0, 1, 0}
	k := Vector3{0, 0, 1}
	ni := Vector3{-1, 0, 0}

	assert := assert.New(t)
	assert.Equal(k, Cross(&i, &j))
	assert.Equal(i, Cross(&j, &k))
	assert.Equal(j, Cross(&k, &i))
	assert.Equal(ni, Cross(&k, &j))
}

func TestNewSubtract(t *testing.T) {
	v1 := Vector3{0, 0, 0}
	v2 := Vector3{1, 2, 3}
	r := NewSubtract(&v1, &v2)

	assert.Equal(t, r, Vector3{-1, -2, -3})
}

func TestNegate(t *testing.T) {
	v1 := Vector3{1, 2, 3}
	v1.Negate()

	assert.Equal(t, Vector3{-1, -2, -3}, v1)
}
