package etcddemo

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"testing"
	"time"
)

func TestWatch(t *testing.T) {
	cfg := clientv3.Config{
		Endpoints: endpoints,
	}
	cli, err := clientv3.New(cfg)
	if err != nil {
		panic(err)
	}

	defer cli.Close()

	ctx := context.Background()
	key := "mykey"
	//watchChan := cli.Watch(ctx, key, clientv3.WithKeysOnly())
	watch := clientv3.NewWatcher(cli)
	watchChan := watch.Watch(ctx, key, clientv3.WithKeysOnly())
	watch.RequestProgress(ctx)

	go func() {
		for {
			select {
			case watchResp := <-watchChan:
				for _, event := range watchResp.Events {
					log.Printf("【goroutine1】 type:%v key:%v value:%v", event.Type, string(event.Kv.Key), string(event.Kv.Value))
				}
			default:
				log.Println("【goroutine1】 no event")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	go func() {
		for {
			select {
			case watchResp := <-watchChan:
				for _, event := range watchResp.Events {
					log.Printf("【goroutine2】 type:%v key:%v value:%v", event.Type, string(event.Kv.Key), string(event.Kv.Value))
				}
			default:
				log.Println("【goroutine2】 no event")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	i := 0
	cli.Put(ctx, key, fmt.Sprintf("testvalue i=%d", i))

	go func() {
		log.Println("start loop put")
		for {
			i++
			log.Printf("put i=%d", i)
			cli.Put(ctx, key, fmt.Sprintf("testvalue i=%d", i))
			time.Sleep(time.Second * 3)
		}
	}()

	make(chan struct{}) <- struct{}{}
}
