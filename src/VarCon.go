package src

import "fmt"

func init() {
	err := ReadKeys()
	if err != nil {
		fmt.Println(err)
		return
	}
	go ResetKeys()
}

var (
	DefKey = "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"
	DefP1  = Message{
		Role:    "user",
		Content: DefPro,
	}
	DefP2 = Message{
		Role:    "assistant",
		Content: "好的，请问您需要我翻译哪篇论文的标题呢？我会用|分隔所提供的五个翻译版本",
	}
	DefPro = "I need you to play the role of a highly professional scholar and complete the task of translating the title of the paper into Chinese. Then I will provide you with several papers, please read the information I give you and read the original text, and translate the title of the paper first word by word, then embellish the result of the translation from an academic point of view before exporting it. The content of the papers will be given in the form xxx={xxx}, but you only need to return your translated titles briefly in Chinese, replacing the English part of the content with the harmonic part before you send it, please also give five different translations separated by |."
	/*"format:@article{xxxxxxyearxxxxxx,\n  title={xxxxxx}," +
	"\n  author={xxxxxx},\n  journal={xxxxxx},\n  volume={xxxxxx},\n  number={xxxxxx},\n  pages={xxxxxx}," +
	"\n  year={xxxxxx},\n  publisher={xxxxxx}\n ...={...}\n}"
	I need you to play the role of a highly professional scholar and perform Chinese translation tasks of Paper title. I will then provide you with several papers, so please read the information I have given you and read the original paper, and then provide an academic perspective on the Chinese translation of The content of the paper will be given in the form of xxx={xxx} called 'key ' fo
	r the former and 'value' for the latter, but you will only need to return the title of your translation short and only in chinese, and before you send it, replace the English part of the content with a harmonic sound part
	*/
	Address     = "127.0.0.1:18889"
	AefaultName = "world"
)

//"我需要你扮演一个专业水平高超的学者,接下来我会提供多份论文,请你细细揣摩阅读原文后,联系对应和相关学科知识提供学术角度的翻译"
//论文内容会以xxx={xxx}的形式给出前者称为"key"后者称为"value",{}内是需要翻译的内容,但注意你只需要翻译key为title,abstract,keyword"
//"中的value部分,并按以下格式给出"
