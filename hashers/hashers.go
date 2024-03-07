package hashers

type StringHasher string

func (s StringHasher) Hash() uint64 {
	var h uint64
	for i := range s {
		h *= 1099511628211
		h ^= uint64(s[i])
	}
	return h
}
