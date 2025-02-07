package linkedlist

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestLengthEmpty(t *testing.T) {
	ll := New()
	require.Equal(t, 0, ll.Length())
}

func TestLenghtFull(t *testing.T) {
	ll := New()
	ll.Push("1")
	ll.Push("2")

	require.Equal(t, 2, ll.Length())
}

func TestPopEmpty(t *testing.T) {
	ll := New()

	lastElement, exists := ll.Pop()
	require.Zero(t, lastElement)
	require.False(t, exists)
}

func TestPopFull(t *testing.T) {
	ll := New()

	ll.Push("1")
	ll.Push("2")

	lastElement, exists := ll.Pop()
	require.Equal(t, "2", lastElement)
	require.True(t, exists)
}

func TestPushHeadEmpty(t *testing.T) {
	ll := New()

	ll.PushHead("1")
	require.Equal(t, 1, ll.Length())

	lastElement, exists := ll.Pop()
	require.Equal(t, "1", lastElement)
	require.True(t, exists)
}

func TestPushHeadFull(t *testing.T) {
	ll := New()

	ll.Push("1")
	ll.Push("2")
	ll.Push("3")

	ll.PushHead("5")
	require.Equal(t, 4, ll.Length())

	lastElement, exists := ll.Pop()
	require.Equal(t, "3", lastElement)
	require.True(t, exists)
}
// // добавить элемент через pushHead
// // проверить, что длина равна 1, если до этого был пустой список
// // сделать pop и убедиться, что вернулся тот элемент, который мы добавили, и что он там есть