package schema

type ProductOrSectorSpecificRule struct {

	// A ProductOrSectorSpecificRule MUST include the property operator with the value conforming to data type ProductOrSectorSpecificRuleOperator.
	// ruleNames
	//
	// Mandatory
	Operator PCROperator `json:"operator"`

	// A ProductOrSectorSpecificRule MUST include the property ruleNames with value the non-empty set of
	// rules applied from the specified operator.
	// NonEmptyStringVector
	//
	// Mandatory
	RuleNames []string `json:"ruleNames"`

	// If the value of property 'operator' is Other,
	// a ProductOrSectorSpecificRule MUST include the property otherOperatorName
	// with value the name of the operator.
	// In this case, the operator declared MUST NOT be included in the definition of ProductOrSectorSpecificRuleOperator.
	// If the value of property operator is NOT Other, the property otherOperatorName of a ProductOrSectorSpecificRule
	// MUST be undefined.
	OtherOperatorName string `json:"otherOperatorName,omitempty"`
}
