# wordle

A simple wordle package in Go.

Written on a bored Tuesday which would have been spent better doing homework.

Want to see it in action? Check out [discordle](https://github.com/vidhanio/discordle).

## Example

```go
package main

import (
    "fmt"

    "github.com/vidhanio/wordle"
)

func main() {
	// all words which are valid
	dictionary := []string{"hello", "world", "among", "us", "sus", "tasks"}

	// words which can be used as the word in the wordle
	// this is because some valid words are near-impossible to guess, although they are valid words
	common := []string{"hello", "among"}

	// Create a new wordle with:
	// - A word length of 5
	// - A guess count of 6
	// with the validWords and commonWords
	w, err := wordle.New(5, 6, dictionary, common)
	if err != nil {
		panic(err)
	}

	guessResult, err := w.Guess("among")
	if err != nil {
		// errors returned from (*wordle).Guess are non-fatal, so we can continue while telling the user the error
		fmt.Println(err)
	}

	fmt.Println(guessResult)
}
```
