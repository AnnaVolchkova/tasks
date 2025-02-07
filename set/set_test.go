package set

import (
	"testing"
	"github.com/stretchr/testify/require"
)

func TestPutUnique(t *testing.T) {
	s := New()
	
	values := []string{"key1", "key2", "key3"}

	for _, value := range values{
		s.Put(value)
	}

	for _, value := range values{
		result := s.Exists(value)

		require.True(t, result)
	}

	require.Equal(t, len(values), s.Length())
}

func TestDuplicates (t *testing.T) {
	s := New()

	values := []string{"key1", "key2", "key2"}

	for _, value := range values{
		s.Put(value)
	}

	require.True(t, s.Exists("key1"))
	require.True(t, s.Exists("key2"))

	require.Equal(t, 2, s.Length())
}

func TestZeroLength (t *testing.T) {
	s := New()

	require.Zero(t, s.Length())
}

func TestDelete (t *testing.T) {
	s := New()

	values := []string{"key1", "key2", "key3"}

	for _, value := range values{
		s.Put(value)
	}

	s.Delete("key3")

	require.False(t, s.Exists("key3"))

	require.Equal(t, 2, s.Length())
}