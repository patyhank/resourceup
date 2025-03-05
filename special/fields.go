package special

type Banner struct {
	Color string `json:"color"`
}

func (b Banner) SpecialType() string {
	return "minecraft:banner"
}

type Bed struct {
	Texture string `json:"texture"`
}

func (b Bed) SpecialType() string {
	return "minecraft:bed"
}

type Chest struct {
	Texture  string  `json:"texture"`
	Openness float64 `json:"openness,omitempty"`
}

func (b Chest) SpecialType() string {
	return "minecraft:bed"
}

type Conduit struct {
}

func (b Conduit) SpecialType() string {
	return "minecraft:conduit"
}

type DecoratedPot struct {
}

func (b DecoratedPot) SpecialType() string {
	return "minecraft:decorated_pot"
}

type Head struct {
	Kind      HeadKind `json:"kind"`
	Texture   string   `json:"texture,omitempty"`
	Animation float64  `json:"animation,omitempty"`
}

func (b Head) SpecialType() string {
	return "minecraft:head"
}

type Shield struct {
}

func (b Shield) SpecialType() string {
	return "minecraft:shield"
}

type ShulkerBox struct {
	Texture string `json:"texture"`

	Openness    float64 `json:"openness,omitempty"`
	Orientation string  `json:"orientation,omitempty"`
}

func (b ShulkerBox) SpecialType() string {
	return "minecraft:shulker_box"
}

//standing_sign.,hanging_sign,trident

type StandingSign struct {
	WoodType string `json:"wood_type"`
	Texture  string `json:"texture,omitempty"`
}

func (b StandingSign) SpecialType() string {
	return "minecraft:standing_sign"
}

type HangingSign struct {
	WoodType string `json:"wood_type"`
	Texture  string `json:"texture,omitempty"`
}

func (b HangingSign) SpecialType() string {
	return "minecraft:hanging_sign"
}

type Trident struct {
}

func (b Trident) SpecialType() string {
	return "minecraft:trident"
}
