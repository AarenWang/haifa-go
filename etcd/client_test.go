package etcddemo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"go.etcd.io/etcd/clientv3"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 4 * time.Second
	endpoints      = []string{"192.168.31.17:2379"}
)

func TestTLSClient(t *testing.T) {
	var homedir = os.Getenv("HOME")
	var etcdCert = homedir + "/dev/biz-workspace/service-solution/certs/client.crt"
	var etcdCertKey = homedir + "/dev/biz-workspace/service-solution/certs/client.key"
	var etcdCa = homedir + "/dev/biz-workspace/service-solution/certs/ca.crt"

	cert, err := tls.LoadX509KeyPair(etcdCert, etcdCertKey)
	if err != nil {
		return
	}

	caData, err := ioutil.ReadFile(etcdCa)
	if err != nil {
		return
	}

	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(caData)

	_tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{cert},
		RootCAs:      pool,
	}

	cfg := clientv3.Config{
		Endpoints: endpoints,
		TLS:       _tlsConfig,
	}

	cli, err := clientv3.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer cli.Close()

	key1, value1 := "testkey1", "value"

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, key1, value1)
	cancel()
	if err != nil {
		log.Println("Put failed. ", err)
	} else {
		log.Printf("Put {%s:%s} succeed\n", key1, value1)
	}

	rsp, err := cli.Get(context.Background(), key1)
	if err != nil {
		log.Println("Get failed. ", err)
	} else {
		for i, kv := range rsp.Kvs {
			log.Printf("Get index %d , key=%s, value= %s\n", i, kv.Key, kv.Value)
		}
	}

	log.Println("Done!")
}
