package main

import "testing"

func Test_updateMessage(t *testing.T) {

	wg.Add(2)
	go updateMessage("Goodbye! cruel world")
	go updateMessage("x")
	wg.Wait()

	if msg != "Goodbye! cruel world" {
		t.Error("incorrect value in msg")
	}

}
