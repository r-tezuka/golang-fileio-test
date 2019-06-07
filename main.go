package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	OutputHoge("./file/input.txt")
}

func OutputHoge(targetFile string) {

	//インプット元の指定
	r, _ := os.Open(targetFile)
	defer r.Close()

	//アウトプット先の指定
	w, _ := os.Create("./file/output.txt")
	defer w.Close()

	//本処理。
	AddText(r, w)
}

//本処理。インプットファイルを一行ずつ読み、アウトプットファイルに追記。
//ファイル操作と本処理を分けることで、本処理をTestingパッケージでテストしやすくする。
func AddText(r io.Reader, w io.Writer) {
	sc := bufio.NewScanner(r)
	buf := bufio.NewWriter(w)
	for sc.Scan() {
		buf.WriteString(sc.Text())
		buf.WriteString("\n")
	}

	//アウトプットファイルの最終行にhogeを追加
	buf.WriteString("hoge")
	buf.Flush()
}
