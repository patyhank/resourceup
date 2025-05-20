package condition

import (
	"encoding/json"
	"github.com/patyhank/resourceup/item_model"
	"reflect"
)

type ConditionContainer struct {
	TrueModel  item_model.Model
	FalseModel item_model.Model
}

type ConditionType interface {
	ConditionType() string
}

type Condition struct {
	ConditionType
}

func (s *Condition) MarshalJSON() ([]byte, error) {
	condition := s.ConditionType

	selectValue := reflect.ValueOf(condition)
	valueField := selectValue.Elem().FieldByName("Value")
	onTrue := valueField.FieldByName("TrueModel").Interface()
	onFalse := valueField.FieldByName("FalseModel").Interface()

	data := map[string]any{
		"type":     s.Type(),
		"property": condition.ConditionType(),
		"on_true":  onTrue,
		"on_false": onFalse,
	}

	return json.Marshal(data)
}

func (s *Condition) Type() string {
	return "minecraft:condition"
}
