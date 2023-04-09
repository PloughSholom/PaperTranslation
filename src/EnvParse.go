package src

import (
	"flag"
	"os"
	"strconv"
)

var (
	//DefKey = "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"
	DefPT1 = Message{
		Role:    "system",
		Content: TitlePrompt,
	}
	DefPT2 = Message{
		Role:    "assistant",
		Content: "好的，请问您需要我翻译哪篇论文的标题呢？我会用|分隔所提供的五个翻译版本",
	}
	DefPA1 = Message{
		Role:    "system",
		Content: AbstractPrompt,
	}
	DefPA2 = Message{
		Role:    "assistant",
		Content: "好的，请问您需要我翻译哪篇论文的摘要呢？",
	}
	TitlePrompt    = "I need you to play the role of a highly professional scholar and complete the task of translating the title of the paper into Chinese. Then I will provide you with several papers, please read the information I give you and read the original text, and translate the title of the paper first word by word, then embellish the result of the translation from an academic point of view before exporting it. The content of the papers will be given in the form xxx={xxx}, but you only need to return your translated titles briefly in Chinese, replacing the English part of the content with the harmonic part before you send it, please also give five different translations separated by |."
	AbstractPrompt = "I need you to play the role of a highly professional scholar and complete the task of translating the abstract of the paper into Chinese. Then I will provide you with several papers, please read the information I give you and read the original text, and translate the abstract of the paper first word by word, then embellish the result of the translation from an academic point of view before exporting it. The content of the papers will be given in the form xxx={xxx}, but you only need to return your translated abstract in Chinese, replacing the English part of the content with the harmonic part before you send it"
)
var (
	Address2Py = "pyweb:50051"
	Port       int
	Port2Py    int
	ModelGPT   = "gpt-3.5-turbo"
	GPTurl     = "https://service-6hpy0xnm-1317247263.sg.apigw.tencentcs.com/release/v1/chat/completions"
	Keys       = "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"
	Keych      = make(chan string)
	Keytem     = []string{}
	RespTime   int64
)

func EnvParse() {
	PORT := flag.Int("PORT", 8888, "")
	PORT2PY := flag.Int("PORT2PY", 50051, "")
	MODELGPT := flag.String("MODELGPT", "gpt-3.5-turbo", "")
	GPTURL := flag.String("GPTURL", "https://service-6hpy0xnm-1317247263.sg.apigw.tencentcs.com/release/v1/chat/completions", "")
	KEYS := flag.String("KEYS", "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", "用,分隔")
	RESPTIME := flag.Int64("RESPTIME", 60, "因为gpt延迟太高了")
	flag.Parse()
	Port = *PORT
	Port2Py = *PORT2PY
	Address2Py = "pyweb:" + strconv.Itoa(Port2Py)
	ModelGPT = *MODELGPT
	GPTurl = *GPTURL
	Keys = *KEYS
	RespTime = *RESPTIME
	if os.Getenv("PORT") != "" {
		Port, _ = strconv.Atoi(os.Getenv("PORT"))
	}
	if os.Getenv("PORT2PY") != "" {
		Port2Py, _ = strconv.Atoi(os.Getenv("PORT2PY"))
		Address2Py = "pyweb:" + strconv.Itoa(Port2Py)
	}
	if os.Getenv("MODELGPT") != "" {
		ModelGPT = os.Getenv("MODELGPT")
	}
	if os.Getenv("GPTURL") != "" {
		GPTurl = os.Getenv("GPTURL")
	}
	if os.Getenv("KEYS") != "" {
		Keys = os.Getenv("KEYS")
	}
	if os.Getenv("RESPTIME") != "" {
		RespTime, _ = strconv.ParseInt(os.Getenv("RESPTIME"), 10, 64)
	}

}
