package item_model

import "errors"

type Model interface {
	Type() string
}

var ErrInvalidType = errors.New("invalid type value: ")

type FieldValidator[T any] interface {
	IsValid(data T) bool
}

func ValidType[T any](data T) bool {
	validator, ok := any(data).(FieldValidator[T])
	if !ok {
		return true
	}

	return validator.IsValid(data)
}
func ValidTypes[T any](data []T) bool {
	for _, datum := range data {
		validator, ok := any(datum).(FieldValidator[T])
		if !ok {
			continue
		}
		if !validator.IsValid(datum) {
			return false
		}
	}
	return true
}

//type ItemModel struct {
//	Model
//}
//
//func (i *ItemModel) MarshalJSON() ([]byte, error) {
//	m := make(map[string]any)
//	m["type"] = i.Model.Type()
//
//	rTyp := reflect.TypeOf(i.Model)
//	rVal := reflect.ValueOf(i.Model)
//
//	for i := 0; i < rTyp.NumField(); i++ {
//		field := rTyp.Field(i)
//		value := rVal.Field(i)
//
//		m[field.Name] = value.Interface()
//	}
//
//	return json.Marshal(m)
//}
