package main

import (
)
import "github.com/sas/tqry/app/docx"

func main() {
	r, err := docx.ReadDocxFile("E:/GOPATH/src/testgo/word_docx/禁毒特情物建审批表_template.docx")
	if err != nil {
		panic(err)
	}
	docx1 := r.Editable()
	docx1.Replace("{_{41}_}", "fd1323216546545sa", -1)
	docx1.ReplaceImage(1,"E:/GOPATH/src/testgo/word_docx/QQ20111.png")
	docx1.WriteToFile("E:/new_result_1.docx")

	r.Close()
}
