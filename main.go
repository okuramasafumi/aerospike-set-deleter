// Package main
package main

import (
	"flag"
	as "github.com/aerospike/aerospike-client-go"
	"log"
	"os"
)

func main() {
	var (
		host      string
		port      int
		namespace string
		set       string

		count int
	)
	flag.StringVar(&host, "host", "127.0.0.1", "host of Aerospike")
	flag.IntVar(&port, "port", 3000, "port of Aerospike")
	flag.StringVar(&namespace, "namespace", "test", "namespace of Aerospike to which the set you want to delete belongs")
	flag.StringVar(&set, "set", "", "set of Aerospike you want to delete")
	flag.Parse()

	log.SetOutput(os.Stdout)

	if flag.NFlag() == 0 {
		printUsage()
		return
	}

	log.Printf("Host: %v", host)
	log.Printf("Port: %v", port)
	log.Printf("Namespace: %v", namespace)
	log.Printf("Set: %v", set)

	if set == "" {
		log.Fatalln("You must provide set")
	}

	client, err := as.NewClient(host, port)
	if err != nil {
		log.Fatalln(err)
	}
	scanPolicy := as.NewScanPolicy()
	scanPolicy.IncludeBinData = false

	rs, err := client.ScanAll(scanPolicy, namespace, set)
	if err != nil {
		log.Fatalln(err)
	}
	for result := range rs.Results() {
		if result.Err != nil {
			log.Fatalln(result.Err)
		}
		success, err := client.Delete(nil, result.Record.Key)
		if err != nil {
			log.Fatalln(err)
		}
		if success {
			count++
		}
	}

	log.Printf("Deleted %v records from set %v", count, set)
}

func printUsage() {
	flag.Usage()
}
