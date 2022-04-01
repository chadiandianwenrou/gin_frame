package router

import (
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	r := gin.Default()
	//跨域配置
	//r.Use(cors.Cors())

	//JWT配置
	//authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
	//	Realm:       config.Conf.Http.Domain,
	//	Key:         []byte(config.Conf.Http.SecretKey),
	//	PayloadFunc: cors.PayloadFunc_,
	//	IdentityHandler: cors.IdentityHandlerFunc_,
	//	CookieName: "token",
	//	SendCookie: true,
	//	SecureCookie: false,
	//	CookieHTTPOnly: true,
	//	CookieDomain: config.Conf.Http.Domain,
	//	TokenLookup: "cookie:token, header: token, query: token",
	//	TimeFunc: time.Now,
	//})

	//if err != nil {
	//	log.Error("JWT Error:" + err.Error())
	//}

	r = NoAuthRouter(r)
	
	//全局404配置
	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	//r   = FsRouter(r,authMiddleware)
	//r	= MailRouter(r,authMiddleware)
	//r	= SmsTextRouter(r,authMiddleware)
	//r	= SmsVoiceRouter(r,authMiddleware)

	return r
}
