# Corpus
Battery-apple expects a newline seperated sets of words. These can be had from the places found below, however a base corpus is included

## Words From the Internet

[Google 10000 word english corpus](https://github.com/first20hours/google-10000-english)

## Included Word List

A word list by Paul Noll, [available on his website](http://www.paulnoll.com/Books/Clear-English/English-3000-common-words.html). 
http://www.paulnoll.com/Books/Clear-English/English-3000-common-words.html

### How to recreate

1) Copy & paste words from [the website](http://www.paulnoll.com/Books/Clear-English/English-3000-common-words.html).
2) Remove double spaces and leading tabs:

    sed "s/^[ \t]*//" -i corpus.txt
    sed "/^$/d" -i corpus.txt

3) Remove the undesirable word (eg contractions, proper nouns). This may or may not have been a good idea.

    sed "/'/d" -i corpus.txt
    sed "/[A-Z]/d" -i corpus.txt

4) Remove smaller words:

    sed '/^\S\S$/d' -i corpus.txt
    sed '/^\S\S\s$/d' -i corpus.txt
