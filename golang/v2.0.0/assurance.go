package schema

import "time"

type Assurance struct {

	// A boolean flag indicating whether the CarbonFootprint has been assured
	// in line with Pathfinder Framework requirements (section 5).
	// Assurance and verification undertaken by independent verifiers can help establish whether
	// PCFs have been accounted for in compliance with the Pathfinder Framework and relevant
	// standards, sectoral guidance, PCRs, and accompanying methods.
	Assurance bool `json:"assurance,omitempty"`

	// Level of granularity of the emissions data assured, with value equal to:
	// - corporate level: for corporate level
	// - product line: for product line
	// - PCF system: for PCF System
	// - product level: for product level
	//
	// This property MAY be undefined only if the kind of assurance was not performed.
	Coverage Coverage `json:"coverage,omitempty"`

	// Level of assurance applicable to the PCF, with value equal to:
	// - limited: for limited assurance
	// - reasonable: for reasonable assurance
	//
	// This property MAY be undefined only if the kind of assurance was not performed.
	Level AssuranceLevel `json:"level,omitempty"`

	// Boundary of the assurance, with value equal to
	// - Gate-to-Gate for Gate-to-Gate
	// - Cradle-to-Gate for Cradle-to-Gate.
	//
	// This property MAY be undefined only if the kind of assurance was not performed.
	Boundary AssuranceBoundary `json:"boundary,omitempty"`

	// The non-empty name of the independent third party engaged to undertake the assurance.
	//
	// Mandatory
	ProviderName string `json:"providerName,omitempty"`

	// The date at which the assurance was completed
	//
	// Optional
	CompletedAt time.Time `json:"completedAt,omitempty"`

	// Name of the standard against which the PCF was assured.
	StandardName Standard `json:"standardName,omitempty"`

	// Any additional comments that will clarify the interpretation of the assurance.
	// The value of this property MAY be the empty string.
	Comments string `json:"comments,omitempty"`
}
