package log_r

import (
	"context"
	"github.com/go-logr/logr"
	"log"
	"testing"
)

func Test1(t *testing.T) {
	log.Println("Test1")

	logger := logr.FromContextOrDiscard(context.Background())

	name := "Jack"

	//logger.Debug("This is Debug Level Log ", "user_name= ", name)

	logger.Info("This is Info Level Log ", "user_name= ", name)

	//logger.Warn("This is Warn Level Log ", "user_name= ", name)

	logger.Error(nil, "This is Error Level Log ", "user_name= ", name)

	//logger.Fatal("This is Fatal Level Log ", "user_name= ", name)

	logger.Info("setting target", "value", "100")

}
