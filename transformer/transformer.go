package transformer

import (
	"reflect"
	"unsafe"
)

// constants
const (
	// ArrIndexValue index for take out value from []interface{}
	ArrIndexValue = 1
	// ArrIndexValue index for take out type from []interface{}
	ArrIndexType = 0
)

// ITransformed functionality contracts
type ITransformed interface {
	ToStruct(to interface{})
	ToProtoc(to interface{})
}

// Transformed struct
type Transformed struct {
	From interface{}
}

// autoAppliedCopies applied every name, type and value from an interface{} to targeted interface{}
func autoAppliedCopies(to interface{}, copies map[string][]interface{}) {
	b := reflect.ValueOf(to).Elem()
	for k, v := range copies {
		switch v[ArrIndexType] {
		case "bool":
			b.FieldByName(k).SetBool(v[ArrIndexValue].(bool))
		case "string":
			b.FieldByName(k).SetString(v[ArrIndexValue].(string))
		case "byte":
			b.FieldByName(k).SetBytes(v[ArrIndexValue].([]byte))

		case "float32":
			val := v[ArrIndexValue].(float32)
			b.FieldByName(k).SetFloat(float64(val))
		case "float64":
			val := v[ArrIndexValue].(float64)
			b.FieldByName(k).SetFloat(val)

		case "int":
			val := v[ArrIndexValue].(int)
			b.FieldByName(k).SetInt(int64(val))
		case "int8":
			val := v[ArrIndexValue].(int8)
			b.FieldByName(k).SetInt(int64(val))
		case "int16":
			val := v[ArrIndexValue].(int16)
			b.FieldByName(k).SetInt(int64(val))
		case "int32":
			val := v[ArrIndexValue].(int32)
			b.FieldByName(k).SetInt(int64(val))
		case "int64":
			val := v[ArrIndexValue].(int64)
			b.FieldByName(k).SetInt(val)

		case "uint":
			val := v[ArrIndexValue].(uint)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint8":
			val := v[ArrIndexValue].(uint8)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint16":
			val := v[ArrIndexValue].(uint16)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint32":
			val := v[ArrIndexValue].(uint32)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint64":
			val := v[ArrIndexValue].(uint64)
			b.FieldByName(k).SetUint(val)
		case "uintptr":
			val := v[ArrIndexValue].(uintptr)
			b.FieldByName(k).SetUint(uint64(val))
		}
	}
}

func (t Transformed) ToStruct(to interface{}) {
	e := reflect.ValueOf(t.From).Elem()
	b := reflect.ValueOf(to).Elem()

	valueCopied := map[string][]interface{}{}
	for i := 0; i < e.NumField(); i++ {
		ref := e.Field(i)
		refIdentify := e.Type().Field(i)

		varName := refIdentify.Name
		if varName == "Id" {
			varName = "ID"
		}
		if b.FieldByName(varName).CanSet() {
			varType := ref.Type().Name()
			unsafeReflect := reflect.NewAt(ref.Type(), unsafe.Pointer(ref.UnsafeAddr())).Elem().Interface()
			varValue := unsafeReflect
			valueCopied[varName] = []interface{}{varType, varValue}
		}
	}
	autoAppliedCopies(to, valueCopied)
}

func (t Transformed) ToProtoc(to interface{}) {
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
	autoAppliedCopies(to, valueCopied)
}
