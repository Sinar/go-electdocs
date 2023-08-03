package main

import "fmt"

func main() {
	fmt.Println("Welcome to PRU15 DUN Ager!!!")
	Run()
}

func Run() {
	ProcessN9()
}

func ProcessN9() {
	// Open up file named SELANGOR.csv that is in testdata
	// First sweep to gather all the PARs + DUNs
	// Second round over the created map; the array of Candidates
	//	For each DUN, find exact match; and fuzzy options for each candidate
	//	Pick the best options for each .. 2 - 4 candidates ..
	//		leftovers fight for scraps!
	LoadTSVCandidates("N9")
}
