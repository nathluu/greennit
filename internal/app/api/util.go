package api

import (
	"github.com/globalsign/mgo"
	"github.com/nathluu/greennit/internal/pkg/db/mongodb"
	"sync"
)

var (
	session *mgo.Session
	sessionOnce sync.Once
)

func dialDefaultMongoDB(conf *mongodb.Config) (*mgo.Session, error) {
	var err error
	sessionOnce.Do(func() {
		session, err = mongodb.Dial(conf)
	})

	if err != nil {
		return nil, err
	}

	s := session.Clone()
	return s, nil
}