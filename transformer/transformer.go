package transformer

import (
	"log"
	"reflect"
)

type Transformed struct {
	From interface{}
}

func autoConf(to interface{}, copies map[string][]interface{}) {
	b := reflect.ValueOf(to).Elem()
	for k, v := range copies {
		switch v[0] {
		case "bool":
			b.FieldByName(k).SetBool(v[1].(bool))
		case "string":
			b.FieldByName(k).SetString(v[1].(string))
		case "byte":
			b.FieldByName(k).SetBytes(v[1].([]byte))

		case "float32":
			val := v[1].(float32)
			b.FieldByName(k).SetFloat(float64(val))
		case "float64":
			val := v[1].(float64)
			b.FieldByName(k).SetFloat(val)

		case "int":
			val := v[1].(int)
			b.FieldByName(k).SetInt(int64(val))
		case "int8":
			val := v[1].(int8)
			b.FieldByName(k).SetInt(int64(val))
		case "int16":
			val := v[1].(int16)
			b.FieldByName(k).SetInt(int64(val))
		case "int32":
			val := v[1].(int32)
			b.FieldByName(k).SetInt(int64(val))
		case "int64":
			val := v[1].(int64)
			b.FieldByName(k).SetInt(val)

		case "uint":
			val := v[1].(uint)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint8":
			val := v[1].(uint8)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint16":
			val := v[1].(uint16)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint32":
			val := v[1].(uint32)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint64":
			val := v[1].(uint64)
			b.FieldByName(k).SetUint(val)
		case "uintptr":
			val := v[1].(uintptr)
			b.FieldByName(k).SetUint(uint64(val))
		}
	}
}

func (t Transformed) TransformerToProto(to interface{}) {
	e := reflect.ValueOf(t.From).Elem()

	valueCopied := map[string][]interface{}{}
	for i := 0; i < e.NumField(); i++ {
		ref := e.Type().Field(i)
		varName := ref.Name
		if varName == "ID" {
			varName = "Id"
		}
		varType := ref.Type.Name()
		varValue := e.Field(i).Interface()
		valueCopied[varName] = []interface{}{varType, varValue}
	}
	log.Println(valueCopied)
	autoConf(to, valueCopied)
}
