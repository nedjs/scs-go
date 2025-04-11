package main

import (
	"fmt"
	"sort"
	"sync"
)

type ScsOptions struct {
	Parallel      bool
	Deterministic bool
}

type Letter struct {
	Index   int
	Word    string
	Value   string
	Links   []*Link
	linksMu sync.Mutex
}

type Link struct {
	A *Letter
	B *Letter
}

func createLink(a *Letter, b *Letter) *Link {
	if a.Word == b.Word {
		panic(fmt.Sprintf("Invalid link between same word %s", a.Word))
	}
	link := &Link{A: a, B: b}

	a.linksMu.Lock()
	a.Links = append(a.Links, link)
	a.linksMu.Unlock()

	b.linksMu.Lock()
	b.Links = append(b.Links, link)
	b.linksMu.Unlock()
	return link
}

func (l *Link) isForWord(word string) bool {
	return l.A.Word == word || l.B.Word == word
}

func (l *Link) indexRel(word string) int {
	if l.A.Word == word {
		return l.A.Index
	}
	if l.B.Word == word {
		return l.B.Index
	}
	return -1
}

func (l *Link) OpposingSide(word string) *Letter {
	if l.A.Word == word {
		return l.B
	}
	if l.B.Word == word {
		return l.A
	}
	panic(fmt.Sprintf("Invalid word %s for link", word))
}

type Linking struct {
	Words     [][]*Letter
	WordsDict map[string][]*Letter
}

func (linking *Linking) AddWord(word string, wg *sync.WaitGroup) {

	if _, ok := linking.WordsDict[word]; ok {
		return
	}

	wordChars := make([]*Letter, len(word))
	for i := 0; i < len(word); i++ {
		wordChars[i] = &Letter{
			Index: i,
			Word:  word,
			Value: string(word[i]),
		}
	}
	linking.WordsDict[word] = wordChars
	linking.Words = append(linking.Words, wordChars)

	for _, existingWord := range linking.Words {
		if existingWord[0].Word == word {
			continue
		}
		if wg != nil {
			wg.Add(1)
			existingWord := existingWord
			go func() {
				linking.createLinksForLCS(existingWord, wordChars)
				wg.Done()
			}()
		} else {
			linking.createLinksForLCS(existingWord, wordChars)
		}
	}
}

func (linking *Linking) createLinksForLCS(str1, str2 []*Letter) {
	n, m := len(str1), len(str2)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := n - 1; i >= 0; i-- {
		for j := m - 1; j >= 0; j-- {
			if str1[i].Value == str2[j].Value {
				dp[i][j] = 1 + dp[i+1][j+1]
			} else {
				if dp[i+1][j] > dp[i][j+1] {
					dp[i][j] = dp[i+1][j]
				} else {
					dp[i][j] = dp[i][j+1]
				}
			}
		}
	}

	x, y := 0, 0
	lcs := ""
	for x < n && y < m {
		if str1[x].Value == str2[y].Value {
			lcs += str1[x].Value
			x++
			y++
		} else if dp[x+1][y] >= dp[x][y+1] {
			x++
		} else {
			y++
		}
	}

	x, y = 0, 0
	for _, c := range lcs {
		for x < n && str1[x].Value != string(c) {
			x++
		}
		for y < m && str2[y].Value != string(c) {
			y++
		}
		createLink(str1[x], str2[y])
		x++
		y++
	}
}

func createLinks(words []string, wg *sync.WaitGroup) *Linking {
	linking := &Linking{
		Words:     [][]*Letter{},
		WordsDict: map[string][]*Letter{},
	}
	for _, word := range words {
		linking.AddWord(word, wg)
	}
	if wg != nil {
		wg.Wait()
	}
	return linking
}

func walkLinks(linking *Linking) string {
	indices := map[string]int{}
	lookingAtLinks := map[*Link]bool{}
	lookingAtWords := map[string]bool{}

	var walk func([]*Letter, int, int) string

	walk = func(letters []*Letter, toIndex int, depth int) string {
		word := letters[0].Word

		if toIndex == indices[word] {
			return ""
		}

		lookingAtWords[word] = true
		left := ""

		for indices[word] < toIndex {
			letter := letters[indices[word]]
			walkableLetters := []*Letter{}
			walkableLinks := []*Link{}

			for _, link := range letter.Links {
				char := link.OpposingSide(word)
				if !lookingAtWords[char.Word] {
					valid := true
					for l := range lookingAtLinks {
						// check if we are processing a link currently behind where our indices are,
						// if so then this link is for in-between that link and our head. This should not be processed
						// and should get processed forward in a later iteration
						if l.isForWord(char.Word) && l.indexRel(char.Word) < indices[char.Word] {
							valid = false
							break
						}
					}
					if valid {
						walkableLinks = append(walkableLinks, link)
						walkableLetters = append(walkableLetters, char)
					}
				}
			}

			for _, v := range walkableLinks {
				lookingAtLinks[v] = true
			}

			leftBuf := ""
			for _, v := range walkableLetters {
				nextWord := linking.WordsDict[v.Word]
				if indices[v.Word] <= v.Index {
					buf := walk(nextWord, v.Index, depth+1)
					leftBuf += buf
					indices[v.Word]++
				}
			}

			for _, v := range walkableLinks {
				delete(lookingAtLinks, v)
			}

			leftBuf += letter.Value
			left += leftBuf

			if indices[word] == letter.Index {
				indices[word]++
			}
		}

		lookingAtWords[word] = false
		return left
	}

	buffer := ""
	for _, word := range linking.Words {
		indices[word[0].Word] = 0
	}
	for _, word := range linking.Words {
		buffer += walk(word, len(word), 0)
	}
	return buffer
}

func scs(words []string, options ScsOptions) string {
	if options.Deterministic {
		sort.Strings(words)
	}
	var wg *sync.WaitGroup
	if options.Parallel {
		wg = &sync.WaitGroup{}
	}

	linking := createLinks(words, wg)
	return walkLinks(linking)
}

func validate(value string, words []string) (bool, []string) {
	invalidWords := []string{}
	for _, word := range words {
		offset := 0
		for i := 0; i < len(word); i++ {
			char := word[i]
			found := false
			for !found && offset < len(value) {
				if value[offset] == char {
					found = true
				}
				offset++
			}
			if !found && offset >= len(value) && i <= len(word) {
				invalidWords = append(invalidWords, word)
				break
			}
		}
	}
	return len(invalidWords) == 0, invalidWords
}
