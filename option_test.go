package fu

import (
	"gotest.tools/assert"
	"testing"
)

func Test_Select(t *testing.T) {

	type Option0 (int)
	type Option1 (int)
	type OptionS (string)

	a := []interface{}{OptionS("text"), Option0(0), Option1(1), "hello world", 1001}

	assert.Equal(t, Select(Option0(1), a), Option0(0))
	assert.Equal(t, Select(Option1(0), a), Option1(1))
	assert.Equal(t, Select(OptionS(""), a), OptionS("text"))
	assert.Equal(t, Select("", a), "hello world")
	assert.Equal(t, Select(0, a), 1001)
}
