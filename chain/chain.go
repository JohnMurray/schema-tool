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
	"bufio"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"
)

// Reference that uniquely identifies a type
type AlterRef string

// Direction of either Up or Down that alter can represent
type Direction int

const (
	Up Direction = iota
	Down
)

type Alter struct {
	Ref       AlterRef
	BackRef   AlterRef
	Direction Direction
	FileName  string
}

type AlterGroup struct {
	Up   *Alter
	Down *Alter
}

// Scan a given directory and return a mapping of AlterRef to AlterGroup
// objects. The objects returned are un-validated aside from meta-data
// parsing.
func ScanDirectory(dir string) (map[AlterRef]AlterGroup, error) {
	stat, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !stat.IsDir() {
		return nil, errors.New(fmt.Sprintf("Path '%s' is not a directory", dir))
	}

	files, err := ioutil.ReadDir(dir)
	for _, f := range files {
		if f.IsDir() {
			// only process top-level of dir
			continue
		}
		if isAlterFile(f.Name()) {
			fmt.Printf("Filename: %s\n", f.Name())
			// todo: use path-like util to concat these
			filePath := dir + "/" + f.Name()

			if header, err := readHeader(dir + "/" + f.Name()); err != nil {
				return nil, err
			} else {
				if _, err = parseMeta(header, filePath); err != nil {
					return nil, err
				}
			}
		}
	}

	return nil, nil
}

// Check if the file is an "alter" by seeing if the name confirms to
// what we expect.
func isAlterFile(name string) bool {
	var filenameRegex = regexp.MustCompile(`^(\d+)-([^-]+-)+(up|down).sql$`)
	return filenameRegex.MatchString(name)
}

// Read the first N lines of an alter file that represent the "header." This is
// the bit of stuff that contains all the meta-data required in alters.
func readHeader(filePath string) ([]string, error) {
	var headerRegex = regexp.MustCompile(`^--`)
	lines := make([]string, 256)

	if file, err := os.Open(filePath); err != nil {
		return lines, err
	} else {
		// clone file after we return
		defer file.Close()

		// read line by line
		scanner := bufio.NewScanner(file)
		i := 0
		for scanner.Scan() {
			if i == 256 {
				return lines, errors.New(`Header lines (continuous block of lines starting with '--')
				exceeds 256. Please add a blank line in-between the meta-data and any
				comment lines that may follow.`)
			}
			line := scanner.Text()
			if headerRegex.MatchString(line) {
				lines[i] = line
				i++
			} else {
				// hit non-header line, we're done
				return lines, nil
			}
		}

		if err = scanner.Err(); err != nil {
			return lines, err
		}
	}

	return lines, nil
}

// Parse the meta-information from the file and return an Alter object.
// Returns error if meta cannot be obtained or required information is
// missing.
func parseMeta(lines []string, filePath string) (*Alter, error) {
	// expect meta-lines to be single-line and in the form of
	//   "-- key: value"
	// regex checks for extraneous whitespace
	var metaEntryRegex = regexp.MustCompile(`^--\s*([^\s]+)\s*:(.+)\s*$`)

	var alter = new(Alter)

	for _, line := range lines {
		if matches := metaEntryRegex.FindStringSubmatch(line); len(matches) == 3 {
			// 3 matches means we're good to go
			key := strings.ToLower(strings.TrimSpace(matches[1]))
			value := strings.TrimSpace(matches[2])

			fmt.Printf("%s:    %s\n", key, value)

			switch key {
			case "ref":
				if !isValidRef(value) {
					return nil, errors.New("Invalid 'ref' value found in " + filePath)
				}
				alter.Ref = AlterRef(value)
			case "backref":
				if !isValidRef(value) {
					return nil, errors.New(fmt.Sprintf("Invalid 'backref' value found in '%s'", filePath))
				}
				alter.BackRef = AlterRef(value)
			case "direction":
				value_lower := strings.ToLower(value)
				if value_lower == "up" {
					alter.Direction = Up
				} else if value_lower == "down" {
					alter.Direction = Down
				} else {
					return nil, errors.New(fmt.Sprintf("Invalid direction '%s' found in '%s'", value_lower, filePath))
				}
			case "require-env":
			case "skip-env":
			default:
				// TODO: warn of unknown meta-data property
			}
		}
	}

	// TODO: validate all required meta-data fields are present

	return alter, nil
}

// Validate that the ref is a valid identifier
func isValidRef(ref string) bool {
	var refRegex = regexp.MustCompile(`^[\da-zA-Z]+$`)
	return refRegex.MatchString(ref)
}
