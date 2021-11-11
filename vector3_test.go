package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
