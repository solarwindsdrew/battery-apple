package main

import (
    "io/ioutil"
    "fmt"
    "flag"
    "math/rand"
    "time"
    "strings"
    "hash/fnv"
)


// Retrieves a list of words from a newline seperated file
// Returns an array of single words
func loadWordList(corpusFile string) []string {

  corpus, err := ioutil.ReadFile(corpusFile)
  if err != nil {
      panic(err)
  }
  words := strings.Split(string(corpus), "\n")

  return words
}

func randomWord(wordList []string, r *rand.Rand) string {

  q := wordList[r.Intn(len(wordList))]

  return q
}

// Uses magic to create a hash of a string
func hash(s string) int64 {

  h := fnv.New64a()
  h.Write([]byte(s))
  return int64(h.Sum64())
}

// Creates the randomization source 
func randSource(s string) rand.Source {

  if (s != "") {
  // if there's a seed flag, hash it, then use it
    seed := hash(s)
    return rand.NewSource(seed)
  } else {
  // if there's no seed flag just use something generic
    return rand.NewSource(time.Now().Unix())
  }
} 

func main() {

  separatorPtr := flag.String("separator", "-", "character(s) to place between words")
  pwdLengthPtr := flag.Int("pwdLength", 4, "number of words in password")
  corpusPtr := flag.String("corpus", "corpus/corpus.txt", "newline sepearted file of words to use")
  seedPtr := flag.String("seed", "", "seed for random generation, should allow repeat password generation")

  flag.Parse()

  var words []string
  wordList := loadWordList(*corpusPtr)

  var s rand.Source = randSource(*seedPtr)
  r := rand.New(s) // initialize local pseudorandom generator

  for i := 0; i < *pwdLengthPtr; i++ {
    words = append(words, randomWord(wordList, r))
  }

  password := strings.Join(words, *separatorPtr)
  fmt.Println(password)

}