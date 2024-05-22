package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func Capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	result := string(unicode.ToUpper(rune(s[0])))
	for i := 1; i < len(s); i++ {
		if unicode.IsSpace((rune(s[i-1]))) {
			result += string(unicode.ToUpper(rune(s[i])))
		} else {
			result += string(s[i])
		}
	}
	return result
}

func RemoveAtIndex(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		fmt.Println("indx out of range")
		return slice
	}
	return append(slice[:index], slice[index+1:]...)
}

func RemovePlusTwoIndex(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		fmt.Println("indx out of range")
		return slice
	}
	return append(slice[:index], slice[index+2:]...)
}

func HandlePunctuations(s []string) []string {
	puncs := []string{",", ".", "!", "?", ":", ";"}

	// move punctuation to the previous word if it is in the middle of a string
	for i, word := range s {
		for _, punc := range puncs {
			if i > 0 && string(word[0]) == punc && string(word[len(word)-1]) != punc { // added =
				s[i-1] += punc
				s[i] = word[1:]
			}
		}
	}

	// move punctuations to the previous word if it is at the end of the string
	for i, word := range s {
		for _, punc := range puncs {
			if i > 0 && string(word[0]) == punc && (s[len(s)-1] == s[i]) {
				s[i-1] += word
				s = s[:len(s)-1]
			}
		}
	}

	// remove punctuation if it is in the middle of a string and not part of a group
	for i, word := range s {
		for _, punc := range puncs {
			if i > 0 && string(word[0]) == punc && string(word[len(word)-1]) == punc && s[i] != s[len(s)-1] {
				s[i-1] += word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	// handle apostrophes
	var count int
	count = 0
	for i, word := range s {
		if word == "'" && count == 0 {
			count += 1
			if i < len(s)-1 {
				s[i+1] = word + s[i+1]
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	for i, word := range s {
		if word == "'" {
			if i > 0 {
				s[i-1] = s[i-1] + word
				s = append(s[:i], s[i+1:]...)
			}
		}
	}

	return s
}

func main() {
	// Open the file
	file, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	// Create a new scanner to read from the file
	scanner := bufio.NewScanner(file)
	// Initialize a slice to hold all words from the file
	var sampleSlice []string
	// Read line by line
	for scanner.Scan() {
		line := scanner.Text()
		sampleSlice = strings.Split(string(line), " ") // Split line into words
		// sampleSlice = append(sampleSlice, words...) // Append words to sampleSlice
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}

	vowels := []string{"a", "e", "i", "o", "u", "A", "E", "I", "O", "U"}
	// Iterate over sampleSlice to modify specific values
	for indx, val := range sampleSlice {
		if val == "(cap)" && indx > 0 { // Ensure we're not at the beginning of the slice
			valMod := Capitalize(sampleSlice[indx-1]) // Capitalize the previous word
			sampleSlice[indx-1] = valMod              // Replace the previous word with capitalized version
			sampleSlice = RemoveAtIndex(sampleSlice, indx)

		} else if val == "(up)" && indx > 0 {
			valMod := strings.ToUpper(sampleSlice[indx-1]) // Capitalize the previous word
			sampleSlice[indx-1] = valMod
			sampleSlice = RemoveAtIndex(sampleSlice, indx)

		} else if val == "(up," {
			valMod := strings.Trim(string(sampleSlice[indx+1]), sampleSlice[indx+1][1:])
			number, _ := strconv.Atoi(string(valMod))
			for i := 1; i <= number; i++ {
				sampleSlice[indx-i] = strings.ToUpper(sampleSlice[indx-i])
			}
			sampleSlice = RemovePlusTwoIndex(sampleSlice, indx)

		} else if val == "(low," {
			valMod := strings.Trim(string(sampleSlice[indx+1]), sampleSlice[indx+1][1:])
			number, _ := strconv.Atoi(string(valMod))
			for i := 1; i <= number; i++ {
				sampleSlice[indx-i] = strings.ToLower(sampleSlice[indx-i])
			}
			sampleSlice = RemovePlusTwoIndex(sampleSlice, indx)

		} else if val == "(cap," {
			valMod := strings.Trim(string(sampleSlice[indx+1]), sampleSlice[indx+1][1:])
			number, _ := strconv.Atoi(string(valMod))
			for i := 1; i <= number; i++ {
				sampleSlice[indx-i] = Capitalize(sampleSlice[indx-i])
			}
			sampleSlice = RemovePlusTwoIndex(sampleSlice, indx)

		} else if val == "(hex)" {
			valMod := sampleSlice[indx-1]
			decimalValue, err := strconv.ParseInt(valMod, 16, 64)
			sampleSlice[indx-1] = strconv.Itoa(int(decimalValue))
			sampleSlice = RemoveAtIndex(sampleSlice, indx)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		} else if val == "(bin)" {
			valMod := sampleSlice[indx-1]
			decimalValue, err := strconv.ParseInt(valMod, 2, 64)
			sampleSlice[indx-1] = strconv.Itoa(int(decimalValue))
			sampleSlice = RemoveAtIndex(sampleSlice, indx)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
		} else if val == "a" && indx > 0 && indx+1 < len(sampleSlice) {
			nextWord := sampleSlice[indx+1]
			for _, char := range vowels {
				if strings.HasPrefix(nextWord, char) {
					sampleSlice[indx] = val + "n"
				}
			}
		}
	}

	word := strings.Join(HandlePunctuations(sampleSlice), " ") + "\n"
	file, err = os.Create("result.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.WriteString((word))
	if err != nil {
		panic(err)
	}
	// fmt.Println(word)
	// Check for any errors during scanning
}
