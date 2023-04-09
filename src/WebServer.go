package src

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

type Req struct {
	Content string `p:"content"`
	Num     int    `p:"num"`
	Option  int    `p:"option"`
}
type Res struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func TitleReqInJson(r *ghttp.Request) {
	req := new(Req)
	err := r.Parse(&req)
	if err != nil {
		temres := Res{
			Code: 1, Error: err.Error(),
		}
		r.Response.WriteJsonExit(temres)
	} else {
		temB := new(BibStruct)
		temB.MapVal = make(map[string]bool)
		temB.MapVal["title"] = true
		temB.Title = req.Content
		temB.MapVar = make(map[string]any)
		temB.MapVar["title"] = &temB.Title
		tem := SendQuestToGPTAndReceive(BibToCQ(*temB, DefPT1, DefPT2, 200), req.Num, req.Option, "title")
		//tem = strings.Replace(tem, "\n", "", -1)
		tem = strings.Replace(tem, "《", "", -1)
		tem = strings.Replace(tem, "》", "", -1)
		temres := Res{
			Code: 0, Data: tem,
		}
		//fmt.Println(tem)
		r.Response.WriteJsonExit(temres)
	}
}
func AbstractReqInJson(r *ghttp.Request) {
	req := new(Req)
	err := r.Parse(&req)
	if err != nil {
		temres := Res{
			Code: 1, Error: err.Error(),
		}
		r.Response.WriteJsonExit(temres)
	} else {
		temB := new(BibStruct)
		temB.MapVal = make(map[string]bool)
		temB.MapVal["title"] = true
		temB.Title = req.Content
		temB.MapVar = make(map[string]any)
		temB.MapVar["title"] = &temB.Title
		tem := SendQuestToGPTAndReceive(BibToCQ(*temB, DefPA1, DefPA2, 2000), req.Num, req.Option, "abstract")
		tem = strings.Replace(tem, "《", "", -1)
		tem = strings.Replace(tem, "》", "", -1)
		temres := Res{
			Code: 0, Data: tem,
		}
		r.Response.WriteJsonExit(temres)
	}
}
func HomoReaqinJson(r *ghttp.Request) {
	req := new(Req)
	err := r.Parse(&req)
	if err != nil {
		temres := Res{
			Code: 1, Error: err.Error(),
		}
		r.Response.WriteJsonExit(temres)
	} else {
		tem := DialToGrpc(req.Content)
		temres := Res{
			Code: 0, Data: tem,
		}
		r.Response.WriteJsonExit(temres)
	}
}
func EchoinJson(r *ghttp.Request) {
	req := new(Req)
	err := r.Parse(&req)
	if err != nil {
		temres := Res{
			Code: 1, Error: err.Error(),
		}
		r.Response.WriteJsonExit(temres)
	} else {
		tem := req.Content
		temres := Res{
			Code: 0, Data: tem,
		}
		r.Response.WriteJsonExit(temres)
	}
}
func GoWebServer() {
	s := g.Server()
	s.Group("/json", func(group *ghttp.RouterGroup) {
		group.POST("/title", TitleReqInJson)
		group.POST("/abstract", AbstractReqInJson)
		group.POST("/homo", HomoReaqinJson)
		group.POST("/echo", EchoinJson)
	})
	s.SetPort(Port)
	s.Run()
}
