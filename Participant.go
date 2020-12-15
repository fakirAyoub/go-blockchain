package main

//participant
type Participant struct {
	publicAddress  string
	privateAddress string
	coinAmont      int
}

func NewParticipant(
	publicAddress string, privateAddress string, coinAmont int) *Participant {
	p := new(Participant)
	p.publicAddress = publicAddress
	p.privateAddress = privateAddress
	p.coinAmont = coinAmont
	return p
}

// mettre sur linux -> ssh keygen pour génère une adresse private et public
