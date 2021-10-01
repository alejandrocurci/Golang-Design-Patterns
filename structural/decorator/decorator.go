package decorator

// it lets you attach new behaviors to objects by placing this objects inside special wrappers
// both target objects and decorators follow the same interface

// example: pizzas which are enhanced/decorated with special ingredients
// pizza is a component, both concrete components and decorators need to implement this interface
type pizza interface {
	getPrice() int
}

// COMPARISON TO SIMILAR PATTERNS
// decorator vs adapter:
// adapter -> changes the interface of an existing object
// decorator -> enhances the object without changing the interface

// decorator vs proxy:
// decorator -> wraps a type dynamically at runtime,
// the decoration may or may not be there or it may be composed of one or many types
// proxy -> wraps a type at compile time, it's a way to access some type
