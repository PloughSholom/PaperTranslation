package src

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"time"
	"unicode"
)

var (
	ModelGPT3_5 = "gpt-3.5-turbo"
	GPTURL      = "https://service-6hpy0xnm-1317247263.sg.apigw.tencentcs.com/v1/chat/completions"
	//GPTURL = "https://chat-gpt.aurorax.cloud/v1/chat/completions"
)

// exp GPT现用请求块
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}
type Usage struct {
	Prompt_tokens     int64 `json:"prompt_tokens"`
	Completion_tokens int64 `json:"completion_tokens"`
	Total_tokens      int64 `json:"total_tokens"`
}
type CQuest struct {
	apiKey      string
	prompt      string
	Messages    []Message `json:"messages"`
	Model       string    `json:"model"`
	Max_tokens  int64     `json:"max_tokens"`
	Temperature float64   `json:"temperature"`
	N           int       `json:"n"`
	User        string    `json:"user"`
}
type Choices struct {
	Message Message `json:"message"`
}
type CResponse struct {
	Object  string    `json:"object"`
	Id      string    `json:"id"`
	Choices []Choices `json:"choices"`
	Error   error     `json:"error"`
	Usage   Usage     `json:"usage"`
}

// todo 新建请求块
func NewCquest() *CQuest {
	temCQ := new(CQuest)
	temCQ.apiKey = ""
	temCQ.N = 1
	temCQ.prompt = DefPro
	temCQ.Max_tokens = 200
	temCQ.Temperature = 0.3
	temCQ.User = ""
	temCQ.Model = ModelGPT3_5
	temCQ.Messages = make([]Message, 0, 60)
	return temCQ
}
func BibToCQ(bib BibStruct) *CQuest {
	temCQ := NewCquest()
	fmt.Println(bib.Title)
	s := ""
	for k, _ := range bib.MapVar {
		if _, ok := bib.MapVar[k].(*int); ok {
			//s += k + "=:" + strconv.Itoa(*i)
		}
		if ss, ok := bib.MapVar[k].(*string); ok {
			//s += k + "=:" + *ss
			s += *ss
		}
		s += "\n"
	}
	temCQ.Messages = append(temCQ.Messages, DefP1, DefP2)
	temCQ.Messages = append(temCQ.Messages, Message{
		Role: "user", Content: s,
	})
	return temCQ
}
func SendQuestToGPTAndReceive(cq *CQuest, num int, op int) string {
	if num < 1 {
		num = 1
	}
	var reqnum = 0
	remess := ""
	messnum := 0
	for {
		reqnum++
		for {
			if cq.apiKey != "" {
				break
			}
			cq.apiKey = getKey()
		}
		hf := new(HeadFun)
		hf.AppendHead("Authorization", fmt.Sprintf("%s %s", "Bearer", cq.apiKey))
		hf.AppendHead("Content-Type", "application/json")
		all, err := PostJson(cq, GPTURL, "POST", hf)
		if err != nil {
			fmt.Println(err.Error())
		}
		//fmt.Println("收到gpt消息:\n" + string(all) + "\n")
		temRE := *new(CResponse)
		temRE.Choices = make([]Choices, 1, 10)
		err = json.Unmarshal(all, &temRE)
		//fmt.Println(temRE.Choices[0].Message.Content)
		//fmt.Println(string(all))
		if strings.Contains(temRE.Choices[0].Message.Content, "|") == false {
			continue
		}
		Mess := strings.Split(temRE.Choices[0].Message.Content, "|")
		for _, v := range Mess {
			if CheckVal(v, reqnum) {
				messnum++
				if op == 1 {
					temv := ReplaceS(v)
					remess += temv
					remess += ","
					//fmt.Println(temv)
				}
				//fmt.Println(v)
				remess += v
				if num != messnum {
					remess += ",\n"
				}
				if num == messnum {
					fmt.Println(remess)
					return remess
				}
			}
		}
		fmt.Println("再次请求")
		cq.Messages[0].Content = strings.Replace(cq.Messages[0].Content, "only in chinese", "only in chinese!!", 1)
	}
}

var keych = make(chan string)
var Keytem = []string{}

func ResetKeys() {
	for {
		close(keych)
		keych = make(chan string)
		for _, v := range Keytem {
			keych <- v
		}
		fmt.Println("已重置")
		time.Sleep(time.Second * 3)
	}
}
func getKey() string {
	return <-keych
}
func ReadKeys() error {
	fileHanle, err := os.OpenFile("./keys.txt", os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	defer func(fileHanle *os.File) {
		err = fileHanle.Close()
		if err != nil {

		}
	}(fileHanle)
	readBytes, err := ioutil.ReadAll(fileHanle)
	if err != nil {
		return err
	}
	results0 := strings.Split(string(readBytes), "\n")
	for _, i := range results0 {
		i = strings.Replace(i, "\r", "", -1)
		if len(i) > 0 {
			Keytem = append(Keytem, i)
		}
	}
	if len(Keytem) < 1 {
		return errors.New("keytem不能为空")
	}
	return nil
}
func dealS(s string) string {

	return s
}
func CheckVal(s string, num int) bool {
	if len(s) < 1 {
		return false
	}
	if strings.Contains(s, "翻译") && num < 3 {
		return false
	}
	var Ccount, Pcount, Lcount float64
	for _, v := range s {
		if unicode.Is(unicode.Han, v) {
			Ccount++
		}
		if unicode.IsPunct(v) {
			Pcount++
		}
		if unicode.IsLetter(v) {
			Lcount++
		}
	}
	Lcount = Lcount - Ccount
	//fmt.Println(Ccount, Pcount, Lcount)
	if Lcount/(Ccount+Lcount) > 0.1+float64(num)*0.05 {
		return false
	}

	return true
}
