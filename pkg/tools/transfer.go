package tools

import (
	"github.com/pkg/errors"
	"reflect"
	"strconv"
	"strings"
)

func StrToMap(s string) (out map[int]int) {
	out = make(map[int]int)
	tmpList := strings.Split(s, ",")
	for _, v := range tmpList {
		tmpid, _ := strconv.Atoi(v)
		out[tmpid] = tmpid
	}
	return out
}

func StrToInt(s string) (out []int) {
	tmpList := strings.Split(s, ",")
	for _, v := range tmpList {
		tmpid, _ := strconv.Atoi(v)
		out = append(out, tmpid)
	}
	return out
}

func IntSliceToStringSlice(target []int) (out []string) {
	for _, v := range target {
		tmp := strconv.Itoa(v)
		out = append(out, tmp)
	}
	return out
}

func StringToInt64(str string) (out []int64) {
	tmpList := strings.Split(str, ",")
	for _, v := range tmpList {
		tmpId, _ := strconv.ParseInt(v, 10, 64)
		out = append(out, tmpId)
	}
	return out
}

func IntSliceToStrSlice(data interface{}) (out []string, err error) {
	err = errors.New("unsupported type")
	switch reflect.TypeOf(data).Kind() {
	case reflect.Slice:
		err = nil
		s := reflect.ValueOf(data)
		for i := 0; i < s.Len(); i++ {
			var (
				tmpData int64
				val     = s.Index(i).Interface()
			)
			switch val.(type) {
			case int8:
				tmpData = int64(val.(int8))
			case int16:
				tmpData = int64(val.(int16))
			case int32:
				tmpData = int64(val.(int32))
			case int64:
				tmpData = val.(int64)
			case int:
				tmpData = int64(val.(int))
			}
			tmpStr := strconv.FormatInt(tmpData, 10)
			out = append(out, tmpStr)
		}
	}
	return out, err
}
