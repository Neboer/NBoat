package cookieauth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func SetIdentity(server *gin.Engine, authString string) {
	server.GET("/"+authString, func(context *gin.Context) {
		context.SetCookie("identity", authString, 0, "/", "", false, false)
		context.Abort()
	})
}

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

func IsAuthed(ctx *gin.Context) bool {
	value, exists := ctx.Get("identity")
	if exists && value == true {
		return true
	} else {
		return false
	}
}

// 仅允许站长访问的api。限制用户访问以block开头的url，禁止用户对这个URL使用一切方法。
func OnlyAllowAuthor(blockUrlPathPrefix string) gin.HandlerFunc {
	return func(context *gin.Context) {
		// 受管制页面
		if strings.HasPrefix(context.Request.URL.Path, blockUrlPathPrefix) {
			isAuth := context.GetBool("identity")
			if isAuth == false {
				context.String(400, "access denied")
				context.Abort()
			} else {
				context.Next()
			}
		} else {
			context.Next()
		}
	}
}
