package main

import (
	log "github.com/sirupsen/logrus"
)

func main() {
	log.WithFields(log.Fields{
		"animal": "Walrus",
	}).Info("Testing BLAHBLAH")
}
