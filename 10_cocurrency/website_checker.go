package main

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		/*
			By giving each anonymous function a parameter for the url - u - and then calling the anonymous function with
			the url as the argument, we make sure that the value of u is fixed as the value of url for the iteration
			of the loop that we're launching the goroutine in. u is a copy of the value of url, and so can't be changed.
		*/
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
