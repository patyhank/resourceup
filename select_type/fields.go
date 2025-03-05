package select_type

type CustomModelData struct {
	Index int32                               `json:"index,omitempty"`
	Value CaseContainer[CustomModelDataValue] `json:"-"`
}

func (CustomModelData) SelectType() string {
	return "minecraft:custom_model_data"
}

type TrimMaterial struct {
	Value CaseContainer[NamespaceID] `json:"-"`
}

func (TrimMaterial) SelectType() string {
	return "minecraft:trim_material"
}

type MainHand struct {
	Value CaseContainer[MainHandValue] `json:"-"`
}

func (MainHand) SelectType() string {
	return "minecraft:main_hand"
}

type LocalTime struct {
	Timezone string                `json:"timezone,omitempty"`
	Locale   string                `json:"locale,omitempty"`
	Pattern  string                `json:"pattern"`
	Value    CaseContainer[string] `json:"-"`
}

func (LocalTime) SelectType() string {
	return "minecraft:local_time"
}

type DisplayContext struct {
	Value CaseContainer[DisplayContextValue] `json:"-"`
}

func (DisplayContext) SelectType() string {
	return "minecraft:display_context"
}

type ContextEntityType struct {
	Value CaseContainer[string] `json:"-"`
}

func (ContextEntityType) SelectType() string {
	return "minecraft:context_entity_type"
}

type ContextDimension struct {
	Value CaseContainer[string] `json:"-"`
}

func (ContextDimension) SelectType() string {
	return "minecraft:context_dimension"
}

type Component struct {
	Value CaseContainer[string] `json:"-"`
}

func (Component) SelectType() string {
	return "minecraft:component"
}

type ChargeType struct {
	Value CaseContainer[ChargeTypeValue] `json:"-"`
}

func (ChargeType) SelectType() string {
	return "minecraft:charge_type"
}

type BlockState struct {
	BlockStateProperty string                `json:"block_state_property,omitempty"`
	Value              CaseContainer[string] `json:"-"`
}

func (BlockState) SelectType() string {
	return "minecraft:block_state"
}
