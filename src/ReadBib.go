package src

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"
)

type BibStruct struct {
	Article        string          `json:"article"`
	Title          string          `json:"title"`
	TranslateTitle string          `json:"translatetitle"`
	Author         string          `json:"author"`
	Journal        string          `json:"journal"`
	Volume         int             `json:"volume"`
	Number         int             `json:"number"`
	Pages          string          `json:"pages"`
	Year           int             `json:"year"`
	Publisher      string          `json:"publisher"`
	Doi            string          `json:"doi"`
	Url            string          `json:"url"`
	Keywords       string          `json:"keywords"`
	Abstract       string          `json:"abstract"`
	Organization   string          `json:"organization"`
	Booktitle      string          `json:"booktitle"`
	Issn           string          `json:"issn"`
	Type           string          `json:"type"`
	MapVar         map[string]any  `json:"mapVar"`
	MapVal         map[string]bool `json:"mapVal"`
	path           string          `json:"path"`
	other          string          `json:"other"`
}
type OnceBibFile struct {
	Text [][]string
	Path string
}

var basePath string

func CreatBibStruct(path string) [][]BibStruct {
	l := strings.LastIndex(path, "/")
	basePath = path[l+1:]
	temB := readbib(path)
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}
	temBB := make([][]BibStruct, 0, 10)
	for _, v := range temB {
		//fmt.Println(i)
		wg.Add(1)
		temch := make(chan []BibStruct)
		go tileFileToBibsstr(v, temch)
		//fmt.Println(<-temch)
		go addBibStruct(&temBB, temch, &mu, &wg)
	}
	wg.Wait()
	time.Sleep(0)
	return temBB
}
func addBibStruct(Bibs *[][]BibStruct, temch chan []BibStruct, mu *sync.RWMutex, wg *sync.WaitGroup) {
	mu.Lock()
	temB := <-temch
	*Bibs = append(*Bibs, temB)
	mu.Unlock()
	wg.Done()
}

func addOnceBibFile(Bibs *[]OnceBibFile, temch chan [][]string, path string, mu *sync.RWMutex, wg *sync.WaitGroup) {
	mu.Lock()
	*Bibs = append(*Bibs, OnceBibFile{
		Text: <-temch, Path: path,
	})
	mu.Unlock()
	wg.Done()
}
func readbib(path string) []OnceBibFile {
	paths := make([]string, 0, 10)
	Bibs := make([]OnceBibFile, 0, 10)
	muOnceBib := sync.RWMutex{}
	wg := sync.WaitGroup{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		paths = append(paths, path)
		temPath := strings.Replace(path, basePath, basePath+"Translate", -1)
		if IsDir(path) {
			CreateDir(temPath)
		}
		if strings.HasSuffix(path, ".bib") {
			fmt.Println(path + "!")
			wg.Add(1)
			temch := make(chan [][]string)
			tempath := path
			go readFile(tempath, temch)
			go addOnceBibFile(&Bibs, temch, tempath, &muOnceBib, &wg)
		}
		return nil
	})
	wg.Wait()
	if err != nil {
		fmt.Println(err)
	}
	return Bibs
}
func readFile(path string, temch chan [][]string) (error, [][]string) {
	fileHanle, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		return err, nil
	}
	defer fileHanle.Close()
	readBytes, err := ioutil.ReadAll(fileHanle)
	if err != nil {
		return err, nil
	}
	results0 := strings.Split(string(readBytes), "@")
	results := make([][]string, 0)
	for _, results1 := range results0 {
		if len(results1) < 1 {
			continue
		}
		results1 = "@" + results1
		results2 := strings.Split(results1, "\n")
		results = append(results, results2)
	}
	//fmt.Printf("read result:%v", results)

	/*for _, v := range results {
		for _, vv := range v {
			fmt.Println(len(vv))
			fmt.Println(vv)
			fmt.Println("~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~")
		}
		fmt.Println("_____________________________________")
	}*/
	temch <- results
	return nil, results
}

