package auth

import "github.com/gin-gonic/gin"

// 鉴权。校验cookie是否和访问的站点相匹配，将是否为校验的请求信息直接绑定到了请求对象身上。
func AuthenticGate(secretCookieString string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		identityValue, parseError := ctx.Cookie("identity")
		if parseError != nil || identityValue != secretCookieString {
			ctx.Set("identity", false)
		} else {
			ctx.Set("identity", true)
			ctx.Next()
		}
	}
}

// 仅允许站长访问的api
func OnlyAllowAuthor() gin.HandlerFunc {
	return func(context *gin.Context) {
		isAuth, _ := context.Get("identity")
		if isAuth == false {
			context.AbortWithStatus(400)
		} else {
			context.Next()
		}
	}
}
