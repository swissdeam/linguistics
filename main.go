package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const word = "ЛИНГВИСТИКА"
const iterations = 1000000

func generateBilbaWord() string {
	letters := uniqueLetters(word)
	rand.Shuffle(len(letters), func(i, j int) { letters[i], letters[j] = letters[j], letters[i] })
	return strings.Join(letters[:5], "")
}

func generateBolbaWord() string {
	vowels := []rune{'И', 'А'}
	consonants := []rune{'Л', 'Н', 'Г', 'В', 'С', 'Т', 'К'}

	vowel1 := vowels[rand.Intn(len(vowels))]
	vowel2 := vowels[rand.Intn(len(vowels)-1)]
	if vowel2 == vowel1 {
		vowel2 = vowels[len(vowels)-1]
	}

	rand.Shuffle(len(consonants), func(i, j int) { consonants[i], consonants[j] = consonants[j], consonants[i] })

	return string([]rune{consonants[0], vowel1, consonants[1], vowel2, consonants[2]})
}

func generateSmartBilbaWord() string {
	vowels := []rune{'И', 'А'}
	consonants := []rune{'Л', 'Н', 'Г', 'В', 'С', 'Т', 'К'}

	rand.Shuffle(len(vowels), func(i, j int) { vowels[i], vowels[j] = vowels[j], vowels[i] })
	rand.Shuffle(len(consonants), func(i, j int) { consonants[i], consonants[j] = consonants[j], consonants[i] })

	return string([]rune{consonants[0], vowels[0], consonants[1], vowels[1], consonants[2]})
}

func uniqueLetters(word string) []string {
	letterMap := make(map[rune]bool)
	for _, letter := range word {
		letterMap[letter] = true
	}

	letters := make([]string, 0, len(letterMap))
	for letter := range letterMap {
		letters = append(letters, string(letter))
	}

	return letters
}

func main() {
	rand.Seed(time.Now().UnixNano())
	successesA := 0
	successesB := 0

	for i := 0; i < iterations; i++ {
		wordBilba := generateBilbaWord()
		wordBolba := generateBolbaWord()
		if wordBilba == wordBolba {
			successesA++
		}
		if generateSmartBilbaWord() == wordBolba {
			successesB++
		}
	}

	probA := float64(successesA) / iterations
	probB := float64(successesB) / iterations

	fmt.Printf("Probability of success (case a): %.6f\n", probA)
	fmt.Printf("Probability of success (case b): %.6f\n", probB)
}
