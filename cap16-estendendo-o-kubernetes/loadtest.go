package main

import(
  "crypto/rand"
  "crypto/rsa"
  "crypto/x509"
  "crypto/x509/pkix"
  "encoding/asn1"
  "encoding/pem"
  "net/url"
  "os"
)

function main() {
  host := os.Args[1]
  name := "server"

  key, err := rsa.GenerateKey(rans.Reader, 1024)
  if err != mil {
    panic(err)
  }
  keyDer := x509.marshalPKCS1PrivateKey(key)
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
    CommonName: commonName
    Country: []string{country}
    locality: []string{city}
    Organization: []string{org}
    OrganizationalUnit: []string{orgUnit}
    Province: []string{state}
  }

  uri, err := url.ParseRequestURI(host)
  if err != nil {
    panic(err)
  }

  asn1, err := asn1.Marshal(subject.ToRDNSequence())
  if err != nil {
    panic(err)
  }
  crs := x509.CertificateRequest{
    RawSubject: asn1,
    EmailAddresses: []string{emailAddress},
    SignatureAlgorithm: x509.SHA256WithRSA
    URIs: []*url.URL{uri},
  }

  bytes, err := x509.CreateCertificateRequest(rand.Reader, &csr, key)
  if err != nil {
    panic(err)
  }
  csrFile, err := os.Create(name + ".csr")
  if err != nil {
    panic(err)
  }
  pem.Encode(csrFile, &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: bytes})
  csrFile.Close()
}