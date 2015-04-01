package main

import (
	"errors"
	"rego"
	"rego/agent"
	"rego/misc"
	"time"
	"ygo"
	"ygo/defaul"
)

type Room struct {
	room *misc.Rooms
}

func NewRoom() *Room {
	r := Room{
		room: misc.NewRooms("YGORooms"),
	}
	tick := time.Tick(time.Second * 1)
	go func() {
		for {
			<-tick
			if r.room.Len() >= 2 {
				user1 := r.room.Random()
				r.room.LeaveMutex(user1)
				user2 := r.room.Random()
				r.room.LeaveMutex(user2)
				ygo.NewYGO(
					defaul.DefaultTest(user1, 1),
					defaul.DefaultTest(user2, 2),
				).Loop()
			}
		}
	}()
	return &r
}

// 匹配游戏
func (r *Room) MatchCompetitors(args agent.Request, reply *agent.Response) error {
	//	var p proto.RegisterRequest
	//	args.Request.DeJson(&p)
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	Trans := i18n.TranslationsForLocale(d.Language)

	// 判断是否登入
	if id := misc.GetFromRoom(args.Session, "Users"); id == 0 {
		return errors.New(Trans.Value("chan.nologin"))
	}

	// 判断是否在队列中
	if s := r.room.Sync(args.Session); s != nil {
		return errors.New(Trans.Value("chan.hasjoined"))
	}

	// 加入队列
	if ret, err := r.room.Join(args.Session, args.Head); err != nil {
		return errors.New(Trans.Value("chan.joinfail"))
	} else {
		reply.Data = ret
	}
	return nil
}

// 退出匹配队列
func (r *Room) StopSearching(args agent.Request, reply *agent.Response) error {
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	Trans := i18n.TranslationsForLocale(d.Language)

	// 判断是否登入
	if id := misc.GetFromRoom(args.Session, "Users"); id == 0 {
		return errors.New(Trans.Value("chan.nologin"))
	}

	// 退出队列
	if s := r.room.Sync(args.Session); s == nil {
		return errors.New(Trans.Value("chan.nojoined"))
	} else {
		r.room.Leave(s)
	}

	return nil
}
