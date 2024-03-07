package dynamichm

import (
	"fmt"
	"strings"

	"ythm"
)

type DynamicHashmap[K ythm.Hasher, V any] struct {
	len     int
	buckets []*item[K, V]
}

type item[K ythm.Hasher, V any] struct {
	key  K
	val  V
	next *item[K, V]
}

func (h *DynamicHashmap[K, V]) getBucketIdx(k K, _cap int) int {
	hash := k.Hash()
	return int(hash) & (_cap - 1)
}

func (h *DynamicHashmap[K, V]) resize() {
	threshold := int(float64(len(h.buckets)) * 0.85 * 6.5)
	if h.len < threshold {
		return
	}

	newCapacity := len(h.buckets) * 2
	newBuckets := make([]*item[K, V], newCapacity)

	for _, b := range h.buckets {
		if b == nil {
			continue
		}

		for cur := b; cur != nil; cur = cur.next {
			newIdx := h.getBucketIdx(b.key, newCapacity)
			newItem := &item[K, V]{key: b.key, val: b.val}
			cur := newBuckets[newIdx]
			if cur == nil {
				newBuckets[newIdx] = newItem
				return
			}

			for cur.next != nil {
				cur = cur.next
			}

			cur.next = newItem
		}
	}

	h.buckets = newBuckets
}

func (h *DynamicHashmap[K, V]) Insert(k K, v V) {
	h.resize()
	b := h.getBucketIdx(k, len(h.buckets))
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

func (h *DynamicHashmap[K, V]) Get(k K) (V, bool) {
	var zero V

	b := h.getBucketIdx(k, len(h.buckets))

	for cur := h.buckets[b]; cur != nil; cur = cur.next {
		if cur.key == k {
			return cur.val, true
		}
	}

	return zero, false
}

func (h *DynamicHashmap[K, V]) Delete(k K) {
	b := h.getBucketIdx(k, len(h.buckets))
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

func (h *DynamicHashmap[K, V]) Len() int { return h.len }

func (h *DynamicHashmap[K, V]) String() string {
	var output strings.Builder
	output.WriteByte('{')
	output.WriteByte('\n')
	for _, b := range h.buckets {
		if b == nil {
			continue
		}

		output.WriteString(fmt.Sprintf("\t%v: %v\n", b.key, b.val))

		for cur := b.next; cur != nil; cur = cur.next {
			output.WriteString(fmt.Sprintf("\t%v: %v\n", cur.key, cur.val))
		}
	}
	output.WriteByte('}')
	return output.String()
}

func NewHM[K ythm.Hasher, V any]() *DynamicHashmap[K, V] {
	return &DynamicHashmap[K, V]{
		len:     0,
		buckets: make([]*item[K, V], 2),
	}
}
