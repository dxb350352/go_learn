package main
import (
	"io/ioutil"
	"crypto/x509"
	"fmt"
)

func main() {
	var cert *x509.Certificate
	pem, err := ioutil.ReadFile("E:/GOPATH_GO/src/github.com/sas/sasslvpn_web/certs/zgf/ZuiGaoFaCA.crl")
	if err != nil {
		panic(err)
	}
	certList, err := x509.ParseCRL(pem)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, revoked := range certList.TBSCertList.RevokedCertificates {
		if cert.SerialNumber.Cmp(revoked.SerialNumber) == 0 {
			fmt.Println("Serial number match: intermediate is revoked.")
		}
	}
}
