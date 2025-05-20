package range_dispatch_type

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/patyhank/resourceup/item_model"
	"strconv"
)

type CustomModelDataValue float64

func StringToCustomModelData(data string) CustomModelDataValue {
	val, err := strconv.ParseFloat(data, 64)
	if err != nil {
		return 0
	}
	return CustomModelDataValue(val)
}

type ZeroToOneDataValue float64

func (a ZeroToOneDataValue) IsValid(data ZeroToOneDataValue) bool {
	return data >= 0.0 && data <= 1.0
}

type CompassTarget string

const (
	CompassTargetSpawn     CompassTarget = "spawn"
	CompassTargetLodestone CompassTarget = "lodestone"
	CompassTargetRecovery  CompassTarget = "recovery"
	CompassTargetNone      CompassTarget = "none"
)

func (c CompassTarget) IsValid(data CompassTarget) bool {
	validVal := []CompassTarget{CompassTargetSpawn, CompassTargetLodestone, CompassTargetRecovery, CompassTargetNone}
	for _, val := range validVal {
		if data == val {
			return true
		}
	}
	return false
}

func (c CompassTarget) MarshalJSON() ([]byte, error) {
	if !c.IsValid(c) {
		return nil, errors.Join(item_model.ErrInvalidType, fmt.Errorf("%T %v", c, c))
	}

	return json.Marshal(c)
}

type TimeSource string

const (
	TimeSourceDayTime   TimeSource = "daytime"
	TimeSourceMoonPhase TimeSource = "moon_phase"
	TimeSourceRandom    TimeSource = "random"
)

func (t TimeSource) IsValid(data TimeSource) bool {
	validVal := []TimeSource{TimeSourceDayTime, TimeSourceMoonPhase, TimeSourceRandom}
	for _, val := range validVal {
		if data == val {
			return true
		}
	}
	return false
}
