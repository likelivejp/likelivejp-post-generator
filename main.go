package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "postgen"
	app.Usage = "Hugo Generate New post cli tool"
	app.Version = "0.1"

	app.Action = func(context *cli.Context) error {
		// ディレクトリ生成
		t := time.Now()
		p := t.Format("2006/01/02")
		dir := "content/posts/" + p
		if err := os.MkdirAll(dir, 0777); err != nil {
			fmt.Println(err)
		}

		// ファイル生成
		f := dir + "/" + context.Args().Get(0) + ".md"
		file, _ := os.OpenFile(f, os.O_CREATE|os.O_WRONLY, 0777)
		defer file.Close()

		// Front matter追加
		var fm []string
		fm = append(fm, "---")
		fm = append(fm, "title: ")
		fm = append(fm, "created_at: "+t.Format("2006-01-02"))
		fm = append(fm, "category: ")
		fm = append(fm, "tags: ")
		fm = append(fm, "  - ")
		fm = append(fm, "image: ")
		fm = append(fm, "author: ")
		fm = append(fm, "---")
		fm = append(fm, "")
		fmt.Fprintln(file, strings.Join(fm, "\n"))

		return nil
	}

	app.Run(os.Args)
}
