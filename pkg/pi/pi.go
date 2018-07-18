// Copyright 2018 The OpenPitrix Authors. All rights reserved.
// Use of this source code is governed by a Apache license
// that can be found in the LICENSE file.

package pi

import (
	"strings"
	"sync"

	"openpitrix.io/openpitrix/pkg/config"
	"openpitrix.io/openpitrix/pkg/db"
	"openpitrix.io/openpitrix/pkg/etcd"
	"openpitrix.io/openpitrix/pkg/logger"
)

type globalCfgWatcher func(*config.GlobalConfig)

type Pi struct {
	cfg              *config.Config
	globalCfg        *config.GlobalConfig
	globalCfgWatcher []globalCfgWatcher
	Db               *db.Database
	Etcd             *etcd.Etcd
}

var global *Pi
var mutex sync.RWMutex
var globalMutex sync.RWMutex

func NewPi(cfg *config.Config) *Pi {
	p := &Pi{cfg: cfg}
	p.openDatabase()
	p.openEtcd()
	p.watchGlobalCfg()

	if p.Db != nil {
		p.Db.UpdateHook = GetUpdateHook(p)
		p.Db.DeleteHook = GetDeleteHook(p)
		p.Db.InsertHook = GetInsertHook(p)
	}

	return p
}

func SetGlobalPi(cfg *config.Config) {
	globalMutex.Lock()
	global = NewPi(cfg)
	globalMutex.Unlock()
}

func Global() *Pi {
	globalMutex.RLock()
	defer globalMutex.RUnlock()
	return global
}

func (p *Pi) GlobalConfig() (globalCfg *config.GlobalConfig) {
	mutex.RLock()
	globalCfg = p.globalCfg
	mutex.RUnlock()
	return
}

func (p *Pi) setGlobalCfg(globalCfg *config.GlobalConfig) {
	mutex.Lock()
	p.globalCfg = globalCfg
	for _, cb := range p.globalCfgWatcher {
		go cb(globalCfg)
	}
	mutex.Unlock()
}

func (p *Pi) ThreadWatchGlobalConfig(cb globalCfgWatcher) {
	p.globalCfgWatcher = append(p.globalCfgWatcher, cb)
}

func (p *Pi) watchGlobalCfg() *Pi {
	watcher := make(config.Watcher)

	go func() {
		err := config.WatchGlobalConfig(p.Etcd, watcher)
		if err != nil {
			logger.Critical("failed to watch global config")
			panic(err)
		}
	}()

	globalCfg := <-watcher
	p.setGlobalCfg(globalCfg)
	logger.Debug("Pi got global config: [%+v]", p.globalCfg)

	go func() {
		for globalCfg := range watcher {
			p.setGlobalCfg(globalCfg)
			logger.Debug("Global config update to [%+v]", globalCfg)
		}
	}()

	return p
}

func (p *Pi) openDatabase() *Pi {
	if p.cfg.Mysql.Disable {
		return p
	}
	dbSession, err := db.OpenDatabase(p.cfg.Mysql)
	if err != nil {
		logger.Critical("failed to connect mysql")
		panic(err)
	}
	p.Db = dbSession
	return p
}

func (p *Pi) openEtcd() *Pi {
	endpoints := strings.Split(p.cfg.Etcd.Endpoints, ",")
	e, err := etcd.Connect(endpoints, config.EtcdPrefix)
	if err != nil {
		logger.Critical("failed to connect etcd")
		panic(err)
	}
	p.Etcd = e
	return p
}
