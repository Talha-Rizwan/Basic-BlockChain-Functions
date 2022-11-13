package assignment02

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"crypto/sha256"
)
var Transection_id_counter = 0



type Transaction struct {
	TransactionID string
	Sender        string
	Receiver      string
	Amount        int
}

type Block struct {
	Nonce int
	BlockData        []Transaction
	PrevPointer *Block
	PrevHash    string
	CurrentHash string
}

type Blockchain struct {
	ChainHead *Block
}

func GenerateNonce(blockData []Transaction) int {
	fmt.Println(blockData)

	return rand.Intn(100)
}

func CalculateHash(blockData []Transaction, nonce int) string {
	dataString := ""
	for i := 0; i < len(blockData); i++ {
		dataString += (blockData[i].TransactionID + blockData[i].Sender +
		blockData[i].Receiver + strconv.Itoa(blockData[i].Amount)) + strconv.Itoa(nonce)
	}
	return fmt.Sprintf("%x", sha256.Sum256([]byte(dataString)))
}

func NewBlock(blockData []Transaction, chainHead *Block) *Block {
	b := new(Block)
	b.Nonce = GenerateNonce(blockData)
	b.BlockData = blockData
	b.PrevPointer = chainHead
	b.CurrentHash = CalculateHash(b.BlockData, b.Nonce)

	if chainHead==nil {
		b.PrevHash = ""

	}else {
		b.PrevHash=chainHead.CurrentHash
	}

	return b
}

func ListBlocks(chainHead *Block) {
	
	fmt.Println("displaying the complete blockchian")


	currentNode := chainHead
	if currentNode == nil {
		fmt.Println("nothing to display")
		return
	}
	fmt.Println( strings.Repeat("=", 25))
	fmt.Printf("%+v\n", *currentNode)
	for currentNode.PrevPointer != nil {
		currentNode = currentNode.PrevPointer
		fmt.Println( strings.Repeat("=", 25))	
		fmt.Printf("%+v\n", *currentNode)
	}
}

func DisplayTransactions(blockData []Transaction) {
	for index, val := range blockData {
		fmt.Printf("%s Transection :  %d %s\n", strings.Repeat("=", 25), index, strings.Repeat("=", 25))

		fmt.Printf(" Transection id: %s \n Sender :  %s \n Receiver :  %s \n Amount:  %d \n \n ",val.TransactionID,val.Sender,val.Receiver,val.Amount)
	}
}

func NewTransaction(sender string, receiver string, amount int) Transaction {
	Transection_id_counter++
	b := new(Transaction)
	b.TransactionID = strconv.Itoa(Transection_id_counter)
	b.Sender = sender
	b.Receiver = receiver
	b.Amount = amount
return *b
}

