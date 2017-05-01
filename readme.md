# Battery-Apple 
A silly (xkcd style password generator)[https://xkcd.com/936/] in go.


## Build

Install (go via the offical instructions)[https://golang.org/doc/install]

Build:

    go build

## Usage



## Corpus how to recreate

1) Copy & paste words from (this handy site)[http://www.paulnoll.com/Books/Clear-English/English-3000-common-words.html]. I did it manually, but I really should have used (this list)[https://github.com/first20hours/google-10000-english
].
2) Remove double spaces and leading tabs:

    sed "s/^[ \t]*//" -i corpus.txt
    sed "/^$/d" -i corpus.txt

3) Remove the undesirable word (eg contractions, proper nouns). This may or may not have been a good idea.

    sed "/'/d" -i corpus.txt
    sed "/[A-Z]/d" -i corpus.txt

4) Remove smaller words:

    sed '/^\S\S$/d' -i corpus.txt
    sed '/^\S\S\s$/d' -i corpus.txt


## Sources

[XKCD password strength](https://xkcd.com/936/)
[XKCD password generator by preshing](http://preshing.com/20110811/xkcd-password-generator/)
[Go Lang command line flags](https://gobyexample.com/command-line-flags)
[A Tour of Go: flow control](https://tour.golang.org/flowcontrol/1)