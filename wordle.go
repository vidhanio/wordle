package wordle

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

type GuessType int

const (
	GuessTypeCorrect GuessType = iota
	GuessTypeWrongPosition
	GuessTypeWrong
)

type void struct{}

type Wordle struct {
	word []rune // Correct word

	dictionary map[string]void // All valid words

	guessesAllowed int           // Number of guesses allowed
	wordLength     int           // Length of word
	guesses        [][]rune      // Guesses
	guessTypes     [][]GuessType // Character guesses

	cancelled bool
}

func New(wordLength, guessesAllowed int, dictionary, common []string) (*Wordle, error) {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// Initialize Wordle
	w := &Wordle{
		dictionary:     make(map[string]void),
		guessesAllowed: guessesAllowed,
		wordLength:     wordLength,
		guesses:        make([][]rune, 0),
	}

	for _, word := range dictionary {

		word = strings.ToLower(word)
		newWord := make([]rune, 0)
		for _, c := range word {
			if isLowercase(c) {
				newWord = append(newWord, c)
			}
		}

		if len(newWord) == wordLength {
			w.dictionary[string(newWord)] = void{}
		}
	}

	cleanedCommon := make([]string, 0)
	for _, word := range common {
		if len(word) != wordLength {
			continue
		}

		word = strings.ToLower(word)
		newWord := make([]rune, 0)
		for _, c := range word {
			if isLowercase(c) {
				newWord = append(newWord, c)
			}
		}

		if len(newWord) == wordLength {
			cleanedCommon = append(cleanedCommon, string(newWord))
			w.dictionary[string(newWord)] = void{}
		}
	}

	if len(cleanedCommon) == 0 {
		return nil, fmt.Errorf("invalid word length: no words with length %d in common words", wordLength)
	}

	randWord := cleanedCommon[r.Intn(len(cleanedCommon))]

	w.word = []rune(randWord)

	return w, nil
}

func (w *Wordle) Guess(guess string) ([]GuessType, error) {
	if w.Done() {
		return nil, fmt.Errorf("game is done")
	}

	guess = strings.ToLower(guess)
	guessRunes := []rune(guess)

	if guessRunes == nil || len(guessRunes) != w.wordLength {
		return nil, fmt.Errorf("invalid word length: wanted %d, got %d", w.wordLength, len(guessRunes))
	}

	if _, ok := w.dictionary[guess]; !ok {
		return nil, fmt.Errorf("invalid word: %s was not found in the dictionary", guess)
	}

	charGuesses := make([]GuessType, w.wordLength)

	for i, g := range guessRunes {
		if g < 'a' || g > 'z' {
			return nil, fmt.Errorf("invalid word: %c is not a letter", g)
		}

		if g == w.word[i] {
			charGuesses[i] = GuessTypeCorrect
		} else if contains(w.word, g) {
			charGuesses[i] = GuessTypeWrongPosition
		} else {
			charGuesses[i] = GuessTypeWrong
		}
	}

	w.guesses = append(w.guesses, guessRunes)
	w.guessTypes = append(w.guessTypes, charGuesses)

	return charGuesses, nil
}

func (w *Wordle) Word() string {
	if w.Done() {
		return string(w.word)
	}

	return ""
}

func (w *Wordle) WordLength() int {
	return w.wordLength
}

func (w *Wordle) Guesses() []string {
	guesses := make([]string, len(w.guesses))

	for i, g := range w.guesses {
		guesses[i] = string(g)
	}

	return guesses
}

func (w *Wordle) GuessTypes() [][]GuessType {
	return w.guessTypes
}

func (w *Wordle) GuessesLeft() int {
	return w.guessesAllowed - len(w.guesses)
}

func (w *Wordle) Cancel() {
	w.cancelled = true
}

func (w *Wordle) Cancelled() bool {
	return w.cancelled
}

func (w *Wordle) Won() bool {
	if len(w.guesses) == 0 {
		return false
	}

	return equal(w.guesses[len(w.guesses)-1], w.word)
}

func (w *Wordle) Lost() bool {
	return len(w.guesses) >= w.guessesAllowed
}

func (w *Wordle) Done() bool {
	return w.Won() || w.Lost() || w.Cancelled()
}
