package proxy

// a proxy controls access to the original object,
// allowing you to perform sth either before or after the request gets through to the original object
// the proxy implements the same interface than the real object, which makes it interchangeable

// example: nginx acts as a proxy for an application server
// it provides controlled access, rate limiting and request caching

// example 2: a cache can act as a proxy to access the real database

type server interface {
	handleRequest(url, method string) (code int, response string)
}
