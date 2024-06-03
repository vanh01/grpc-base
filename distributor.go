package grpcbase

import (
	"encoding/json"
	"reflect"
)

type GEventMessage struct {
	DataType       string
	Data           string
	ReturnDataType string
	ConsumerType   string
}

type GEventResult struct {
	DataType string
	Data     string
}

var TypeRegistry = make(map[string]reflect.Type)

func RegisterType(request interface{}) {
	t := reflect.TypeOf(request)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	TypeRegistry[t.Name()] = t
}

func RegisterTypeWithKey(key string, request interface{}) {
	t := reflect.TypeOf(request)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	TypeRegistry[key] = t
}

func RegisterGenericType(request interface{}) {
	t := reflect.TypeOf(request)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	f1 := t.Field(0).Type
	f2 := t.Field(1).Type
	TypeRegistry[f1.Name()] = f1
	TypeRegistry[f2.Name()] = f2
	TypeRegistry[f1.Name()+f2.Name()] = t
}

func GetRequest[T any, K any](request T) GEventMessage {
	data, err := json.Marshal(request)
	if err != nil {
		return GEventMessage{}
	}
	var k K
	dataType := reflect.TypeOf(request).Name()
	returnDataType := reflect.TypeOf(k).Name()
	return GEventMessage{
		DataType:       dataType,
		ReturnDataType: returnDataType,
		ConsumerType:   dataType + returnDataType,
		Data:           string(data),
	}
}

func SendMessage(request GEventMessage) GEventResult {
	consumer := reflect.New(TypeRegistry[request.ConsumerType])

	data := reflect.New(TypeRegistry[request.DataType]).Interface()
	err := json.Unmarshal([]byte(request.Data), &data)
	if err != nil {
		return GEventResult{
			DataType: reflect.TypeOf(err).Name(),
			Data:     err.Error(),
		}
	}

	value := consumer.MethodByName("Consume").Call([]reflect.Value{reflect.ValueOf(data).Elem()})

	result := value[0]

	resJson, err := json.Marshal(result.Interface())
	if err != nil {
		return GEventResult{
			DataType: reflect.TypeOf(err).Name(),
			Data:     err.Error(),
		}
	}

	return GEventResult{
		DataType: result.Type().Name(),
		Data:     string(resJson),
	}
}
