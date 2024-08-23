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

	// if there are arguments, use them as p1 and p2
	args := flag.Args()
	if *p1 == "" && len(args) > 0 {
		*p1 = args[0]
	}
	if *p2 == "" && len(args) > 1 {
		*p2 = args[1]
	}

	// if args are missing, print default
	if *p1 == "" || *p2 == "" {
		flag.PrintDefaults()
		return
	}

	nums1, nums2, err := parseBothArguments(*p1, *p2)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		return
	}

	cmd.Run(nums1, nums2, os.Stdout)
}

func parseBothArguments(p1, p2 string) ([]int, []int, error) {
	nums1, err1 := parseSingleArgument(p1)
	nums2, err2 := parseSingleArgument(p2)

	if err1 != nil {
		return nil, nil, fmt.Errorf("error: parsing p1: %v", err1)
	}
	if err2 != nil {
		return nil, nil, fmt.Errorf("error parsing p2: %v", err2)
	}
	return nums1, nums2, nil
}

func parseSingleArgument(arg string) ([]int, error) {
	// if arg is a range, like 21-100
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
