# Hints

## Create CA (Won't need this if already have valid verified CA)
`openssl genrsa -out ca.key 2048`

## Create CA certificate OR get valid/verified CA from authorized agency
`openssl req -new -x509 -key ca.key -out ca.crt`

## Create Server key
`openssl genrsa -out server.key 2048`

## Create Server certificate
`openssl req -new -sha256 -key server.key -out server.csr` 

watch your server CN (common-name), this will be used for server host http client(your hostname will be OK for testing)

## Sign Server certificate with CA
`openssl x509 -req -days 365 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt` 

## Create Client key
`openssl genrsa -out client.key 2048`

## Create Client certificate
`openssl req -new -sha256 -key client.key -out client.csr`

## Sign Client certificate with CA
`openssl x509 -req -days 365 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt`