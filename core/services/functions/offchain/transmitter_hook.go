package offchain

type SignedReport struct {
	Report []byte
	Rs     [][32]byte
	Ss     [][32]byte
	Vs     [32]byte
}

type TransmitterHook interface {
	Transmit(signedReport *SignedReport)
	ReportChannel() chan *SignedReport
}

type transmitterHook struct {
	chReports chan *SignedReport
}

func NewTransmitterHook() TransmitterHook {
	return &transmitterHook{chReports: make(chan *SignedReport)}
}

func (t *transmitterHook) Transmit(signedReport *SignedReport) {
	t.chReports <- signedReport
}

func (t *transmitterHook) ReportChannel() chan *SignedReport {
	return t.chReports
}
