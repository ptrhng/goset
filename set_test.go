package goset_test

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ptrhng/goset"
)

func TestSetNew(t *testing.T) {
	set := goset.New[int]()
	assert.True(t, set.IsEmpty())
}

func TestSetFrom(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		set := goset.From[int](nil)
		assert.Zero(t, set.Size())
		assert.True(t, set.IsEmpty())
	})

	t.Run("some", func(t *testing.T) {
		set := goset.From([]string{"Apple", "Banana"})
		assert.Equal(t, set.Size(), 2)
		assert.True(t, set.Contains("Apple"))
		assert.True(t, set.Contains("Banana"))
	})

	t.Run("duplicate", func(t *testing.T) {
		set := goset.From([]string{"Apple", "Banana", "Apple"})
		assert.Equal(t, set.Size(), 2)
		assert.True(t, set.Contains("Apple"))
		assert.True(t, set.Contains("Banana"))
	})
}

func TestSetContains(t *testing.T) {
	set := goset.From[string]([]string{"Apple", "Banana"})
	assert.True(t, set.Contains("Apple"))
	assert.True(t, set.Contains("Banana"))
	assert.False(t, set.Contains("Mango"))
}

func TestSetAdd(t *testing.T) {
	t.Run("one", func(t *testing.T) {
		set := goset.New[string]()
		assert.True(t, set.Add("Apple"))
		assert.Equal(t, set.Size(), 1)
		assert.True(t, set.Contains("Apple"))
	})

	t.Run("some", func(t *testing.T) {
		set := goset.New[string]()
		assert.True(t, set.Add("Apple"))
		assert.True(t, set.Add("Banana"))
		assert.Equal(t, set.Size(), 2)
		assert.True(t, set.Contains("Apple"))
		assert.True(t, set.Contains("Banana"))
	})

	t.Run("re-add", func(t *testing.T) {
		set := goset.New[string]()
		assert.True(t, set.Add("Apple"))
		assert.False(t, set.Add("Apple"))
	})
}

func TestSetRemove(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		set := goset.New[string]()
		assert.False(t, set.Remove("Apple"))
		assert.True(t, set.IsEmpty())
	})

	t.Run("one", func(t *testing.T) {
		set := goset.From[string]([]string{"Apple"})
		assert.True(t, set.Remove("Apple"))
		assert.True(t, set.IsEmpty())
	})

	t.Run("not present", func(t *testing.T) {
		set := goset.From[string]([]string{"Apple"})
		assert.False(t, set.Remove("Banana"))
		assert.Equal(t, set.Size(), 1)
	})

	t.Run("re-remove", func(t *testing.T) {
		set := goset.From[string]([]string{"Apple"})
		assert.True(t, set.Remove("Apple"))
		assert.False(t, set.Remove("Apple"))
		assert.True(t, set.IsEmpty())
	})
}

func TestSetClear(t *testing.T) {
	tests := []struct {
		name  string
		input []string
	}{
		{
			name:  "empty",
			input: []string{},
		},
		{
			name:  "non-empty",
			input: []string{"Apple", "Banana"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := goset.From(tt.input)
			set.Clear()
			assert.True(t, set.IsEmpty())
		})
	}
}

func TestSetRange(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		set := goset.From[string]([]string{"Apple", "Banana"})
		visited := make([]string, 0, 2)
		set.Range(func(item string) bool {
			visited = append(visited, item)
			return false
		})
		sort.Strings(visited)

		assert.Equal(t, []string{"Apple", "Banana"}, visited)
	})

	t.Run("stop", func(t *testing.T) {
		set := goset.From[string]([]string{"Apple", "Banana"})
		var visited []string
		set.Range(func(item string) bool {
			visited = append(visited, item)
			return true
		})

		assert.Equal(t, 1, len(visited))
		assert.True(t, visited[0] == "Apple" || visited[0] == "Banana")
	})

	t.Run("empty", func(t *testing.T) {
		set := goset.New[string]()
		count := 0
		set.Range(func(item string) bool {
			count++
			return false
		})

		assert.Zero(t, count)
	})
}

func TestSetSlice(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{
			name:  "empty",
			input: []string{},
			want:  []string{},
		},
		{
			name:  "some",
			input: []string{"Apple", "Banana"},
			want:  []string{"Apple", "Banana"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			set := goset.From(tt.input)
			got := set.Slice()
			sort.Strings(got)
			assert.Equal(t, tt.want, got)
		})
	}
}
