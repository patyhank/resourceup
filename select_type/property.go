package select_type

import "fmt"

type ChargeTypeValue string

const (
	ChargeTypeNone   ChargeTypeValue = "none"
	ChargeTypeRocket ChargeTypeValue = "rocket"
	ChargeTypeArrow  ChargeTypeValue = "arrow"
)

func (c ChargeTypeValue) IsValid(data ChargeTypeValue) bool {
	validValues := []ChargeTypeValue{ChargeTypeNone, ChargeTypeRocket, ChargeTypeArrow}
	for _, validValue := range validValues {
		if validValue == data {
			return true
		}
	}
	return false
}

type DisplayContextValue string

const (
	DisplayContextNone                 DisplayContextValue = "none"
	DisplayContextThirdPersonLeftHand  DisplayContextValue = "thirdperson_lefthand"
	DisplayContextThirdPersonRightHand DisplayContextValue = "thirdperson_righthand"
	DisplayContextFirstPersonLeftHand  DisplayContextValue = "firstperson_lefthand"
	DisplayContextFirstPersonRightHand DisplayContextValue = "firstperson_righthand"
	DisplayContextHead                 DisplayContextValue = "head"
	DisplayContextGUI                  DisplayContextValue = "gui"
	DisplayContextGround               DisplayContextValue = "ground"
	DisplayContextFixed                DisplayContextValue = "fixed"
)

func (d DisplayContextValue) IsValid(data DisplayContextValue) bool {
	validValues := []DisplayContextValue{DisplayContextNone, DisplayContextThirdPersonLeftHand, DisplayContextThirdPersonRightHand, DisplayContextFirstPersonLeftHand, DisplayContextFirstPersonRightHand, DisplayContextHead, DisplayContextGUI, DisplayContextGround, DisplayContextFixed}
	for _, validValue := range validValues {
		if validValue == data {
			return true
		}
	}
	return false
}

type MainHandValue string

const (
	MainHandLeft  MainHandValue = "left"
	MainHandRight MainHandValue = "right"
)

func (m MainHandValue) IsValid(data MainHandValue) bool {
	validValues := []MainHandValue{MainHandLeft, MainHandRight}
	for _, validValue := range validValues {
		if validValue == data {
			return true
		}
	}
	return false
}

type NamespaceID string

type CustomModelDataValue string

func IntToCustomModelData(data int) CustomModelDataValue {
	return CustomModelDataValue(fmt.Sprintf("%d", data))
}

func FloatToCustomModelData(data float64) CustomModelDataValue {
	return CustomModelDataValue(fmt.Sprintf("%.0f", data))
}
