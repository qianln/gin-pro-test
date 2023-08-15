package authorization

import (
	"github.com/gin-gonic/gin"
)

type HeaderParams struct {
	Authorization string `header:"Authorization" binding:"required,min=20"`
}

func CheckTokenAuth() gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Next()
	}
	//
	//	headerParams := HeaderParams{}
	//
	//	//  推荐使用 ShouldBindHeader 方式获取头参数
	//	if err := context.ShouldBindHeader(&headerParams); err != nil {
	//		response.TokenErrorParam(context, consts.JwtTokenMustValid+err.Error())
	//		return
	//	}
	//	token := strings.Split(headerParams.Authorization, " ")
	//	if len(token) == 2 && len(token[1]) >= 20 {
	//		//tokenIsEffective := web.NewUsersServices().IsEffective(token[1])
	//		if tokenIsEffective {
	//			signKey := system.Config.GetString("Token.JwtTokenSignKey")
	//			if customToken, err := my_jwt.NewMyJwt(signKey).ParseToken(token[1]); err == nil {
	//				key := system.Config.GetString("Token.BindContextKeyName")
	//				// token验证通过，同时绑定在请求上下文
	//				context.Set(key, customToken)
	//			}
	//			context.Next()
	//		} else {
	//			response.ErrorTokenAuthFail(context)
	//		}
	//	} else {
	//		response.ErrorTokenBaseInfo(context)
	//	}
	//}
	//context.Next()
}
