package main

import (
	"errors"
	"fmt"

	//"fmt"
	"service/proto"
	"time"

	"github.com/wzshiming/base"
	"github.com/wzshiming/server/agent"

	cards "github.com/wzshiming/ygo_cards"
	ygo "github.com/wzshiming/ygo_core"

	//"ygo/defaul"
)

var cardBag = cards.CardBag_test

type Room struct {
	roomMatc  *agent.Room
	roomHall  *agent.Room
	roomGame  *agent.Room
	gameList  map[string]*ygo.YGO
	saveQuery map[string]*base.EncodeBytes
	timer     *base.Timer
}

func NewRoom() *Room {
	h := agent.NewRoom("Hall")
	m := agent.NewRoom("Matc")
	g := h.GetChild("Game")
	r := Room{
		roomMatc:  m,
		roomHall:  h,
		roomGame:  g,
		gameList:  map[string]*ygo.YGO{},
		saveQuery: map[string]*base.EncodeBytes{},
		timer:     base.NewTimer(time.Second),
	}
	go r.timer.Start()
	go r.games()
	return &r
}

func (r *Room) gameReady(sesss []*agent.Session) {
	uniq := sesss[0].ToUint()
	room := r.roomGame.GetChild(fmt.Sprint(uniq))

	for _, v := range sesss {
		v.Mutex(func() {
			r.roomHall.ToChild(v, "Game")
			code := r.roomMatc.Leave(v)
			room.Join(v, nil)
			v.Push(map[string]string{
				"status": "init",
			}, code.Head)
			v.Data.Set("gameYGO", room.Name())
			//base.INFO(v.Rooms.Data())
			return
		})
	}

	game := ygo.NewYGO(room)
	for _, v := range sesss {

		vc := v

		var id uint
		vc.Data.Get("userid", &id)
		var name string
		vc.Data.Get("username", &name)
		var deck proto.Deck
		vc.Data.Get("deck", &deck)

		yd := ygo.NewDeck()
		for _, c := range deck.Main {
			yd.MainAddCard(c.Index, c.Size)
		}
		//		p := game.GetPlayer(vc)
		//		//game.InitForPlayer()
		//		p.Name = name
		//		p.Id = id
		//		p.Decks = yd

		game.InitForPlayer(vc, id, name, yd)
	}

	r.gameList[room.Name()] = game
	game.SetCardVer(cardBag)
	r.timer.NewNode(time.Second*10, func() {
		if !room.IsReady() {
			game.GameOver()
		}
	})
}

//func (r *Room) gameReady(sesss []*agent.Session) {
//	uniq := sesss[0].ToUint()
//	room := r.roomGame.GetChild(fmt.Sprint(uniq))

//	for _, v := range sesss {
//		v.Mutex(func() {
//			r.roomHall.ToChild(v, "Game")
//			code := r.roomMatc.Leave(v)
//			room.Join(v, nil)
//			v.Push(map[string]string{
//				"status": "init",
//			}, code.Head)
//			v.Data.Set("gameYGO", room.Name())
//			//base.INFO(v.Rooms.Data())
//			return
//		})
//	}
//	game := ygo.NewYGO(room)
//	r.gameList[room.Name()] = game
//	game.CardVer = cardBag
//	r.timer.NewNode(time.Second*10, func() {
//		if !room.IsReady() {
//			game.GameOver()
//		}
//	})
//}

func (r *Room) gamePlay(yg *ygo.YGO) {
	yg.Loop()
	r.gameOver(yg)
}

func (r *Room) gameOver(yg *ygo.YGO) {

	yg.Room.ForEach(func(s *agent.Session) {
		s.Mutex(func() {
			r.roomGame.ToParent(s)
			yg.Room.Leave(s)
			s.Data.Del("gameYGO")
		})
	})
	delete(r.gameList, yg.Room.Name())
}

