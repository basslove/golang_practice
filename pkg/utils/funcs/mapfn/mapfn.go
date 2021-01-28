package mapfn

import (
	"reflect"
)

// mapに対する関数tips
// ** sub routine作ったらシバく

func Keys(m interface{}) []string {
	mapKeys := reflect.ValueOf(m).MapKeys()
	keys := make([]string, 0, len(mapKeys))

	for _, valueKey := range mapKeys {
		keys = append(keys, valueKey.Interface().(string))
	}

	return keys
}

func KeyIsExisted(target string, m interface{}) bool{
	mapKeys := reflect.ValueOf(m).MapKeys()

	for _, valueKey := range mapKeys {
		if target == valueKey.Interface().(string) {
			return true
		}
	}

	return false
}