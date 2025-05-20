package model

import (
	"encoding/json"
	"github.com/patyhank/resourceup/item_model"
)

type Model struct {
	Model string
	Tints []any
}

func (m *Model) Type() string {
	return "minecraft:model"
}

func (m *Model) MarshalJSON() ([]byte, error) {
	ma := make(map[string]any)
	ma["type"] = m.Type()
	ma["model"] = m.Model
	if m.Tints != nil {
		ma["tints"] = m.Tints
	}
	return json.Marshal(ma)
}

type Composite struct {
	Models []item_model.Model
}

func (m *Composite) MarshalJSON() ([]byte, error) {
	ma := make(map[string]any)
	ma["type"] = m.Type()
	ma["models"] = m.Models
	return json.Marshal(ma)
}

func (m *Composite) Type() string {
	return "minecraft:model"
}

type Empty struct {
}

func (m *Empty) MarshalJSON() ([]byte, error) {
	ma := make(map[string]any)
	ma["type"] = m.Type()
	return json.Marshal(ma)
}

func (m *Empty) Type() string {
	return "minecraft:empty"
}