// check
func newBib() *BibStruct {
	temB := BibStruct{
		Article:        *new(string),
		Title:          *new(string),
		Author:         *new(string),
		Journal:        *new(string),
		Volume:         *new(int),
		Number:         *new(int),
		Pages:          *new(string),
		Year:           *new(int),
		Publisher:      *new(string),
		Doi:            *new(string),
		Url:            *new(string),
		Keywords:       *new(string),
		Abstract:       *new(string),
		Organization:   *new(string),
		Booktitle:      *new(string),
		TranslateTitle: *new(string),
		Type:           *new(string),
	}
	temB.MapVar = make(map[string]any)
	temB.MapVar["article"] = &temB.Article
	temB.MapVar["author"] = &temB.Author
	temB.MapVar["title"] = &temB.Title
	temB.MapVar["journal"] = &temB.Journal
	temB.MapVar["volume"] = &temB.Volume
	temB.MapVar["number"] = &temB.Number
	temB.MapVar["pages"] = &temB.Pages
	temB.MapVar["year"] = &temB.Year
	temB.MapVar["publisher"] = &temB.Publisher
	temB.MapVar["doi"] = &temB.Doi
	temB.MapVar["url"] = &temB.Url
	temB.MapVar["keywords"] = &temB.Keywords
	temB.MapVar["abstract"] = &temB.Abstract
	temB.MapVar["organization"] = &temB.Organization
	temB.MapVar["booktitle"] = &temB.Booktitle
	temB.MapVar["issn"] = &temB.Issn
	temB.MapVal = make(map[string]bool)
	temB.MapVal["translatetitle"] = true
	//fmt.Printf("%v!%v!\n", &temB.Title, temB.mapVar["title"])
	return &temB
}
func tileFileToBibsstr(file OnceBibFile, temchh chan []BibStruct) []BibStruct {
	temB := make([]BibStruct, 0, 100)
	muBibStr := sync.RWMutex{}
	wg := sync.WaitGroup{}
	for _, v := range file.Text {
		temch := make(chan *BibStruct)
		wg.Add(1)
		go tilestrToOnceBibs(v, temch, file.Path)
		go addOnceBibStruct(&temB, temch, &muBibStr, &wg)
	}
	//time.Sleep(time.Second)
	wg.Wait()
	//fmt.Println(temB[0].Article)
	temchh <- temB
	return temB
}
func addOnceBibStruct(Bibs *[]BibStruct, temch chan *BibStruct, mu *sync.RWMutex, wg *sync.WaitGroup) {
	mu.Lock()
	temb := <-temch
	*Bibs = append(*Bibs, *temb)
	mu.Unlock()
	defer wg.Done()
}
func tilestrToOnceBibs(s []string, temch chan *BibStruct, path string) {
	temB := newBib()
	temB.path = path
	for _, v := range s {
		if len(v) <= 1 {
			continue
		}
		var l, r int
		if strings.Contains(v, "@") {
			l = strings.Index(v, "{")
			r = strings.LastIndex(v, ",")
			sub1 := v[l+1 : r]
			sub2 := v[1:l]
			temB.Article = sub1
			temB.Type = sub2
		} else {
			l = strings.Index(v, "{")
			r = strings.LastIndex(v, "}")
			if r <= l {
				temB.other += v
			}
			key := v[:l]
			//fmt.Printf("%v %v\n", r, v)
			value := v[l+1 : r]
			key = strings.Replace(key, " ", "", -1)
			key = strings.Replace(key, "=", "", -1)
			if i, ok := temB.MapVar[key].(*int); ok {
				//fmt.Printf("%v!%v!%v!\n", i, temB.mapVar[key], &temB.Title)
				*i, _ = strconv.Atoi(value)
			}
			if ss, ok := temB.MapVar[key].(*string); ok {
				//fmt.Printf("%v!%v!%v!\n", ss, temB.mapVar[key], &temB.Title)
				*ss = value
			}
			temB.MapVal[key] = true
		}
	}
	temch <- temB
}

func BibStrTobibFile(Bibs *[][]BibStruct) {
	wg := sync.WaitGroup{}
	for _, v := range *Bibs {
		wg.Add(1)
		tv := v
		go WriteFile(&tv, &wg)
	}
	wg.Wait()
}
func WriteFile(Bibs *[]BibStruct, wg *sync.WaitGroup) {
	defer wg.Done()
	Path := (*Bibs)[0].path
	fmt.Println(Path)
	index := strings.LastIndex(Path, "\\")
	temPath := strings.Replace(Path[:index], basePath, basePath+"Translate", -1)
	CreateDir(temPath)
	Path = temPath + Path[index:]
	Path = strings.Replace(Path, ".bib", "TranslateToChinese.bib", 1)
	//Path = strings.Replace(Path, "PaperStore", "TranslatePaperStore", 1)
	os.Remove(Path)
	file, err := os.OpenFile(Path, os.O_CREATE, 0)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	for _, v := range *Bibs {
		vt := reflect.TypeOf(v)
		vv := reflect.ValueOf(v)
		tems := "@"
		tems += v.Type + "{" + v.Article + ",\n"
		_, err = file.WriteString(tems)
		if err != nil {
			return
		}
		for i := 0; i < vv.NumField(); i++ {
			tagContent := vt.Field(i).Tag.Get("json")
			pre := tagContent + "="
			fieldVal := vv.Field(i)
			if v.MapVal[tagContent] != true {
				continue
			}
			switch fieldVal.Kind() {
			case reflect.String:
				val := fieldVal.String()
				_, err = file.WriteString(pre + "{" + val + "},\n")
				if err != nil {
					return
				}
			case reflect.Int:
				val := fieldVal.Int()
				_, err = file.WriteString(pre + "{" + strconv.Itoa(int(val)) + "},\n")
				if err != nil {
					return
				}
			}

		}
		_, err = file.WriteString("}\n\n\n")
		/*tem, _ := json.Marshal(v)
		tems := string(tem)
		tems = strings.Replace(tems, ",", ",\n", -1)
		//fmt.Println(tems)*/
		//_, err = file.WriteString(tems + "\n\n")
		if err != nil {
			return
		}
	}
}
func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

// 创建文件夹
func CreateDir(path string) {
	_exist, _err := HasDir(path)
	if _err != nil {
		fmt.Printf("获取文件夹异常 -> %v\n", _err)
		return
	}
	if _exist {
		fmt.Println("文件夹已存在！")
	} else {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			fmt.Printf("创建目录异常 -> %v\n", err)
		} else {
			fmt.Println("创建成功!")
		}
	}
}
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {

		return false
	}
	return s.IsDir()

}
