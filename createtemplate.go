package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
)

const template = `package main

import (
	"fmt"
)

func main(){
	fmt.Println("hello")
}`

var (
	help = flag.String("", "help", "./createtemplate <filename>")
)

func main() {
	flag.Parse()

	if len(os.Args) < 2 {
		fmt.Println("ファイル名を指定して下さい。")
		return
	}
	fileName := os.Args[1]

	isGo, _ := regexp.MatchString(".go$", fileName)
	if !isGo {
		fileName = fileName + ".go"
	}

	_, err := os.Open(fileName)
	if err == nil {
		fmt.Printf("ファイル[%s]が重複しています。",fileName)

		for {
			fmt.Printf("上書きしてよろしいですか? (y or n) : ")

			str := ""
			fmt.Scanf("%s\n", &str)
			if str == "n" {
				return
			} else if str == "y" {
				break
			} else {
				fmt.Printf("yかnを入力してください。")
			}
		}
	}

	f, _ := os.Create(fileName) // *File
	f.Chdir()                   //カレントディレクトリにファイルを作成する
	f.WriteString(template)     //templateを書き込む

}
