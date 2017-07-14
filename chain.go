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
