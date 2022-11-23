package car

import (
	"github.com/labstack/echo"
	"net/http"
)

func RouteInit() *echo.Echo {
	e := echo.New()
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "ok")
	})
	policyHistory := &PolicyHistory{}
	e.GET("/policyHistories", policyHistory.Get)
	e.GET("/policyHistories/:id", policyHistory.FindById)
	e.POST("/policyHistories", policyHistory.Persist)
	e.PUT("/policyHistories/:id", policyHistory.Put)
	e.DELETE("/policyHistories/:id", policyHistory.Remove)
	return e
}
