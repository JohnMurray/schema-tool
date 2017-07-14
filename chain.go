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
package main

import (
	"os"
)

type AlterRef string
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

func ScanDirectory() (map[AlterRef]AlterGroup, error) {
	_, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	return nil, nil
}

func parseFileName(name string) (*Alter, error) {
	return nil, nil
}
