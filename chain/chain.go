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
			continue
		}
		if isAlterFile(f.Name()) {
			fmt.Printf("Filename: %s\n", f.Name())
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
func readHeader(filePath string) ([256]string, error) {
	var headerRegex = regexp.MustCompile(`^--`)
	var lines [256]string

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
func parseMeta(lines [256]string, filePath string) (*Alter, error) {
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

			switch key {
			case "ref":
			case "backref":
			case "direction":
			case "require-env":
			case "skip-env":
			default:
				// TODO: warn of unknown meta-data property
			}
		}
	}

	return alter, nil
}
