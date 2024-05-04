package main

import "sandbox_golang/experiments/0001/golang"

func main() {
	conf := golang.Config{
		MutantsDir:       "./gambit_out/mutants",
		OriginalFilePath: "src/Counter.sol",
		OriginalFileName: "Counter.sol",
		BackupFile:       "src/Backup.sol",
	}

	golang.Slow(conf)
}
