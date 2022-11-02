A repository for small scale Go performance comparisons I'm curious about. 

Currently only one test is implemented, a comparison of min(x,y,z) == 0 vs (x == 0 or y == 0 or z == 0) inspired by [this article](http://0x80.pl/notesen/2021-03-11-any-word-is-zero.html). 

It appears that, in Go at least, the simple comparison without min() is faster.