package internal

import (
	"fmt"
	"reflect"
)

var kvBase KeyValue

func ShowValueOfKey(key string) bool {
	valuesArray, check := CheckKeyExists(key)
	if !check {
		return false
	}
	fmt.Println("The attributes present for this key", key, " are : ")
	for _, val := range valuesArray {
		fmt.Println("Attribute Key : ", val.attrKey)
		fmt.Println("Attribute Value : ", val.attrValue)
	}
	return true
}

//check if key exists within the map
func CheckKeyExists(key string) ([]valueAttribute, bool) {
	kvMap := kvBase.GetKVMap()
	if val, ok := kvMap[key]; ok {
		return val, true
	}
	return nil, false
}

//check if attributetype exists in attributeTyep map
func CheckAttributeTypeExists(attrKey string) (reflect.Kind, bool) {
	kvMap := kvBase.GetAttrTypeMap()
	if val, ok := kvMap[attrKey]; ok {
		return val, true
	}
	return 0, false
}

func DeleteKeyFromMap(key string) {
	kvMap := kvBase.GetKVMap()
	if _, ok := kvMap[key]; ok {
		delete(kvMap, key)
	}
}
