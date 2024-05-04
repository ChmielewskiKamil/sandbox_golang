package golang

import (
	"fmt"
	"os/exec"
	"testing"
)

// Benchmark: 14.928s
func BenchmarkSlow(b *testing.B) {
	err := generateMutants()
	if err != nil {
		b.Fatal(err)
	}

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

func generateMutants() error {
	cmd := exec.Command("gambit", "mutate", "--filename", "src/Counter.sol", "--contract", "Counter")
	// cmd.Dir = "/Users/kamilchmielewski/Projects/sandbox_golang/experiments/0001"
	cmd.Dir = "../"
	println(cmd.Dir)
	err := cmd.Run()
	if err != nil {
		return err
	}

	fmt.Println("Mutants generated")
	return nil
}
