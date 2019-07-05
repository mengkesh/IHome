package handler

import (
	"context"
	"encoding/json"
	"net/http"
	GETAREA "IHome/GetArea/proto/example"
	//example "github.com/micro/examples/template/srv/proto/example"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/micro/go-grpc"
	"IHome/IhomeWeb/models"
)

func ExampleCall(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//模板
	//// decode the incoming request as json
	//var request map[string]interface{}
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	//// call the backend service
	//exampleClient := example.NewExampleService("go.micro.srv.template", client.DefaultClient)
	//rsp, err := exampleClient.Call(context.TODO(), &example.Request{
	//	Name: request["name"].(string),
	//})
	//if err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
	//
	//// we want to augment the response
	//response := map[string]interface{}{
	//	"msg": rsp.Msg,
	//	"ref": time.Now().UnixNano(),
	//}
	//
	//// encode and write the response as json
	//if err := json.NewEncoder(w).Encode(response); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}
}
func GetArea(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Println("获取地域信息服务  GetArea /api/v1.0/areas")
	//由于request为空，所以不需要这一项
	//var request map[string]interface{}
	//if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	//	http.Error(w, err.Error(), 500)
	//	return
	//}

	// call the backend service
	//创建grpc服务，并初始化
	cli:=grpc.NewService()
	cli.Init()
	exampleClient := GETAREA.NewExampleService("go.micro.srv.GetArea"/*与GetArea的main中保持一致*/, cli.Client()/*使用cli.出来的client*/)
	rsp, err := exampleClient.GetArea(context.TODO(), &GETAREA.Request{
		//由于request为空，所以此项不填
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	//接收返回值
	errno:=rsp.Errno
	errmsg:=rsp.Errmsg
	var area_data []models.Area
	for _,v:=range rsp.Data{
		tmp:=models.Area{Id:int(v.Aid),Name:v.Aname}
		area_data=append(area_data,tmp)
	}
	response := map[string]interface{}{
		"errno": errno,
		"errmsg":errmsg ,
		"data":area_data,
	}
	//设置给前端返回的格式
	w.Header().Set("Content-Type","application/json")
	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}
