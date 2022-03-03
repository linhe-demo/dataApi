package tools

import (
	"dataApi/app"
	"reflect"
	"strings"
)

func GetIndexSlice(target int, slice []int) (index int) {
	position := -1
	if len(slice) > 0 {
		num := app.FirstIndex
		for _, v := range slice {
			if v == target {
				position = num
			}
			num++
		}
	}
	return position
}

func InArray(val interface{}, array interface{}) (exists bool, index int) {
	exists = false
	index = -1
	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)
		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) == true {
				index = i
				exists = true
				return
			}
		}
	}
	return
}

func ArrayUnique(a interface{}) (ret []interface{}) {
	va := reflect.ValueOf(a)
	for i := 0; i < va.Len(); i++ {
		if i > 0 && reflect.DeepEqual(va.Index(i-1).Interface(), va.Index(i).Interface()) {
			continue
		}
		ret = append(ret, va.Index(i).Interface())
	}
	return ret
}

func RemoveDuplicate(list []int) (out []int) {
	for _, i := range list {
		if len(out) == 0 {
			out = append(out, i)
		} else {
			for k, v := range out {
				if i == v {
					break
				}
				if k == len(out)-1 {
					out = append(out, i)
				}
			}
		}
	}
	return out
}

// IntersectionInt 获取两个slice的交集
func IntersectionInt(slice1 []int, slice2 []int) (outSlice []int) {
	if len(slice1) == 0 || len(slice2) == 0 {
		return outSlice
	}
	for _, li1 := range slice1 {
		for _, li2 := range slice2 {
			if li1 == li2 {
				outSlice = append(outSlice, li1)
				continue
			}
		}
	}
	return outSlice
}

func InArrayStr(ele string, target []string, ignoreCase bool) bool {
	if ignoreCase {
		ele = strings.ToLower(ele)
	}
	for _, v := range target {
		if ignoreCase {
			v = strings.ToLower(v)
		}
		if v == ele {
			return true
		}
	}
	return false
}

func InArrayInt64(ele int64, target []int64) bool {
	for _, v := range target {
		if v == ele {
			return true
		}
	}
	return false
}

func InArrayInt(ele int, target []int) bool {
	for _, v := range target {
		if v == ele {
			return true
		}
	}
	return false
}
