package main

type RollerChangeType int64

const (
	New RollerChangeType = iota
	Changed
	Deleted
)

type RollerChange struct {
	change   RollerChangeType
	srcPath  string
	destPath string
	apply    bool
}
