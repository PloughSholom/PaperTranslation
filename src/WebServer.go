package src

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"strings"
)

type TitleReq struct {
	Content string `p:"content"`
	Num     int    `p:"num"`
	Option  int    `p:"option"`
}
type TitleRes struct {
	Code  int         `json:"code"`
	Error string      `json:"error"`
	Data  interface{} `json:"data"`
}

func TitleReqInJson(r *ghttp.Request) {
	req := new(TitleReq)
	err := r.Parse(&req)
	if err != nil {
		temres := TitleRes{
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
		tem := SendQuestToGPTAndReceive(BibToCQ(*temB), req.Num, req.Option)
		//tem = strings.Replace(tem, "\n", "", -1)
		tem = strings.Replace(tem, "《", "", -1)
		tem = strings.Replace(tem, "》", "", -1)
		temres := TitleRes{
			Code: 0, Data: tem,
		}
		//fmt.Println(tem)
		r.Response.WriteJsonExit(temres)
	}
}
func GoWebServer() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/json", TitleReqInJson)
	})
	s.SetPort(8888)
	s.Run()
}
