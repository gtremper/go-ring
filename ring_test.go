package ring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicRing(t *testing.T) {
	assert := assert.New(t)
	var r IntRing

	for i := 0; i < 3; i++ {
		r.Add(i)
	}
	for i := 0; i < 3; i++ {
		assert.Equal(i, r.Pop())
	}

	for i := 0; i < 10; i++ {
		r.Add(i)
	}
	for i := 0; i < 6; i++ {
		assert.Equal(i, r.Pop())
	}
	for i := 10; i < 25; i++ {
		r.Add(i)
	}
	for i := 6; i < 25; i++ {
		assert.Equal(i, r.Pop())
	}
}
