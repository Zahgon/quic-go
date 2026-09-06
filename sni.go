package quic

const (
	extTypeSNI = 0
	extTypeECH = 0xfe0d
)

func findSNIAndECH(data []byte) (sniPos, sniLen, echPos int, err error) {
	_ = "STUB: not implemented"
	return 0, 0, 0, nil
}
