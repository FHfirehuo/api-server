package user

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/logger"
	"apiserver/util"
	"github.com/gin-gonic/gin"
)

func Create(c *gin.Context) {
	logger.Info("User Create function called. X-Request-Id: %s", util.GetReqID(c))

	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}
	u := model.UserModel{
		Username: r.Username,
		Password: r.Password,
	}
	if err := u.Validate(); err != nil {
		handler.SendResponse(c, err, nil)
		return
	}

	if err := u.Create(); err != nil {
		handler.SendResponse(c, err, nil)
	}

	rsp := CreateResponse{
		Username: r.Username,
	}

	handler.SendResponse(c, nil, rsp)
}
