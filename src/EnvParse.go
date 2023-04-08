package src

import (
	"flag"
	"os"
	"strconv"
)

var (
	DefKey = "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"
	DefPT1 = Message{
		Role:    "user",
		Content: TitlePrompt,
	}
	DefPT2 = Message{
		Role:    "assistant",
		Content: "好的，请问您需要我翻译哪篇论文的标题呢？我会用|分隔所提供的五个翻译版本",
	}
	DefPA1 = Message{
		Role:    "user",
		Content: AbstractPrompt,
	}
	DefPA2 = Message{
		Role:    "assistant",
		Content: "好的，请问您需要我翻译哪篇论文的标题呢？我会用|分隔所提供的两个翻译版本",
	}
	TitlePrompt    = "I need you to play the role of a highly professional scholar and complete the task of translating the title of the paper into Chinese. Then I will provide you with several papers, please read the information I give you and read the original text, and translate the title of the paper first word by word, then embellish the result of the translation from an academic point of view before exporting it. The content of the papers will be given in the form xxx={xxx}, but you only need to return your translated titles briefly in Chinese, replacing the English part of the content with the harmonic part before you send it, please also give five different translations separated by |."
	AbstractPrompt = "I need you to play the role of a highly professional scholar and complete the task of translating the abstract of the paper into Chinese. Then I will provide you with several papers, please read the information I give you and read the original text, and translate the abstract of the paper first word by word, then embellish the result of the translation from an academic point of view before exporting it. The content of the papers will be given in the form xxx={xxx}, but you only need to return your translated abstract in Chinese, replacing the English part of the content with the harmonic part before you send it, please also give two different translations separated by |."
)
var (
	Address2Py = "pyweb:50051"
	Port       int
	Port2Py    int
	ModelGPT   = "gpt-3.5-turbo"
	GPTurl     = "https://api.openai.com/v1/chat/completions"
	Keys       string
	Keych      = make(chan string)
	Keytem     = []string{}
)

func EnvParse() {
	PORT := flag.Int("PORT", 8888, "")
	PORT2PY := flag.Int("PORT2PY", 50051, "")
	MODELGPT := flag.String("MODELGPT", "gpt-3.5-turbo", "")
	GPTURL := flag.String("GPTURL", "https://api.openai.com/v1/chat/completions", "")
	KEYS := flag.String("KEYS", "", "用,分隔")
	flag.Parse()
	Port = *PORT
	Port2Py = *PORT2PY
	Address2Py = "pyweb:" + strconv.Itoa(Port2Py)
	ModelGPT = *MODELGPT
	GPTurl = *GPTURL
	Keys = *KEYS
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

}
