package chain

import (
	"fmt"
	"os"
	"testing"
)

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

func TestReadHeader(t *testing.T) {
	// TODO: finish test
	if dir, err := os.Getwd(); err != nil {
		t.Fail()
	} else {
		fmt.Println(dir + "/../test/chains/")
	}
}

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
