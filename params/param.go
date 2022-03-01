package params

import (
	"fmt"
	"github.com/pkg/errors"
	"net/http"
	"reflect"
	"strconv"
	"strings"
)

const (
	defaultMaxMemory = 32 << 20 // 32 MB
)

var data = make(map[string]interface{})

// Unpack 从 HTTP 请求 req 的参数中提取数据填充到 ptr 指向结构体的各个字段
func Unpack(req *http.Request, ptr interface{}) error {
	if err := req.ParseForm(); err != nil {
		return err
	}

	switch req.Method {
	case "GET":
		if err := analysisGet(req, ptr); err != nil {
			return errors.Wrap(err, "param fail")
		}
	case "POST":
		if req.Form != nil {
			if err := analysisGet(req, ptr); err != nil {
				return errors.Wrap(err, "param fail")
			}
		} else {

			if req.Form == nil {
				_ = req.ParseForm()
				_ = req.ParseMultipartForm(defaultMaxMemory)
			}

			for k, v := range req.Form {
				if len(v) > 0 {
					SetMData(k, v[0])
				}
			}
			fmt.Print(data)
			//_ = req.ParseForm()
			//_ = req.ParseMultipartForm(defaultMaxMemory)
			//if err := analysisGet(req, ptr); err != nil {
			//	return errors.Wrap(err, "param fail")
			//}
		}
		//fmt.Println(req.Form)
		//fmt.Println(req.PostForm)
		//fmt.Println(req.Body)
	default:
		return errors.Wrap(nil, "Sorry, only GET and POST methods are supported.")
	}

	return nil
}

func analysisGet(req *http.Request, ptr interface{}) error {
	// 创建字段映射表，键为有效名称
	fields := make(map[string]reflect.Value)
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		fieldInfo := v.Type().Field(i)
		tag := fieldInfo.Tag
		name := tag.Get("http")
		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		fields[name] = v.Field(i)
	}

	// 对请求中的每个参数更新结构体中对应的字段
	for name, values := range req.Form {
		f := fields[name]
		if !f.IsValid() {
			continue // 忽略不能识别的 HTTP 参数
		}

		for _, value := range values {
			if f.Kind() == reflect.Slice {
				elem := reflect.New(f.Type().Elem()).Elem()
				if err := populate(elem, value); err != nil {
					return errors.Wrapf(err, "param analysis failed %s", name)
				}
				f.Set(reflect.Append(f, elem))
			} else {
				if err := populate(f, value); err != nil {
					return errors.Wrapf(err, "param analysis failed %s", name)
				}
			}
		}
	}
	return nil
}

func populate(v reflect.Value, value string) error {
	switch v.Kind() {
	case reflect.String:
		v.SetString(value)
	case reflect.Int:
		i, err := strconv.ParseInt(value, 10, 64)
		if err != nil {
			return err
		}
		v.SetInt(i)
	case reflect.Bool:
		b, err := strconv.ParseBool(value)
		if err != nil {
			return err
		}
		v.SetBool(b)
	default:
		return fmt.Errorf("unsupported kind %s", v.Type())
	}
	return nil
}

func SetMData(name string, val interface{}) {
	data[name] = val
}
