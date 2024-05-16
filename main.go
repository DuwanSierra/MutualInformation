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
	pathFile1 := "audio1.weba"
	pathFile2 := "audio3.weba"
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
	maxEntropy := math.Min(entropy1, entropy2)
	percentage := (mutualInformation / maxEntropy) * 100
	fmt.Printf("Entropy of %s: %.2f\n", pathFile1, entropy1)
	fmt.Printf("Entropy of %s: %.2f\n", pathFile2, entropy2)
	fmt.Printf("Joint Entropy: %.2f\n", jointEntropy)
	fmt.Printf("Mutual Information: %.2f%%\n", percentage)
}
