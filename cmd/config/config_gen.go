package main

import "pulsar_mate_go/pkg/pulsar"

func main() {
	err := pulsar.Config()
	if err != nil {
		panic(err)
	}
}
