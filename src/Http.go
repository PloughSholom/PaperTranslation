package src

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/PloughSholom/PaperTranslation/src/proto"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

// todo 一个请求封装,传入一个请求块,URL,和操作,返回一个未解析的[]byte
type HeadAddFunc interface {
	HeadAddFunc(r *http.Request)
}
type HeadFun struct {
	KeysValue []struct {
		Key   string
		Value string
	}
}

func (h *HeadFun) HeadAddFunc(r *http.Request) {
	for _, hf := range h.KeysValue {
		r.Header.Set(hf.Key, hf.Value)
	}

}
func (h *HeadFun) AppendHead(key string, value string) {
	h.KeysValue = append(h.KeysValue, struct {
		Key   string
		Value string
	}{Key: key, Value: value})
}
func PostJson(v any, URL string, methon string, headadd HeadAddFunc) ([]byte, error) {
	temreq, err := json.Marshal(v)
	temh := headadd.(*HeadFun).KeysValue[0]
	if err != nil {
		return nil, err
	}
	//fmt.Println("开始请求:\n" + string(temreq))
	//fmt.Printf(string(temreq) + "\n")
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, time.Second*time.Duration(RespTime)) //设置建立连接超时
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(time.Second * time.Duration(RespTime))) //设置发送接受数据超时
				return conn, nil
			},
			ResponseHeaderTimeout: time.Second * time.Duration(RespTime),
		},
	}
	body := bytes.NewReader(temreq)
	req, err := http.NewRequest(methon, URL, body)
	if err != nil {
		return nil, err
	}
	//todo 编写头部,做json标识
	headadd.HeadAddFunc(req)
	//req.Header.Set("Authorization", fmt.Sprintf("%s %s", "Bearer", "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"))
	//req.Header.Set("Content-Type", "application/json")
	//todo 得到一个Response
	t1 := time.Now()
	resp, err := client.Do(req)
	t2 := time.Now()
	fmt.Println(time.Duration(t2.Sub(t1)).Seconds())
	if err != nil {
		return nil, err
		// handle err
	}
	defer resp.Body.Close()
	all, err := io.ReadAll(resp.Body)
	fmt.Println(temh.Value)
	return all, err
}

func DialToGrpc(s string) []string {
	conn, err := grpc.Dial(Address2Py, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewMyServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Gethomo(ctx, &proto.HomoReq{
		Message: s,
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//fmt.Printf("gprc result: %+v", r.Message)
	return r.Message
}
