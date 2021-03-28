package transformer

import (
	"reflect"
	"unsafe"
)

// constants
const (
	// arrIndexValue index for take out value from []interface{}
	arrIndexValue = 1
	// arrIndexType index for take out type from []interface{}
	arrIndexType = 0
)

type Transformed struct {
	From interface{}
}

// StructTransformed struct
type StructTransformed Transformed

// ProtocTransformed struct
type ProtocTransformed Transformed

// autoAppliedCopies applied every name, type and value from an interface{} to targeted interface{}
func autoAppliedCopies(to interface{}, copies map[string][]interface{}) {
	b := reflect.ValueOf(to).Elem()
	for k, v := range copies {
		switch v[arrIndexType] {
		case "bool":
			b.FieldByName(k).SetBool(v[arrIndexValue].(bool))
		case "string":
			b.FieldByName(k).SetString(v[arrIndexValue].(string))
		case "byte":
			b.FieldByName(k).SetBytes(v[arrIndexValue].([]byte))

		case "float32":
			val := v[arrIndexValue].(float32)
			b.FieldByName(k).SetFloat(float64(val))
		case "float64":
			val := v[arrIndexValue].(float64)
			b.FieldByName(k).SetFloat(val)

		case "int":
			val := v[arrIndexValue].(int)
			b.FieldByName(k).SetInt(int64(val))
		case "int8":
			val := v[arrIndexValue].(int8)
			b.FieldByName(k).SetInt(int64(val))
		case "int16":
			val := v[arrIndexValue].(int16)
			b.FieldByName(k).SetInt(int64(val))
		case "int32":
			val := v[arrIndexValue].(int32)
			b.FieldByName(k).SetInt(int64(val))
		case "int64":
			val := v[arrIndexValue].(int64)
			b.FieldByName(k).SetInt(val)

		case "uint":
			val := v[arrIndexValue].(uint)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint8":
			val := v[arrIndexValue].(uint8)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint16":
			val := v[arrIndexValue].(uint16)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint32":
			val := v[arrIndexValue].(uint32)
			b.FieldByName(k).SetUint(uint64(val))
		case "uint64":
			val := v[arrIndexValue].(uint64)
			b.FieldByName(k).SetUint(val)
		case "uintptr":
			val := v[arrIndexValue].(uintptr)
			b.FieldByName(k).SetUint(uint64(val))
		}
	}
}

func Protoc(protocStub interface{}) StructTransformed {
	return StructTransformed{From: protocStub}
}

func Struct(strct interface{}) ProtocTransformed {
	return ProtocTransformed{From: strct}
}

func (t StructTransformed) Struct(to interface{}) {
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

func (t ProtocTransformed) Protoc(to interface{}) {
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
