package main

/*
中间件（Middleware）是 Web 框架中的一种机制，允许开发者在请求处理的前后插入自定义的逻辑。中间件通常用于处理与业务无关的通用功能，例如：

记录日志

身份验证

请求超时处理

错误恢复

性能监控

中间件的核心思想是：将非业务逻辑从业务逻辑中分离出来，提高代码的可复用性和可维护性。

中间件的设计关键点
插入点：

中间件的插入点应该在框架接收到请求后，执行用户定义的业务逻辑（Handler）之前或之后。

插入点离用户太近（如在 Handler 中手动调用）会导致代码冗余；插入点太底层（如在框架内部）会导致中间件逻辑复杂。

输入：

中间件的输入通常是请求的上下文（Context），包含请求和响应的信息。

通过 Context，中间件可以访问和修改请求、响应的内容。

执行顺序：

中间件应该支持链式调用，即一个中间件可以调用下一个中间件，最终调用业务逻辑（Handler）。

通过 Next() 方法，中间件可以在业务逻辑执行前后插入自定义逻辑。


*/
import (
	"log"
	"net/http"
	"time"
)

// Context 封装请求和响应的上下文
type Context struct {
	Writer   http.ResponseWriter
	Request  *http.Request
	handlers []HandlerFunc // 中间件和业务逻辑
	index    int           // 当前执行的中间件索引
}

// HandlerFunc 定义中间件和业务逻辑的类型
type HandlerFunc func(*Context)

// Next 执行下一个中间件或业务逻辑
func (c *Context) Next() {
	c.index++
	if c.index < len(c.handlers) {
		c.handlers[c.index](c)
	}
}

// Logger 中间件：记录请求处理时间
func Logger() HandlerFunc {
	return func(c *Context) {
		start := time.Now()
		log.Printf("Request started: %s %s", c.Request.Method, c.Request.URL.Path)
		c.Next() // 调用下一个中间件或业务逻辑
		log.Printf("Request completed: %s %s, Time: %v", c.Request.Method, c.Request.URL.Path, time.Since(start))
	}
}

// Engine 是 Web 框架的核心
type Engine struct {
	router map[string]HandlerFunc
}

// ServeHTTP 实现 http.Handler 接口
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	handlers := []HandlerFunc{Logger()} // 添加中间件
	if handler, ok := e.router[r.URL.Path]; ok {
		handlers = append(handlers, handler) // 添加业务逻辑
	}
	c := &Context{Writer: w, Request: r, handlers: handlers, index: -1}
	c.Next() // 开始执行中间件链
}

// NewEngine 创建一个新的 Engine
func NewEngine() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// GET 注册 GET 请求的路由
func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router[path] = handler
}

func main() {
	e := NewEngine()
	e.GET("/", func(c *Context) {
		c.Writer.Write([]byte("Hello, World!"))
	})
	log.Fatal(http.ListenAndServe(":8080", e))
}
