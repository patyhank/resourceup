package range_dispatch_type

type BundleFullness struct {
	Value RangeContainer[float64] `json:"-"`
}

func (b BundleFullness) RangeDispatchType() string {
	return "minecraft:bundle/fullness"
}

type Compass struct {
	Value RangeContainer[ZeroToOneDataValue] `json:"-"`

	Target CompassTarget `json:"target,omitempty"`
	Wobble bool          `json:"wobble,omitempty"`
}

func (b Compass) RangeDispatchType() string {
	return "minecraft:compass"
}

type Cooldown struct {
	Value RangeContainer[ZeroToOneDataValue] `json:"-"`
}

func (b Cooldown) RangeDispatchType() string {
	return "minecraft:cooldown"
}

type Count struct {
	Value RangeContainer[float64] `json:"-"`

	Normalize bool `json:"normalize,omitempty"`
}

func (b Count) RangeDispatchType() string {
	return "minecraft:count"
}

type CrossbowPull struct {
	Value RangeContainer[float64] `json:"-"`
}

func (b CrossbowPull) RangeDispatchType() string {
	return "minecraft:crossbow/pull"
}

type Damage struct {
	Value RangeContainer[float64] `json:"-"`

	Normalize bool `json:"normalize,omitempty"`
}

func (b Damage) RangeDispatchType() string {
	return "minecraft:damage"
}

type Time struct {
	Value RangeContainer[ZeroToOneDataValue] `json:"-"`

	Source TimeSource `json:"source"`
	Wobble bool       `json:"wobble,omitempty"`
}

func (b Time) RangeDispatchType() string {
	return "minecraft:time"
}

type UseCycle struct {
	Value RangeContainer[float64] `json:"-"`

	Period float64 `json:"period,omitempty"`
}

func (b UseCycle) RangeDispatchType() string {
	return "minecraft:use_cycle"
}

type UseDuration struct {
	Value RangeContainer[float64] `json:"-"`

	Remaining bool `json:"remaining,omitempty"`
}

func (b UseDuration) RangeDispatchType() string {
	return "minecraft:use_duration"
}

type CustomModelData struct {
	Value RangeContainer[CustomModelDataValue] `json:"-"`

	Index int32 `json:"index,omitempty"`
}

func (b CustomModelData) RangeDispatchType() string {
	return "minecraft:custom_model_data"
}
