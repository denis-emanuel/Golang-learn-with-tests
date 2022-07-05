package concurrency

// WebsiteChecker checks a url, returning a bool.
type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

// /CheckWebsites takes a WebsiteChecker and a slice of urls and returns  a map.
// of urls to the result of checking each url with the WebsiteChecker function.
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	// chan result -> the type of channel, a channel of 'result'
	resultChannel := make(chan result)

	for _, url := range urls {
		//anonymous function:
		go func(u string) {
			//send to channel statement:
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		//receive from the channel statement:
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}