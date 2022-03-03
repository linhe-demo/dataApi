package tools

import (
	"math/rand"
	"time"
)

//生成若干个不重复的随机数
func GenerateRandomNumber(start int, end int, count int) (nums map[int]int) {
	//范围检查
	if end < start || (end-start) < count {
		return nil
	}
	nums = make(map[int]int)
	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for len(nums) < count {
		//生成随机数
		num := r.Intn((end - start)) + start
		//查重
		if _, ok := nums[num]; !ok {
			nums[num] = num
		}
	}
	return nums
}

//生成一个随机数
func GenerateRandomSingleNumber(start int, end int, count int) (num int) {
	//范围检查
	if end < start || (end-start) < count {
		return num
	}

	//随机数生成器，加入时间戳保证每次生成的随机数不一样
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	num = r.Intn((end - start)) + start
	return num
}
