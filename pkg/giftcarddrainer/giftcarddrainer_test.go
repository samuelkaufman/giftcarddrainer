package giftcarddrainer

import (
	"bytes"
	"testing"
)

func TestGiftCardDrainer(t *testing.T) {
	var testInput = `
Candy Bar, 500
Paperback Book, 700
Detergent, 1000
Headphones, 1400
Earmuffs, 2000
Bluetooth Stereo, 6000
`
	drainer := New(bytes.NewReader([]byte(testInput)), 2500)
	bestPair := drainer.Run()
	if bestPair[0].Id != "Candy Bar" || bestPair[1].Id != "Earmuffs" {
		t.Error("failed to get correct pair")
	}
	drainer = New(bytes.NewReader([]byte(testInput)), 2300)
	bestPair = drainer.Run()
	if bestPair[0].Id != "Paperback Book" || bestPair[1].Id != "Headphones" {
		t.Error("failed to get correct pair")
	}
	drainer = New(bytes.NewReader([]byte(testInput)), 10000)
	bestPair = drainer.Run()
	if bestPair[0].Id != "Earmuffs" || bestPair[1].Id != "Bluetooth Stereo" {
		t.Error("failed to get correct pair")
	}

	drainer = New(bytes.NewReader([]byte(testInput)), 1100)
	bestPair = drainer.Run()
	if bestPair[0] != nil {
		t.Error("bestPair should not be defined")
	}
}
