package ythm

type Hasher interface {
	comparable
	Hash() uint64
}

type Hashmap[K Hasher, V any] interface {
	Insert(K, V)
	Get(K) (V, bool)
	Delete(K)
	Len() int
}
