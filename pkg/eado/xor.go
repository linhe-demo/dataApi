package eado

func Xor(str []byte, key []byte) []byte {
	for _, kv := range key {
		str = func(b []byte, k byte) (ret []byte) {
			for _, bv := range b {
				ret = append(ret, k^bv)
			}
			return ret
		}(str, kv)
	}
	return str
}
