package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/one2n-go-bootcamp/game-of-pig/cmd"
)

func main() {
	p1 := flag.String("p1", "", "Strategy of First Player")
	p2 := flag.String("p2", "", "Strategy of Second Player")
	flag.Parse()

	if *p1 == "" || *p2 == "" {
		flag.PrintDefaults()
		return
	}

	nums1, err1 := parseArgument(*p1)
    nums2, err2 := parseArgument(*p2)

    if err1 != nil {
        fmt.Printf("Error parsing p1: %v\n", err1)
        return
    }
    if err2 != nil {
        fmt.Printf("Error parsing p2: %v\n", err2)
        return
    }

    cmd.Run(nums1, nums2, os.Stdout)
}



func parseArgument(arg string) ([]int, error) {
    if strings.Contains(arg, "-") {
        parts := strings.Split(arg, "-")
        if len(parts) != 2 {
            return nil, fmt.Errorf("error: invalid range format")
        }
        start, err1 := strconv.Atoi(parts[0])
        end, err2 := strconv.Atoi(parts[1])
        if err1 != nil || err2 != nil {
            return nil, fmt.Errorf("error: invalid number in range")
        }
        if start > end {
            return nil, fmt.Errorf("error: start of range is greater than end")
        }
        var result []int
        for i := start; i <= end; i++ {
            result = append(result, i)
        }
        return result, nil
    } else {
        num, err := strconv.Atoi(arg)
        if err != nil {
            return nil, fmt.Errorf("error: invalid number format")
        }
        return []int{num}, nil
    }
}