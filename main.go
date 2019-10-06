package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type WordList []struct {
	En string
	Tr string
}

func main() {
	wordList := readFile()

	for {
		fmt.Println("Test : 1")
		fmt.Println("Exam : 2")
		fmt.Println("Add  : 3")

		_, selectMenu := readInput("")

		switch selectMenu {
		case 0:
			cleanTerminal()
			test(wordList)
			cleanTerminal()
		case 1:
			cleanTerminal()
			exam(wordList)
			cleanTerminal()
		case 2:
			cleanTerminal()
			add()
			cleanTerminal()
		default:
			cleanTerminal()
		}
	}

}

func cleanTerminal() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func readFile() WordList {
	var wordList WordList
	words, _ := os.Open("word-list.json")
	defer words.Close()
	byteValue, _ := ioutil.ReadAll(words)
	json.Unmarshal(byteValue, &wordList)
	return wordList
}

func readValue(wordList WordList, number int) (string, string) {
	for i := range wordList {
		if i == number {
			return wordList[i].En, wordList[i].Tr
		}
	}
	return "none", "none"
}

func randValueGenerator(max int) int {
	rand.Seed(time.Now().UnixNano())
	return (rand.Intn((max-1)-0+1) + 0)
}

func readInput(text string) (string, int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text)
	input, _ := reader.ReadString('\n')
	input = strings.Replace(input, "\n", "", -1)
	intInput, _ := strconv.Atoi(input)
	return input, intInput - 1
}

func exam(wordList WordList) {
	var trueScore int
	_, questionCount := readInput("How many questions do you want ? : ")

	for i := 0; i <= questionCount; i++ {
		number := randValueGenerator(len(wordList))
		enWord, trWord := readValue(wordList, number)

		fmt.Println()
		fmt.Println()
		fmt.Println(enWord)
		fmt.Println("--------------------------------")

		trInputAnswer, _ := readInput("Enter answer: ")

		trWordSplit := strings.Split(trWord, ",")
		for i := range trWordSplit {
			if trWordSplit[i] == trInputAnswer {
				fmt.Println("Well done !   ", enWord, "  --->  ", trWord)
				fmt.Println("--------------------------------")
				fmt.Println()
				fmt.Println()
				trueScore++
				break
			}
		}
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("--------------------------------")
	fmt.Println("Statistics -> ", "Total:", questionCount+1, " | ", "True:", trueScore, " | ", "False:", questionCount+1-trueScore)
	readInput("Press enter to continue")
}

func test(wordList WordList) {
	cleanTerminal()
	//var trueScore int
	_, questionCount := readInput("How many questions do you want ? : ")

	for i := 0; i <= questionCount; i++ {
		var enArray [4]string
		var trArray [4]string

		for j := 0; j < 4; j++ {
			enArray[j], trArray[j] = readValue(wordList, randValueGenerator(4))
		}

		fmt.Println()
		fmt.Println()
		fmt.Println(enArray[randValueGenerator(4)])
		fmt.Println("--------------------------------")
		fmt.Println("A) ", trArray[randValueGenerator(4)])
		fmt.Println("B) ", trArray[randValueGenerator(4)])
		fmt.Println("C) ", trArray[randValueGenerator(4)])
		fmt.Println("D) ", trArray[randValueGenerator(4)])

		trInputAnswer, _ := readInput("Enter answer: ")
		fmt.Println(trInputAnswer)

		// trWordSplit := strings.Split(trWord, ",")
		// for i := range trWordSplit {
		// 	if trWordSplit[i] == trInputAnswer {
		// 		fmt.Println("Well done !   ", enWord, "  --->  ", trWord)
		// 		fmt.Println("--------------------------------")
		// 		fmt.Println()
		// 		fmt.Println()
		// 		trueScore++
		// 		break
		// 	}
		// }
	}

	readInput("Press enter to continue")
}

func add() {
	for {
		fmt.Println("English : 1")
		fmt.Println("Turkish : 2")
		fmt.Println("Menu    : 3")

		_, selectMenu := readInput("")

		if selectMenu == 0 {
			cleanTerminal()
			en, _ := readInput("")
			fmt.Println(en)
			tr, _ := readInput("")
			fmt.Println(tr)
			cleanTerminal()
		} else if selectMenu == 1 {
			cleanTerminal()
		} else if selectMenu == 2 {
			break
		} else {
			cleanTerminal()
		}
	}
}
