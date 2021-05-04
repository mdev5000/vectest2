package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mdev5000/vectest2/ajax"
	"net/http"
)

func Serve() (*echo.Echo, error) {
	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/public", "public")

	// Routes
	e.POST("/api/first", api)
	e.GET("/main/*", root)

	// Start server
	return e, e.Start(":1323")
}

func api(c echo.Context) error {
	return c.JSON(http.StatusOK, &ajax.User{Name: "bob"})
}

func root(c echo.Context) error {
	return c.HTML(http.StatusOK, `
<!DOCTYPE html>
<!-- Polyfill for the old Edge browser -->
<script src="https://cdn.jsdelivr.net/npm/text-encoding@0.7.0/lib/encoding.min.js"></script>
<script src="/public/wasm/wasm_exec.js"></script>
<script>
    (async () => {
        const resp = await fetch('/public/wasm/main.wasm');
        if (!resp.ok) {
            const pre = document.createElement('pre');
            pre.innerText = await resp.text();
            document.body.appendChild(pre);
            return;
        }
        const src = await resp.arrayBuffer();
        const go = new Go();
        const result = await WebAssembly.instantiate(src, go.importObject);
        go.argv = [];
        go.run(result.instance);
    })();
</script>
`)
}
