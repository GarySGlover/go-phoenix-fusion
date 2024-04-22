package matrix_test

import (
	"slices"
	"testing"

	"github.com/GarySGlover/go-phoenix-fusion/pkg/joystick"
	"github.com/GarySGlover/go-phoenix-fusion/pkg/matrix"
)

var joy = joystick.NewFouwWay(0, 0, 2, 3, 1, 4)

func TestEnableSwitchesTable(t *testing.T) {
	t.Parallel()

	scan := matrix.NewScanColumn([]bool{false, false, false, false, false})

	var empty matrix.ScanColumn
	tests := []struct {
		name      string
		button    string
		initState matrix.ScanColumn
		enable    int8
	}{
		{"up on existing should be true", "up", scan, joy.Up},
		{"down on existing should be true", "down", scan, joy.Down},
		{"left on existing should be true", "left", scan, joy.Left},
		{"right on existing should be true", "right", scan, joy.Right},
		{"push on existing should be true", "push", scan, joy.Push},
		{"up on new should be true", "up", empty, joy.Up},
		{"down on new should be true", "down", empty, joy.Down},
		{"left on new should be true", "left", empty, joy.Left},
		{"right on new should be true", "right", empty, joy.Right},
		{"push on new should be true", "push", empty, joy.Push},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ans := test.initState.Press(test.enable)
			if ans[test.enable] != true {
				t.Errorf("%s not pressed", test.button)
			}
		})
	}

	if slices.Equal(scan, empty) {
		t.Error("scan and empty initState shouldn't match")
	}
}
