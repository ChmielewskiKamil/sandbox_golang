package golang

import (
	"testing"
)

// Benchmark: 14.928s
func BenchmarkSlow(b *testing.B) {
	conf := Config{
		MutantsDir:       "../gambit_out/mutants",
		OriginalFilePath: "../src/Counter.sol",
		OriginalFileName: "Counter.sol",
		BackupFile:       "../src/Backup.sol",
	}
	for i := 0; i < b.N; i++ {
		Slow(conf)
	}
}
