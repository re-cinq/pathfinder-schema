package schema

import (
	"encoding/json"
	"errors"
)

type Coverage string

// Error parsing the Coverage
var ErrCoverageParse = errors.New("unsupported Coverage")

var coverageList = map[string]Coverage{
	"corporate level": Corporate,
	"product line":    ProductLine,
	"PCF system":      PCF,
	"product level":   Product,
}

const (
	// Corporate level of granularity of the emissions data assured
	Corporate Coverage = "corporate level"

	// Product line level of granularity of the emissions data assured
	ProductLine Coverage = "product line"

	// PCF system level of granularity of the emissions data assured
	PCF Coverage = "PCF system"

	// Product level of granularity of the emissions data assured
	Product Coverage = "product level"
)

func (u Coverage) String() string {
	return string(u)
}

func (u *Coverage) UnmarshalJSON(data []byte) error {
	var value string
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}

	if status, ok := coverageList[value]; !ok {
		return ErrCoverageParse
	} else {
		*u = status
	}

	return nil
}
