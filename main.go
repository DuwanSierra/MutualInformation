package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func entropy(filename string) (float64, error) {
	bytes, err := os.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	counts := make(map[int]int)
	for _, b := range bytes {
		counts[int(b)]++
	}

	var entropy float64
	for _, count := range counts {
		p := float64(count) / float64(len(bytes))
		entropy -= p * math.Log2(p)
	}

	return entropy, nil
}

func jointEntropy(filename1, filename2 string) (float64, error) {
	bytes1, err := os.ReadFile(filename1)
	if err != nil {
		return 0, err
	}

	bytes2, err := os.ReadFile(filename2)
	if err != nil {
		return 0, err
	}

	// Pad the smaller file with zeros
	if len(bytes1) < len(bytes2) {
		padding := make([]byte, len(bytes2)-len(bytes1))
		bytes1 = append(bytes1, padding...)
	} else if len(bytes2) < len(bytes1) {
		padding := make([]byte, len(bytes1)-len(bytes2))
		bytes2 = append(bytes2, padding...)
	}

	counts := make(map[int]int)
	for i := range bytes1 {
		jointByte := int(bytes1[i])<<8 | int(bytes2[i])
		counts[jointByte]++
	}

	var entropy float64
	for _, count := range counts {
		p := float64(count) / float64(len(bytes1))
		entropy -= p * math.Log2(p)
	}

	return entropy, nil
}

func main() {
	pathFile1 := "input_image_2.jpg"
	pathFile2 := "input_image_3.jpg"
	entropy1, err := entropy(pathFile1)
	if err != nil {
		log.Fatal(err)
	}

	entropy2, err := entropy(pathFile2)
	if err != nil {
		log.Fatal(err)
	}

	jointEntropy, err := jointEntropy(pathFile1, pathFile2)
	if err != nil {
		log.Fatal(err)
	}

	mutualInformation := entropy1 + entropy2 - jointEntropy
	//Convert mutual information to percentage inside the range [0, 100]
	mutualInformation = mutualInformation / math.Max(entropy1, entropy2) * 100
	fmt.Println("Mutual Information:", mutualInformation, "%")
}
