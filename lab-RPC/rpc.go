
/*
 * Copyright (c) 2019  郑建勋(jonson)
 * License: https://creativecommons.org/licenses/by-nc-sa/4.0/
 * go语言交流3群：713385260
 */

package main

import (
	"bytes"
	"encoding/gob"
	"log"
	"math/rand"
	"reflect"
	"strings"
	"sync"
	"time"
)

//回复
type replyMsg struct {
	ok    bool
	reply []byte
}

//请求信息
type reqMsg struct {
	endname  interface{} // name of sending ClientEnd
	svcMeth  string      // e.g. "Raft.AppendEntries"
	argsType reflect.Type
	args     []byte
	replych  chan replyMsg
}

//客服端
type ClientEnd struct {
	endname interface{}
	ch      chan reqMsg
}

//远程调用
//@svcMeth "Raft.AppendEntries"
//args参数的序列化
//reply 回复
// send an RPC, wait for the reply.
// the return value indicates success; false means that
// no reply was received from the server.
func (e *ClientEnd) Call(svcMeth string, args interface{}, reply interface{}) bool {
	req := reqMsg{}
	req.endname = e.endname
	req.svcMeth = svcMeth
	req.argsType = reflect.TypeOf(args)
	req.replych = make(chan replyMsg)

	qb := new(bytes.Buffer)
	qe := gob.NewEncoder(qb)
	qe.Encode(args)

	req.args = qb.Bytes()

	e.ch <- req

	rep := <-req.replych

	if rep.ok {
		rb := bytes.NewBuffer(rep.reply)
		rd := gob.NewDecoder(rb)
		if err := rd.Decode(reply); err != nil {
			log.Fatalf("ClientEnd.Call():decode reply:%v\n", err)
		}
		return true

	} else {
		return false
	}
}

//服务器
type Server struct {
	mu sync.Mutex
	//服务器连接服务
	services map[string]*Service
	//连接过的rpc数量
	count int
}

//新建服务器
func MakeServer() *Server {
	rs := &Server{}

	rs.services = map[string]*Service{}
	return rs
}
//增加服务
func (rs *Server) AddService(svc *Service) {
	rs.mu.Lock()

	defer rs.mu.Unlock()

	rs.services[svc.name] = svc
}
//获取数量
func (rs *Server) GetCount() int {
	rs.mu.Lock()

	defer rs.mu.Unlock()

	return rs.count
}

//服务器处理请求
func (rs *Server) dispatch(req reqMsg) replyMsg {
	rs.mu.Lock()
	rs.count++
	//分离"Raft.AppendEntries"为 Raft 和 AppendEntries
	dot := strings.LastIndex(req.svcMeth, ".")

	serviceName := req.svcMeth[:dot]

	methodName := req.svcMeth[dot+1:]

	service, ok := rs.services[serviceName]

	rs.mu.Unlock()

	if ok {
		//指定service进行信息处理
		return service.dispatch(methodName, req)

	} else {
		choices := []string{}

		for k, _ := range rs.services {
			choices = append(choices, k)
		}

		log.Fatalf("labrpc.Service.dispatch:unkonw service %v in %v.%v;expecting one of %v\n",
			serviceName, serviceName, methodName, choices)

		return replyMsg{false, nil}

	}
}
//指定service进行信息处理
func (svc *Service) dispatch(methodname string, req reqMsg) replyMsg {
	if method, ok := svc.methods[methodname]; ok {

		args := reflect.New(req.argsType)

		//反序列化

		ab := bytes.NewBuffer(req.args)
		ad := gob.NewDecoder(ab)

		ad.Decode(args.Interface())

		// 为返回分配空间
		replyType := method.Type.In(2)
		replyType = replyType.Elem()
		replyv := reflect.New(replyType)

		//调用函数 func with receiver as first argument
		function := method.Func
		function.Call([]reflect.Value{svc.rcvr, args.Elem(), replyv})

		//序列化
		rb := new(bytes.Buffer)
		re := gob.NewEncoder(rb)
		re.EncodeValue(replyv)

		return replyMsg{true, rb.Bytes()}

	} else {
		choices := []string{}

		for k, _ := range svc.methods {
			choices = append(choices, k)
		}

		log.Fatalf("labrpc.Service.dispatch:unkonw service %v in %v;expecting one of %v\n",
			methodname, req.svcMeth, choices)

		return replyMsg{false, nil}
	}
}

type Service struct {
	name    string
	rcvr    reflect.Value
	typ     reflect.Type
	methods map[string]reflect.Method
}

func MakeService(rcvr interface{}) *Service {
	svc := &Service{}

	svc.typ = reflect.TypeOf(rcvr)
	svc.rcvr = reflect.ValueOf(rcvr)

	svc.name = reflect.Indirect(svc.rcvr).Type().Name() //获取注册服务的名字

	svc.methods = map[string]reflect.Method{}

	for m := 0; m < svc.typ.NumMethod(); m++ {
		method := svc.typ.Method(m)
		mtype := method.Type //方法的类型
		mname := method.Name //方法的名字

		if method.PkgPath != "" || //大写？
			mtype.NumIn() != 3 || //3个参数
			mtype.In(2).Kind() != reflect.Ptr || //最后一个参数是否是指针
			mtype.NumOut() != 0 {

		} else {
			svc.methods[mname] = method
		}
	}

	return svc

}

