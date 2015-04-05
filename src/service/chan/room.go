package main

import (
	"errors"
	"rego"
	//"fmt"
	"rego/agent"
	"rego/misc"
	"service/proto"
	"time"
	"ygo"
	"ygo/defaul"
)

type Room struct {
	room     *misc.Rooms
	gameList map[uint]*ygo.YGO
}
type dataUniq struct {
	Uniq uint `json:"__YGOGame__,string"`
}

func NewRoom() *Room {
	r := Room{
		room:     misc.NewRooms("YGORooms"),
		gameList: make(map[uint]*ygo.YGO),
	}
	go r.Games()
	return &r
}

func (r *Room) YGOGame(sesss ...*agent.Session) {
	uniq := sesss[0].ToUint()
	room := misc.NewRooms("YGO")
	for _, v := range sesss {
		head := r.room.Head(v)
		v.LockSession()
		r.room.Leave(v)
		res, err := room.Join(v, head)
		if err != nil {
			v.UnlockSession(nil)
		} else {
			v.UnlockSession(&agent.Response{
				Head: head,
				Data: rego.SumJson(rego.EnJson(dataUniq{uniq}), res),
				Response: rego.EnJson(map[string]string{
					"status": "init",
				}),
			})
		}
	}
	r.gameList[uniq] = ygo.NewYGO(room)
	//	room.BroadcastPush(map[string]string{
	//		"status": "init",
	//	}, nil)
}

func (r *Room) Games() {
	tick := time.Tick(time.Second * 1)
	for {
		<-tick
		if r.room.Len() >= 2 {
			go r.YGOGame(r.room.GroupFromSize(2)...)
		}
	}
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

	//	// 判断是否登入
	//	if id := misc.GetFromRoom(args.Session, "Users"); id == 0 {
	//		return errors.New(Trans.Value("chan.nologin"))
	//	}

	// 判断是否在队列中
	if id := misc.GetFromRoom(args.Session, "YGORooms"); id != 0 {
		return errors.New(Trans.Value("chan.hasjoined"))
	}

	// 判断是否在游戏中
	if id := misc.GetFromRoom(args.Session, "YGOGame"); id != 0 {
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

func (r *Room) SelectDeck(args agent.Request, reply *agent.Response) error {
	reply.Response = rego.EnJson(proto.SelectDeckResponse{
		Deck: defaul.DefaultDeck,
	})

	return nil
}

func (r *Room) GameInit(args agent.Request, reply *agent.Response) error {
	var d0 struct {
		Language string
	}
	args.Session.Data.DeJson(&d0)
	Trans := i18n.TranslationsForLocale(d0.Language)

	var d dataUniq
	args.Session.Data.DeJson(&d)
	if d.Uniq == 0 {
		return errors.New(Trans.Value("chan.nogame1"))
	}
	game := r.gameList[d.Uniq]
	if game == nil {
		return errors.New(Trans.Value("chan.nogame2"))
	}
	sess := game.Room.Sync(args.Session)
	if sess == nil {
		return errors.New(Trans.Value("chan.nogame3"))
	}
	gi := proto.GameInitResponse{}
	game.Room.ForEach(func(se *agent.Session) {
		gi.Users = append(gi.Users, proto.PlayerInit{
			Deck: defaul.DefaultDeck,
			Hp:   8000,
			Name: "no name",
		})
		if se == sess {
			gi.Index = len(gi.Users) - 1
		}
	})
	reply.Response = rego.EnJson(gi)

	return nil
}

func (r *Room) Game(args agent.Request, reply *agent.Response) error {
	var d0 struct {
		Language string
	}
	args.Session.Data.DeJson(&d0)
	Trans := i18n.TranslationsForLocale(d0.Language)
	var d dataUniq
	args.Session.Data.DeJson(&d)
	if d.Uniq == 0 {
		return errors.New(Trans.Value("chan.nogame"))
	}
	game := r.gameList[d.Uniq]
	if game == nil {
		return errors.New(Trans.Value("chan.nogame"))
	}
	sess := game.Room.Sync(args.Session)
	if sess == nil {
		return errors.New(Trans.Value("chan.nogame"))
	}
	game.Room.SetHead(sess, args.Head)
	sess.Readlist() <- args.Request
	return nil
}
