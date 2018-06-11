package utils

import (
	"reflect"
	"encoding/json"
)

const (
	VIAJSON = "json"
	VIAREFLECT = "reflect"
)

func StructToMap(obj interface{},toType string) map[string]interface{} {
	data := make(map[string]interface{})
	switch toType {
	case VIAREFLECT:
		elem := reflect.ValueOf(&obj).Elem()
        relType := elem.Type()
        for i := 0; i < relType.NumField(); i++ {
            data[relType.Field(i).Name] = elem.Field(i).Interface()
        }        
	case VIAJSON:
	default:
		res, _ := json.Marshal(obj)
		json.Unmarshal(res, &data)
	}
    
    return data
}

func StructToMap2(obj interface{},) map[string]interface{} {
	data := make(map[string]interface{})
	res, _ := json.Marshal(obj)
	json.Unmarshal(res, &data)
	
    
    return data
}