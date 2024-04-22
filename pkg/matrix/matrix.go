package matrix

type (
	ScanColumn         []bool
	ScanColumnSequence []ScanColumn
)

func NewScanColumn(data []bool) ScanColumn {
	s := make(ScanColumn, len(data))
	_ = copy(s, data)

	return s
}

func (scanColumn ScanColumn) with(pos int8, value bool) ScanColumn {
	var scan ScanColumn
	if len(scanColumn) <= int(pos) || scanColumn == nil {
		scan = make(ScanColumn, int(pos+1))
		_ = copy(scan, scanColumn)
	} else {
		scan = scanColumn
	}

	scan[pos] = value

	return scan
}

func (scanColumn ScanColumn) Press(pos int8) ScanColumn {
	return scanColumn.with(pos, true)
}

func (scanColumn ScanColumn) Release(pos int8) ScanColumn {
	return scanColumn.with(pos, false)
}

func (scanColumn ScanColumn) withSeq(pos []int8, value bool) ScanColumn {
	s := scanColumn
	for _, p := range pos {
		s = scanColumn.with(p, value)
	}

	return s
}

func (scanColumn ScanColumn) PressSeq(pos []int8) ScanColumn {
	return scanColumn.withSeq(pos, true)
}

func (scanColumn ScanColumn) ReleaseSeq(pos []int8) ScanColumn {
	return scanColumn.withSeq(pos, false)
}
