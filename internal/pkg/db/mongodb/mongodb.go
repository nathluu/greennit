package mongodb

import (
	"github.com/globalsign/mgo"
	"github.com/nathluu/greennit/internal/pkg/log"
	"time"
)

type (
	Config struct {
		Addrs    []string      `mapstructure:"addresses"`
		Database string        `mapstructure:"database"`
		Username string        `mapstructure:"username"`
		Password string        `mapstructure:"password"`
		Timeout  time.Duration `mapstructure:"timeout"`
	}
)

func Dial(conf *Config) (*mgo.Session, error) {
	log.Debugf("dailing to target mongoDB at %v, database %s\n", conf.Addrs, conf.Database)
	sess, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:    conf.Addrs,
		Database: conf.Database,
		Username: conf.Username,
		Password: conf.Password,
		Timeout:  conf.Timeout,
	})

	if err != nil {
		return nil, err
	}

	sess.SetMode(mgo.Monotonic, true)
	log.Debugf("success dialing monfoDB at %v\n", conf.Addrs)
	return sess, nil
}
