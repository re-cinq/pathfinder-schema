package schema

import (
	"encoding/json"
	"errors"
)

// liter: for unit liter
// kilogram: for unit kilogram
// cubic meter: for cubic meter
// kilowatt hour: for kilowatt hour
// megajoule: for megajoule
// ton kilometer: for ton kilometer
// square meter: for square meter
type DeclaredUnit string

// Error parsing the DeclaredUnit
var ErrDeclaredUnitParse = errors.New("unsupported DeclaredUnit")

var units = map[string]DeclaredUnit{
	"liter":         Liter,
	"kilogram":      KiloGram,
	"cubic meter":   CubicMeter,
	"kilowatt hour": KiloWattHour,
	"megajoule":     MegaJoule,
	"ton kilometer": TonKilometer,
	"square meter":  SquareMeter,
}

const (
	// for unit liter
	Liter DeclaredUnit = "liter"

	// for unit kilogram
	KiloGram DeclaredUnit = "kilogram"

	// for cubic meter
	CubicMeter DeclaredUnit = "cubic meter"

	// hour for kilowatt hour
	KiloWattHour DeclaredUnit = "kilowatt hour"

	// for megajoule
	MegaJoule DeclaredUnit = "megajoule"

	// ton kilometer
	TonKilometer DeclaredUnit = "ton kilometer"

	// square meter: for square meter
	SquareMeter DeclaredUnit = "square meter"
)

func (u DeclaredUnit) String() string {
	return string(u)
}

func (u DeclaredUnit) ShortString() string {
	switch u {
	case Liter:
		return "L"
	case KiloGram:
		return "Kg"
	case CubicMeter:
		return "m^3"
	case KiloWattHour:
		return "KWh"
	case MegaJoule:
		return "MJ"
	case TonKilometer:
		return "tkm"
	case SquareMeter:
		return "sq m"
	}
	return u.String()
}

func (u *DeclaredUnit) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if unit, ok := units[value]; !ok {
		return ErrDeclaredUnitParse
	} else {
		*u = unit
	}

	return nil
}
