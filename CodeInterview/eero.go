/*
Can you help us write code to select the best wifi channel for a network to use?
We have provided the code below which includes simple implementations of most of the helper functions;
you should fill in the function ChannelSelection() to tie them all together to calculate the clearest channel and send that to the rest of the network.
Please be conscientious of handling any error that could occur (based on the function signatures, not the given simple implementations)
Our goal is to be able to go from scanning the available channels to broadcasting the new channel choice to the network in less than 250 milliseconds.

Starter code*/

package main

import (
	"fmt"
	"time"
)

type WiFiChannel int

const goal = 250 * time.Millisecond

func main() {
	fmt.Println("Starting")
	ChannelSelection()
}

func ChannelSelection() {

	start := time.Now()

	// TODO: select the best channel based on the current and historical
	// scan data. Once the channel is selected, we should broadcast it to
	// to the rest of the network (so everyone knows the new channel) and
	// store the process scanned data, so we can use it next time.

	scan := ScanChannels()
	data, err := LoadHistoricalScanData()
	if err != nil {
		fmt.Println("Cannot load historical data...")
		return
	}

	scanResult := ProcessScan(scan)
	bestCh := PickBestChannel(scanResult, data)

	err = BroadcastNewChannel(bestCh)
	if err != nil {
		fmt.Println("Cannot broadcast for new channel...")
		return
	}
	err = StoreScanResults(scanResult)
	if err != nil {
		fmt.Println("Cannot store scan result...")
		return

	}

	if duration := time.Since(start); duration <= goal {

		fmt.Println("Good job! Processing scan took", duration)

	} else {

		fmt.Println("Can you make it faster? Processing scan took", duration)

	}

}

// ScanChannels scans a portion of the wifi channel space and returns a raw

// dump of the captured data.

func ScanChannels() []byte {
	return []byte{0x2a, 0x42, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06}

}

// LoadHistoricalScanData loads the last scan results from our local database.

func LoadHistoricalScanData() (map[WiFiChannel]int, error) {
	time.Sleep(100 * time.Millisecond)
	return map[WiFiChannel]int{
		// WiFiChannel: how good the channel is (higher is better)
		1: 4,
		3: 7, // the best
		6: 2, // the worst
		9: 4,
	}, nil

}

// ProcessScan takes raw scan data and converts it to a map of channel IDs to
// to a quality score.
func ProcessScan(scan []byte) map[WiFiChannel]int {
	time.Sleep(100 * time.Millisecond)
	return map[WiFiChannel]int{
		// WiFiChannel: how good the channel is (higher is better)
		1: 3,
		3: 9,
		6: 10, // the best
		9: 1,  // the worst
	}
}

func PickBestChannel(newData, historicalData map[WiFiChannel]int) WiFiChannel {
	// TODO: Phase 2: implement some selection algorithm here.
	resultMap := make(map[WiFiChannel]int)

	for k, v := range historicalData {
		resultMap[k] = (v + newData[k]) / 2
	}

	best := 0
	var bestCh WiFiChannel
	for k, v := range resultMap {
		if v > best {
			best = v
			bestCh = k
		}
	}

	return bestCh

}

// StoreScanResults stores the given scan results in our local database.

func StoreScanResults(results map[WiFiChannel]int) error {
	time.Sleep(100 * time.Millisecond)
	return nil
}

// BroadcastNewChannel sends the new channel selection to the rest of the network.

func BroadcastNewChannel(newChannel WiFiChannel) error {

	fmt.Println("Broadcasting new channel", newChannel)
	time.Sleep(100 * time.Millisecond)
	return nil
}
