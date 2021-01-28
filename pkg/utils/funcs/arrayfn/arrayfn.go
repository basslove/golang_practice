package arrayfn

import (
	"reflect"
)

// 配列に対する関数tips
// ** sub routine作ったらシバく

func IsIncluded(target interface{}, slice interface{}) bool{
	valueOfSlice := reflect.ValueOf(slice)

	for i := 0; i < valueOfSlice.Len(); i++ {
		switch target.(type) {
		case int:
			if target.(int) == valueOfSlice.Index(i).Interface().(int) {
				return true
			}
		case string:
			if target.(string) == valueOfSlice.Index(i).Interface().(string) {
				return true
			}
		}
	}

	return false
}

func IsUniq(slice interface{}) bool {
	valueOfSlice := reflect.ValueOf(slice)
	if valueOfSlice.Len() == 0 {
		return true
	}

	switch slice.(type) {
	case []string:
		encountered := map[string]bool{}
		for i := 0; i < valueOfSlice.Len(); i++ {
			value := valueOfSlice.Index(i).Interface().(string)
			if !encountered[value] {
				encountered[value] = true
			} else {
				return false
			}
		}
		return true
	case []int:
		encountered := map[int]bool{}
		for i := 0; i < valueOfSlice.Len(); i++ {
			value := valueOfSlice.Index(i).Interface().(int)
			if !encountered[value] {
				encountered[value] = true
			} else {
				return false
			}
		}
		return true
	default:
		return false
	}
}

func NumUniq(slice []int) []int {
	newSlice := make([]int, 0, len(slice))
	encountered := map[int]bool{}

	for i := 0; i < len(slice); i++ {
		if !encountered[slice[i]] {
			encountered[slice[i]] = true
			newSlice = append(newSlice, slice[i])
		}
	}

	return newSlice
}

func StrUniq(slice []string) []string {
	newSlice := make([]string, 0, len(slice))
	encountered := map[string]bool{}

	for i := 0; i < len(slice); i++ {
		if !encountered[slice[i]] {
			encountered[slice[i]] = true
			newSlice = append(newSlice, slice[i])
		}
	}

	return newSlice
}