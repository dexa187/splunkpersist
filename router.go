package splunkpersist

type fn func(Request) Response

// Router calls func based on path
type Router struct {
	routes map[string]fn
}

// NewRouter Create new routeer
func NewRouter() *Router {
	return &Router{
		routes: make(map[string]fn),
	}
}

//Add new path
func (r *Router) Add(path string, handle fn) {
	r.routes[path] = handle
}
