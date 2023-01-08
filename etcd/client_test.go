package etcddemo

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	clientv3 "go.etcd.io/etcd/client/v3"
	"io/ioutil"
	"log"
	"os"
	"testing"
	"time"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 4 * time.Second
	//endpoints      = []string{"192.168.31.17:2379"}
	endpoints = []string{"127.0.0.1:2379"}
)

func TestLease(t *testing.T) {

	cfg := clientv3.Config{
		Endpoints: endpoints,
	}

	cli, err := clientv3.New(cfg)

	if err != nil {
		//log.Fatal(err)
		panic(err)
	}

	defer cli.Close()

	// 创建一个3秒的租约
	lease := clientv3.NewLease(cli)
	grantResp, err := lease.Grant(context.Background(), 3) // 3秒后过期
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}
	leaseID := grantResp.ID

	// 5秒后过期的key
	key := "testkey"
	value := "testvalue"

	// 设置key的过期时间
	_, err = cli.Put(context.Background(), key, value, clientv3.WithLease(clientv3.LeaseID(leaseID)))
	if err != nil {
		//log.Fatal(err)
		panic(err)
	}

	// 保持租约
	lease.TimeToLive(context.Background(), clientv3.LeaseID(leaseID), clientv3.WithAttachedKeys())

	// 等等超过3秒
	time.Sleep(4 * time.Second)

	// 5秒后过期的key
	rsp, err := cli.Get(context.Background(), key)
	if err != nil {
		//log.Println("Get failed. ", err)
		panic(err)
	} else {
		for i, kv := range rsp.Kvs {
			log.Printf("Get index %d , key=%s, value= %s\n", i, kv.Key, kv.Value)
		}
	}

	log.Println("Done!")
}

func TestTLSClient(t *testing.T) {
	var homedir = os.Getenv("HOME")
	var etcdCert = homedir + "/dev/biz-workspace/service-solution/certs/client.crt"
	var etcdCertKey = homedir + "/dev/biz-workspace/service-solution/certs/client.key"
	var etcdCa = homedir + "/dev/biz-workspace/service-solution/certs/ca.crt"

	// 加载客户端证书
	cert, err := tls.LoadX509KeyPair(etcdCert, etcdCertKey)
	if err != nil {
		return
	}

	// 加载 CA 证书
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
		TLS:       _tlsConfig, // Client.Config设置 TLS
	}

	cli, err := clientv3.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	defer cli.Close()

	key1, value1 := "testkey1", "testvalue"
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
