package schema

import "encoding/json"

// DataModelExtension additional schema
type DataModelExtension struct {
	// Version of the spec
	// Mandatory
	SpecVersion string `json:"specVersion"`

	// DataSchema
	// Mandatory
	DataSchema string `json:"dataSchema"`

	// Data
	// Mandatory
	Data json.RawMessage `json:"data"`
}
