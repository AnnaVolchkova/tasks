package binarytree

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestBinarytreeEmpty(t *testing.T) {
	tree := New()
	require.Equal(t, 0, tree.ElementsNumber())
}

func TestBinarytreeFull(t *testing.T) {
	tree := New()

	tree.Insert(15)
	tree.Insert(30)
	tree.Insert(45)
	tree.Insert(60)
	tree.Insert(75)
	require.Equal(t, 5, tree.ElementsNumber())
}