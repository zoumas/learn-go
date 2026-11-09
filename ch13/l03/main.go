package main

import "fmt"

func waitForDBs(numDBs int, dbChan chan struct{}) {
	for range numDBs {
		<-dbChan
	}
}

// don't touch below this line

func getDBsChannel(numDBs int) (chan struct{}, *int) {
	count := 0
	ch := make(chan struct{})

	go func() {
		for i := range numDBs {
			count++ // before the send: the receive makes this write visible to waitForDBs' caller
			ch <- struct{}{}
			fmt.Printf("Database %v is online\n", i+1)
		}
	}()

	return ch, &count
}
