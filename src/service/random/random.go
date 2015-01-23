package main

import (
	"gooo/configs"
	"gooo/encoder"
	"gooo/handeln"
	"gooo/helper"
	"gooo/protocol"
	"gooo/router"
	//"log"
	"math/rand"
	//"reflect"
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
		Error:    "",
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
		Clients: []uint64{args.Session.Uniq.Id},
	}
	var t struct {
		Size int
	}
	encoder.Decode(args.Session.Data, &t)

	var re connprtc.SendResponse

	for i := 0; i != p.Size; i++ {
		time.Sleep(time.Duration(p.Space))
		ret := randprtc.RandResponse{
			Rands: []int{<-r.ch, t.Size},
		}
		res, _ := encoder.Encode(ret)
		b.Data = res
		r.call.CallBySession(args.Session, "Connect.Send", b, &re)
	}

	*reply = protocol.RpcResponse{
		Data: &map[string]interface{}{
			"Size": t.Size + 1,
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
	h := handeln.NewHandeln()
	c := NewRandom()
	h.Register(c)
	h.Register(handeln.NewStatus(h))
	helper.EchoPortInfo(configs.Name, configs.Port)
	h.Start(configs.Port)
}
