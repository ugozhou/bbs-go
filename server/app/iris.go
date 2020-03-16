package app

import (
	"bbs-go/middleware"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/go-resty/resty/v2"

	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	"github.com/kataras/iris/v12/mvc"
	"github.com/mlogclub/simple"
	"github.com/sirupsen/logrus"

	"bbs-go/common/config"
	"bbs-go/controllers/api"

	"bbs-go/controllers/admin"
)

func InitIris() {
	app := iris.New()
	app.Logger().SetLevel("warn")
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		MaxAge:           600,
		AllowedMethods:   []string{iris.MethodGet, iris.MethodPost, iris.MethodOptions, iris.MethodHead, iris.MethodDelete, iris.MethodPut},
		AllowedHeaders:   []string{"*"},
	}))
	app.AllowMethods(iris.MethodOptions)

	app.OnAnyErrorCode(func(ctx iris.Context) {
		path := ctx.Path()
		var err error
		if strings.Contains(path, "/api/admin/") {
			_, err = ctx.JSON(simple.JsonErrorCode(ctx.GetStatusCode(), "Http error"))
		}
		if err != nil {
			logrus.Error(err)
		}
	})

	app.Any("/", func(i iris.Context) {
		_, _ = i.HTML("<h1>wkycloudApp</h1>")
	})
	//配置页面路由，目前分为/api,和/api/admin/
	/*
		设计为/admin,需要登陆才能访问的数据，被要求登陆。
		页面应该根据用户角色来确定，是否有权限修改和查看数据内容。
	*/
	// api
	//配置控制器组，
	mvc.Configure(app.Party("/api"), func(m *mvc.Application) {
		m.Router.Use(middleware.LoginAuth)
		//m.Party("/topic").Handle(new(api.TopicController))
		m.Party("/article").Handle(new(api.ArticleController))
		m.Party("/project").Handle(new(api.ProjectController))
		m.Party("/user").Handle(new(api.UserController))
		m.Party("/tag").Handle(new(api.TagController))
		m.Party("/config").Handle(new(api.ConfigController))
		m.Party("/upload").Handle(new(api.UploadController))

		m.Party("/waiter").Handle(new(api.WaiterController))
		m.Party("/assert").Handle(new(api.UserAssertController))
		m.Party("/assertlog").Handle(new(api.UserAssertLogController))
		m.Party("/notice").Handle(new(api.NoticeController))
		m.Party("/user-score").Handle(new(api.UserScoreController))
		m.Party("/user-score-log").Handle(new(api.UserScoreLogController))
	})

	// api
	//配置控制器组，
	mvc.Configure(app.Party("/"), func(m *mvc.Application) {
		m.Party("/login").Handle(new(api.LoginController))
		m.Party("/captcha").Handle(new(api.CaptchaController))
	})

	// admin
	mvc.Configure(app.Party("/api/admin"), func(m *mvc.Application) {
		m.Router.Use(middleware.AdminAuth)
		m.Party("/common").Handle(new(admin.CommonController))
		m.Party("/user").Handle(new(admin.UserController))
		m.Party("/third-account").Handle(new(admin.ThirdAccountController))
		m.Party("/tag").Handle(new(admin.TagController))
		m.Party("/article").Handle(new(admin.ArticleController))
		m.Party("/article-tag").Handle(new(admin.ArticleTagController))
		m.Party("/sys-config").Handle(new(admin.SysConfigController))
		m.Party("/user-score").Handle(new(admin.UserScoreController))
		m.Party("/user-score-log").Handle(new(admin.UserScoreLogController))
	})

	app.Get("/api/img/proxy", func(i iris.Context) {
		url := i.FormValue("url")
		resp, err := resty.New().R().Get(url)
		i.Header("Content-Type", "image/jpg")
		if err == nil {
			_, _ = i.Write(resp.Body())
		} else {
			logrus.Error(err)
		}
	})

	server := &http.Server{Addr: ":" + config.Conf.Port}
	handleSignal(server)
	err := app.Run(iris.Server(server), iris.WithConfiguration(iris.Configuration{
		DisableStartupLog:                 false,
		DisableInterruptHandler:           false,
		DisablePathCorrection:             false,
		EnablePathEscape:                  false,
		FireMethodNotAllowed:              false,
		DisableBodyConsumptionOnUnmarshal: false,
		DisableAutoFireStatusCode:         false,
		EnableOptimizations:               true,
		TimeFormat:                        "2006-01-02 15:04:05",
		Charset:                           "UTF-8",
	}))
	if err != nil {
		logrus.Error(err)
		os.Exit(-1)
	}
}

func handleSignal(server *http.Server) {
	c := make(chan os.Signal) //申明一个Signal 通道
	signal.Notify(c, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)

	go func() {
		s := <-c
		logrus.Infof("got signal [%s], exiting now", s)
		if err := server.Close(); nil != err {
			logrus.Errorf("server close failed: " + err.Error())
		}

		simple.CloseDB()

		logrus.Infof("Exited")
		os.Exit(0)
	}()
}
