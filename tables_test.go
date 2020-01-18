package tables

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddLine(t *testing.T) {
	a := New()
	a.AddLine("data1", "data2")
	s := a.String()
	assert.Equal(t, s, "data1  data2\n", "strings should be equal")
}

func TestAddLineTwo(t *testing.T) {
	a := New()
	a.AddLine("data1", "data2")
	a.AddLine("data3", "data4")
	s := a.String()
	assert.Equal(t, s, "data1  data2\ndata3  data4\n", "strings should be equal")
}

func TestAddHeader(t *testing.T) {
	a := New()
	a.AddHeader("data1", "data2")
	s := a.String()
	assert.Equal(t, s, "data1  data2\n", "strings should be equal")
}

func TestAddHeaderWithLines(t *testing.T) {
	a := New()
	a.AddHeader("data1", "data2")
	a.AddLine("item1", "item2")
	s := a.String()
	assert.Equal(t, s, "data1  data2\nitem1  item2\n", "strings should be equal")
}

func TestAddHeaderWithHook(t *testing.T) {
	a := New()
	a.AddHeader("data1", "data2")
	a.AddHooks(func(s string) string {
		return "a  a"
	},
	)
	a.AddLine("item1", "item2")
	s := a.String()
	assert.Equal(t, s, "data1  data2\na  a", "strings should be equal")
}
