package condition

//  One of minecraft:broken, minecraft:bundle/has_selected_item, minecraft:carried, minecraft:component, minecraft:damaged, minecraft:extended_view, minecraft:fishing_rod/cast, minecraft:has_component, minecraft:keybind_down, minecraft:selected, minecraft:using_item, minecraft:view_entity or minecraft:custom_model_data.

type Broken struct {
	Value ConditionContainer
}

func (b Broken) ConditionType() string {
	return "minecraft:broken"
}

type BundleHasSelectedItem struct {
	Value ConditionContainer
}

func (b BundleHasSelectedItem) ConditionType() string {
	return "minecraft:bundle/has_selected_item"
}

type Carried struct {
	Value ConditionContainer
}

func (c Carried) ConditionType() string {
	return "minecraft:carried"
}

type Component struct {
	Predicate      string `json:"predicate"`
	ComponentValue string `json:"value"`
	Value          ConditionContainer
}

func (c Component) ConditionType() string {
	return "minecraft:component"
}

type Damaged struct {
	Value ConditionContainer
}

func (d Damaged) ConditionType() string {
	return "minecraft:damaged"
}

type ExtendedView struct {
	Value ConditionContainer
}

func (e ExtendedView) ConditionType() string {
	return "minecraft:extended_view"
}

type FishingRodCast struct {
	Value ConditionContainer
}

func (f FishingRodCast) ConditionType() string {
	return "minecraft:fishing_rod/cast"
}

type HasComponent struct {
	Component     string `json:"component"`
	IgnoreDefault bool   `json:"ignore_default,omitempty"`
	Value         ConditionContainer
}

func (h HasComponent) ConditionType() string {
	return "minecraft:has_component"
}

type KeybindDown struct {
	Key   string `json:"keybind"`
	Value ConditionContainer
}

func (k KeybindDown) ConditionType() string {
	return "minecraft:keybind_down"
}

type Selected struct {
	Value ConditionContainer
}

func (s Selected) ConditionType() string {
	return "minecraft:selected"
}

type UsingItem struct {
	Value ConditionContainer
}

func (u UsingItem) ConditionType() string {
	return "minecraft:using_item"
}

type ViewEntity struct {
	Value ConditionContainer
}

func (v ViewEntity) ConditionType() string {
	return "minecraft:view_entity"
}

type CustomModelData struct {
	Index int `json:"index"`
	Value ConditionContainer
}

func (c CustomModelData) ConditionType() string {
	return "minecraft:custom_model_data"
}
