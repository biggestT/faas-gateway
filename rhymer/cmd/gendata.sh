#! /bin/sh

wget -O - -o /dev/null https://svn.spraakdata.gu.se/sb/fnplusplus/pub/parole.txt \
  | awk '{print $1, $3}' \
  | uniq \
  > data/wordlist.txt

