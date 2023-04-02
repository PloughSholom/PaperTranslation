package main

import "github.com/PloughSholom/PaperTranslation/src"

/*
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
	for i, v := range temB {
		for ii, vv := range v {
			wg.Add(1)
			ti := i
			tii := ii
			vvv := vv
			go func() {
				defer wg.Done()
				tem := src.SendQuestToGPTAndReceive(src.BibToCQ(vvv))
				tem = strings.Replace(tem, "\n", "", -1)
				tem = strings.Replace(tem, "《", "", -1)
				tem = strings.Replace(tem, "》", "", -1)
				temB[ti][tii].TranslateTitle = tem
			}()
			time.Sleep(time.Second) //请求访问上限
		}
	}
	wg.Wait()
	src.BibStrTobibFile(&temB)
}*/

func main() {
	src.GoWebServer()
}
