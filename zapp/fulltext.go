package zapp

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"google.golang.org/appengine/log"
	"google.golang.org/appengine/search"
)

type Fulltext struct {
	Ngram     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func SplitNgrams(text string, n int) []string {
	sepText := strings.Split(text, "")
	var ngrams []string
	if len(sepText) < n {
		return nil
	}
	for i := 0; i < (len(sepText) - n + 1); i++ {
		ngrams = append(ngrams, strings.Join(sepText[i:i+n], ""))
	}
	return ngrams
}

func SplitNgramsRange(text string, size int) (ngrams []string) {
	for i := 0; i <= size; i++ {
		ngrams = append(ngrams, SplitNgrams(text, i)...)
	}
	return
}

func WordToSplittedWords(word string) []string {
	qSize := len(strings.Split(word, ``))

	var splits []string
	switch qSize {
	case 0:
		splits = []string{}
	case 1, 2, 3:
		splits = []string{word}
	default:
		splits = SplitNgrams(word, 3)
		splits = append(splits, word[len(word)-3:len(word)])
	}

	uniq := UniqString(splits)
	return uniq
}

func SearchByGAEFulltext(ctx context.Context, idx string, words []string) []int {
	values := []string{}

	for _, word := range words {
		//　文字を指定字単位で分割し配列にする ngram
		ngrams := WordToSplittedWords(word)
		for _, s := range ngrams {
			values = append(values, fmt.Sprintf(`Ngram="%s"`, s))
		}
	}
	query := strings.Join(values, ` AND `)

	searchAPIIndex, _ := search.Open(idx)
	iterator := searchAPIIndex.Search(ctx, query, &search.SearchOptions{IDsOnly: true})
	log.Debugf(ctx, `%s`, values)
	log.Debugf(ctx, `%s`, query)
	// iterator := searchAPIIndex.Search(ctx, query, &search.SearchOptions{})

	var IDs []int
	for {
		sid, err := iterator.Next(nil)
		if err == search.Done {
			break
		} else if err != nil {
			break
		}
		ID, _ := strconv.Atoi(sid)
		IDs = append(IDs, ID)
	}
	return IDs
}

func QueryToSearchwords(q string) (searchWords []string){
	// 半角・全角で区切る
	for _, word := range strings.Split(q, ` `) { // 半角
		for _, w := range strings.Split(word, `　`) { // 全角
			if w != `` {
				searchWords = append(searchWords, w)
			}
		}
	}
	return
}


func FindBySeachGAE(ctx context.Context, idx string, q string) (anyIDs []int) {
	q = strings.Trim(q, ` `) // 前後の空白を除く
	searchWords := QueryToSearchwords(q)
	anyIDs = SearchByGAEFulltext(ctx, idx, searchWords)
	return
}
