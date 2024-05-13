package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/mvlipka/bhe-code-exercise/pkg/primes"
	"github.com/mvlipka/bhe-code-exercise/pkg/primes/calculators"
	"log"
	"time"
)

func main() {
	calculatorFlag := flag.String("method", "eratosthenes", "the calculation method used to generate primes")
	indexFlag := flag.Int64("index", 100, "the index of prime in which to calculate")
	timeoutFlag := flag.Duration("timeout", -1, "the duration, in seconds, at which the program will timeout and cancel generating a prime")

	flag.Parsed()

	var calculator calculators.Calculator
	switch *calculatorFlag {
	case "eratosthenes":
		calculator = calculators.NewEratosthenesCalculator()
	}

	ctx := context.Background()
	var cancel context.CancelFunc
	if *timeoutFlag >= 1 {
		ctx, cancel = context.WithTimeout(context.Background(), *timeoutFlag*time.Second)
		defer cancel()
	}

	generator := primes.NewGenerator(calculator)

	result, err := generator.GetPrimeAtIndex(ctx, *indexFlag)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(fmt.Sprintf("The %d prime number is %d", *indexFlag, result))
}
