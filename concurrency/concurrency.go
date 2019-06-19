package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, websites []string) map[string]bool {

	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range websites {
		url := url
		go func() { //// FIXME:
			//results[u] = wc(u)
			resultChannel <- result{url, wc(url)}
		}()
	}

	for i := 0; i < len(websites); i++ {
		result := <-resultChannel
		results[result.string] = result.bool
	}
	return results
}
