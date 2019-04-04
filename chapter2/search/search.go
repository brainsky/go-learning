package search

import (
	"log"
	"sync"
)

var matchers = make(map[string]Matcher)

func Run(searchTerm string) {
	//get  sources list that need to be search
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	//create no buffer channel, receieve result
	results := make(chan *Result)

	// 构造一个 waitGroup,以便处理所有的数据源
	var waitGroup sync.WaitGroup

	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)

	}

	go func() {
		// 等候所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式,通知 Display 函数
		close(results)
	}()

	Display(results)

}

func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
