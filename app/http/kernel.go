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

func BindContext(validatorInterface IValidator, prefix string, context *gin.Context) *gin.Context {
	var tempJson interface{}
	if tmpBytes, err1 := json.Marshal(validatorInterface); err1 == nil {
		if err2 := json.Unmarshal(tmpBytes, &tempJson); err2 == nil {
			if value, ok := tempJson.(map[string]interface{}); ok {
				for k, v := range value {
					context.Set(prefix+k, v)
				}
				curDateTime := time.Now().Format(system.DateFormat)
				context.Set(prefix+"created_at", curDateTime)
				context.Set(prefix+"updated_at", curDateTime)
				context.Set(prefix+"deleted_at", curDateTime)
				return context
			}
		}
	}
	return nil
}
