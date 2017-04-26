# Corpus how to

1) Copy past words from this site: http://www.paulnoll.com/Books/Clear-English/English-3000-common-words.html
2) remove double spaces and leading tabs:

    sed "s/^[ \t]*//" -i corpus.txt
    sed "/^$/d" -i corpus.txt

3) Remove the undesirable word (eg contractions, proper nouns)

    sed "/'/d" -i corpus.txt
    sed "/[A-Z]/d" -i corpus.txt