package schema

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testUnit struct {
	Unit DeclaredUnit `json:"unit"`
}

func TestDeclaredUnitParseLiter(t *testing.T) {

	testData := `{
		"unit": "liter"
	}`

	var unit testUnit
	err := json.Unmarshal([]byte(testData), &unit)
	assert.Nil(t, err)

	assert.Equal(t, unit.Unit, Liter)

}

func TestDeclaredUnitParseKilo(t *testing.T) {

	testData := `{
		"unit": "kilogram"
	}`

	var unit testUnit
	err := json.Unmarshal([]byte(testData), &unit)
	assert.Nil(t, err)

	assert.Equal(t, unit.Unit, KiloGram)

}

func TestDeclaredUnitParseCubicMeter(t *testing.T) {

	testData := `{
		"unit": "cubic meter"
	}`

	var unit testUnit
	err := json.Unmarshal([]byte(testData), &unit)
	assert.Nil(t, err)

	assert.Equal(t, unit.Unit, CubicMeter)

}

func TestDeclaredUnitParseKWh(t *testing.T) {

	testData := `{
		"unit": "kilowatt hour"
	}`

	var unit testUnit
	err := json.Unmarshal([]byte(testData), &unit)
	assert.Nil(t, err)

	assert.Equal(t, unit.Unit, KiloWattHour)

}

func TestDeclaredUnitParseMegajoule(t *testing.T) {

	testData := `{
		"unit": "megajoule"
	}`

	var unit testUnit
	err := json.Unmarshal([]byte(testData), &unit)
	assert.Nil(t, err)

	assert.Equal(t, unit.Unit, MegaJoule)

}

func TestDeclaredUnitParseTonKm(t *testing.T) {

	testData := `{
		"unit": "ton kilometer"
	}`

	var unit testUnit
	err := json.Unmarshal([]byte(testData), &unit)
	assert.Nil(t, err)

	assert.Equal(t, unit.Unit, TonKilometer)

}

func TestDeclaredUnitParseToSqm(t *testing.T) {

	testData := `{
		"unit": "square meter"
	}`

	var unit testUnit
	err := json.Unmarshal([]byte(testData), &unit)
	assert.Nil(t, err)

	assert.Equal(t, unit.Unit, SquareMeter)

}
