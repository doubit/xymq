package nameserver

import (
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"strings"
	"testing"
	"time"
)

type ParseErr struct {
	Index int
	Word  string
	Error error
}

func (e *ParseErr) String() string {
	return fmt.Sprintf("Error")
}

func TestParse(t *testing.T) {
	f, err := os.Create("cpuprofile.txt")
	if err != nil {
		t.Errorf("create profile file failed")
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			_, ok = r.(error)
			if !ok {
				t.Errorf("pkg: $v", r)
			}
		}
	}()
	testSrc := []string{
		"1 2 3 4 5",
		"100 50 24 12.5 6.25",
		"2 + 2 = 4",
		"1 XIAOYAO TEST",
		"",
	}

	fields := strings.Fields(testSrc[0])
	fields2numbers(fields)
	return
}

func fields2numbers(fields []string) (numbers []int) {
	if len(fields) == 0 {
		panic("no words to parse")
	}
	f := os.File
	defer f.Close()
	defer func() {
       err, _ = recover()
       if (err) {
       	 fmt.Printf("error : %v", err)
       }
	}
	for idx, field := range fields {
		num, err := strconv.Atoi(field)
		if err != nil {
			panic(&ParseErr{idx, field, err})
		}
		numbers = append(numbers, num)
	}
	return
}
