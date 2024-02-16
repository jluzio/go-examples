package main

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/tour/tree"
)

func main() {
	// main_concurrency()
	// main_channel()
	// main_buffered_channels()
	// main_buffered_channels_deadlock()
	// main_range_and_close()
	// main_select()
	// main_select_default()
	// exercise_equivalent_binary_trees()
	// main_mutex()
	exercise_webcrawler()
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main_concurrency() {
	/*
		A goroutine is a lightweight thread managed by the Go runtime.
		go f(x, y, z): starts a new goroutine running f(x, y, z)
	*/
	go say("world")
	say("hello")
}

/*
Channels
Channels are a typed conduit through which you can send and receive values with the channel operator, <-.

ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and

	// assign value to v.

(The data flows in the direction of the arrow.)

Like maps and slices, channels must be created before use:

ch := make(chan int)
By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.

The example code sums the numbers in a slice, distributing the work between two goroutines. Once both goroutines have completed their computation, it calculates the final result.
*/
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}
func main_channel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

/*
Buffered Channels
Channels can be buffered. Provide the buffer length as the second argument to make to initialize a buffered channel:

ch := make(chan int, 100)
Sends to a buffered channel block only when the buffer is full. Receives block when the buffer is empty.

Modify the example to overfill the buffer and see what happens.
*/
func main_buffered_channels() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func main_buffered_channels_deadlock() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	ch <- 3
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

/*
Range and Close
A sender can close a channel to indicate that no more values will be sent. Receivers can test whether a channel has been closed by assigning a second parameter to the receive expression: after

v, ok := <-ch
ok is false if there are no more values to receive and the channel is closed.

The loop for i := range c receives values from the channel repeatedly until it is closed.

Note: Only the sender should close a channel, never the receiver. Sending on a closed channel will cause a panic.

Another note: Channels aren't like files; you don't usually need to close them. Closing is only necessary when the receiver must be told there are no more values coming, such as to terminate a range loop.
*/

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main_range_and_close() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci_select(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
func main_select() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci_select(c, quit)
}

func main_select_default() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}

// type Tree struct {
// 	Left  *Tree
// 	Value int
// 	Right *Tree
// }

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for i := 1; i <= 10; i++ {
		v1, v2 := <-ch1, <-ch2
		if v1 != v2 {
			return false
		}
	}
	return true
}

func exercise_equivalent_binary_trees() {
	// ch := make(chan int)
	// go Walk(tree.New(1), ch)
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(<-ch)
	// }

	fmt.Printf("%v = %v\n", "1 and 1", Same(tree.New(1), tree.New(1)))
	fmt.Printf("%v = %v\n", "1 and 2", Same(tree.New(1), tree.New(2)))
}

// SafeCounter is safe to use concurrently.
type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

// Inc increments the counter for the given key.
func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

// Value returns the current value of the counter for the given key.
func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}
func main_mutex() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type FetchJob struct {
	status string
	url    string
	depth  int
	data   FetchedData
}
type FetchedData struct {
	body string
	urls []string
	err  error
}

type Crawler struct {
	cache      map[string]FetchJob
	cacheMutex sync.Mutex
}

var crawler = Crawler{
	make(map[string]FetchJob),
	sync.Mutex{},
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	ch := make(chan FetchJob)
	go CrawlJob(url, depth, fetcher, ch)
	<-ch
	close(ch)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func CrawlJob(url string, depth int, fetcher Fetcher, ch chan FetchJob) {
	// fmt.Printf("processing :: url:%v | depth:%v\n", url, depth)
	if depth <= 0 {
		fmt.Printf("%v - Aborted by depth\n", url)
		ch <- FetchJob{status: "Aborted by depth"}
	}

	crawler.cacheMutex.Lock()
	cachedJob, ok := crawler.cache[url]
	if ok {
		fmt.Printf("Cache Hit: %v\n", url)
		crawler.cacheMutex.Unlock()
		ch <- cachedJob
		return
	}

	crawler.cache[url] = FetchJob{status: "pending", url: url, depth: depth}
	crawler.cacheMutex.Unlock()

	body, urls, err := fetcher.Fetch(url)
	crawler.cacheMutex.Lock()
	job := crawler.cache[url]
	job.data = FetchedData{body, urls, err}
	crawler.cacheMutex.Unlock()

	if err != nil {
		fmt.Println(err)
		ch <- job
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	childrenCh := make(chan FetchJob, len(urls))
	for _, u := range urls {
		go CrawlJob(u, depth-1, fetcher, childrenCh)
	}
	for range urls {
		<-childrenCh
		// subJob := <-childrenCh
		// fmt.Printf("subJob: %v\n", subJob)
	}

	ch <- job
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

func exercise_webcrawler() {
	Crawl("https://golang.org/", 4, fetcher)
}
