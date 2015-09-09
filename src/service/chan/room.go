package main

import (
	"errors"

	//"fmt"
	"service/proto"
	"time"

	"github.com/wzshiming/base"
	"github.com/wzshiming/server/agent"
	"github.com/wzshiming/server/rooms"
	cards "github.com/wzshiming/ygo_cards"
	ygo "github.com/wzshiming/ygo_core"
	//"ygo/defaul"
)

var cardBag = cards.CardBag_test

type Room struct {
	room      *rooms.Rooms
	gameList  map[uint]*ygo.YGO
	saveQuery map[string]*base.EncodeBytes
}
type dataUniq struct {
	Uniq uint `json:"__YGOGame__,string"`
}

func NewRoom() *Room {
	r := Room{
		room:      rooms.NewRooms("YGORooms"),
		gameList:  make(map[uint]*ygo.YGO),
		saveQuery: make(map[string]*base.EncodeBytes),
	}
	go r.Games()
	return &r
}

func (r *Room) YGOGame(sesss ...*agent.Session) {
	uniq := sesss[0].ToUint()
	room := rooms.NewRooms("YGOGame")
	for _, v := range sesss {
		head := r.room.Head(v)
		v.LockSession()
		r.room.Leave(v)
		res, err := room.Join(v, head)
		if err != nil {
			v.UnlockSession(nil)
			return
		} else {
			v.UnlockSession(&agent.Response{
				Head: head,
				Data: base.SumJson(base.EnJson(dataUniq{uniq}), res),
				Response: base.EnJson(map[string]string{
					"status": "init",
				}),
			})
		}
	}
	game := ygo.NewYGO(room)
	r.gameList[uniq] = game
	game.CardVer = cardBag
	game.Loop()
	for _, v := range sesss {
		room.LeaveMutex(v)
	}
}

func (r *Room) Games() {
	tick := time.Tick(time.Second * 1)
	for {
		<-tick
		r.room.BroadcastPush(map[string]string{"info": "匹配对手中请等待..."}, func(sess *agent.Session) {
			r.room.Leave(sess)
		})
		if r.room.Len() >= 2 {
			if g := r.room.GroupFromSize(2); g != nil {
				go r.YGOGame(g...)
			}
		}
	}
}

// 搜索卡牌
func (r *Room) CardFind(args agent.Request, reply *agent.Response) error {
	//var d struct {
	//	Language string
	//}

	var gr struct {
		Query string `json:"query"`
	}
	args.Request.DeJson(&gr)
	b := r.saveQuery[gr.Query]
	if b == nil {
		b = base.EnJson(cardBag.Find(gr.Query, true))
		r.saveQuery[gr.Query] = b
	}

	//args.Session.Data.DeJson(&d)
	//Trans := i18n.TranslationsForLocale(d.Language)
	reply.Response = b
	return nil
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
	//	if id := rooms.GetFromRoom(args.Session, "Users"); id == 0 {
	//		return errors.New(Trans.Value("chan.nologin"))
	//	}

	// 判断是否在队列中
	if id := rooms.GetFromRoom(args.Session, "YGORooms"); id != 0 {
		return errors.New(Trans.Value("chan.hasjoined"))
	}

	// 判断是否在游戏中
	if id := rooms.GetFromRoom(args.Session, "YGOGame"); id != 0 {
		return errors.New(Trans.Value("chan.hasjoined"))
	}

	// 储存牌组
	id := uint64(rooms.GetFromRoom(args.Session, "Users"))
	deck := GetDeck(id)
	//rego.INFO(deck)
	if len(deck) == 0 || len(deck[0].Main) == 0 {
		return errors.New(Trans.Value("chan.deckerr"))
	}
	// 加入队列
	if ret, err := r.room.Join(args.Session, args.Head); err != nil {
		return errors.New(Trans.Value("chan.joinfail"))
	} else {
		reply.Data = base.SumJson(ret, base.EnJson(struct {
			Deck proto.Deck `json:"deck"`
		}{deck[0]}))
	}

	return nil
}

// 退出匹配队列
func (r *Room) StopMatch(args agent.Request, reply *agent.Response) error {
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	Trans := i18n.TranslationsForLocale(d.Language)

	// 判断是否登入
	if id := rooms.GetFromRoom(args.Session, "Users"); id == 0 {
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

func (r *Room) GameBC(args agent.Request, bc func(*ygo.YGO, *agent.Session) error) error {
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
	return bc(game, sess)
}

func (r *Room) GameRegister(args agent.Request, reply *agent.Response) error {
	return r.GameBC(args, func(game *ygo.YGO, sess *agent.Session) error {
		game.Room.SetHead(sess, args.Head)
		return nil
	})
}

func (r *Room) GameGetDeck(args agent.Request, reply *agent.Response) error {
	//	var d struct {
	//		Language string
	//	}
	//args.Session.Data.DeJson(&d)
	//Trans := i18n.TranslationsForLocale(d.Language)

	// 判断是否登入
	id := uint64(rooms.GetFromRoom(args.Session, "Users"))
	//if id := rooms.GetFromRoom(args.Session, "Users"); id == 0 {
	//return errors.New(Trans.Value("chan.nologin"))
	//}
	// 获取卡组
	odecks := GetDeck(id)

	//rego.INFO(odecks)
	reply.Response = base.EnJson(odecks)
	return nil

}

func (r *Room) GameSetDeck(args agent.Request, reply *agent.Response) error {
	var d struct {
		Language string
	}
	args.Session.Data.DeJson(&d)
	//Trans := i18n.TranslationsForLocale(d.Language)

	gr := proto.Deck{}
	args.Request.DeJson(&gr)
	// 判断是否登入
	id := uint64(rooms.GetFromRoom(args.Session, "Users"))
	//if id := rooms.GetFromRoom(args.Session, "Users"); id == 0 {
	//return errors.New(Trans.Value("chan.nologin"))
	//}
	//gr.UserId = id

	if len(gr.Main) == 0 && len(gr.Extra) == 0 && len(gr.Side) == 0 {
		DelDeck(id, gr.Name)
	} else {
		SetDeck(id, gr.Name, gr)
		gr.UpdatedAt = time.Now()
	}

	return nil
}

func (r *Room) GameDelDeck(args agent.Request, reply *agent.Response) error {
	return r.GameBC(args, func(game *ygo.YGO, sess *agent.Session) error {
		return nil
	})
}

func (r *Room) GameCardActionSelectable(args agent.Request, reply *agent.Response) error {
	return r.GameBC(args, func(game *ygo.YGO, sess *agent.Session) error {
		ar := proto.SelectableRequest{}
		args.Request.DeJson(&ar)
		game.GetPlayer(sess).AddCode(ar.Uniq, ar.Method)
		return nil
	})
}
