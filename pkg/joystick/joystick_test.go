package joystick_test

import (
	"reflect"
	"testing"

	"github.com/GarySGlover/go-phoenix-fusion/pkg/joystick"
	"github.com/GarySGlover/go-phoenix-fusion/pkg/matrix"
)

var joy = joystick.NewFouwWay(0, 0, 2, 3, 1, 4)

func TestNormaliseScanColumnTable(t *testing.T) {
	t.Parallel()

	line := matrix.NewScanColumn([]bool{false, false, false, false, false})

	tests := []struct {
		name   string
		button string
		enable int8
	}{
		{"up forces push disabled", "up", joy.Up},
		{"down forces push disabled", "down", joy.Down},
		{"left forces push disabled", "left", joy.Left},
		{"right forces push disabled", "right", joy.Right},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			ans := joy.NormaliseScanColumn(line.PressSeq([]int8{test.enable, joy.Push}))
			if ans[joy.Push] == true {
				t.Errorf("push should be disabled when %s pressed", test.button)
			}

			line.Release(test.enable)
		})
	}
}

func TestNormaliseScanState(t *testing.T) {
	t.Parallel()

	unpushed := matrix.NewScanColumn([]bool{false, false, false, false, false})
	pushed := unpushed.Press(joy.Push)
	scan := matrix.ScanColumnSequence{
		pushed,
		pushed.Press(joy.Up),
		pushed.Press(joy.Down),
		pushed.Press(joy.Left),
		pushed.Press(joy.Right),
	}
	expected := matrix.ScanColumnSequence{
		pushed,
		unpushed.Press(joy.Up),
		unpushed.Press(joy.Down),
		unpushed.Press(joy.Left),
		unpushed.Press(joy.Right),
	}
	joys := make([]joystick.FourWay, 5)

	for i := range joys {
		joys[i] = joystick.NewFouwWay(int8(i), 0, 2, 3, 1, 4)
	}

	result := joystick.NormaliseState(scan, joys)

	if !reflect.DeepEqual(expected, result) {
		t.Errorf("Expected normalised state to be %v, got %v", expected, result)
	}
}