type Network struct {
	mu sync.Mutex

	reliable       bool //可靠
	longDelay      bool //延迟
	longreordering bool // 延迟回应

	ends        map[interface{}]*ClientEnd // 客户端名字  --> 客户端结构体
	enabled     map[interface{}]bool
	servers     map[interface{}]*Server     // 服务器名字 ————》服务器结构体
	connections map[interface{}]interface{} //客户端---》服务器
	endch       chan reqMsg
}

//初始化网络
func MakeNetWork() *Network {
	rn := &Network{}
	rn.reliable = true
	rn.ends = map[interface{}]*ClientEnd{}
	rn.enabled = map[interface{}]bool{}
	rn.servers = map[interface{}]*Server{}
	rn.connections = map[interface{}]interface{}{}
	rn.endch = make(chan reqMsg)

	go func() {
		for xreq := range rn.endch {
			go rn.ProcessReq(xreq)
		}
	}()

	return rn

}

//处理请求  模拟延迟
func (rn *Network) ProcessReq(req reqMsg) {

	enable, servername, server, reliable, longreordering	 := rn.ReadEndnameInfo(req.endname)

	if enable && servername != nil && server != nil {

		if reliable == false {
			// 短暂延迟 short delay
			ms := (rand.Int() % 27)
			time.Sleep(time.Duration(ms) * time.Millisecond)
		}

		if reliable == false && (rand.Int()%1000) < 100 {
			// 模拟超时 drop the request, return as if timeout
			req.replych <- replyMsg{false, nil}
			return
		}

		ech := make(chan replyMsg)

		go func() {
			r := server.dispatch(req)

			ech <- r
		}()

		var reply replyMsg

		replyOk := false
		serverDead := false

		for replyOk == false && serverDead == false {
			select {
			case reply = <-ech:
				replyOk = true
			case <-time.After(100 * time.Millisecond):
				serverDead = rn.IsServerDead(req.endname, servername, server)
			}
		}

		serverDead = rn.IsServerDead(req.endname, servername, server)

		if replyOk == false || serverDead == true {
			req.replych <- replyMsg{false, nil}
		}else if reliable == false && (rand.Int()%1000) < 100 {
			// drop the reply, return as if timeout   返回超时
			req.replych <- replyMsg{false, nil}
		} else if longreordering == true && rand.Intn(900) < 600 {
			// delay the response for a while  延迟返回
			ms := 200 + rand.Intn(1+rand.Intn(2000))
			time.Sleep(time.Duration(ms) * time.Millisecond)
			req.replych <- reply
		}  else {
			req.replych <- reply
		}
	} else {
		req.replych <- replyMsg{false, nil}
	}

}

//连接服务器与客户端
func (rn *Network) Connet(endname interface{}, servername interface{}) {

	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.connections[endname] = servername

}

func (rn *Network) Enable(endname interface{}, enabled bool) {

	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.enabled[endname] = enabled
}

func (rn *Network) Reliable(Yes bool) {

	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.reliable = Yes
}

func (rn *Network) LongReodering(Yes bool) {

	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.longreordering = Yes
}

func (rn *Network) LongDelay(Yes bool) {

	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.longDelay = Yes
}

//删除服务器
func (rn *Network) DeleteServer(servername interface{}) {

	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.servers[servername] = nil
}

//连接服务器与客户端
func (rn *Network) MakeEnd(endname interface{}) *ClientEnd {

	rn.mu.Lock()
	defer rn.mu.Unlock()

	if _, ok := rn.ends[endname]; ok {
		log.Fatalf("Make End: %v already exist\n", endname)
	}

	e := &ClientEnd{}

	e.endname = endname
	e.ch = rn.endch
	rn.ends[endname] = e

	rn.enabled[endname] = false
	rn.connections[endname] = nil

	return e
}

//连接服务器与客户端
func (rn *Network) IsServerDead(endname interface{}, servername interface{}, server *Server) bool {

	rn.mu.Lock()
	defer rn.mu.Unlock()

	if rn.enabled[endname] == false || rn.servers[servername] != server {
		return true
	}

	return false
}

func (rn *Network) GetCOunt(servername interface{}) int {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	svr := rn.servers[servername]

	return svr.GetCount()
}

func (rn *Network) ReadEndnameInfo(endname interface{}) (enable bool,
	servername interface{}, server *Server, reliable bool, longreordering bool) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	enable = rn.enabled[endname]

	servername = rn.connections[endname]

	if servername != nil {
		server = rn.servers[servername]
	}

	reliable = rn.reliable
	longreordering = rn.longreordering

	return

}

func (rn *Network) AddServer(servername interface{}, rs *Server) {
	rn.mu.Lock()
	defer rn.mu.Unlock()

	rn.servers[servername] = rs
}
