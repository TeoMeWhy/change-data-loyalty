package tables

import (
	"fmt"
	"migrator/configs"
	"migrator/models/schemas"
	"os"
	"testing"

	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/assert"
)

type pairTest struct {
	name          string
	expected      []schemas.CustomerSchema
	expectedError error
	value         string
}

var config = &configs.Config{
	QueryBaseFolder: "testdata/query",
	DataFolder:      "testdata/data",
}

func TestReadValidData(t *testing.T) {

	db, err := mockDB()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	tests := []pairTest{
		{
			name:          "Test 1: Read all data",
			expected:      makeTables(),
			expectedError: nil,
			value:         "customer",
		}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			customerTable, err := NewTableRepository[schemas.CustomerSchema](tt.value, "points", db, config)
			if err != nil {
				t.Fatalf("Failed to create table repository: %v", err)
			}

			err = customerTable.ReadData()
			assert.ErrorIs(t, err, tt.expectedError)
			assert.Equal(t, tt.expected, customerTable.Data)
		})
	}

}

func TestReadInvalidData(t *testing.T) {

	db, err := mockDB()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	tests := []pairTest{
		{
			name:          "Test 1: Read not exists data",
			expected:      []schemas.CustomerSchema{},
			expectedError: fmt.Errorf("no such table: not_exists"),
			value:         "not_exists",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			customerTable, err := NewTableRepository[schemas.CustomerSchema](tt.value, "points", db, config)
			if err != nil {
				t.Fatalf("Failed to create table repository: %v", err)
			}

			err = customerTable.ReadData()

			assert.Error(t, err, tt.expectedError)
			assert.Equal(t, tt.expected, customerTable.Data)
		})
	}
}

func TestToString(t *testing.T) {

	db, err := mockDB()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	tests := []struct {
		name          string
		expected      string
		expectedError error
		value         string
	}{
		{
			name:          "Test 1: data to string",
			expected:      makeString(),
			expectedError: nil,
			value:         "customer",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			customerTable, err := NewTableRepository[schemas.CustomerSchema](tt.value, "points", db, config)
			if err != nil {
				t.Fatalf("Failed to create table repository: %v", err)
			}

			err = customerTable.ReadData()
			if err != nil {
				t.Fatalf("Failed to read data: %v", err)
			}

			result := customerTable.ToString()
			assert.Equal(t, tt.expected, result)

		})
	}
}

func TestWriteData(t *testing.T) {

	db, err := mockDB()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}

	os.MkdirAll("testdata", os.ModePerm)

	tests := []struct {
		name          string
		expectedError error
		value         string
	}{
		{
			name:          "Test 1: Write csv file",
			expectedError: nil,
			value:         "customer",
		},
	}

	for _, tt := range tests {

		t.Run(tt.name, func(t *testing.T) {

			customerTable, err := NewTableRepository[schemas.CustomerSchema](tt.value, "points", db, config)
			if err != nil {
				t.Fatalf("Failed to create table repository: %v", err)
			}

			err = customerTable.ReadData()
			if err != nil {
				t.Fatalf("Failed to read data: %v", err)
			}

			err = customerTable.WriteData()
			assert.ErrorIs(t, err, tt.expectedError)

		})
	}

	os.RemoveAll("testdata")

}
