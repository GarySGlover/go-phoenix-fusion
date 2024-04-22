// Package joystick implements joystick support for Phoenix Fusion.
package joystick

import "github.com/GarySGlover/go-phoenix-fusion/pkg/matrix"

// FourWay represents a digital joystick control with four directions.
type FourWay struct {
	Col, Up, Down, Left, Right, Push int8
}

// EightWay represents a digital joystick control with eight directions.
type EightWay struct {
	Col, Up, Down, Left, Right, Push, UpRight, DownRight, DownLeft, UpLeft int8
}

// NewFourWay creates a new FourWay joystick.
func NewFouwWay(col, up, down, left, right, push int8) FourWay {
	return FourWay{col, up, down, left, right, push}
}

// NewEightWay creates a new EightWay joystick.
func NewEightWay(col, up, down, left, right, push, upRight, downRight, downLeft, upLeft int8) EightWay {
	return EightWay{col, up, down, left, right, push, upRight, downRight, downLeft, upLeft}
}

// Extend returns an EightWay joystick.
// The returned joystick is a extended with additional combination directions.
func (j FourWay) Extend(upRight, downRight, downLeft, upLeft int8) EightWay {
	return EightWay{j.Col, j.Up, j.Down, j.Left, j.Right, j.Push, upRight, downRight, downLeft, upLeft}
}

// NormaliseScanColumn returns a new copy of ScanColumn.
// The returned ScanColumn is normalised to ensure that Push is pressed when no other button is pressed.
func (j FourWay) NormaliseScanColumn(line matrix.ScanColumn) matrix.ScanColumn {
	if line == nil {
		return line
	}

	if line[j.Up] || line[j.Down] || line[j.Left] || line[j.Right] {
		line = line.Release(j.Push)
	}

	return line
}

// NormaliseState returns a new ScanColumnSequence.
func NormaliseState(scan matrix.ScanColumnSequence, joys []FourWay) matrix.ScanColumnSequence {
	for _, joy := range joys {
		scan[joy.Col] = joy.NormaliseScanColumn(scan[joy.Col])
	}

	return scan
}
