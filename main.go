package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
)

func entropy(filename string) (float64, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return 0, err
	}

	var bits []int
	for _, b := range bytes {
		for i := 7; i >= 0; i-- {
			bit := (b >> uint(i)) & 1
			bits = append(bits, int(bit))
		}
	}

	counts := make(map[int]int)
	for _, bit := range bits {
		counts[bit]++
	}

	var entropy float64
	for _, count := range counts {
		p := float64(count) / float64(len(bits))
		entropy -= p * math.Log2(p)
	}

	return entropy, nil
}

func jointEntropy(filename1, filename2 string) (float64, error) {
	bytes1, err := ioutil.ReadFile(filename1)
	if err != nil {
		return 0, err
	}

	bytes2, err := ioutil.ReadFile(filename2)
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

	var jointBits []int
	for i := range bytes1 {
		for j := 7; j >= 0; j-- {
			bit1 := (bytes1[i] >> uint(j)) & 1
			bit2 := (bytes2[i] >> uint(j)) & 1
			jointBit := int(bit1<<1 | bit2) // Convert jointBit to int
			jointBits = append(jointBits, jointBit)
		}
	}

	counts := make(map[int]int)
	for _, bit := range jointBits {
		counts[bit]++
	}

	var entropy float64
	for _, count := range counts {
		p := float64(count) / float64(len(jointBits))
		entropy -= p * math.Log2(p)
	}

	return entropy, nil
}

func main() {
	pathFile1 := "input_image.jpg"
	pathFile2 := "input_image_2.jpg"
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
	/*
		Una comprensión correcta:
			Variables similares⟹MI es máximo.
			Sin cambios en las variables⟹MI es mínimo (MI es cero)
	*/
	fmt.Println("Entropy 1:", entropy1)
	fmt.Println("Entropy 2:", entropy2)
	fmt.Println("Mutual Information:", mutualInformation)
}
