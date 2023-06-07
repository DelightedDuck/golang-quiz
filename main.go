package main
import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"time"
	"html"
)

type questionData struct {
	Text      string   `json:"question"`
	Wrong	[]string   `json:"incorrect_answers"`
	Correct   string   `json:"correct_answer"`
}

var question questionData

func clearScreen() {
	fmt.Printf("\033c")
}

func shufStrings(strings []string, rng *rand.Rand) {
	rng.Seed(time.Now().UnixNano())
	n := len(strings)
	for i := n - 1; i > 0; i-- {
		j := rng.Intn(i + 1)
		strings[i], strings[j] = strings[j], strings[i]
	}
}

func pickQuestion(rng *rand.Rand) {
	f := "questions.json"
	fileData, err := ioutil.ReadFile(f)
	if err != nil {
		fmt.Println("Failed to open file:", err)
		return
	} 

	// Parse the JSON data into a struct
	var jsonData struct {
		Results []questionData `json:"results"`
	}
	
	// Parse the JSON data into a slice of Question structs
	err = json.Unmarshal(fileData, &jsonData)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	rand.Seed(time.Now().UnixNano())
	num := rng.Intn(len(jsonData.Results))
	randomQuestion := jsonData.Results[num]
	
	question.Correct = html.UnescapeString(randomQuestion.Correct)
	question.Text = html.UnescapeString(randomQuestion.Text)
	question.Wrong = make([]string, len(randomQuestion.Wrong))
	
	for i, wrong := range randomQuestion.Wrong {
		question.Wrong[i] = html.UnescapeString(wrong)
	}
	shufStrings(question.Wrong, rng)
} 

func askQuestion(rng *rand.Rand) {
	var i int 
	var x string
	for {
		fmt.Println("Press Enter to continue...")
		fmt.Scanln(&x)
		if x == "" {
			clearScreen()
			pickQuestion(rng)
			rng.Seed(time.Now().UnixNano())
			idk := rand.Intn(4)
			f := 1
			answer := 0
			print := false

			// Print Question	
			fmt.Println(question.Text, ":")

			// Print answers in random order
			for _, q := range question.Wrong {
				if idk == f {
					fmt.Println(f,") ", question.Correct)
					answer = f
					f++
					print = true	
				}	
				fmt.Println(f, ") ", q)
				f++
			}
			
			if !print {
				fmt.Println(f, ") ", question.Correct)
				answer = f
			}			

			// Ask for users answer
			fmt.Println(" What is your answer?: ")
			fmt.Scanln(&i)
			if answer == i {
				fmt.Println("Correct!")
			} else {
				fmt.Println("Wrong...")
			}
		} else {
			x = ""
		}
		time.Sleep(1 * time.Second)
		clearScreen()
	}

}

func main() {
	clearScreen()

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))
	askQuestion(rng)	
}
