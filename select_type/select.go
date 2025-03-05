package select_type

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"resourceup/item_model"
)

type Value[T ~string] struct {
	When []T

	Model item_model.Model
}

func (v *Value[T]) MarshalJSON() ([]byte, error) {
	m := make(map[string]any)

	if len(v.When) == 1 {
		if item_model.ValidType(v.When[0]) {
			m["when"] = v.When[0]
		} else {
			return nil, errors.New(fmt.Sprintf("invalid %T value: %#v", v.When[0], v.When[0]))
		}
	} else {
		if item_model.ValidTypes(v.When) {
			m["when"] = v.When
		} else {
			return nil, errors.New(fmt.Sprintf("invalid %T value: %#v", v.When[0], v.When))
		}
	}

	m["model"] = v.Model

	return json.Marshal(m)
}

type CaseContainer[T ~string] struct {
	Case []Value[T]

	Fallback item_model.Model
}

type SelectType interface {
	SelectType() string
}

type Select struct {
	SelectType
}

func (s *Select) MarshalJSON() ([]byte, error) {
	selectType := s.SelectType

	selectValue := reflect.ValueOf(selectType)
	valueField := selectValue.FieldByName("Value")
	cases := valueField.FieldByName("Case").Interface()
	fallback := valueField.FieldByName("Fallback").Interface()

	data := map[string]any{
		"type":     s.Type(),
		"property": selectType.SelectType(),
		"cases":    cases,
		"fallback": fallback,
	}

	return json.Marshal(data)
}

func (s *Select) Type() string {
	return "minecraft:select"
}
