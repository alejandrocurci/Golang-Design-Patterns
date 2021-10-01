package proxy

// nginx is the proxy, it controls access to the main application
// it implements the same interface server
// the proxy wraps the main object
type nginx struct {
	app         *application
	maxRequest  int
	rateLimiter map[string]int
}

func newProxy() *nginx {
	return &nginx{
		app:         &application{},
		maxRequest:  5,
		rateLimiter: make(map[string]int),
	}
}

// handleRequest checks whether the rate limit has been reached
// then it delegates the request to the actual application
func (n *nginx) handleRequest(url, method string) (code int, response string) {
	if !n.checkRateLimit(url) {
		return 403, "Not allowed"
	}
	return n.app.handleRequest(url, method)
}

// checkRateLimit returns false when the rate limit has been reached
func (n *nginx) checkRateLimit(url string) bool {
	if n.rateLimiter[url] == n.maxRequest {
		return false
	}
	n.rateLimiter[url]++
	return true
}
