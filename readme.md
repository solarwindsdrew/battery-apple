# Battery-Apple 
A silly (xkcd style password generator)[https://xkcd.com/936/] in go.


## Build

Install (go via the offical instructions)[https://golang.org/doc/install]

Build:

    go build

## Usage

Generate a 4 word random password from the built in corpus:

    ./battery-apple

    -or-

    ./battery-apple \
    -separator '-' \
    -pwdLength 4 \
    -corpus './corpus/corpus.txt'

Flags:
* separator: character between words
* pwdLength: Number of words
* corpus: Newline delimited file of words
* seed: Randomizer seed. This allows repeatable results. Use at your own risk.



## Sources

[XKCD password strength](https://xkcd.com/936/)
[XKCD password generator by preshing](http://preshing.com/20110811/xkcd-password-generator/)
[Go Lang command line flags](https://gobyexample.com/command-line-flags)
[A Tour of Go: flow control](https://tour.golang.org/flowcontrol/1)