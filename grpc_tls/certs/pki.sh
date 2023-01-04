#!/bin/bash -x
export NAME1=grpc-server
export ADDRESS1=192.168.31.17


days=3650

cat > openssl.conf << EOF
[req]
req_extensions = v3_req
distinguished_name = req_distinguished_name
[req_distinguished_name]
[ v3_req ]
keyUsage = critical, digitalSignature, keyEncipherment
extendedKeyUsage = serverAuth, clientAuth
subjectAltName = @alt_names
[alt_names]
DNS.1 = $NAME1
IP.1 = 127.0.0.1
IP.2 = $ADDRESS1
EOF


# 准备 CA 证书
[ -f ca.key ] || openssl genrsa -out ca.key 2048
[ -f ca.crt ] || openssl req -x509 -new -nodes -key ca.key -subj "/CN=grpc-ca" -days ${days} -out ca.crt

# 创建 grpc client 证书
[ -f client.key ] || openssl genrsa -out client.key 2048
[ -f client.csr ] || openssl req -new -key client.key -subj "/CN=grpc-client" -out client.csr -config openssl.conf
[ -f client.crt ] || openssl x509 -req -sha256 -in client.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out client.crt -days ${days} -extensions v3_req  -extfile openssl.conf


[ -f server.key ] || openssl genrsa -out server.key 2048
[ -f server.csr ] || openssl req -new -key server.key -subj "/CN=grpc-server" -out server.csr -config openssl.conf
[ -f server.crt ] || openssl x509 -req -sha256 -in server.csr -CA ca.crt -CAkey ca.key -CAcreateserial -out server.crt -days ${days} -extensions v3_req  -extfile openssl.conf