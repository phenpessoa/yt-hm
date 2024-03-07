package statichm

import (
	"ythm"
)

type StaticHashmap[K ythm.Hasher, V any] struct {
	len     int
	buckets []*item[K, V]
}

type item[K ythm.Hasher, V any] struct {
	key  K
	val  V
	next *item[K, V]
}

func (h *StaticHashmap[K, V]) getBucketIdx(k K) int {
	hash := k.Hash()
	return int(hash) & (len(h.buckets) - 1)
}

func (h *StaticHashmap[K, V]) Insert(k K, v V) {
	b := h.getBucketIdx(k)
	cur := h.buckets[b]
	if cur == nil {
		h.buckets[b] = &item[K, V]{key: k, val: v}
		h.len++
		return
	}

	prev := h.buckets[b]
	for ; cur != nil; cur = cur.next {
		if cur.key == k {
			cur.val = v
			return
		}
		prev = cur
	}
	prev.next = &item[K, V]{key: k, val: v}
	h.len++
}

func (h *StaticHashmap[K, V]) Get(k K) (V, bool) {
	var zero V

	b := h.getBucketIdx(k)

	for cur := h.buckets[b]; cur != nil; cur = cur.next {
		if cur.key == k {
			return cur.val, true
		}
	}

	return zero, false
}

func (h *StaticHashmap[K, V]) Delete(k K) {
	b := h.getBucketIdx(k)
	cur := h.buckets[b]

	if cur == nil {
		return
	}

	var prev *item[K, V]
	for ; cur != nil; cur = cur.next {
		if cur.key != k {
			prev = cur
			continue
		}

		if prev == nil {
			h.buckets[b] = cur.next
		} else {
			prev.next = cur.next
		}

		h.len--
		return
	}
}

func (h *StaticHashmap[K, V]) Len() int { return h.len }

func NewHM[K ythm.Hasher, V any]() *StaticHashmap[K, V] {
	return &StaticHashmap[K, V]{
		len:     0,
		buckets: make([]*item[K, V], 1024),
	}
}
