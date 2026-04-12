package middleware

import (
	"bytes"
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

func decryptAES(cipherTextBase64 string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(cipherTextBase64)
}

// 加密解密中间件：仅处理POST/PUT请求的解密
func DecryptMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 1. 仅处理POST/PUT方法，其他方法直接放行
		method := c.Request.Method
		if method != http.MethodPost && method != http.MethodPut {
			c.Next()
			return
		}

		// 2. 读取原始请求体
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "读取请求体失败"})
			c.Abort()
			return
		}
		// 读取后必须关闭原始Body
		defer c.Request.Body.Close()

		// 3. 解密请求体（空请求体直接放行）
		if len(bodyBytes) > 0 {
			decryptedBytes, err := decryptAES(string(bodyBytes))
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": "解密失败：" + err.Error()})
				c.Abort()
				return
			}
			panic(decryptedBytes)

			// 4. 重置请求体：将解密后的内容放回Request.Body
			// 注意：必须用io.NopCloser包装，因为Body需要实现ReadCloser接口
			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
			// 重置Content-Length（避免后续读取长度异常）
			c.Request.ContentLength = int64(len(bodyBytes))
		}

		// 5. 放行，执行后续处理器
		c.Next()
	}
}
