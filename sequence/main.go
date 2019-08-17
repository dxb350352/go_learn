package main

import (
	"fmt"
	"github.com/jackiedong168/sequence"
	"strings"
)

func main() {
	keywords := `__index=sas_gaxz_audit_2* loAAAgin=00000ABC 000ABC000 NOT request_path="/audit/search" AND request_path=/exact/search`
	fmt.Println(keywords)
	//fmt.Println(readableFileSize(1023))
	//Prepare("", "", "", keywords, "", "", 1512230400000, 1514995199999)
	scan := sequence.NewScanner()
	lparser := sequence.NewLucParser()
	seq1, _ := scan.Scan(keywords)
	seq2, _ := lparser.Parse(seq1)
	m := map[string]bool{"AND": true, "NOT": true, "OR": true}
	var l sequence.Sequence
	for _, token := range seq2 {
		if !m[token.Op] {
			token.Op = strings.ToLower(token.Op)
		}
		token.Value = strings.ToLower(token.Value)
		l = append(l, token)
	}
	fmt.Println(lparser.String(l, "="))
}
