package proxy

// application is the real object, it implements the server interface
type application struct{}

func (a *application) handleRequest(url, method string) (code int, response string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}
	if url == "/create/user" && method == "POST" {
		return 202, "User created"
	}
	return 404, "Not found"
}
