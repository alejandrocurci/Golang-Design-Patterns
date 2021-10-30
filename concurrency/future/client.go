package future

import (
	"errors"
	"fmt"
)

func main() {
	future := &promise{}

	// set success and failure cases
	future.success(func(s string) {
		fmt.Println("Success: ", s)
	}).fail(func(err error) {
		fmt.Println("Error: ", err)
	})

	// execute some function asynchronously
	// it will trigger successFn
	future.execute(func() (string, error) {
		return "Hello success", nil
	})

	// execute some function asynchronously
	// it will trigger failFn
	future.execute(func() (string, error) {
		return "", errors.New("sth went wrong")
	})

}
