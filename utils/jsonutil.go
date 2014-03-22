package utils

import (
	"encoding/json"
	"fmt"
)

func JSONstringToMap(str string) map[string]interface{} {
	var f interface{}
	err := json.Unmarshal([]byte(str), &f)
	if(err != nil){
	    fmt.Printf("%s\n", "转换JSON出错！");
	}
	m := f.(map[string]interface{})
	return m
}

func MapToJSONstring(m map[string]interface{}) string {
	maps, _ := json.Marshal(m)
	return string(maps)
}

