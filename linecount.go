package linecount

import (
	"bytes"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type CountsExperiment struct {
	Runs        int          `json:"runs"`        // number of duplicate trials using the same variables
	Count       int          `json:"count"`       // number of runs with varying conditions
	N           int          `json:"n"`           // number of measurements (bytes) per run
	KeepContent bool         `json:"keepContent"` // whether to keep the content or only a summary
	Content     []SampleData `json:"content"`     // set of returned sample content (stored if keepContent == true)
	Avg         float64      `json:"Avg"`         // average of all counts
}

func (t *CountsExperiment) Lines() float64 {
	return t.Avg
}

// Lines returns the number of newline characters in b.
func Lines(b []byte) (n int) {
	for _, c := range b {
		if c == '\n' {
			n += 1
		}
	}
	return
}

func randByte() byte {
	return byte(rand.Intn(255))
}

func SampleRun(trials, n int) (int, []byte) {
	return sampleRunStringsBuilder(n)
}

func sampleRunStringsBuilder(s int) (int, []byte) {
	sb := strings.Builder{}
	defer sb.Reset()

	if s < 1 {
		s = defaultSampleSize
	}

	for i := 0; i < s; i++ {
		sb.WriteByte(randByte())
	}

	b := []byte(sb.String())
	n := Lines(b)

	return n, b
}

func sampleRunBytesBuffer(s int) (int, []byte) {
	buf := bytes.Buffer{}
	defer buf.Reset()

	if s < 1 {
		s = defaultSampleSize
	}

	for i := 0; i < s; i++ {
		buf.WriteByte(randByte())
	}

	b := buf.Bytes()
	n := Lines(b)

	return n, b
}

func Sample() {

	n, buf := SampleRun(defaultTrialCount, defaultSampleSize)

	os.WriteFile("sample", buf, 0644)
	fmt.Println("linecount sample:")
	fmt.Printf("count: %v\n", n)
	fmt.Println("")
	// fmt.Print(string(buf))
}

func init() {
	// set the random number generator seed
	rand.Seed(time.Now().UTC().UnixNano())
}
