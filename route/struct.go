package route

import "net/http"

type Route struct {
	// Request handler for the route.
	handler http.Handler
	// If true, this route never matches: it is only used to build URLs.
	buildOnly bool
	// The name used to build URLs.
	name string
	// Error resulted from building a route.
	err error

	// "global" reference to all named routes
	namedRoutes map[string]*Route

	// config possibly passed in from `Router`
	routeConf
}

type routeConf struct {
	// If true, "/path/foo%2Fbar/to" will match the path "/path/{var}/to"
	useEncodedPath bool

	// If true, when the path pattern is "/path/", accessing "/path" will
	// redirect to the former and vice versa.
	strictSlash bool

	// If true, when the path pattern is "/path//to", accessing "/path//to"
	// will not redirect
	skipClean bool

	// If true, the http.Request context will not contain the Route.
	omitRouteFromContext bool

	// Manager for the variables from host and path.
	regexp routeRegexpGroup

	// List of matchers.
	matchers []matcher

	// The scheme used when building URLs.
	buildScheme string

	buildVarsFunc BuildVarsFunc
}

type matcher interface {
	Match(*http.Request, *RouteMatch) bool
}

type routeRegexpGroup struct {
	host    *routeRegexp
	path    *routeRegexp
	queries []*routeRegexp
}

type Router struct {
	// Configurable Handler to be used when no route matches.
	// This can be used to render your own 404 Not Found errors.
	NotFoundHandler http.Handler

	// Configurable Handler to be used when the request method does not match the route.
	// This can be used to render your own 405 Method Not Allowed errors.
	MethodNotAllowedHandler http.Handler

	// Routes to be matched, in order.
	routes []*Route

	// Routes by name for URL building.
	namedRoutes map[string]*Route

	// If true, do not clear the request context after handling the request.
	//
	// Deprecated: No effect, since the context is stored on the request itself.
	KeepContext bool

	// Slice of middlewares to be called after a match is found
	middlewares []middleware

	// configuration shared with `Route`
	routeConf
}
