package generic

import (
	"fmt"
	"testing"
)

type Store[T any] interface {
	Save(entity Entity[T])
	Load(entity Entity[T])
}

type Entity[T any] struct {
	data T
}

type DBStore struct {
}

func (d *DBStore) Save(entity Entity[int]) {
	fmt.Println("save to db")
}

func (d *DBStore) Load(entity Entity[int]) {
	fmt.Println("load from db")
}

func Test1(t *testing.T) {
	db := &DBStore{}
	e := Entity[int]{data: 1}
	db.Save(e)

}
