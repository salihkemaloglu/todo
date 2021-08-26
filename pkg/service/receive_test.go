package service

import (
	"testing"

	"github.com/salihkemaloglu/todo/pkg/model"
	"github.com/salihkemaloglu/todo/pkg/util/config"
	"github.com/stretchr/testify/require"
)

func TestReceive(t *testing.T) {

	c := config.LoadConfig("/home/salihkemaloglu/go/src/github.com/salihkemaloglu/todo/config.yml")
	o := model.Object{
		ObjectIDs: []int{57, 70, 87, 80, 73},
	}
	err := Receive(o, c)

	require.Equal(t, nil, err)
}
