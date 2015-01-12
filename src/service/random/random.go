package main

import (
	"gooo/configs"
	"gooo/encoder"
	"gooo/helper"
	"gooo/protocol"
	"gooo/router"
	"log"
	"math/rand"
	"reflect"
	connprtc "service/connect/protocol"
	randprtc "service/random/protocol"
	"time"
)

type Random struct {
	conf *configs.Configs
	call *router.CallServer
	//helper.Control
	r  *rand.Rand
	ch chan int
}

func NewRandom() (random *Random) {
	random = &Random{
		r:  rand.New(rand.NewSource(time.Now().UnixNano())),
		ch: make(chan int, 1024),
	}
	go func() {
		for {
			//time.Sleep(1)
			random.ch <- random.r.Intn(100)
		}
	}()
	return
}

func (r *Random) Range100(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p randprtc.RandRequest
	encoder.Decode(args.Request, &p)
	if p.Size == 0 {
		p.Size = 1
	}
	t := make([]int, p.Size)
	for i := 0; i != p.Size; i++ {
		t[i] = <-r.ch
	}
	ret := randprtc.RandResponse{
		Rands: t,
	}
	res, _ := encoder.Encode(ret)
	*reply = protocol.RpcResponse{
		Error:    0,
		Response: res,
	}
	//log.Println(t)
	return nil
}

func (r *Random) Range100Spacing(args protocol.RpcRequest, reply *protocol.RpcResponse) error {
	var p randprtc.SpacingRequest
	encoder.Decode(args.Request, &p)
	if p.Size == 0 {
		p.Size = 1
	}
	b := connprtc.SendRequest{
		Clients: []uint{args.Session.Uniq.Id},
	}
	sess := args.Session.Data
	vv := reflect.ValueOf(sess["size"])
	log.Println(vv)
	log.Println(args.Id, sess)
	log.Println(sess["size"])
	tmp, _ := sess["size"].(float64)
	rss := int(tmp)
	log.Println(rss)
	var re connprtc.SendResponse
	//args.Session.Uniq.Server
	for i := 0; i != p.Size; i++ {
		time.Sleep(time.Duration(p.Space))
		ret := randprtc.RandResponse{
			Rands: []int{<-r.ch, rss},
		}
		res, _ := encoder.Encode(ret)
		b.Data = res
		r.call.CallBySession(args.Session, "Connect.Send", b, &re)
		//sc.Call("Connect.Send", b, &re)
		//log.Println("ret",b)
	}

	*reply = protocol.RpcResponse{
		Error: 0,
		Data: &map[string]interface{}{
			"size": rss + 1,
		},
	}
	//log.Println(t)
	return nil
}

func (r *Random) Init(args protocol.InitRequest, reply *int) error {
	if args.State == 1 {
		r.conf = &args.Conf
		//r.conf.StartConnect()
		r.call = router.NewCallServer("Connect", r.conf)
	}
	return nil
}

func main() {
	defer helper.Recover()
	h := helper.NewHandeln()
	c := NewRandom()
	h.Register(c)
	h.Register(helper.NewStatus(h))
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
