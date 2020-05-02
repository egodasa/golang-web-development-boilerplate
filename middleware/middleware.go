// Packacge middleware untuk menampung fungsi yang berhubungan dengan middleware
package middleware

import (
  hpl "belajar-ech0-framework/helper"
  "os"
  "strings"

  gin "github.com/gin-gonic/gin"
)

func ValidasiTokenJwt(c *gin.Context) {
  var cekToken bool
  token := c.Request.Header.Get("Authorization")
  if token != "" {
    tokenJwt := strings.Split(token, " ")
    if len(tokenJwt) == 2 {
      if tokenJwt[0] == "Bearer" {
        cekToken = hpl.ValidateJwtToken(tokenJwt[1], os.Getenv("JWT_KEY"))
      }
    }
  }

  if cekToken == true {
    c.Next()
  } else {
    c.AbortWithStatusJSON(403, map[string]interface{}{
      "code":    403,
      "message": "Token not valid!",
    })
  }

}
