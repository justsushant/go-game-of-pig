package main

import (
	"flag"
	"fmt"
	"os"
	"slices"
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

	nums1, nums2 = removeCommonElements(nums1, nums2)
	return nums1, nums2, nil
}

func parseSingleArgument(arg string) ([]int, error) {
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

func removeCommonElements(nums1, nums2 []int) ([]int, []int) {
	var slice2 []int

	for _, num := range nums2 {
		if !slices.Contains(nums1, num) {
			slice2 = append(slice2, num)
		}
	}

	return nums1, slice2
}
