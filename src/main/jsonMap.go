package main

import (
	"encoding/json"
	"strings"
	//	"github.com/bitly/go-simplejson" // for json get
)

//把两层嵌套结构的json格式的数据组转成map(map中不含interface结构)
func NoInterfaceJsonToMap(input string) (map[string]map[string]interface{}, error) {
	result := make(map[string]map[string]interface{})
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
func MapToJson(input map[string]interface{}) (string, error) {
	result, err := json.Marshal(input)
	if err != nil {
		//		panic(err)
		return "", err
	}
	return string(result), nil
}

func MapMapToJson(input map[string]map[string]interface{}) (string, error) {
	result, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func JsonToMap(input string) (map[string]interface{}, error) {
	result := make(map[string]interface{})
	err := json.Unmarshal([]byte(input), &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func BoltKeyValueToJson(key, value string, delimeter string) (string, error) {
	keys := []string{key}
	values := []string{value}
	return BoltKeyValuesToJson(keys, values, delimeter)
}
func BoltKeyValuesToJson(keys, values []string, delimeter string) (string, error) {
	mapResult := make(map[string]interface{})
	for i := range keys {
		key := strings.Split(keys[i], delimeter)
		value := values[i]
		cur := mapResult
		for j := range key {
			if j == len(key)-1 {
			} else if j == len(key)-2 {
				if cur[key[j]] == nil {
					cur[key[j]] = map[string]string{key[len(key)-1]: value}
				} else {
					cur[key[j]].(map[string]string)[key[len(key)-1]] = value
				}
			} else {
				if cur[key[j]] == nil {
					cur[key[j]] = make(map[string]interface{})
				}
				cur = cur[key[j]].(map[string]interface{})
			}
		}
	}
	return MapToJson(mapResult)
}
