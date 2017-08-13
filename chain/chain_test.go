// <--
// Copyright © 2017 John Murray <me@johnmurray.io>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// -->
package chain

import (
	"fmt"
	"os"
	"testing"
)

//------------------------------------------------------------------------------
// ScanDirectory tests
//------------------------------------------------------------------------------
func TestScanNonExistantDir(t *testing.T) {
	if _, err := ScanDirectory("/dev/null/abcdefg"); err == nil {
		t.Fail()
	}
}

func TestScanNonDirFile(t *testing.T) {
	if _, err := ScanDirectory("/dev/null"); err == nil {
		t.Fail()
	}
}

//------------------------------------------------------------------------------
// readHeader tests
//------------------------------------------------------------------------------
func TestReadHeaderNormalFile(t *testing.T) {
	dir, _ := os.Getwd()
	alter := dir + "/../test/chains/normal-chain/1234-init-up.sql"
	header, err := readHeader(alter)

	if err != nil {
		t.Fail()
	}

	// check that it stopped after reading first two lines
	if header[0] == "" || header[1] == "" {
		t.Fail()
	}
	if header[2] != "" {
		t.Fail()
	}
}

func TestReadHeaderNormalFileHeaderOnly(t *testing.T) {
	dir, _ := os.Getwd()
	alter := dir + "/../test/chains/normal-chain/1234-init-down.sql"
	header, err := readHeader(alter)

	if err != nil {
		t.Fail()
	}

	// check that it stopped after reading first two lines
	if header[0] == "" || header[1] == "" {
		t.Fail()
	}
	if header[2] != "" {
		t.Fail()
	}
}

func TestReaderHeaderNonExistantFile(t *testing.T) {
	// TODO: use custom error types to check the type of error returned
	dir, _ := os.Getwd()
	alter := dir + "/../test/chains/normal-chain/1234-dont-exist.up"
	if _, err := readHeader(alter); err == nil {
		t.Fail()
	}
}

func TestHeaderTooLarge(t *testing.T) {
	// TODO: use custom error types to check the type of error returned
	dir, _ := os.Getwd()
	alter := dir + "/../test/chains/invalid-headers/1234-init-up.sql"
	if _, err := readHeader(alter); err == nil {
		t.Fail()
	}
}

//------------------------------------------------------------------------------
// isAlterFile tests
//------------------------------------------------------------------------------
func TestAlterFilenameCheck(t *testing.T) {
	// positive assertions
	if !isAlterFile("1234-ABC-1234-some-update-up-down-blah-up.sql") {
		t.Fail()
	}
	if !isAlterFile("1234-ABC-1234-some-update-up-down-blah-down.sql") {
		t.Fail()
	}
	if !isAlterFile("1234-short-up.sql") {
		t.Fail()
	}
	if !isAlterFile("1234-i.has.dots-up.sql") {
		t.Fail()
	}

	// negative assertions
	if isAlterFile("1234-ABC-1234-some-uprade-up.sql.bak") {
		t.Fail()
	}
	if isAlterFile("ABC-1234-some-uprade-up.sql") {
		t.Fail()
	}
	if isAlterFile("1234-up.sql") {
		t.Fail()
	}
	if isAlterFile("1234-down.sql") {
		t.Fail()
	}
}

//------------------------------------------------------------------------------
// parseMeta tests
//------------------------------------------------------------------------------

//------------------------------------------------------------------------------
// isValidRef tests
//------------------------------------------------------------------------------

func TestIsValidRef(t *testing.T) {
	var testData = []*struct {
		ref   string
		valid bool
	}{
		{ref: "hello", valid: true},
		{ref: "1234567890", valid: true},
		{ref: "abc1234def", valid: true},
		{ref: "", valid: false},
		{ref: " 1234 ", valid: false},
	}

	for _, test := range testData {
		if isValidRef(test.ref) != test.valid {
			if test.valid {
				fmt.Printf("Failed to accept valid ref: '%s'\n", test.ref)
			}
			t.Fail()
		}
	}
}
