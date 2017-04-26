package main

import (
    "io/ioutil"
    "fmt"
    "math/rand"
    "time"
    "strings"
)

func main() {
    var corpusFile string = "corpus.txt"
        

    // read the whole file at once
    corpus, err := ioutil.ReadFile(corpusFile)
    if err != nil {
        panic(err)
    }

    words := strings.Split(string(corpus), "\n")


    s := rand.NewSource(time.Now().Unix())
    r := rand.New(s) // initialize local pseudorandom generator 
    q := words[r.Intn(len(words))]

    fmt.Printf("%s\n", q)
}