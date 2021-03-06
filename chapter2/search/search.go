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
		/* Goroutine 函数怎么传入参数？ 为什么只传入了Matcher, Feed这两个参数呢？ 因为matcher、feed每次调用值不同，所以要当变量传进来。
		定义值传递和指针传递有什么区别吗？ 为什么map不用指针，feed类型用指针呢？
		make map返回的是指针类型的，map的值也是一个指针类型。值传递只是原来的值的一个副本，指针传递也是传的一个指针的副本，通过指针的指针去修改它的值。
		为什么 searchTerm、results可以直接获取匿名函数外面的参数。
		searchTerm和results是使用闭包获取函数外部的变量。
		什么是闭包呢？
		可以包含自由（未绑定到特定对象）变量的代码块，这些变量不在这个代码块内或者
		任何全局上下文中定义，而是在定义代码块的环境中定义。
		给自由变量绑定特定语法环境，使其完成闭合。例如： x是一个自由变量， f(x)（x+2）也是自由变量，
		所以f(x)可以看成变量 y 作为参数传递，y的结果是在使用f(x)的代码环境有关的。
		如：f z() {
			x = 2
			f(x) (x+2) (x)  //把x传入到f（x）中。这时候f(x)才有了确切的定义。
		}
		*/
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

/* 注册器相当于在matchers map 中保存一个数据
 *
 */
func Register(feedType string, matcher Matcher) {
	if _, exists := matchers[feedType]; exists {
		log.Fatalln(feedType, "Matcher already registered")
	}

	log.Println("Register", feedType, "matcher")
	matchers[feedType] = matcher
}
