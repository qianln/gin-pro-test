package http

import (
	"encoding/json"
	"gin-pro/app/core/system"
	"github.com/gin-gonic/gin"
	"time"
)

type IValidator interface {
	CheckParams(context *gin.Context)
}

func BindContext(validatorInterface IValidator, context *gin.Context) *gin.Context {
	var tempJson interface{}
	if tmpBytes, err1 := json.Marshal(validatorInterface); err1 == nil {
		if err2 := json.Unmarshal(tmpBytes, &tempJson); err2 == nil {
			if value, ok := tempJson.(map[string]interface{}); ok {
				for k, v := range value {
					context.Set(system.ParamPrefix+k, v)
				}
				curDateTime := time.Now().Format(system.DateFormat)
				context.Set(system.ParamPrefix+"created_at", curDateTime)
				context.Set(system.ParamPrefix+"updated_at", curDateTime)
				context.Set(system.ParamPrefix+"deleted_at", curDateTime)
				return context
			}
		}
	}
	return nil
}
