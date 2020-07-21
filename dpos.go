package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)


type Block struct {
	Index     int
	PreHash   string
	HashCode  string
	BMP       int
	validator string
	TimeStamp string
}

var Blockchain []Block

func GenerateNextBlock(oldBlock Block, BMP int, adds string) Block {
	var newBlock Block
	newBlock.Index = oldBlock.Index + 1
	newBlock.PreHash = oldBlock.HashCode
	newBlock.BMP = BMP
	newBlock.TimeStamp = time.Now().String()
	newBlock.validator = adds
	newBlock.HashCode = GenerateHashValue(newBlock)
	return newBlock
}

func GenerateHashValue(block Block) string {
	var hashCode = block.PreHash + block.validator + block.TimeStamp +
		strconv.Itoa(block.Index) + strconv.Itoa(block.BMP)

	var sha = sha256.New()
	sha.Write([]byte(hashCode))
	hashed := sha.Sum(nil)
	return hex.EncodeToString(hashed)
}

var delegate = []string{"aaa", "bbb", "ccc", "dddd"}

func RandDelegate() {
	rand.Seed(time.Now().Unix())
	var r = rand.Intn(3)
	t := delegate[r]
	delegate[r] = delegate[3]
	delegate[3] = t
}

func main() {
	fmt.Println(delegate)

	var firstBlock Block
	Blockchain = append(Blockchain, firstBlock)
	var n = 0
	ch1 := make(chan bool)
	ch2 := make(chan bool)

	go func() {
	flag:
		<-ch1
		count := 0
		for {
			count++
			time.Sleep(time.Second * 3)
			var nextBlock = GenerateNextBlock(firstBlock, 1, delegate[n])
			n++
			n = n % len(delegate)
			firstBlock = nextBlock
			Blockchain = append(Blockchain, nextBlock)
			fmt.Println(Blockchain)
			fmt.Println(count)
			if count == 10 {
				count = 0
				ch2 <- true
				goto flag
			}
		}
	}()

	go func() {
		for {

			RandDelegate()
			fmt.Println("", delegate)
			ch1 <- true
			<-ch2

		}

	}()

	for {

	}
}
