# go-http

golang http server

基于 [gin-gonic/gin](https://github.com/gin-gonic/gin)

## 使用例子

```go

func TestRunServer(t *testing.T) {
	engine := New()

	var handler = func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}

	RegisterRoutes(engine, []*Route{NewRoute("GET", "ping", handler)})

	if err := RunServer(engine); err != nil {
		panic(err)
	}
}

```

## todo

日志/健康...