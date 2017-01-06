package message

import (
	"testing"
)

func TestSetWord(t *testing.T) {
	inLow := uint16(0xEEFF)
	inHigh := uint16(0xBBBB)

	out := &UInt32{}
	//out.SetWord(uint32(0xBBBBEEFF))
	out.Low = inLow
	out.High = inHigh

	if !AssertEqual(t, inLow, out.Low) || !AssertEqual(t, inHigh, out.High) {
		return
	}

	if !AssertEqual(t, uint32(0xBBBBEEFF), out.Word()) {
		return
	}
}

func TestQuadToHighLow(t *testing.T) {
	in := uint64(0xFFFFFFFF0000000A)

	var u UInt64
	u.SetQuad(in)
	if !AssertEqual(t, uint32(in), u.Low.Word()) {
		return
	}

	if !AssertEqual(t, uint32(in>>32), u.High.Word()) {
		return
	}

	if !AssertEqual(t, in, u.Quad()) {
		return
	}
}

func TestHighLowToQuad(t *testing.T) {
	inHigh := uint16(0xff)
	inLow := uint16(0xaa)

	u := UInt64{
		High: UInt32{High: inHigh},
		Low:  UInt32{Low: inLow},
	}

	if !AssertEqual(t, (uint64(inHigh)<<48)+uint64(inLow), u.Quad()) {
		return
	}
}
