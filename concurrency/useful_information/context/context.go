package context

import (
	"context"
	"fmt"
	"time"
)

// using the context data type is the idiomatic way to pass "common scoped data" within our application:
// two main purposes:
// 1- cancellation signals to terminate operations performed by goroutines (to prevent our system to do unnecessary work)
// 2- pass information required at every function call invoked by an operation (request ID of an HTTP request call)

// there are two sides to context cancellation
// 1- listening for the cancellation event
// context.Done() returns a channel that receives an empty struct{} every time the context receives  cancellation event
// we need to wait on "<-ctx.Done()"
// 2- emitting the cancellation event
// it can be triggered through the function cancel() from a context.WithCancel(), WithTimeout(), WithDeadline()

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	ctx, cancel = context.WithTimeout(ctx, 300*time.Millisecond)
	ctx, cancel = context.WithDeadline(ctx, time.Date(2021, 3, 18, 0, 0, 0, 0, time.UTC))
	defer cancel() // call this function to emit the cancellation signal
	// cancelling the parent context after finishing the process frees resources allocated for timers

	key := "request_id"
	ctx = context.WithValue(context.Background(), key, 81921)
	id := ctx.Value(key)
	fmt.Println("Request ID: ", id)
	// only use this feature to pass down operation-scoped information (contextual or common data)
	// the value cannot be modified, a new context.WithValue() needs to be created instead
}

// parent and child contexts -> they build a node tree -> if the parent gets cancelled, every child gets cancelled as well
// if this is not the intended behaviour, instead of deriving a context, a new one should be created with context.Background()

// BEST PRACTICES
/*
### context.Background() should be used only at the highest level, as the root of all derived contexts
### context.TODO() should be used where not sure what to use or if the current function will be updated to use context in future
### context cancellations are advisory, the functions may take time to clean up and exit
### context.Value() should be used very rarely, it should never be used to pass in optional parameters.
	This makes the API implicit and can introduce bugs. Instead, such values should be passed in as arguments.
### Don’t store contexts in a struct, pass them explicitly in functions, preferably, as the first argument.
### Never pass nil context, instead, use a context.TODO() if you are not sure what to use.
### The Context struct does not have a cancel method because only the function that derives the context should cancel it.
*/
