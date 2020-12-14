package main

import (
	"bytes"
	"crypto"
	"crypto/sha1"
	"encoding/binary"
	"fmt"
	"reflect"

	uuid "github.com/nu7hatch/gouuid"
)

type User struct {
	publicAddress  string
	privateAddress string
	name           string
	coinAmount     uint64
}

type Transaction struct {
	hash     []byte
	sender   User
	receiver User
	amount   uint64
}

type Bloc struct {
	hash         []byte
	transactions []Transaction
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
	add1 := getString(tobytes[:])
	fmt.Printf("%x", add1)

	var bin_buf2 bytes.Buffer
	binary.Write(&bin_buf2, binary.BigEndian, u4)
	tobytes2 := sha1.Sum(bin_buf2.Bytes())
	fmt.Printf("%x", tobytes2)
	add2 := getString(tobytes2[:])
	fmt.Printf("%x", add2)

	user := User{publicAddress: add1, privateAddress: add2, name: _name, coinAmount: _coinAmount}
	return &user
}

func Hash(objs ...interface{}) []byte {
	digester := crypto.SHA1.New()
	for _, ob := range objs {
		fmt.Fprint(digester, reflect.TypeOf(ob))
		fmt.Fprint(digester, ob)
	}
	return digester.Sum(nil)
}

func newTransaction(_sender *User, _receiver *User, _amount uint64) (*Transaction, string, *User, *User) {
	if _sender.coinAmount >= _amount {
		fmt.Printf("Transaction Hash : %x\n", Hash(_sender.privateAddress, _receiver.privateAddress))
		var transactionHash = Hash(_sender.privateAddress, _receiver.privateAddress)
		_sender.coinAmount -= _amount
		_receiver.coinAmount += _amount
		transaction := Transaction{hash: transactionHash, sender: *_sender, receiver: *_receiver, amount: _amount}

		return &transaction, "", _sender, _receiver
	} else {
		print("Not enough funds.")
		return nil, "Not enough funds.", _sender, _receiver
	}
}

func newBloc(_transactions []Transaction) Bloc {
	// Faire merkel puis rajouter le hash
	bloc := Bloc{hash: hash, transactions: _transactions}
	return bloc
}

func merkleTree(transactions []Transaction) {
	if len(transactions)%2 == 0 && (len(transactions)/2)%2 == 0 {

	}
}

func main() {
	var ledger []*Transaction

	usr1 := newUser("Helder", 132)
	print("\n")
	usr2 := newUser("Ayoub le bg", 1000)
	print("\n")
	tr1, err, usr1, usr2 := newTransaction(usr1, usr2, 50)
	tr2, err, usr1, usr2 := newTransaction(usr1, usr2, 50)
	print("\n")
	print(err)
	print("\n")
	ledger = append(ledger, tr1)
	ledger = append(ledger, tr2)
	print("\n")
	print(usr2.coinAmount)
	print("\n")
	print(len(ledger))
	print("\n")
	print(usr1.coinAmount)
	print("\n")
	//bloc1 =
}
