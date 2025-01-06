package main

import (
	"html/template"
	"log"
	"net/http"
	"path"
	"strings"
	"time"
)

// Context 封装了请求和响应的上下文
type Context struct {
	Writer  http.ResponseWriter // 用于向客户端发送 HTTP 响应
	Request *http.Request       // 客户端的 HTTP 请求
	Params  map[string]string   // URL 参数（未在此代码中使用）
}

// HandlerFunc 定义了处理函数的类型，接收一个 Context 对象作为参数
type HandlerFunc func(*Context)

// Engine 是 Web 框架的核心结构
type Engine struct {
	router        map[string]HandlerFunc // 路由映射表，存储路径和对应的处理函数
	staticPrefix  string                 // 静态文件服务的前缀
	htmlTemplates *template.Template     // HTML 模板
}

// ServeHTTP 实现了 http.Handler 接口，用于处理 HTTP 请求
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 如果请求路径以静态文件前缀开头，则提供静态文件服务
	if strings.HasPrefix(r.URL.Path, e.staticPrefix) {
		file := path.Join(".", r.URL.Path) // 拼接文件路径
		http.ServeFile(w, r, file)         // 提供静态文件
		return
	}

	// 查找路由映射表中是否存在对应的处理函数
	if handler, ok := e.router[r.URL.Path]; ok {
		c := &Context{Writer: w, Request: r} // 创建 Context 对象
		handler(c)                           // 调用处理函数
	} else {
		http.NotFound(w, r) // 如果路径未注册，返回 404 Not Found
	}
}

// NewEngine 创建一个新的 Engine 实例
func NewEngine() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)} // 初始化路由映射表
}

// GET 注册一个 GET 请求的路由
func (e *Engine) GET(path string, handler HandlerFunc) {
	e.router[path] = handler // 将路径和处理函数添加到路由映射表
}

// Static 注册静态文件服务
func (e *Engine) Static(prefix string, root string) {
	e.staticPrefix = prefix // 设置静态文件前缀
	// 使用 http.FileServer 提供静态文件服务，并去掉 URL 中的前缀
	http.Handle(prefix, http.StripPrefix(prefix, http.FileServer(http.Dir(root))))
}

// LoadHTMLGlob 加载 HTML 模板文件
func (e *Engine) LoadHTMLGlob(pattern string) {
	// 使用 template.Must 加载模板文件，如果失败则 panic
	e.htmlTemplates = template.Must(template.ParseGlob(pattern))
}

// HTML 渲染 HTML 模板并发送响应
func (e *Engine) HTML(c *Context, name string, data interface{}) {
	c.Writer.Header().Set("Content-Type", "text/html") // 设置响应头为 HTML 类型
	// 执行模板渲染，并将结果写入 ResponseWriter
	if err := e.htmlTemplates.ExecuteTemplate(c.Writer, name, data); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError) // 如果渲染失败，返回 500 错误
		c.Writer.Write([]byte("Internal Server Error"))
	}
}

func main() {
	e := NewEngine() // 创建 Engine 实例

	// 注册静态文件服务，将 /assets 路径映射到 ./static 目录
	e.Static("/assets", "./static")

	// 加载 HTML 模板文件
	// 使用相对路径（如果工作目录正确）
	// e.LoadHTMLGlob("templates/*")
	// 使用绝对路径（确保路径正确）
	e.LoadHTMLGlob("G:/GoProject/Gee/web/day6/demo/templates/*")

	// 注册根路径 "/" 的处理函数
	e.GET("/", func(c *Context) {
		// 渲染 index.html 模板，并传递数据
		e.HTML(c, "index.html", map[string]string{
			"Title": "Welcome to Gee Web Framework",           // 页面标题
			"Time":  time.Now().Format("2006-01-02 15:04:05"), // 当前时间
		})
	})
	// 启动 HTTP 服务器，监听 8080 端口
	log.Fatal(http.ListenAndServe(":8080", e))
}
