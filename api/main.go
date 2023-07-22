package main

import (
	"fmt"
	"github.com/MyOrg/KeyValueStore/internal"
)

var KeyValueBase internal.KeyValue

func main() {
	KeyValueBase.Initialize()
	KeyValueBase.AddValueToKV("delhi", "pollution_level", "high")

	// this will throw error as previously this ttr_key was set as string and integer is being passed now
	KeyValueBase.AddValueToKV("delhi", "pollution_level", 7)

	KeyValueBase.AddValueToKV("jakarta", "latitude", -6)
	KeyValueBase.AddValueToKV("bangalore", "latitude", 12.94)
	KeyValueBase.AddValueToKV("bangalore", "longitude", 77.64)
	KeyValueBase.AddValueToKV("bangalore", "pollution_level", "high")

	fmt.Println(KeyValueBase.GetKVMap())
}
