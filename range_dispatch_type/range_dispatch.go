package range_dispatch_type

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/patyhank/resourceup/item_model"
	"reflect"
)

type Value[T ~float64] struct {
	Threshold T `json:"threshold"`

	Model item_model.Model `json:"model"`
}

type RangeContainer[T ~float64] struct {
	Entries []Value[T]

	Fallback item_model.Model
}

type RangeDispatchType interface {
	RangeDispatchType() string
}

func (v *Value[T]) MarshalJSON() ([]byte, error) {
	m := make(map[string]any)

	if !item_model.ValidType(v.Threshold) {
		return nil, errors.Join(item_model.ErrInvalidType, fmt.Errorf("%T %v", v.Threshold, v.Threshold))
	}

	m["threshold"] = v.Threshold
	m["model"] = v.Model

	return json.Marshal(m)
}

type RangeDispatcher struct {
	RangeDispatchType
}

func (s *RangeDispatcher) MarshalJSON() ([]byte, error) {
	selectType := s.RangeDispatchType

	selectValue := reflect.ValueOf(selectType)

	valueField := selectValue.Elem().FieldByName("Value")
	entries := valueField.FieldByName("Entries").Interface()
	fallback := valueField.FieldByName("Fallback").Interface()

	data := map[string]any{
		"type":     s.Type(),
		"property": selectType.RangeDispatchType(),
		"entries":  entries,
		"fallback": fallback,
	}

	return json.Marshal(data)
}

func (s *RangeDispatcher) Type() string {
	return "minecraft:range_dispatch"
}
