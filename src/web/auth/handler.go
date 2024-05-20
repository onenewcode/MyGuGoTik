package auth

import (
	"GuGoTik/src/constant/config"
	"GuGoTik/src/constant/strings"
	"GuGoTik/src/extra/tracing"
	"GuGoTik/src/rpc/auth"
	grpc2 "GuGoTik/src/utils/grpc"
	"GuGoTik/src/utils/logging"
	"GuGoTik/src/web/models"
	"GuGoTik/src/web/utils"
	"github.com/gin-gonic/gin"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"github.com/sirupsen/logrus"
	"net/http"
)

var Client auth.AuthServiceClient

func LoginHandle(c *gin.Context) {
	var req models.LoginReq
	// “star” 创建跨度和上下文。包含新创建的跨度的上下文。
	_, span := tracing.Tracer.Start(c.Request.Context(), "LoginHandler")
	// 	End 完成 Span。跨度被认为是完整的，可以了 在此方法之后，通过遥测管道的其余部分传递
	defer span.End()
	logging.SetSpanWithHostname(span)
	logger := logging.LogService("GateWay.Login").WithContext(c.Request.Context())
	// 参数校验
	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, models.LoginRes{
			StatusCode: strings.GateWayParamsErrorCode,
			StatusMsg:  strings.GateWayParamsError,
			UserId:     0,
			Token:      "",
		})
		return
	}
	// 客户端 进行登陆
	res, err := Client.Login(c.Request.Context(), &auth.LoginRequest{
		Username: req.UserName,
		Password: req.Password,
	})
	if err != nil {
		logger.WithFields(logrus.Fields{
			"Username": req.UserName,
		}).Warnf("Error when trying to connect with AuthService")
		c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
		return
	}

	logger.WithFields(logrus.Fields{
		"Username": req.UserName,
		"Token":    res.Token,
		"UserId":   res.UserId,
	}).Infof("User log in")

	c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
}

func RegisterHandle(c *gin.Context) {
	var req models.RegisterReq
	_, span := tracing.Tracer.Start(c.Request.Context(), "LoginHandler")
	defer span.End()
	logger := logging.LogService("GateWay.Register").WithContext(c.Request.Context())

	if err := c.ShouldBindQuery(&req); err != nil {
		c.JSON(http.StatusOK, models.RegisterRes{
			StatusCode: strings.GateWayParamsErrorCode,
			StatusMsg:  strings.GateWayParamsError,
			UserId:     0,
			Token:      "",
		})
		return
	}

	res, err := Client.Register(c.Request.Context(), &auth.RegisterRequest{
		Username: req.UserName,
		Password: req.Password,
	})

	if err != nil {
		logger.WithFields(logrus.Fields{
			"Username": req.UserName,
		}).Warnf("Error when trying to connect with AuthService")
		c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
		return
	}

	logger.WithFields(logrus.Fields{
		"Username": req.UserName,
		"Token":    res.Token,
		"UserId":   res.UserId,
	}).Infof("User register in")

	c.Render(http.StatusOK, utils.CustomJSON{Data: res, Context: c})
}

func init() {
	conn := grpc2.Connect(config.AuthRpcServerName)
	Client = auth.NewAuthServiceClient(conn)
}
