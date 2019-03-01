package httprouter

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

type Request struct {
	Method        string
	Header        http.Header
	URL           *url.URL
	ContentLength int64
	Host          string
	RequestURI    string
	RemoteAddr    string
	Data          map[string][]string
}

func (self *Request) Get(name string) []string {
	return self.Data[name]
}

func (self *Request) GetString(name string) string {
	strs := self.Get(name)
	if len(strs) > 0 {
		return strs[0]
	}
	return ""
}

func (self *Request) GetInt(name string) (int, error) {
	strs := self.Get(name)
	if len(strs) == 0 {
		return 0, fmt.Errorf("")
	}
	if len(strs) > 0 {
		return strconv.Atoi(strs[0])
	}

	return 0, nil
}

func (self *Request) GetIntSlice(name string) (ins []int, err error) {
	strs := self.Get(name)

	if len(strs) > 0 {
		ins = make([]int, len(strs))
		var val int
		for _, str := range strs {
			val, err = strconv.Atoi(str)
			if err != nil {
				return nil, err
			}
			ins = append(ins, val)
		}
	}

	return ins, err
}

func (self *Request) GetFloat32(name string) (float32, error) {
	strs := self.Get(name)
	if len(strs) > 0 {
		f, err := strconv.ParseFloat(strs[0], 32)

		if err != nil {
			return 0, err
		}
		return float32(f), nil
	}

	return 0, nil
}

//func (self *Request) GetFloat32Slice(name string) (fs []float32, err error) {
//	strs := self.Get(name)
//	if len(strs) > 0 {
//
//	}
//}

//func (self *Request) GetInt(name string) (int, error) {
//	if self.Data != nil {
//		if val, ok := self.Data[name]; ok {
//			if intVal, ok := val.(int); ok {
//				return intVal, nil
//			}
//			return 0, fmt.Errorf("cannot convert %s value to int", name)
//		}
//		return 0, fmt.Errorf("%s is not in request data", name)
//	}
//	return 0, errors.New("request data is not inited")
//}
//
//func (self *Request) GetFloat32(name string) (float32, error) {
//	if self.Data != nil {
//		if val, ok := self.Data[name]; ok {
//			if fVal, ok := val.(float32); ok {
//				return fVal, nil
//			}
//			return 0, fmt.Errorf("cannot convert %s value to float32", name)
//		}
//		return 0, fmt.Errorf("%s is not in request data", name)
//	}
//	return 0, errors.New("request data is not inited")
//}
//
//func (self *Request) GetFloat64(name string) (float64, error) {
//	if self.Data != nil {
//		if val, ok := self.Data[name]; ok {
//			if fVal, ok := val.(float64); ok {
//				return fVal, nil
//			}
//			return 0, fmt.Errorf("cannot convert %s value to float64", name)
//		}
//		return 0, fmt.Errorf("%s is not in request data", name)
//	}
//	return 0, errors.New("request data is not inited")
//}
//
//func (self *Request) GetString(name string) (string, error) {
//	if self.Data != nil {
//		if val, ok := self.Data[name]; ok {
//			if fVal, ok := val.(string); ok {
//				return fVal, nil
//			}
//			return "", fmt.Errorf("cannot convert %s value to string", name)
//		}
//		return "", fmt.Errorf("%s is not in request data", name)
//	}
//	return "", errors.New("request data is not inited")
//}
