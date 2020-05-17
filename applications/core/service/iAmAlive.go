package service

import (
	repository "github.com/I-Reven/Hexagonal/infrastructures/repository/mongo/core"
)

func TestDatabase () error {
	iAmAlive := repository.IAmAlive{}
	err := iAmAlive.Add()

	return err
}
