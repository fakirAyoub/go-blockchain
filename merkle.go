package main

import (
	"bytes"
	"crypto/sha1"
	"encoding/binary"
	"fmt"

	uuid "github.com/nu7hatch/gouuid"
)

type User struct {
	address    string
	name       string
	coinAmount uint64
}

type Transaction struct {
	sender   User
	receiver User
	amount   uint64
}

func getString(b []byte) string {
	return string(b)
}

func newUser(_name string, _coinAmount uint64) *User {
	u4, err := uuid.NewV4()
	if err != nil {
		fmt.Println("error:", err)

	}

	var bin_buf bytes.Buffer
	binary.Write(&bin_buf, binary.BigEndian, u4)
	tobytes := sha1.Sum(bin_buf.Bytes())
	fmt.Printf("%x", tobytes)
	add := getString(tobytes[:])
	fmt.Printf("%x\n", add)

	//finalTrimmedRandString := strings.Replace(toStrUUID, "-", "", -1)
	user := User{address: add, name: _name, coinAmount: _coinAmount}
	return &user
}

func newTransaction(_sender *User, _receiver *User, _amount uint64) (*Transaction, string, *User, *User) {
	if _sender.coinAmount >= _amount {
		transaction := Transaction{sender: *_sender, receiver: *_receiver, amount: _amount}
		_sender.coinAmount -= _amount
		_receiver.coinAmount += _amount
		return &transaction, "", _sender, _receiver
	} else {
		print("Not enough funds.")
		return nil, "Not enough funds.", _sender, _receiver
	}
}

func merkleTree(transactions []Transaction) {
	if len(transactions)%2 == 0 {

	}
}

func main() {
	fmt.Println("bonjour")
	var ledger []*Transaction

	usr1 := newUser("Helder", 132)
	usr2 := newUser("Ayoub le bg", 1000)

	tr1, err, usr1, usr2 := newTransaction(usr1, usr2, 50)
	print(err)
	ledger = append(ledger, tr1)

	print(usr1.name)
	print(len(ledger))
}
