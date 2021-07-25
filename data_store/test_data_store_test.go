package data_store

import (
	"fmt"
	"os"
	"testing"
)

const testStorage = "./testStorage"

var testDataStore *DataStore

func TestMain(m *testing.M) {
	fmt.Println("Testing data store module")
	var storage = testStorage
	m.Run()
}

// Refactored code, so tests will have to be changed.
