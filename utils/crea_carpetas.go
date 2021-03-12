package main

import (
	"log"
	"os"
	"time"
)

func main() {
	err := os.Mkdir("/Users/fernando/workspace_go/go_examples/utils/test", 0755)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(2 * time.Second)
	os.Remove("/Users/fernando/workspace_go/go_examples/utils/test")

}
