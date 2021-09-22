package prototype

// this pattern allows to create copies of objects
// the responsibility of cloning an object is delegated to the actual object to clone (called prototype)
// the prototype exposes a clone method which returns a copy of its own

type prototype interface {
	clone() prototype
}

// both file and folder types are prototypes
// they can be cloned to create new instances

type file struct {
	name string
}

func (f file) clone() prototype {
	return &file{name: f.name}
}

type folder struct {
	name  string
	files []prototype
}

func (f folder) clone() prototype {
	cloneFiles := make([]prototype, len(f.files))
	for _, c := range f.files {
		cloneFiles = append(cloneFiles, c.clone())
	}
	return &folder{name: f.name, files: cloneFiles}
}
