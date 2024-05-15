package scratch

import (
	"strings"
	"sync"
)

type WordInfo struct {
	TF          int // Text Frequency
	DF          int // Document Frequency
	SearchCount int // Number of times the word was searched
}

type wordIndex struct {
	words map[string]*WordInfo
	mu    sync.Mutex
}

var index *wordIndex = nil

func newWordIndex() *wordIndex {
	if index == nil {
		index = &wordIndex{
			words: make(map[string]*WordInfo),
		}
	}

	return index
}

func (wi *wordIndex) addWord(word string) {
	wi.mu.Lock()
	defer wi.mu.Unlock()

	if _, ok := wi.words[word]; !ok {
		wi.words[word] = &WordInfo{}
	}
	wi.words[word].SearchCount++
}

func (wi *wordIndex) updateWordInfo(word string, tf, df int) {
	wi.mu.Lock()
	defer wi.mu.Unlock()

	info, ok := wi.words[word]
	if !ok {
		info = &WordInfo{}
		wi.words[word] = info
	}
	info.TF = tf
	info.DF = df
}

func (wi *wordIndex) getWordInfo(word string) *WordInfo {
	info, ok := wi.words[word]
	if !ok {
		wi.words[word] = &WordInfo{}
	}

	return info
}

func (wi *wordIndex) searchWord(word string, fileContents map[string]string) *WordInfo {
	wi.mu.Lock()
	defer wi.mu.Unlock()

	wg := sync.WaitGroup{}

	for _, content := range fileContents {
		go func() {
			wg.Add(1)
			defer wg.Done()
			wi.searchWordInFile(word, content)
		}()
	}

	wg.Wait()

	return wi.getWordInfo(word)
}

func (wi *wordIndex) searchWordInFile(word, content string) {
	info := wi.getWordInfo(word)

	tf := strings.Count(content, word)

	wi.mu.Lock()
	info.SearchCount++
	if tf > 0 {
		info.TF += tf
		info.DF += 1
	}

	wi.mu.Unlock()
}

func SearchWord(word string, fileContents map[string]string) (*WordInfo, error) {
	index := newWordIndex()

	return index.searchWord(word, fileContents), nil
}
