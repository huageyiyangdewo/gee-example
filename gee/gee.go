package gee

import (
	"net/http"
)

type HandlerFunc func(c *Context)

type Engine struct {
	router *router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

func (engine *Engine)addRoute(method, pattern string, handler HandlerFunc)  {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *Engine)GET(patter string, handler HandlerFunc)  {
	engine.addRoute("GET", patter, handler)
}

func (engine *Engine)POST(patter string, handler HandlerFunc)  {
	engine.addRoute("POST", patter, handler)
}

func (engine *Engine)Run(add string) (err error) {
	return http.ListenAndServe(add, engine)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}