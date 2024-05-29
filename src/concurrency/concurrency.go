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

type SearchEngine struct {
	index *wordIndex
}

func NewSearchEngine() *SearchEngine {
	return &SearchEngine{
		index: &wordIndex{
			words: make(map[string]*WordInfo),
		},
	}
}

func (e *SearchEngine) SearchWords(words []string, fileContents map[string]string) (map[string]*WordInfo, error) {
	result := make(map[string]*WordInfo)

	wg := sync.WaitGroup{}

	for _, word := range words {
		wg.Add(1)
		go func() {
			defer wg.Done()

			info, _ := e.SearchWord(word, fileContents)

			result[word] = info
		}()
	}

	wg.Wait()

	return result, nil
}

func (e *SearchEngine) SearchWord(word string, fileContents map[string]string) (*WordInfo, error) {
	return e.index.searchWord(word, fileContents), nil
}

type wordIndex struct {
	words map[string]*WordInfo
	mu    sync.Mutex
}

func (wi *wordIndex) searchWord(word string, fileContents map[string]string) *WordInfo {
	info := wi.getWordInfo(word)
	info.SearchCount++
	info.TF = 0
	info.DF = 0

	wg := sync.WaitGroup{}

	for _, content := range fileContents {
		wg.Add(1)

		go func(content string) {
			defer wg.Done()

			tf := wi.countWordOccurrenceInFile(word, content)

			if tf > 0 {
				wi.mu.Lock()

				info.TF += tf
				info.DF++

				wi.mu.Unlock()
			}
		}(content)
	}

	wg.Wait()

	return info
}

func (wi *wordIndex) countWordOccurrenceInFile(word, content string) int {
	return strings.Count(content, word)
}

func (wi *wordIndex) getWordInfo(word string) *WordInfo {
	info, ok := wi.words[word]
	if !ok {
		info = &WordInfo{}
		wi.words[word] = info
	}

	return info
}
