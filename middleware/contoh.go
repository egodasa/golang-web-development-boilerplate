// Packacge middleware untuk menampung fungsi yang berhubungan dengan middleware
package middleware

import (
  "strings"
  hpl "belajar-ech0-framework/helper"

	"github.com/labstack/echo"
)

func CekLogin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
    if c.Request().Header.Get("Authorization") != "" {
      token := strings.Split(c.Request().Header.Get("Authorization"), " ");
      
      if len(token) != 2 {
        return c.JSON(401, struct{code int; message string;}{
          code: 401,
          message: "Token not valid",
        });
      } else {
          if token[0] != "Bearer" && hpl.ValidateJwtToken(token[1]) != true {
            return c.JSON(401, struct{code int; message string;}{
              code: 401,
              message: "Token not valid",
            });
          }
      }
    } else {
      return c.JSON(401, struct{code int; message string;}{
          code: 401,
          message: "Token not valid",
        });
    }
    return next(c)
	}
}
