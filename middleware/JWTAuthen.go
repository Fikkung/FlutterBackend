package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func JWTAuthen() gin.HandlerFunc {
	return func(c *gin.Context) {
		// ดึง secret key จากตัวแปรสภาพแวดล้อม
		hmacSampleSecret := []byte(os.Getenv("JWT_SECRET_KEY"))

		// ดึง Authorization header จาก request
		header := c.Request.Header.Get("Authorization")

		// ดึง JWT จาก header (ลบ "Bearer " ที่อยู่ข้างหน้า)
		tokenString := strings.Replace(header, "Bearer ", "", 1)

		// แยกและตรวจสอบ JWT
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// ตรวจสอบว่าวิธีการเซ็นต์ตัวจริงตรงตามที่คาดหวัง
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}

			// ให้ secret key เพื่อการตรวจสอบ
			return hmacSampleSecret, nil
		})

		// ถ้า JWT ถูกต้อง แยก claim และตั้งค่าใน Gin context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("userId", claims["userId"])
		} else {
			// ถ้า JWT ไม่ถูกต้อง ยุติ request ด้วยสถานะ forbidden
			c.AbortWithStatusJSON(http.StatusOK, gin.H{"status": "forbidden", "message": err.Error()})
		}

		// ดำเนินการต่อไปกับ Middleware/Handler ถัดไป
		c.Next()
	}
}