func (r *Room) games() {
	tick := time.Tick(time.Second * 1)

	var ogn, ohn, omn int

	for {
		<-tick

		gn := len(r.gameList)
		hn := r.roomHall.Len()
		mn := r.roomMatc.Len()

		if gn != ogn || hn != ohn || mn != omn {
			msg := map[string]int{
				"inGameNum":  gn,
				"inHallNum":  hn,
				"inMatchNum": mn,
			}
			r.roomHall.Broadcast(msg)
			ogn = gn
			ohn = hn
			omn = mn
		}

		if mn >= 2 {
			if g := r.roomMatc.GroupFromSize(2); g != nil {
				r.gameReady(g)
			}
		}
	}
}

// 搜索卡牌
//func (r *Room) CardFind(args agent.Request, reply *agent.Response) error {
//	args.Mutex(reply, func() {

//		var gr struct {
//			Query string `json:"query"`
//		}
//		args.Request.DeJson(&gr)
//		b := r.saveQuery[gr.Query]
//		if b == nil {
//			b = base.EnJson(cardBag.Find(gr.Query, true))
//			r.saveQuery[gr.Query] = b
//		}
//		reply.ReplyEncode(b)
//	})
//	return nil
//}

// 搜索卡牌
func (r *Room) CardAll(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		reply.Reply(cardBag.AllIsValid())
	})
	return nil
}

// 匹配游戏
func (r *Room) MatchCompetitors(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		// 判断是否登入
		id0 := uint64(args.Session.RoomsUniq("Users"))
		//		if id0 == 0 {
		//			reply.ReplyError("chan.nologin")
		//			return
		//		}

		// 判断是否在队列中
		if id := r.roomMatc.Uniq(args.Session); id != 0 {
			reply.ReplyError("chan.hasjoined")
			return
		}

		// 判断是否在游戏中
		if id := r.roomGame.Uniq(args.Session); id != 0 {
			reply.ReplyError("chan.hasjoined")
			return
		}

		// 储存牌组
		if id0 == 0 {
			id0 = 1
		}
		deck := GetDef(id0)
		if deck.Id == 0 {
			reply.ReplyError("chan.deckerr")
			return
		}

		args.Session.Data.Set("deck", deck)
		// 加入队列
		r.roomMatc.Join(args.Session, args.Head)
		reply.ReplyError("")
	})
	return nil
}

// 退出匹配队列
func (r *Room) StopMatch(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		// 判断是否登入
		//		if id := args.Session.RoomsUniq("Users"); id == 0 {
		//			reply.ReplyError("chan.nologin")
		//			return
		//		}

		// 退出队列
		r.roomMatc.Leave(args.Session)
		reply.ReplyError("")
	})
	return nil
}

// 大厅右下角 显示人数的
func (r *Room) InfoRegister(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		//if id := args.Session.RoomsUniq("Users"); id == 0 {
		//	reply.ReplyError("chan.nologin")
		//	return
		//}
		r.roomHall.Join(args.Session, args.Head)
		reply.ReplyError("")
	})
	return nil
}

//func (r *Room) GameNewDeck(args agent.Request, reply *agent.Response) error {
//	args.Mutex(reply, func() {

//		gr := proto.Deck{}
//		args.Request.DeJson(&gr)
//		// 判断是否登入
//		id := uint64(args.Session.RoomsUniq("Users"))
//		if id == 0 {
//			reply.ReplyError("chan.nologin")
//			return
//		}
//		gr.UserId = id
//		gr.UpdatedAt = time.Now()
//		NewDeck(id)
//		reply.ReplyError("")
//	})
//	return nil

//}

//保存卡组使用名字做索引
func (r *Room) SaveDeckForName(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		var gr struct {
			Name string       `json:"name"`
			Deck []proto.Card `json:"main"`
		}

		args.Request.DeJson(&gr)
		//base.INFO(gr)
		// 判断是否登入
		id := uint64(args.Session.RoomsUniq("Users"))

		//		if id == 0 {
		//			reply.ReplyError("chan.nologin")
		//			return
		//		}
		if id == 0 {
			id = 1
		}
		SetDeck(id, gr.Name, gr.Deck)
		reply.ReplyError("")
	})
	return nil
}

