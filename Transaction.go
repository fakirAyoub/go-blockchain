package main

//transaction
type Transaction struct {
	ID []byte
	//pointeur vers le sender
	sender []*Participant
	//pointeur vers le receiver
	receiver []*Participantte
	//montant de la transaction
	amount int
}
