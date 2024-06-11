package route

func NewRouter() *Router {
	return &Router{namedRoutes: make(map[string]*Route)}
}
