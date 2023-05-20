package main

/* program to create a melody */

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"math"
	"os"
	"strings"
)

// https://pages.mtu.edu/~suits/notefreqs.html
var Notes = map[byte]int{
	'A': 440,
	'B': 494,
	'C': 523,
	'D': 587,
	'E': 659,
	'F': 699,
	'G': 784,
}

const (
	Duration   = 0.33
	SampleRate = 44100
	Frequency  = 440
	Delay      = 0.15
)

var (
	tau = math.Pi * 2
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("require name of note.dat file and output")
		return
	}
	filePath := os.Args[1]
	readFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	var notes [][]string
	for _, line := range fileLines {
		notes = append(notes, strings.Split(line, " "))
		for _, n := range(notes[len(notes)-1]) {
			fmt.Print(n)
		}
		fmt.Println()
	}


	fmt.Fprintf(os.Stderr, "generating sine wave..\n")
	generate(os.Args[2], notes)
	fmt.Fprintf(os.Stderr, "\ndone\n")
}

func generate(filename string, notes [][]string) {
	nsamps := Duration * SampleRate
	var blob []float64
	var angle float64 = tau / float64(nsamps)
	file := filename
	f, _ := os.Create(file)
	for _, bar := range notes {
		for _, n := range bar {
			var (
				start float64 = 1.0
				end   float64 = 1.0e-4
			)
			decayfac := math.Pow(end/start, 1.0/float64(nsamps))
			for i := 0; i < int(nsamps); i++ {
				sample := math.Sin(angle * float64(Notes[n[0]]) * float64(i))
				sample *= start
				start *= decayfac
				var buf [4]byte
				binary.LittleEndian.PutUint32(buf[:], math.Float32bits(float32(sample)))
				bw, err := f.Write(buf[:])
				blob = append(blob, sample)
				if err != nil {
					panic(err)
				}
				fmt.Printf("\rWrote: %v bytes to %s", bw, file)
			}
		}
	}
}
