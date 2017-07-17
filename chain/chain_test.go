package chain

import (
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
