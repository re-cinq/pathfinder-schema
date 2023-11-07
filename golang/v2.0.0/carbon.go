package schema

import (
	"time"

	"github.com/shopspring/decimal"
)

type CarbonFootprint struct {

	// The unit of analysis of the product.
	// See Data Type DeclaredUnit for further information.
	//
	// Mandatory
	DeclaredUnit string `json:"declaredUnit"`

	// The amount of Declared Units contained within the product to which the PCF is referring to.
	// The value MUST be strictly greater than 0.
	//
	// Mandatory
	UnitaryProductAmount decimal.Decimal `json:"unitaryProductAmount"`

	// The product carbon footprint of the product excluding biogenic CO2 emissions.
	// The value MUST be calculated per declared unit with unit kg of CO2 equivalent per declared unit (kgCO2e / declaredUnit),
	// expressed as a decimal equal to or greater than zero.
	//
	// Mandatory
	PCfExcludingBiogenic decimal.Decimal `json:"pCfExcludingBiogenic"`

	// If present, the product carbon footprint of the product including all biogenic emissions (CO2 and otherwise).
	// The value MUST be calculated per declared unit with unit kg of CO2 equivalent per declared unit (kgCO2e / declaredUnit),
	// expressed as a decimal.
	//
	// Note: the value of this property can be less than 0 (zero).
	//
	// Optional
	PCfIncludingBiogenic decimal.Decimal `json:"pCfIncludingBiogenic,omitempty"`

	// The emissions from fossil sources as a result of fuel combustion, from fugitive emissions,
	// and from process emissions. The value MUST be calculated per declared unit with unit kg of CO2 equivalent
	// per declared unit (kgCO2e / declaredUnit), expressed as a decimal equal to or greater than zero.
	//
	// Mandatory
	FossilGhgEmissions decimal.Decimal `json:"fossilGhgEmissions"`

	// The fossil carbon content of the product (mass of carbon).
	// The value MUST be calculated per declared unit with unit kg Carbon per declared unit (kgC / declaredUnit),
	// expressed as a decimal equal to or greater than zero.
	//
	// Mandatory
	FossilCarbonContent decimal.Decimal `json:"fossilCarbonContent"`

	// The biogenic carbon content of the product (mass of carbon).
	// The value MUST be calculated per declared unit with unit kg Carbon per declared unit (kgC / declaredUnit),
	// expressed as a decimal equal to or greater than zero.
	//
	// Mandatory
	BiogenicCarbonContent decimal.Decimal `json:"biogenicCarbonContent"`

	// If present, emissions resulting from recent (i.e., previous 20 years) carbon stock loss due to
	// land conversion directly on the area of land under consideration.
	// The value of this property MUST include direct land use change (dLUC) where available,
	// otherwise statistical land use change (sLUC) can be used.
	// The value MUST be calculated per declared unit with unit kg of CO2 equivalent per declared unit
	// (kgCO2e / declaredUnit), expressed as a decimal equal to or greater than zero.
	// See Pathfinder Framework (Appendix B) for details.
	//
	// Optional
	DLucGhgEmissions decimal.Decimal `json:"dLucGhgEmissions,omitempty"`

	// If present, GHG emissions and removals associated with land-management-related changes,
	// including non-CO2 sources.
	// The value MUST be calculated per declared unit with unit kg of CO2 equivalent per declared unit (kgCO2e / declaredUnit),
	// expressed as a decimal.
	//
	// Optional now but mandatory from 2025 onwards
	LandManagementGhgEmissions decimal.Decimal `json:"landManagementGhgEmissions,omitempty"`

	// If present, all other biogenic GHG emissions associated with product manufacturing and transport
	// that are not included in dLUC (dLucGhgEmissions), iLUC (iLucGhgEmissions),
	// and land management (landManagementGhgEmissions).
	// The value MUST be calculated per declared unit with unit kg of CO2 equivalent per declared unit
	// (kgCO2e / declaredUnit), expressed as a decimal equal to or greater than zero.
	//
	// Optional
	OtherBiogenicGhgEmissions decimal.Decimal `json:"otherBiogenicGhgEmissions,omitempty"`

	// If present, emissions resulting from recent (i.e., previous 20 years)
	// carbon stock loss due to land conversion on land not owned
	// or controlled by the company or in its supply chain,
	// induced by change in demand for products produced or sourced by the company.
	// The value MUST be calculated per declared unit with unit kg of CO2 equivalent
	// per declared unit (kgCO2e / declaredUnit),
	//  expressed as a decimal equal to or greater than zero.
	// See Pathfinder Framework (Appendix B) for details.
	//
	// Optional
	ILucGhgEmissions decimal.Decimal `json:"iLucGhgEmissions,omitempty"`

	// If present, the Biogenic Carbon contained in the product converted to kilogram of CO2e.
	// The value MUST be calculated per declared unit with unit kgCO2e / declaredUnit expressed
	// as a decimal equal to or less than zero.
	//
	// Optional
	BiogenicCarbonWithdrawal decimal.Decimal `json:"biogenicCarbonWithdrawal,omitempty"`

	// If present, the GHG emissions resulting from aircraft engine usage for the transport of the product.
	// The value MUST be calculated per declared unit with unit kg of CO2 equivalent per declared unit
	// (kgCO2e / declaredUnit), expressed as a decimal equal to or greater than zero.
	//
	// Optional
	AircraftGhgEmissions decimal.Decimal `json:"aircraftGhgEmissions,omitempty"`

	// The IPCC version of the GWP characterization factors used in the calculation of the PCF
	// The Pathfinder Framework provides the methodological framework for studying GHG
	// emissions. Companies shall account for the GHGs identified within the GHG Protocol titled “Required
	// Greenhouse Gases in Inventories; Accounting and Reporting Standard Amendment
	//
	// The list includes:
	// - carbon dioxide (CO2),
	// - methane (CH4),
	// - nitrous oxide (N2O),
	// - hydrofluorocarbons (HFCs),
	// - perfluorinated compounds,
	// - sulfur hexafluoride (SF6),
	// - nitrogen trifluoride (NF3),
	// - perfluorocarbons (PFCs),
	// - fluorinated ethers (HFEs),
	// - perfluoropolyethers (e.g., PFPEs),
	// - chlorofluorocarbons (CFCs),
	// - hydrochlorofluorocarbons (HCFCs)
	//
	// Following common practice, the global warming impact of these gases can be converted into and expressed
	// as CO2e. Their respective characterization factors (100-year GWP, including carbon feedbacks) shall
	// be derived from the latest version of the IPCC Assessment Report publication
	//
	// The value MUST be one of the following:
	//
	// AR6:
	//     for the Sixth Assessment Report of the Intergovernmental Panel on Climate Change (IPCC)
	// AR5:
	//    for the Fifth Assessment Report of the IPCC.
	//
	// The set of characterization factor identifiers will likely change in future revisions.
	// It is recommended to account for this when implementing the validation of this property.
	//
	// Mandatory
	CharacterizationFactors CharacterizationFactor `json:"characterizationFactors"`

	// The cross-sectoral standards applied for calculating or allocating GHG emissions
	//
	// GHG Protocol Product standard: for the GHG Protocol Product standard
	// ISO Standard 14067: for ISO Standard 14067
	// ISO Standard 14044: for ISO Standard 14044
	//
	// Mandatory
	CrossSectoralStandardsUsed []Standard `json:"crossSectoralStandardsUsed"`

	// The product-specific or sector-specific rules applied for calculating or allocating GHG emissions.
	// If no product or sector specific rules were followed, this set MUST be empty.
	// A ProductOrSectorSpecificRule refers to a set of product or sector specific rules published
	// by a specific operator and applied during product carbon footprint calculation.
	//
	// Optional
	ProductOrSectorSpecificRules []ProductOrSectorSpecificRule `json:"productOrSectorSpecificRules,omitempty"`

	// The standard followed to account for biogenic emissions and removals.
	// If defined, the value MUST be one of the following:
	// PEF:
	//    For the EU Product Environmental Footprint Guide
	// ISO:
	//    For the ISO 14067 standard
	// GHGP:
	//    For the Greenhouse Gas Protocol (GHGP) Land sector and Removals Guidance
	// Quantis:
	//    For the Quantis Accounting for Natural Climate Solutions Guidance
	//
	// Optional
	BiogenicAccountingMethodology AccountingMethodology `json:"biogenicAccountingMethodology,omitempty"`

	// The processes attributable to each lifecycle stage.
	// Example: Electricity consumption included as an input in the production phase
	//
	// Mandatory
	BoundaryProcessesDescription string `json:"boundaryProcessesDescription"`

	// he start (including) of the time boundary for which the PCF value is considered to be representative.
	// Specifically, this start date represents the earliest date from which activity data was collected to
	// include in the PCF calculation.
	//
	// See the Pathfinder Framework section 6.1.2.1 for further details.
	//
	// Mandatory
	ReferencePeriodStart time.Time `json:"referencePeriodStart"`

	// The end (excluding) of the time boundary for which the PCF value is considered to be representative.
	// Specifically, this end date represents the latest date from which activity data was
	// collected to include in the PCF calculation.
	//
	// See the Pathfinder Framework section 6.1.2.1 for further details.
	//
	// The time boundary of a PCF refers to the time period for which the PCF value is considered to be representative
	// While PCFs should be calculated on a regular basis to track improvements over time, the resources
	// required to calculate PCFs also need to be considered to ensure companies are able to scale
	// the calculations to a larger number of products. This is especially true for companies that currently rely on
	// manual PCF calculations and that do not yet have an automated calculation approach.
	//
	// PCFs shall therefore have a maximum validity period of up to three years, provided that no major
	// changes to the production process take place within the validity period. Major changes are defined as
	// a variance of 10 percent or more compared to the original PCF. After three years or if the PCF has
	// varied by more than 10 percent, PCF values will no longer be considered representative and shall be
	// recalculated and exchanged
	//
	// Companies that are able to do so are invited to update their PCFs more regularly and may also wish
	// to request suppliers to update their PCF calculations on a more regular basis (e.g., annually) based on
	// contractual agreements.
	//
	// The temporal validity of the PCF calculation will be captured by the reporting period.40 The PCF’s
	// reporting period and date of publication shall always be disclosed. Emissions that were averaged over
	// several years may be reported, e.g., to reduce the effect of revisions, turnarounds, or other untypical
	// production conditions.
	//
	// Mandatory
	ReferencePeriodEnd time.Time `json:"referencePeriodEnd"`

	// If present, a ISO 3166-2 Subdivision Code.
	// See § 4.2.1 Scope of a CarbonFootprint for further details.
	//
	// Example 1:
	//    value for the State of New York in the United States of America:
	//      US-NY
	// Example 2:
	//    value for the department Yonne in France :
	//      FR-89
	//
	// Optional
	GeographyCountrySubdivision string `json:"geographyCountrySubdivision,omitempty"`

	// If present, the value MUST conform to data type ISO3166CC.
	// See § 4.2.1 Scope of a CarbonFootprint for further details.
	//
	// Example value in case the geographic scope is France:
	//     FR
	//
	// Optional
	GeographyCountry string `json:"geographyCountry,omitempty"`

	// If present, the value MUST conform to data type RegionOrSubregion.
	// See § 4.2.1 Scope of a CarbonFootprint for further details.
	// Additionally, see the Pathfinder Framework Section 6.1.2.2.
	//
	// Optional
	GeographyRegionOrSubregion RegionOrSubregion `json:"geographyRegionOrSubregion"`

	// If secondary data was used to calculate the CarbonFootprint,
	// then it MUST include the property secondaryEmissionFactorSources with value the emission
	// factors used for the CarbonFootprint calculation.
	// If no secondary data is used, this property MUST BE undefined.
	//
	// An EmissionFactorDS references emission factor databases (see Pathfinder Framework Section 4.1.3.2).
	// Primary emission factors are also not always available. For instance, suppliers may be unable
	// to provide GHG data for a component required to manufacture the product for which Company X
	// wishes to calculate a PCF.
	// In such scenarios, emission factors from secondary sources should be used (base case).
	//
	// The employment of secondary emission factors shall be compliant with the general quality rules
	// for secondary data sources. To ensure the use of verified and credible secondary emission factors
	// while still allowing for flexibility in the data sources used, the Pathfinder Framework defines a series
	// of safeguards that secondary emission factors shall comply with if they are to be used for the
	// calculation of PCFs:
	//
	// 1. Documentation:
	//    - Data included in the secondary emission factor shall be validated in line with globally recognized LCA principles.
	//    - The emission factor source should ensure transparency by providing information on key methodological (i.e., LCA modeling approach,
	//      aggregation and allocation approach, if any) and data
	//      (time period, geography, technology,representativeness) elements
	// 2. Management and maintenance:
	//    - If life cycle inventory databases are used, they shall be periodically maintained and updated
	//      with the latest data sets.
	//
	// 3. Choice of modeling:
	//    - The modeling of the secondary emission factor shall be consistent with the methodological principles of this Framework
	//      (e.g., attributional approach).
	//
	// Optional
	SecondaryEmissionFactorSources []string `json:"secondaryEmissionFactorSources,omitempty"`

	// The Percentage of emissions excluded from PCF, expressed as a decimal number between 0.0 and 5 including.
	// See Pathfinder Framework.
	//
	// Mandatory
	ExemptedEmissionsPercent Percentage `json:"exemptedEmissionsPercent"`

	// Rationale behind exclusion of specific PCF emissions,
	// CAN be the empty string if no emissions were excluded.
	//
	// Optional
	ExemptedEmissionsDescription string `json:"exemptedEmissionsDescription,omitempty"`

	// A boolean flag indicating whether packaging emissions are included in the
	//  PCF (pCfExcludingBiogenic, pCfIncludingBiogenic).
	//
	// Mandatory
	PackagingEmissionsIncluded bool `json:"packagingEmissionsIncluded"`

	// Emissions resulting from the packaging of the product.
	// If present, the value MUST be calculated per declared unit with unit kg of CO2 equivalent per kilogram
	// (kgCO2e / declared unit), expressed as a decimal equal to or greater than zero.
	// The value MUST NOT be defined if packagingEmissionsIncluded is false.
	//
	// Optional
	PackagingGhgEmissions decimal.Decimal `json:"packagingGhgEmissions,omitempty"`

	// If present, a description of any allocation rules applied and the rationale explaining
	// how the selected approach aligns with Pathfinder
	// Framework rules (see Section 3.3.1.4).
	//
	// Optional
	AllocationRulesDescription string `json:"allocationRulesDescription,omitempty"`

	// If present, the results, key drivers, and a short qualitative description of the uncertainty assessment.
	//
	// Optional
	UncertaintyAssessmentDescription string `json:"uncertaintyAssessmentDescription,omitempty"`

	// Decimal number in percentage from 0 to 100. Examples:
	// 100
	// 23.0
	// 7.183924
	// 0.0
	//
	// The share of primary data in percent. See the Pathfinder Framework Sections 4.2.1 and 4.2.2, Appendix B.
	//
	// Initially, companies shall calculate and report, as part of PCF data exchange, on at least one of the following metrics:
	// - Primary Data Share (PDS): Percentage of PCF emissions that were calculated using primary activity and emissions data
	// - Data Quality Ratings (DQRs): Quantitative score for five data quality indicators based on the data quality matrix
	//
	// From 2025, both metrics shall be reported by companies to ensure continued alignment with
	// the Pathfinder Framework. This will ensure a fuller picture of both the quality of the PCFs and the
	// amount of primary data being used to calculate them. Until 2025, companies should base their
	// initial choice of metric(s) on the relevance to their situation and resources available. For instance, a
	// company calculating a PCF for the first time may not have access to a large amount of primary data and
	// may wish instead to reflect on the accuracy of the secondary sources used to calculate its PCF.
	//
	// Primary data share calculation:
	// To create visibility on the share of primary data in PCF calculations, the PDS in each data set should
	// be determined and exchanged across the value chain. This can be done by calculating the proportion
	// (percentage) of the total GHG emissions (CO 2 e) that is derived using primary data:
	//
	// Formula: Part of PCF based on primary data (CO2e) / PCF (CO2e) = PDS PCF (%)
	//
	// In order for an input to be considered primary data, both the activity and emission factor shall be
	// compliant with the primary data definitions included in Table 5 (see a clarifying example below, Table 8).
	//
	// In order for the upstream emissions’ PDS to be greater than 0, companies would need to request
	// PCFs and their corresponding PDS from their suppliers. Should PDS for relevant components
	// be obtained from upstream suppliers (tier n-1), the total PDS of the PCF should be calculated using
	// a weighted average approach of the material and energy inputs based on their GHG contribution to the
	// studied product’s PCF.
	//
	// To do so, the individual PDSs received from every input supplier (PDSPCFcomponent 1 and
	// PDSPCFcomponent 2) as well as other components, such as energy inputs or direct emissions from
	// production, should be multiplied by their respective relative contribution (in percentage) to the PCF
	// emissions. All weighted PDS components should then be added up to obtain an overarching PDS
	// (PDSPCF product) (Figure 15).
	//
	// Format: Number
	//
	// Optional
	PrimaryDataShare Percentage `json:"primaryDataShare,omitempty"`

	// If present, the Data Quality Indicators (dqi) in accordance with the Pathfinder Framework Sections 4.2.1 and 4.2.3, Appendix B.
	// For reporting periods ending before the beginning of year 2025, at least property primaryDataShare or propery dqi MUST be defined.
	// For reporting periods including the beginning of year 2025 or after, this property MUST be defined.
	//
	// Optional but mandatory from 2025
	Dqi DataQualityIndicators `json:"dqi,omitempty"`

	// If present, the Assurance information in accordance with the Pathfinder Framework.
	//
	// Optional
	Assurance Assurance `json:"assurance,omitempty"`
}
