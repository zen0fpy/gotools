package prototype

import (
	"github.com/stretchr/testify/require"
	"testing"
)

var manager *Manger

type Type1 struct {
	name string
}

func (t *Type1) Clone() Clonable {
	tc := *t
	return &tc
}

type Type2 struct {
	name string
}

func (t *Type2) Clone() Clonable {
	tc := *t
	return &tc
}

func TestClone(t *testing.T) {
	t1 := manager.Get("t1")

	t2 := t1.Clone()
	require.NotEmpty(t, t2, t1)
}

func TestCloneFromManager(t *testing.T) {
	c := manager.Get("t1").Clone()

	t1 := c.(*Type1)
	require.Equal(t, "type1", t1.name)

}

func init() {
	manager = NewProtoManager()

	t1 := &Type1{
		name: "type1",
	}
	manager.Set("t1", t1)
}
