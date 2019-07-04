package main

import (
        "github.com/micro/go-log"
	"net/http"

        "github.com/micro/go-web"
        "github.com/julienschmidt/httprouter"
        _ "IHome/IhomeWeb/models"
)

func main() {
	// 创建web服务
        service := web.NewService(
                web.Name("go.micro.web.IhomeWeb"),
                web.Version("latest"),
                web.Address(":22233"),
        )

	// 初始化服务
        if err := service.Init(); err != nil {
                log.Fatal(err)
        }
    rou:=httprouter.New()
    //映射静态页面
    rou.NotFound=http.FileServer(http.Dir("html"))

	// 注册服务
	service.Handle("/", rou)


	// 运行服务
        if err := service.Run(); err != nil {
                log.Fatal(err)
        }
}
