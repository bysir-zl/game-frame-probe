package app

// 服务器节点

import (
	"github.com/AsynkronIT/protoactor-go/actor"
	"github.com/bysir-zl/game-frame-probe/common/service"
	"sync"
	"fmt"
	"time"
	"github.com/bysir-zl/hubs/core/hubs"
	"github.com/bysir-zl/hubs/core/net/listener"
	"github.com/AsynkronIT/protoactor-go/remote"
	"strings"
	"github.com/bysir-zl/bygo/log"
	"github.com/bysir-zl/bygo/util/discovery"
	"github.com/bysir-zl/game-frame-probe/common/pbgo"
)

const TAG = "agent"

var (
	id   string = "agent/1"
	addr string = "127.0.0.1"
	port int    = 8080
)

type Server struct {
	Server *discovery.Server
	*actor.PID
}

type Servers map[string]*Server
type ServerGroups map[string]Servers

func (p ServerGroups) SelectServer(serverType string) (server *Server, ok bool) {
	lock.RLock()
	defer lock.RUnlock()
	if servers, has := p[serverType]; has {
		if len(servers) != 0 {
			for _, s := range servers {
				server = s
				ok = true
				return
			}
		}
	}
	return
}

type ServerChange struct {
	server *discovery.Server
	change discovery.ServerChange
}

var (
	lock            = sync.RWMutex{}
	stdServerGroups = ServerGroups{}
)

// 
func GetServers(serverType string) (servers map[string]*Server, has bool) {
	lock.RLock()
	if ss, ok := stdServerGroups[serverType]; ok {
		lock.RUnlock()
		if len(ss) != 0 {
			return ss, true
		}
	} else {
		lock.RUnlock()
	}
	return
}

// 服务监听
func OnServerChange(server *discovery.Server, change discovery.ServerChange) {
	id := server.Id
	switch change {
	case discovery.SC_Online:
		serverAddr := fmt.Sprintf("%s:%d", server.Address, server.Port)
		// 不要连接自己
		if serverAddr == fmt.Sprintf("%s:%d", addr, port) {
			break
		}

		pid := actor.NewPID(serverAddr, id)
		// 通知gate
		pid.Request(&pbgo.GateConnectReq{}, stdGatePid)

		serverType := getServerTypeFromId(id)
		if servers, ok := stdServerGroups[serverType]; ok {
			servers[id] = &Server{PID: pid}
		} else {
			stdServerGroups[serverType] = Servers{
				id: {PID: pid},
			}
		}
		log.InfoT(TAG, "server %s is readied", id)

	case discovery.SC_Offline:
		serverType := getServerTypeFromId(id)
		if servers, ok := stdServerGroups[serverType]; ok {
			delete(servers, server.Id)
		}

		log.InfoT(TAG, "server %s offline", id)
	}
}

func getServerTypeFromId(id string) string {
	return strings.Split(id, "/")[0]
}

var stdGatePid *actor.PID

func Run() {
	router := stdRouter
	gatePid, err := serverNode(router)
	if err != nil {
		panic(err)
	}
	stdGatePid = gatePid

	serverCli(gatePid, router)

	// 注册服务
	lease, err := service.Discoverer.RegisterService(&discovery.Server{
		Id:      id,
		Name:    id,
		Address: addr,
		Port:    port,
	})
	if err != nil {
		panic(err)
	}

	// 监听服务变化
	service.Discoverer.WatchServer(OnServerChange)

	for range time.Tick(time.Second * 5) {
		service.Discoverer.UpdateServerTTL(lease)
	}
}

func serverNode(router *Router) (agentPid *actor.PID, err error) {
	agentActor := NewAgentActor(router)
	nodeAddr := fmt.Sprintf("%s:%d", addr, port)

	remote.Start(nodeAddr)
	agentPid, err = actor.SpawnNamed(actor.FromInstance(agentActor), id)
	return
}

func serverCli(agentPid *actor.PID, router *Router) {
	addr := ":8081"

	handler := &ClientHandler{
		agentPid: agentPid,
		router:   router,
	}

	cliServer = hubs.New(addr, listener.NewWs(), handler)
	log.Info("serverCli started on", addr)
	go func() {
		err := cliServer.Run()
		if err != nil {
			log.Error("serverCli err", err)
		}
	}()

	go func() {
		for range time.Tick(2 * time.Second) {
			log.InfoT(TAG, "online client count:", handler.clientCount)
		}
	}()
}

var cliServer *hubs.Server
