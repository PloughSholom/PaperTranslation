package main

import (
	"PaperTranslation/src"
	"fmt"
	"os"
	"strings"
	"sync"
	"time"
)

func main() {
	var Path string
	if len(os.Args) > 1 {
		Path = os.Args[1]
	}
	if len(Path) < 0 {
		Path = "./PaperStore"
	}
	if (strings.Count(Path, "/") < 1 && strings.Count(Path, "\\") < 1) || src.IsDir(Path) == false {
		fmt.Println("路径必须包含一文件夹,更改为./PaperStore")
		Path = "./PaperStore"
	}
	err := src.ReadKeys()
	if err != nil {
		fmt.Println(err)
		return
	}
	go src.ResetKeys()
	temB := src.CreatBibStruct(Path)
	wg := sync.WaitGroup{}
	for _, v := range temB {
		for _, vv := range v {
			wg.Add(1)
			vvv := vv
			go func() {
				defer wg.Done()
				temb := new(src.BibStruct)
				temb.MapVal = make(map[string]bool)
				temb.MapVal["title"] = true
				temb.Title = vvv.Title
				temb.MapVar = make(map[string]any)
				temb.MapVar["title"] = &temb.Title
				tem := src.SendQuestToGPTAndReceive(src.BibToCQ(*temb), 1, 1)
				tem = strings.Replace(tem, "\n", "", -1)
				tem = strings.Replace(tem, "《", "", -1)
				tem = strings.Replace(tem, "》", "", -1)
			}()
			time.Sleep(time.Second) //请求访问上限
		}
	}
	wg.Wait()
	src.BibStrTobibFile(&temB)
}
