package special

import (
	"encoding/json"
	"reflect"
	"strings"
)

type Special struct {
	Model SpecialType
	Base  string
}

type SpecialType interface {
	SpecialType() string
}

func (s *Special) MarshalJSON() ([]byte, error) {
	ma := make(map[string]any)
	{
		ma["type"] = s.Model.SpecialType()

		rTyp := reflect.TypeOf(s.Model).Elem()
		rVal := reflect.ValueOf(s.Model).Elem()
		for i := 0; i < rTyp.NumField(); i++ {
			field := rTyp.Field(i)
			value := rVal.Field(i)

			result, ok := field.Tag.Lookup("json")
			valKey := field.Name
			val := value.Interface()
			if ok {
				valKey = result
				if strings.HasSuffix(result, ",omitempty") {
					valKey = strings.TrimSuffix(result, ",omitempty")

					if reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface()) {
						continue
					}
				}
			}

			ma[valKey] = val
		}
	}

	m := map[string]any{
		"model": ma,
		"type":  s.Type(),
		"base":  s.Base,
	}

	return json.Marshal(m)
}

func (s *Special) Type() string {
	return "minecraft:special"
}
