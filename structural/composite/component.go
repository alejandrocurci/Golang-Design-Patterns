package composite

// this pattern allows composing objects into a tree-like structure and work with it as if it was a singular object
// its great feature is the ability to run methods recursively over the whole tree structure and sum up the results

// Component –> it is the interface which defines the common operations for both the Composite and Leaf objects
// Composite –> it implements Component interface and embeds an array of child Components
// Leaf –> it is the primitive object in the tree. It also implements the Component Interface

type component interface {
	search(string)
}

// example -> tree-like structure composed by files and folders
