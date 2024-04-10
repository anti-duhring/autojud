package processes

func getCourtFromString(court string) Court {
	switch court {
	case "TJPE":
		return COURT_TJPE
	default:
		return COURT_UNKNOWN
	}
}
