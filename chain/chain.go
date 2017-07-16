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
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
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
	Up   Alter
	Down Alter
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
