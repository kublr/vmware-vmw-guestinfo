package bdoor

const (
	BackdoorMagic = uint64(0x564D5868)
)

type BackdoorProto struct {
	AX, BX, CX, DX, SI, DI, BP UInt64
	size                       uint32
}

func bdoor_inout(ax, bx, cx, dx, si, di, bp uint64) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint64) {
	panic("VMWare on arm64 is not supported")
}
func bdoor_hbout(ax, bx, cx, dx, si, di, bp uint64) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint64) {
	panic("VMWare on arm64 is not supported")
}
func bdoor_hbin(ax, bx, cx, dx, si, di, bp uint64) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint64) {
	panic("VMWare on arm64 is not supported")
}
func bdoor_inout_test(ax, bx, cx, dx, si, di, bp uint64) (retax, retbx, retcx, retdx, retsi, retdi, retbp uint64) {
	panic("VMWare on arm64 is not supported")
}
