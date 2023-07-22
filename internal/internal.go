package internal

import (
	"fmt"
	"reflect"
)

type KeyValue struct{}

type valueAttribute struct {
	attrKey   string
	attrValue interface{}
}

var (

	//This map holds the key and all attribute values associated with it
	keyValueStore map[string][]valueAttribute

	//This map holds the attribute key and its datatype
	attributeType map[string]reflect.Kind

	//This map holds the attribute structure and keys array associated with it
	attributeKeyMap map[valueAttribute][]string
)

func (kv *KeyValue) Initialize() {
	kv1 := make(map[string][]valueAttribute)
	kv2 := make(map[valueAttribute][]string)
	kv3 := make(map[string]reflect.Kind)

	keyValueStore = kv1
	attributeKeyMap = kv2
	attributeType = kv3
	fmt.Println("Starting key value operations...")
}

func (kv *KeyValue) AddValueToKV(key string, attributeKey string, attributeValue interface{}) {
	val := valueAttribute{
		attrKey:   attributeKey,
		attrValue: attributeValue,
	}
	attrType, check := CheckAttributeTypeExists(key)

	if !check {
		switch reflect.TypeOf(attributeValue).Kind() {
		case reflect.Int, reflect.String, reflect.Bool, reflect.Float64:
			attributeType[attributeKey] = reflect.TypeOf(attributeValue).Kind()
		default:
			fmt.Println("Type of attribute value is not a valid one")
			return
		}
	} else {
		attributeValueType := reflect.TypeOf(attributeValue).Kind()
		if attrType != attributeValueType {
			fmt.Println("Type of attribute value is different")
			return
		}
	}

	keyValueStore[key] = append(keyValueStore[key], val)
	attributeKeyMap[val] = append(attributeKeyMap[val], key)
}

func (kv *KeyValue) GetKVMap() map[string][]valueAttribute {
	return keyValueStore
}

func (kv *KeyValue) GetAttrTypeMap() map[string]reflect.Kind {
	return attributeType
}

func (kv *KeyValue) FetchValueForKey(key string) {
	if !ShowValueOfKey(key) {
		fmt.Println("The following key doesn't exist in map")
	}
}

func (kv *KeyValue) DeleteKeyValue(key string) {
	_, check := CheckKeyExists(key)
	if !check {
		fmt.Println("The following key doesn't exist in map")
		return
	}
	DeleteKeyFromMap(key)
}

func (kv *KeyValue) StretchKeyValue(attribute valueAttribute) {
	if val, ok := attributeKeyMap[attribute]; ok {
		fmt.Println("The keys associated with this query are : ")
		for i, key := range val {
			fmt.Println("Key", i, " : ", key)
		}
	} else {
		fmt.Println("There is no key matching the following query")
	}
	return
}
