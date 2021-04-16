package repository

import (
	"github.com/1r0npipe/web-broker/pkg/model"
	"github.com/1r0npipe/web-broker/pkg/repository"
)

type MemRepo repository.Storage

// func NewMemRepo() *MemRepo {
// 	return &MemRepo{}
// }

func (mr *MemRepo) New() *MemRepo {
	return &MemRepo{}
}

func (mr *MemRepo) Get(getReq string) (string, error) {
	// TODO: impl
	return "", nil
}

func (mr *MemRepo) Put(putReq *model.PutValue) error {
	// TODO: impl
	return nil
}
