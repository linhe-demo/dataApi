package bloomfilter

import (
	"github.com/bits-and-blooms/bitset"
)

const DefaultHashArrSize = 16 // 默认hash数组大小

var seeds = []uint{9, 13, 33, 43, 53, 63} // hash 函数计算方式种子

type BloomFilter struct {
	set     *bitset.BitSet                        //第三方库
	hashFun [6]func(seed uint, value string) uint //hash 函数
}

func NewBloomFilter() *BloomFilter {
	bf := new(BloomFilter)
	bf.set = bitset.New(DefaultHashArrSize)

	for i := 0; i < len(bf.hashFun); i++ {
		bf.hashFun[i] = createHash()
	}

	return bf
}

func createHash() func(seed uint, value string) uint {
	return func(seed uint, value string) uint {
		var result uint = 0
		for i := 0; i < len(value); i++ {
			result = result*seed + uint(value[i])
		}
		return result & (DefaultHashArrSize - 1)
	}
}

func (b *BloomFilter) Add(value string) {
	for i, f := range b.hashFun {
		b.set.Set(f(seeds[i], value))
	}
}

func (b *BloomFilter) CheckExist(value string) bool {
	for i, f := range b.hashFun {
		if !b.set.Test(f(seeds[i], value)) {
			return false
		}
	}
	return true
}
