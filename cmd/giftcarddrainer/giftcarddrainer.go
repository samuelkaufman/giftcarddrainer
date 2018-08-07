//You have been given a gift card that is about to expire and you want to buy gifts for 2 friends. You want to spend the whole gift card, or if that’s not an option as close to the balance as possible. You have a list of sorted prices for a popular store that you know they both like to shop at. Your challenge is to find two distinct items in the list whose sum is minimally under (or equal to) the gift card balance.
//

package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/samuelkaufman/giftcarddrainer/pkg/giftcarddrainer"
)

var (
	usage = "USAGE: gocarddrainer FILENAME BALANCE"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	if len(os.Args) != 3 {
		log.Fatal(usage)
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	balance, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(err)
	}

	drainer := giftcarddrainer.New(f, balance)
	bestPair := drainer.Run()
	if bestPair[1] == nil {
		fmt.Printf("Not possible")
		return
	}
	fmt.Printf("%s %d, %s %d\n", bestPair[0].Id, bestPair[0].Price, bestPair[1].Id, bestPair[1].Price)

}
