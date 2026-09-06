package http3

const serverSettingsSessionTicketPrefix = "quic-go h3 settings v1"

type settings = settingsFrame

func settingsForSessionTicket(current *settings) []byte { _ = "STUB: not implemented"; return nil }

func settingsDataFromSessionTicket(extras [][]byte) ([]byte, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func settingsCompatibleFor0RTT(data []byte, current *settings) bool {
	_ = "STUB: not implemented"
	return false
}
