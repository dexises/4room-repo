package router

// import "net/http"

// type Handle func(http.ResponseWriter, *http.Request, Params)

// type Param struct {
// 	Key   string
// 	Value string
// }

// type Params []Params

// type Router struct {
// 	routes map[string]*node
// }

// type Endpoint struct {
// 	Path    string
// 	Method  string
// 	Handler http.HandlerFunc
// }

// func (sr *Router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	http.NotFound(w, r)
// }

// func (r *Router) Route(method, path string, handlerFucn http.HandlerFunc) {
// 	e := Endpoint{
// 		Method:  method,
// 		Path:    path,
// 		Handler: handlerFucn,
// 	}
// }

// // GET is a shortcut for router.Handle(http.MethodGet, path, handle)
// func (r *Router) GET(path string, handle Handle) {
// 	r.Handle(http.MethodGet, path, handle)
// }

// // HEAD is a shortcut for router.Handle(http.MethodHead, path, handle)
// func (r *Router) HEAD(path string, handle Handle) {
// 	r.Handle(http.MethodHead, path, handle)
// }

// // OPTIONS is a shortcut for router.Handle(http.MethodOptions, path, handle)
// func (r *Router) OPTIONS(path string, handle Handle) {
// 	r.Handle(http.MethodOptions, path, handle)
// }

// // POST is a shortcut for router.Handle(http.MethodPost, path, handle)
// func (r *Router) POST(path string, handle Handle) {
// 	r.Handle(http.MethodPost, path, handle)
// }

// // PUT is a shortcut for router.Handle(http.MethodPut, path, handle)
// func (r *Router) PUT(path string, handle Handle) {
// 	r.Handle(http.MethodPut, path, handle)
// }

// // PATCH is a shortcut for router.Handle(http.MethodPatch, path, handle)
// func (r *Router) PATCH(path string, handle Handle) {
// 	r.Handle(http.MethodPatch, path, handle)
// }

// // DELETE is a shortcut for router.Handle(http.MethodDelete, path, handle)
// func (r *Router) DELETE(path string, handle Handle) {
// 	r.Handle(http.MethodDelete, path, handle)
// }
