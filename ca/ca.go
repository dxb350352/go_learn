package main

/*
* Copyrights : CNRS
* Author : Oleg Lodygensky
* Acknowledgment : XtremWeb-HEP is based on XtremWeb 1.8.0 by inria :
http://www.xtremweb.net/
* Web : http://www.xtremweb-hep.org
*
* This file is part of XtremWeb-HEP.
*
* XtremWeb-HEP is free software: you can redistribute it and/or modify
* it under the terms of the GNU General Public License as published by
* the Free Software Foundation, either version 3 of the License, or
* (at your option) any later version.
*
* XtremWeb-HEP is distributed in the hope that it will be useful,
* but WITHOUT ANY WARRANTY; without even the implied warranty of
* MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
* GNU General Public License for more details.
*
* You should have received a copy of the GNU General Public License
* along with XtremWeb-HEP. If not, see <http://www.gnu.org/licenses/>.
*
*/


import (
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testgo/ca/common"
)



/**
* This is the standard main function
*/
func main() {

	//	var logger = NewPrefixed("main#main")

	// verbose := flag.Bool("v", false, "enable/disable verbose mode")

	if len(os.Args) < 4 {
		fmt.Printf("FATAL::Usage: %v caCertPath serverCertPath serverKeyPath",
			os.Args[0])
		os.Exit(0)
	}

	//
	// Retrieve CA certificates
	//
	caRootPool, err := ca.PopulateCertPool(os.Args[1])
	if err != nil {
		fmt.Printf("FATAL::" + err.Error())
	}
	fmt.Printf("INFO::CA pool length = %v", len(caRootPool.Subjects()))
	//
	// Retrieve certificate
	//
//	/////////////////////
//	serverCert, err := ca.CerficateFromPEMs(os.Args[2], os.Args[3])
//
//	if err != nil {
//		fmt.Printf("FATAL::" + err.Error())
//	}
//
//	fmt.Printf("INFO::serverCert.Subject = %v", serverCert.Subject)
//	fmt.Printf("INFO::serverCert.Issuer = %v", serverCert.Issuer)
//
//	//
//	// Verify user cert against known CA
//	//
//	vOpts := x509.VerifyOptions{Roots: caRootPool}
//	chains, err := serverCert.Verify(vOpts)
//	if err != nil {
//		fmt.Printf("WARN::failed to parse certificate: " + err.Error())
//	}
//	fmt.Printf("INFO::shains = %v\n", chains)
//	///////////////////////////////////
	ca_b, _ := ioutil.ReadFile(os.Args[2])
	priv_b, _ := ioutil.ReadFile(os.Args[3])
	priv, _ := x509.ParsePKCS1PrivateKey(priv_b)

	cert := tls.Certificate{
		Certificate: [][]byte{ca_b},
		PrivateKey: priv,
	}

	config := tls.Config{
		// RootCAs: caRootPool,
		ClientCAs: caRootPool,
		Certificates: []tls.Certificate{cert},
		//MinVersion: tls.VersionSSL30, //don't use SSLv3,https://www.openssl.org/~bodo/ssl-poodle.pdf
		MinVersion: tls.VersionTLS10,
		//MinVersion: tls.VersionTLS11,
		//MinVersion: tls.VersionTLS12,
		// ClientAuth: tls.VerifyClientCertIfGiven,
		ClientAuth: tls.RequestClientCert,
		// ClientAuth: tls.RequireAnyClientCert,
		// ClientAuth: tls.RequireAndVerifyClientCert,
	}
	config.Rand = rand.Reader
	var a Authenticator

	http.Handle("/", a)
	server := http.Server{Addr: ":4040", TLSConfig: &config}

	// start https
	fmt.Printf("INFO::Listening HTTPS : 4040")

	server.ListenAndServeTLS(os.Args[2], os.Args[3])
}

type Authenticator struct {}

func (a Authenticator) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	//	logger := NewPrefixed("Authenticator#ServHTTP")

	fmt.Printf("DEBUG::len(r.TLS.PeerCertificates) = %v",
		len(r.TLS.PeerCertificates))

	if (len(r.TLS.PeerCertificates) < 1) {
		fmt.Fprint(w, "<p>You are not logged in</p>")
		return
	}

	for i, c := range r.TLS.PeerCertificates {
		fmt.Printf("DEBUG::r.TLS.PeerCertificates[%v].Subject = %v", i,
			c.Subject)
		fmt.Printf("DEBUG::r.TLS.PeerCertificates[%v].Issuer = %v", i, c.Issuer)
		for j, a := range c.EmailAddresses {
			fmt.Printf("DEBUG::r.TLS.PeerCertificates[%v].EmailAddresses[%v] =%v", i, j, a)
		}
	}
	fmt.Printf("DEBUG::len(r.TLS.VerifiedChains) = %v",
		len(r.TLS.VerifiedChains))

	caRootPool, err := ca.PopulateCertPool(os.Args[1])
	if err != nil {
		fmt.Fprint(w, "<p>Internal error : cant retreive CA Rool Pool</p>")
		return
	}
	fmt.Printf("DEBUG::CA pool length = %v", len(caRootPool.Subjects()))

	vOpts := x509.VerifyOptions{Roots: caRootPool}

	userCert := r.TLS.PeerCertificates[0]
	chains, err := userCert.Verify(vOpts)
	if err != nil {
		fmt.Printf("WARN::failed to parse certificate: " + err.Error())
		fmt.Fprint(w, "<p>Certificate error: can't validate your certificate < / p > ")
		return
	}

	fmt.Printf("INFO::shains = %v\n", chains)

	fmt.Fprint(w, "<p>You are logged in as ", userCert.EmailAddresses[0],
		"</p>")
}