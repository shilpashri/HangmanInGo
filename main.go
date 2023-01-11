package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	countries := []string{"india", "singapore", "srilanka", "france", "germany", "korea", "italy",
		"switzerland", "pakistan", "austria", "poland", "myanmar", "nepal", "bhutan", "china", "afghanistan",
		"indonesia", "denmark", "finland", "iceland", "russia", "ukrain", "iran", "iraq", "uzbekinstan", "greece",
		"mexico", "argentina", "croatia", "chile", "brazil", "vietnam", "spain", "netherlands", "norway", "sweden"}

	states := []string{"karnataka", "tamil nadu", "goa", "telangana", "andhra pradesh", "maharashtra", "gujarat",
		"kerala", "madhya pradesh", "chattisgarh", "uttar pradesh", "jharkhand", "uttaranchal", "punjab", "haryana",
		"himachal pradesh", "rajasthan", "west bengal", "orissa", "sikkim", "arunachal pradesh", "assam", "meghalaya",
		"nagaland", "manipur", "mizoram", "tripura", "bihar",
	}

	wordIndex := rand.Intn(len(countries))
	var word string
	var choice int
	fmt.Println("Enter 1 for countries, 2 for states, 3 to exit")
	fmt.Scanf("%d", &choice)
	switch choice {
	case 1:
		word = countries[wordIndex]
	case 2:
		word = states[wordIndex]
	default:
		fmt.Println("Exiting.....")
		os.Exit(1)
	}

	var ch, resultantWord string
	resultantWord = strings.Repeat("-", len(word))

	for i := 5; i >= 0; {
		fmt.Print("Enter a character: ")
		fmt.Scanf("%s", &ch)
		if strings.Contains(resultantWord, ch) {
			fmt.Println("Already entered ", ch)
		}
		ch := strings.ToLower(ch)
		indeces := getAllIndeces(word, ch)
		if len(indeces) == 0 {
			fmt.Println("Character ", ch, "not present, you have", i, "turns left.")
			i--
		} else {
			resultantWord = replaceAtAllIndeces(resultantWord, ch, indeces)
			if !strings.Contains(resultantWord, "-") {
				fmt.Println("you got it right, the country / state was", resultantWord)
				break
			}
			fmt.Println(resultantWord)
		}
		if i == 0 {
			fmt.Println("All turns over, the country to be guessed was: ", word)
		}
	}
}

func replaceAtAllIndeces(in, r string, ii []int) string {
	out := []rune(in)
	ch := []rune(r)
	for i := 0; i < len(ii); i++ {
		out[ii[i]] = ch[0]
	}
	return string(out)
}

func getAllIndeces(word, ch string) []int {
	res := []int{}
	ind := strings.Index(word, ch)
	if ind < 0 {
		return res
	} else {
		lastIndex := strings.LastIndex(word, ch)
		if ind == lastIndex {
			res = append(res, ind)
			return res
		} else {
			for k := 0; k < len(word); k++ {
				if string(word[k]) == ch {
					res = append(res, k)
				}
			}
		}
	}
	return res // dummy
}
