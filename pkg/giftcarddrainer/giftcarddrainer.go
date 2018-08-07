//You have been given a gift card that is about to expire and you want to buy gifts for 2 friends. You want to spend the whole gift card, or if that’s not an option as close to the balance as possible. You have a list of sorted prices for a popular store that you know they both like to shop at. Your challenge is to find two distinct items in the list whose sum is minimally under (or equal to) the gift card balance.
//This approach is O(n!)

package giftcarddrainer

import (
	"encoding/csv"
	"errors"
	"io"
	"log"
	"strconv"
)

var (
	usage = "USAGE: gocarddrainer FILENAME BALANCE"
)

type GiftCardDrainer struct {
	balance   int
	fh        io.ReadSeeker
	csvReader *csv.Reader
	bestPair  [2]*Item
}

func New(fh io.ReadSeeker, balance int) *GiftCardDrainer {
	return &GiftCardDrainer{
		fh:        fh,
		balance:   balance,
		csvReader: csv.NewReader(fh),
	}
}

type Item struct {
	Price int
	Id    string
}

var (
	ErrNoMaxPossible    = errors.New("No max possible")
	ErrInvalidCSV       = errors.New("Invalid CSV Format")
	ErrNextItemNotFound = errors.New("Next item not found")
)

func (g *GiftCardDrainer) newCsvReader() (csvReader *csv.Reader) {
	g.fh.Seek(0, 0)
	csvReader = csv.NewReader(g.fh)
	csvReader.TrimLeadingSpace = true
	csvReader.FieldsPerRecord = 2
	return
}

func recordToItem(record []string, err error) (*Item, error) {
	if err != nil {
		return nil, err
	}
	price, err := strconv.Atoi(record[1])
	if err != nil {
		log.Printf("Incorrectly formatted csv line %+v\n", record)
		return nil, ErrInvalidCSV
	}
	return &Item{
		Price: price,
		Id:    record[0],
	}, nil
}

func scanAndReturnNextItem(csvReader *csv.Reader, curItem *Item) (*Item, error) {

	for item, err := recordToItem(csvReader.Read()); err == nil; item, err = recordToItem(csvReader.Read()) {
		if item.Id == curItem.Id {
			return recordToItem(csvReader.Read())
		}
	}
	return nil, ErrNextItemNotFound
}

func (g *GiftCardDrainer) findMax(curItem *Item, csvReader *csv.Reader) (*Item, error) {
	max := g.balance - curItem.Price
	var maxItem *Item
	for {
		item, err := recordToItem(csvReader.Read())
		if err != nil {
			if err == io.EOF {
				if maxItem != nil {
					return maxItem, nil
				} else {
					return nil, ErrNoMaxPossible
				}
			}
			log.Println(err)
			return nil, err
		}
		if item.Price <= max {
			maxItem = item
			if item.Price == max {
				return maxItem, nil
			}
		} else if item.Price > max {
			if maxItem == nil {
				return nil, ErrNoMaxPossible
			}
			return maxItem, nil
		}
	}
}

func (g *GiftCardDrainer) Run() [2]*Item {
	csvReader := g.newCsvReader()
	var err error
	var testItem *Item
	var curMax int
	testItem, err = recordToItem(csvReader.Read())
	if err != nil {
		log.Fatal(err)
	}
	g.bestPair[0] = testItem
	for {
		maxItem, err := g.findMax(testItem, csvReader)
		if err == ErrNoMaxPossible {
			break
		}
		if err != nil {
			log.Fatal(err)
		} else {
			max := maxItem.Price + testItem.Price
			if max == g.balance {
				return [2]*Item{testItem, maxItem}
			}
			if g.bestPair[1] == nil {
				g.bestPair[1] = maxItem
				curMax = max
			} else if maxItem.Price+testItem.Price > curMax {
				g.bestPair[0] = testItem
				g.bestPair[1] = maxItem
			}
		}
		csvReader = g.newCsvReader()
		testItem, err = scanAndReturnNextItem(csvReader, testItem)
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}
	}
	if g.bestPair[1] == nil {
		return [2]*Item{nil, nil}
	}
	return g.bestPair
}
