package main

import (
	"flag"
	"fmt"
	"github.com/go-mysql-org/go-mysql/replication"
	"os"
)

var (
	filePath string
)

func onParse(b *replication.BinlogEvent) error {
	b.Dump(os.Stdout)
	return nil
}

func main() {
	flag.StringVar(&filePath, "file", "", "file path to binlog")
	flag.Parse()

	if filePath == "" {
		fmt.Println("[ERROR] -file option is not specified")
		os.Exit(1)
	}

	parser := replication.NewBinlogParser()
	parser.ParseFile(filePath, 0, replication.OnEventFunc(onParse))
}
