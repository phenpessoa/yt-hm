package tests

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"

	"ythm"
	"ythm/dynamichm"
	"ythm/hashers"
	"ythm/statichm"
)

func TestStaticHashmap(t *testing.T) {
	m := statichm.NewHM[hashers.StringHasher, string]()
	testHashmapImpl(t, m)
}

func TestDynamicHashmap(t *testing.T) {
	m := dynamichm.NewHM[hashers.StringHasher, string]()
	testHashmapImpl(t, m)
}

func testHashmapImpl(
	t *testing.T,
	m ythm.Hashmap[hashers.StringHasher, string],
) {
	assert := assert.New(t)
	assert.Zero(m.Len())

	m.Insert("pedro", "pessoa")
	res, ok := m.Get("pedro")
	assert.True(ok)
	assert.Equal("pessoa", res)
	assert.Equal(1, m.Len())

	res, ok = m.Get("foo")
	assert.False(ok)
	assert.Zero(res)

	m.Delete("pedro")
	res, ok = m.Get("pedro")
	assert.False(ok)
	assert.Zero(res)
	assert.Zero(m.Len())

	m.Insert("foo", "foo")
	m.Insert("bar", "bar")
	m.Insert("baz", "baz")
	m.Insert("qux", "qux")
	assert.Equal(4, m.Len())

	m.Delete("baz")
	assert.Equal(3, m.Len())

	res, ok = m.Get("foo")
	assert.True(ok)
	assert.Equal("foo", res)

	res, ok = m.Get("bar")
	assert.True(ok)
	assert.Equal("bar", res)

	res, ok = m.Get("baz")
	assert.False(ok)
	assert.Zero(res)

	res, ok = m.Get("qux")
	assert.True(ok)
	assert.Equal("qux", res)

	m.Delete("foo")
	assert.Equal(2, m.Len())
	res, ok = m.Get("foo")
	assert.False(ok)
	assert.Zero(res)

	m.Delete("bar")
	assert.Equal(1, m.Len())
	res, ok = m.Get("bar")
	assert.False(ok)
	assert.Zero(res)

	m.Delete("qux")
	assert.Zero(m.Len())
	res, ok = m.Get("qux")
	assert.False(ok)
	assert.Zero(res)

	const n = 1000

	for i := range n {
		str := strconv.FormatInt(int64(i), 10)
		m.Insert(hashers.StringHasher(str), str)
	}

	assert.Equal(n, m.Len())

	for i := range n {
		str := strconv.FormatInt(int64(i), 10)
		res, ok := m.Get(hashers.StringHasher(str))
		assert.True(ok)
		assert.Equal(str, res)
	}

	for i := range n {
		str := strconv.FormatInt(int64(i), 10)
		m.Delete(hashers.StringHasher(str))
		res, ok := m.Get(hashers.StringHasher(str))
		assert.False(ok)
		assert.Zero(res)
		assert.Equal(n-1-i, m.Len())
	}

	assert.Zero(m.Len())
}
