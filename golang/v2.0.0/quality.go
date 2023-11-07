package schema

import "github.com/shopspring/decimal"

type DataQualityIndicators struct {

	// Percentage of PCF included in the data quality assessment based on the >5% emissions threshold.
	//
	// Format: number
	//
	CoveragePercent Percentage `json:"coveragePercent"`

	// Quantitative data quality rating (DQR) based on the data quality matrix (See Pathfinder Framework Table 9),
	// scoring the technological representativeness of the sources used for PCF calculation
	// based on weighted average of all inputs representing >5% of PCF emissions.
	//
	// Format: number => The value MUST be a decimal between 1 and 3 including.
	//
	TechnologicalDQR decimal.Decimal `json:"technologicalDQR"`

	// Quantitative data quality rating (DQR) based on the data quality matrix (Table 9),
	// scoring the temporal representativeness of the sources used for PCF calculation
	// based on weighted average of all inputs representing >5% of PCF emissions.
	//
	// Format: number => The value MUST be a decimal between 1 and 3 including.
	//
	TemporalDQR decimal.Decimal `json:"temporalDQR"`

	// Quantitative data quality rating (DQR) based on the data quality matrix (Table 9),
	// scoring the geographical representativeness of the sources used
	// for PCF calculation based on weighted average of all inputs representing >5% of PCF emissions.
	//
	// The value MUST be a decimal between 1 and 3 including.
	//
	GeographicalDQR decimal.Decimal `json:"geographicalDQR"`

	// Quantitative data quality rating (DQR) based on the data quality matrix (Table 9),
	// scoring the completeness of the data collected for PCF calculation based on
	// weighted average of all inputs representing >5% of PCF emissions.
	//
	// The value MUST be a decimal between 1 and 3 including.
	//
	CompletenessDQR decimal.Decimal `json:"completenessDQR"`

	// Quantitative data quality rating (DQR) based on the data quality matrix (Table 9),
	// scoring the reliability of the data collected for PCF calculation based on
	// weighted average of all inputs representing >5% of PCF emissions.
	//
	// The value MUST be a decimal between 1 and 3 including.
	//
	ReliabilityDQR decimal.Decimal `json:"reliabilityDQR"`
}
