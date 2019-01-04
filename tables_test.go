package tables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLine(t *testing.T) {
	a := New()
	a.AddLine("data1", "data2")
	s := a.String()
	assert.Equal(t, s, "data1\tdata2\n", "strings should be equal")
}
