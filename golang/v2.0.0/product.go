package schema

import (
	"time"

	"github.com/google/uuid"
	urn "github.com/leodido/go-urn"
)

type ProductFootprint struct {

	// The product footprint identifier
	// Format: uuidv4
	//
	// Mandatory
	Id uuid.UUID `json:"id"`

	// The version of the ProductFootprint data specification with value 2.0.0
	//
	// Mandatory
	SpecVersion string `json:"specVersion"`

	// If defined, MUST be non-empty set of preceding product
	// footprint identifiers without duplicates
	//
	// Optional
	PrecedingPfIds []uuid.UUID `json:"precedingPfIds,omitempty"`

	// The version of the ProductFootprint with value
	// an integer in the inclusive range of 0..2^31-1.
	//
	// Mandatory
	Version int32 `json:"version"`

	// A ProductFootprint MUST include the property created with value
	// the timestamp of the creation of the ProductFootprint.
	//
	// Mandatory
	Created time.Time `json:"created"`

	// A ProductFootprint SHOULD include the property updated
	// with value the timestamp of the ProductFootprint update.
	// A ProductFootprint MUST NOT include this property if
	// an update has never been performed. The timestamp MUST be in UTC.
	//
	// Optional
	Updated time.Time `json:"updated,omitempty"`

	// If defined, the value must be one of the following values:
	// Active:
	//   The default status of a product footprint is Active.
	//   A product footprint with status Active can be used by a data recipients,
	//   e.g. for product footprint calculations.
	// Deprecated:
	//  The product footprint is deprecated and should not be used for e.g.
	//  product footprint calculations by data recipients.
	//
	// Mandatory
	Status Status `json:"status"`

	// if defined, the value should be a message explaining the reason for the current status.
	//
	// Optional
	StatusComment string `json:"statusComment,omitempty"`

	// If defined, the start of the validity period of the ProductFootprint.
	// The validity period is the time interval during which the ProductFootprint is declared as valid
	// for use by a receiving data recipient.
	// The validity period is defined by the properties
	// validityPeriodStart (including) and validityPeriodEnd (excluding).
	// If a validity period is to be specified, then:
	// 1. the value of validityPeriodStart MUST be defined with value greater than or equal to
	//    the value of referencePeriodEnd.
	// 2. the value of validityPeriodEnd MUST be defined with value
	//    a. strictly greater than validityPeriodStart, and
	//    b. less than or equal to referencePeriodEnd + 3 years.
	//
	// Optional
	ValidityPeriodStart time.Time `json:"validityPeriodStart,omitempty"`

	// The end (excluding) of the valid period of the ProductFootprint.
	// See validityPeriodStart for further details.
	//
	// Optional
	ValidityPeriodEnd time.Time `json:"validityPeriodEnd,omitempty"`

	// The name of the company that is the ProductFootprint Data Owner, with value a non-empty
	//
	// Mandatory
	CompanyName string `json:"companyName"`

	// The non-empty set of Uniform Resource Names (URN).
	// Each value of this set is supposed to uniquely identify the ProductFootprint Data Owner.
	// See CompanyIdSet for details.
	// https://wbcsd.github.io/tr/2023/data-exchange-protocol-20230221/#dt-productid-custom
	//
	// Mandatory
	CompanyIds []urn.URN `json:"companyIds"`

	// The free-form description of the product plus other information related to it
	// such as production technology or packaging.
	//
	// Mandatory
	ProductDescription string `json:"productDescription"`

	// Product Ids
	// The non-empty set of ProductIds.
	// Each of the values in the set is supposed to uniquely identify the product.
	// What constitutes a suitable product identifier depends on the product,
	// the conventions, contracts, and agreements between the Data Owner and a Data Recipient
	// and is out of the scope of this specification.
	//
	// Mandatory
	ProductIds []urn.URN `json:"productIds"`

	// A UN Product Classification Code (CPC) that the given product belongs to.
	// UN CPC Code
	//
	// Mandatory
	ProductCategoryCpc string `json:"productCategoryCpc"`

	// The non-empty trade name of the product.
	//
	// Mandatory
	ProductNameCompany string `json:"productNameCompany"`

	// The additional information related to the product footprint.
	// Whereas the property productDescription contains product-level information,
	// comment SHOULD be used for information and instructions related to the
	// calculation of the footprint,
	// or other information which informs the ability to interpret, to audit or to verify
	// the Product Footprint.
	//
	// Mandatory
	Comment string `json:"comment"`

	// The carbon footprint of the given product with value conforming to the data type CarbonFootprint.
	//
	// Mandatory
	Pcf CarbonFootprint `json:"pcf"`

	// If defined, 1 or more data model extensions associated with the ProductFootprint.
	// extensions MUST be encoded as a non-empty JSON Array of DataModelExtension JSON objects.
	// See DataModelExtension for details.
	//
	// Optional
	Extensions []DataModelExtension `json:"extensions,omitempty"`
}
