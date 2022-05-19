package dongle

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecode_String(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "hello world"},
	}

	for index, test := range tests {
		d := newDecode()
		d.dst = []byte(test.input)
		assert.Nil(d.Error)
		assert.Equal(test.expected, fmt.Sprintf("%s", d), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ToString(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected string // 期望值
	}{
		{"", ""},
		{"hello world", "hello world"},
	}

	for index, test := range tests {
		d := newDecode()
		d.dst = []byte(test.input)
		assert.Nil(d.Error)
		assert.Equal(test.expected, d.ToString(), "Current test index is "+strconv.Itoa(index))
	}
}

func TestDecode_ToBytes(t *testing.T) {
	assert := assert.New(t)

	tests := []struct {
		input    string // 输入值
		expected []byte // 期望值
	}{
		{"", []byte("")},
		{"hello world", []byte("hello world")},
	}

	for index, test := range tests {
		d := newDecode()
		d.dst = []byte(test.input)
		assert.Nil(d.Error)
		assert.Equal(test.expected, d.ToBytes(), "Current test index is "+strconv.Itoa(index))
	}
}
