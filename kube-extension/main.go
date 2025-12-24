package main

import (
    "crypto/rand"
    "crypto/rsa"
    "crypto/x509"
    "encoding/pem"
    "fmt"
    "os"
)

func main() {
    host := os.Args[1]
    name := "server"
    key, err := rsa.GenerateKey(rand.Reader, 1024)
    if err != nil {
        fmt.Println("Error generating RSA key:", err)
        return
    }
    keyDer := x509.MarshalPKCS1PrivateKey(key)
    keyBlock := pem.Block{
        Type: "RSA PRIVATE KEY",
        Bytes: keyDer,
    }
    keyFile, err := os.Create(name + ".key")
    if err != nil {
        panic(err)
    }
    pem.Encode(keyFile, &keyBlock)
    keyFile.Close()

    commonName := "myuser"
    emailAddress := "someone@myco.com"
    org := "My Co, Inc."
    orgUnit := "Widget Farmers"
    city := "Seattle"
    state := "WA"
    country := "US"

    subject := pkix.Name{
            CommonName: commonName,
            Country: []string{country},
            Locality: []string{city},
            Organization: []string{org},
            OrganizationalUnit: []string{orgUnit},
            Province: []string{state},

    }
    uri, err := url.ParseRequestURI(host)
        if err != nil {
        panic(err)  
    }
    asn1, err := asn1.Marshal(subject.ToRDNSequence())
    if err != nil {
        panic(err)  
    }
    template := x509.CertificateRequest{
            RawSubject: asn1,
            Subject: subject,
            DNSNames: []string{uri.Host},
            EmailAddresses: []string{emailAddress},
    }
    csr, err := x509.CreateCertificateRequest(rand.Reader, &template, key)
    if err != nil {
        panic(err)  
    }
    csrFile, err := os.Create(name + ".csr")
    if err != nil {
        panic(err)  
    }
    pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: csr})
    csrFile.Close()
}