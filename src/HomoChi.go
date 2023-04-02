package src

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var homodictmap = map[string]string{
	"AA": "阿",
	"AE": "埃",
	"AH": "阿",
	"AO": "奥",
	"AW": "奥",
	"AY": "唉",
	"B":  "波",
	"CH": "起",
	"D":  "德",
	"DH": "泽",
	"EH": "诶",
	"ER": "尔",
	"EY": "诶",
	"F":  "福",
	"G":  "格",
	"HH": "赫",
	"IH": "毅",
	"IY": "毅",
	"JH": "珏",
	"K":  "克",
	"L":  "勒",
	"M":  "麽",
	"N":  "讷",
	"NG": "摁",
	"OW": "欧沃",
	"OY": "欧毅",
	"P":  "帔",
	"R":  "渃",
	"S":  "瑟",
	"SH": "瑟",
	"T":  "忒",
	"TH": "泽",
	"UH": "喔",
	"UW": "沃",
	"V":  "沃",
	"W":  "沃",
	"Y":  "毅",
	"Z":  "泽",
	"ZH": "泽",
}

func DeleteAllNum(s []string) []string {
	tem := regexp.MustCompile("[0-9]*")
	for i, _ := range s {

		s[i] = tem.ReplaceAllString(s[i], "")
	}
	return s
}
func GetHomo(s []string) string {
	s = DeleteAllNum(s)
	ans := ""
	for _, j := range s {
		ans += homodictmap[j]
	}
	fmt.Println(ans)
	return ans
}
func GetWord(s string) []string {
	anss := make([][]rune, 0, 10)
	tems := []rune(s)
	l, r, flag := 0, 0, 0
	for i := 0; i <= len(tems); i++ {
		if i != len(tems) && unicode.IsLetter(tems[i]) && !unicode.Is(unicode.Scripts["Han"], tems[i]) {
			if flag == 0 {
				flag = 1
				l, r = i, i
			} else {
				r = i
			}
		} else {
			if flag == 1 {
				anss = append(anss, tems[l:r+1])
			}
			flag = 0
		}
	}
	ansss := make([]string, 0, 10)
	for _, i := range anss {
		ansss = append(ansss, string(i))
	}
	return ansss
}

func ReplaceS(s string) string {
	anss := s
	tems := GetWord(s)
	for _, j := range tems {
		temj:=strings.ToLower(j)
		temss := DialToGrpc(temj)
		temsss := GetHomo(temss)
		anss = strings.Replace(anss, j, temsss, -1)
	}
	return anss
}