//获取卡组数量和名字
func (r *Room) GetDecksInfo(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		gr := proto.Deck{}
		args.Request.DeJson(&gr)
		// 判断是否登入
		id := uint64(args.Session.RoomsUniq("Users"))
		//		if id == 0 {
		//			reply.ReplyError("chan.nologin")
		//			return
		//		}
		if id == 0 {
			id = 1
		}
		odeck := GetDecks(id)
		//base.INFO(odeck)
		reply.Reply(odeck)
		//reply.ReplyError("")
	})
	return nil
}

//重命名卡组
func (r *Room) RenameDeck(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		gr := proto.Deck{}
		args.Request.DeJson(&gr)
		// 判断是否登入
		//id := uint64(args.Session.RoomsUniq("Users"))
		//		if id == 0 {
		//			reply.ReplyError("chan.nologin")
		//			return
		//		}

		reply.ReplyError("")
	})
	return nil
}

//获取卡组使用名字做索引
func (r *Room) GetDeckForName(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		gr := proto.Deck{}
		args.Request.DeJson(&gr)
		// 判断是否登入
		//id := uint64(args.Session.RoomsUniq("Users"))
		//		if id == 0 {
		//			reply.ReplyError("chan.nologin")
		//			return
		//		}

		reply.ReplyError("")
	})
	return nil
}

//func (r *Room) GameDelDeck(args agent.Request, reply *agent.Response) error {
//	args.Mutex(reply, func() {

//		gr := proto.Deck{}
//		args.Request.DeJson(&gr)
//		// 判断是否登入
//		id := uint64(args.Session.RoomsUniq("Users"))
//		if id == 0 {
//			reply.ReplyError("chan.nologin")
//			return
//		}
//		//DelDeck(id, gr.Name)
//		reply.ReplyError("")
//	})
//	return nil
//}

//func (r *Room) GameGetDeck(args agent.Request, reply *agent.Response) error {
//	args.Mutex(reply, func() {

//		// 判断是否登入
//		id := args.Session.RoomsUniq("Users")
//		if id == 0 {
//			reply.ReplyError("chan.nologin")
//			return
//		}
//		// 获取卡组
//		odecks := GetDecks(uint64(id))
//		reply.Reply(odecks)
//	})
//	return nil

//}

//func (r *Room) GameSetDeck(args agent.Request, reply *agent.Response) error {
//	args.Mutex(reply, func() {

//		gr := proto.Deck{}
//		args.Request.DeJson(&gr)
//		// 判断是否登入
//		id := uint64(args.Session.RoomsUniq("Users"))
//		if id == 0 {
//			reply.ReplyError("chan.nologin")
//			return
//		}
//		gr.UserId = id

//		if len(gr.Main) == 0 && len(gr.Extra) == 0 && len(gr.Side) == 0 {
//			//DelDeck(id, gr.Id)
//		} else {
//			SetDeck(id, gr.Id, gr)
//		}
//		reply.ReplyError("")
//	})
//	return nil
//}

func (r *Room) GameRegister(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		r.gameBC(args, reply, func(game *ygo.YGO, sess *agent.Session) error {
			game.Room.SetHead(sess, args.Head)

			if game.Room.IsReady() {
				go r.gamePlay(game)
			}
			return nil
		})

		reply.ReplyError("")
	})
	return nil
}

func (r *Room) GameCardActionSelectable(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		err := r.gameBC(args, reply, func(game *ygo.YGO, sess *agent.Session) error {
			ar := proto.SelectableRequest{}
			args.Request.DeJson(&ar)
			game.AddCodeForPlayer(sess, ar.Uniq, ar.Method)
			return nil
		})
		if err != nil {
			reply.ReplyError(err.Error())
		}
	})
	return nil
}

func (r *Room) gameBC(args agent.Request, reply *agent.Response, bc func(*ygo.YGO, *agent.Session) error) error {

	var name string

	args.Session.Data.Get("gameYGO", &name)

	if name == "" {
		return errors.New("chan.nogame")
	}
	game := r.gameList[name]
	if game == nil {
		return errors.New("chan.nogame")
	}
	if !game.Room.IsExist(args.Session) {
		return errors.New("chan.nogame")
	}

	return bc(game, args.Session)
}
