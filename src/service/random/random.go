package main

import (
	"gooo"
	"gooo/protocol"
	"math/rand"
	"time"
)

type Random struct {
	conf *gooo.Configs
	call *gooo.CallServer
	//helper.Control
	r  *rand.Rand
	ch chan int
}

func NewRandom(m *Status) (random *Random) {
	random = &Random{
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
		ch:   make(chan int, 1024),
		call: gooo.NewCallServer("Connect", m.Conf),
	}
	go func() {
		for {
			//time.Sleep(1)
			random.ch <- random.r.Intn(100)
		}
	}()
	return
}

func (r *Random) Range100(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.RandRequest
	gooo.Decode(args.Request, &p)
	if p.Size == 0 {
		p.Size = 1
	}
	t := make([]int, p.Size)
	for i := 0; i != p.Size; i++ {
		t[i] = <-r.ch
	}
	ret := protocol.RandResponse{
		Rands: t,
	}
	res, _ := gooo.Encode(ret)
	*reply = gooo.RpcResponse{
		Error:    "",
		Response: res,
	}
	//log.Println(t)
	return nil
}

func (r *Random) Range100Spacing(args gooo.RpcRequest, reply *gooo.RpcResponse) error {
	var p protocol.SpacingRequest
	gooo.Decode(args.Request, &p)
	if p.Size == 0 {
		p.Size = 1
	}
	b := protocol.SendRequest{
		Clients: []uint64{args.Session.Uniq.Id},
	}
	var t struct {
		Size int
	}
	gooo.Decode(args.Session.Data, &t)

	var re protocol.SendResponse

	for i := 0; i != p.Size; i++ {
		time.Sleep(time.Duration(p.Space))
		ret := protocol.RandResponse{
			Rands: []int{<-r.ch, t.Size},
		}
		res, _ := gooo.Encode(ret)
		b.Data = res
		r.call.CallBySession(args.Session, "Connect.Send", b, &re)
	}

	*reply = gooo.RpcResponse{
		Data: &map[string]interface{}{
			"Size": t.Size + 1,
		},
	}
	//log.Println(t)
	return nil
}
