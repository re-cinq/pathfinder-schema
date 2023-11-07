package schema

import (
	"encoding/json"
	"errors"
)

type RegionOrSubregion string

// Error parsing the RegionOrSubregion
var ErrRegionOrSubregionParse = errors.New("unsupported RegionOrSubregion")

var regions = map[string]RegionOrSubregion{
	"Africa":                          Africa,
	"Americas":                        Americas,
	"Asia":                            Asia,
	"Europe":                          Europe,
	"Oceania":                         Oceania,
	"Australia and New Zealand":       AustraliaAndNewZealand,
	"Central Asia":                    CentralAsia,
	"Eastern Asia":                    EasternAsia,
	"Eastern Europe":                  EasternEurope,
	"Latin America and the Caribbean": LatinAmericaAndCaribbean,
	"Melanesia":                       Melanesia,
	"Micronesia":                      Micronesia,
	"Northern Africa":                 NorthernAfrica,
	"Northern America":                NorthernAmerica,
	"Northern Europe":                 NorthernEurope,
	"Polynesia":                       Polynesia,
	"South-eastern Asia":              SouthEasternAsia,
	"Southern Asia":                   SouthernAsia,
	"Southern Europe":                 SouthernEurope,
	"Sub-Saharan Africa":              SubSaharanAfrica,
	"Western Asia":                    WesternAsia,
	"WesternEurope":                   WesternEurope,
}

const (
	// for the UN geographic region Africa
	Africa RegionOrSubregion = "Africa"

	// for the UN geographic region Americas
	Americas RegionOrSubregion = "Americas"

	// for the UN geographic region Asia
	Asia RegionOrSubregion = "Asia"

	// for the UN geographic region Europe
	Europe RegionOrSubregion = "Europe"

	// for the UN geographic region Oceania
	Oceania RegionOrSubregion = "Oceania"

	// for the UN geographic subregion Australia and New Zealand
	AustraliaAndNewZealand RegionOrSubregion = "Australia and New Zealand"

	// for the UN geographic subregion Central Asia
	CentralAsia RegionOrSubregion = "Central Asia"

	// for the UN geographic subregion Eastern Asia
	EasternAsia RegionOrSubregion = "Eastern Asia"

	// for the UN geographic subregion Eastern Europe
	EasternEurope RegionOrSubregion = "Eastern Europe"

	// for the UN geographic subregion Latin America and the Caribbean
	LatinAmericaAndCaribbean RegionOrSubregion = "Latin America and the Caribbean"

	// for the UN geographic subregion Melanesia
	Melanesia RegionOrSubregion = "Melanesia"

	// for the UN geographic subregion Micronesia
	Micronesia RegionOrSubregion = "Micronesia"

	// for the UN geographic subregion Northern Africa
	NorthernAfrica RegionOrSubregion = "Northern Africa"

	// for the UN geographic subregion Northern America
	NorthernAmerica RegionOrSubregion = "Northern America"

	// for the UN geographic subregion Northern Europe
	NorthernEurope RegionOrSubregion = "Northern Europe"

	// for the UN geographic subregion Polynesia
	Polynesia RegionOrSubregion = "Polynesia"

	// for the UN geographic subregion South-eastern Asia
	SouthEasternAsia RegionOrSubregion = "South-eastern Asia"

	// for the UN geographic subregion Southern Asia
	SouthernAsia RegionOrSubregion = "Southern Asia"

	// for the UN geographic subregion Southern Europe
	SouthernEurope RegionOrSubregion = "Southern Europe"

	// for the UN geographic subregion Sub-Saharan Africa
	SubSaharanAfrica RegionOrSubregion = "Sub-Saharan Africa"

	// for the UN geographic subregion Western Asia
	WesternAsia RegionOrSubregion = "Western Asia"

	// for the UN geographic subregion Western Europe
	WesternEurope RegionOrSubregion = "Western Europe"
)

func (u RegionOrSubregion) String() string {
	return string(u)
}

func (u *RegionOrSubregion) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if status, ok := regions[value]; !ok {
		return ErrRegionOrSubregionParse
	} else {
		*u = status
	}

	return nil
}
