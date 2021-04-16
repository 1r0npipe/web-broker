package repository

import (
	"fmt"

	"github.com/1r0npipe/web-broker/pkg/model"
)

type (
	Storager interface {
		New() (interface{}, error)
	}

	Storage struct {
		stor     Storager
		typeStor string
	}
)

var stor Storage

func New(typeStor string) (*Storage, error) {
	switch typeStor {
	case "file":
		return &Storage{stor: stor.stor, typeStor: "file"}, nil
	case "memory":
		return &Storage{stor: stor.stor, typeStor: "memory"}, nil
	case "pg":
		return &Storage{stor: stor.stor, typeStor: "postgres"}, nil
	}
	return nil, fmt.Errorf("can't allocate storage type")
}

func (s *Storage) Get(str string) (string, error) {
	return str, nil
}

func (s *Storage) Put(putReq *model.PutValue) error {
	return nil
}
