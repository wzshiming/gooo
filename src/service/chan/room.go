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
	}
	go r.Games()
	return &r
}

func (r *Room) YGOGame(sesss []*agent.Session) {
	uniq := sesss[0].ToUint()
	room := r.roomGame.GetChild(fmt.Sprint(uniq))

	for _, v := range sesss {
		v.Mutex(func() {
			r.roomHall.ToChild(v, "Game")
			code := r.roomMatc.Leave(v)
			room.Join(v, code.Head)
			v.Data.Set("gameYGO", room.Name())
			base.INFO(v.Rooms.Data())
			return
		})
	}
	time.Sleep(time.Second)

	room.Broadcast(map[string]string{
		"status": "init",
	})
	game := ygo.NewYGO(room)
	r.gameList[room.Name()] = game
	game.CardVer = cardBag
	game.Loop()
	for _, v := range sesss {
		v.Mutex(func() {
			r.roomGame.ToParent(v)
			room.Leave(v)
			v.Data.Del("gameYGO")
		})
	}
	delete(r.gameList, room.Name())
}

func (r *Room) Games() {
	tick := time.Tick(time.Second * 1)
	for {
		<-tick
		size := r.roomMatc.Len()
		msg := map[string]int{
			"inGameNum":  len(r.gameList),
			"inHallNum":  r.roomHall.Len(),
			"inMatchNum": size,
		}
		r.roomHall.Broadcast(msg)
		if size >= 2 {
			if g := r.roomMatc.GroupFromSize(2); g != nil {
				go r.YGOGame(g)
			}
		}
	}
}

// 搜索卡牌
func (r *Room) CardFind(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		var gr struct {
			Query string `json:"query"`
		}
		args.Request.DeJson(&gr)
		b := r.saveQuery[gr.Query]
		if b == nil {
			b = base.EnJson(cardBag.Find(gr.Query, true))
			r.saveQuery[gr.Query] = b
		}
		reply.ReplyEncode(b)
	})
	return nil
}

// 匹配游戏
func (r *Room) MatchCompetitors(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		// 判断是否登入
		id0 := uint64(args.Session.RoomsUniq("Users"))
		if id0 == 0 {
			reply.ReplyError("auth.nologin")
			return
		}

		// 判断是否在队列中
		if id := r.roomMatc.Uniq(args.Session); id != 0 {
			reply.ReplyError("auth.hasjoined")
			return
		}

		// 判断是否在游戏中
		if id := r.roomGame.Uniq(args.Session); id != 0 {
			reply.ReplyError("auth.hasjoined")
			return
		}

		// 储存牌组
		deck := GetDef(id0)
		if deck.Id == 0 {
			reply.ReplyError("auth.deckerr")
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
		if id := args.Session.RoomsUniq("Users"); id == 0 {
			reply.ReplyError("auth.nologin")
			return
		}

		// 退出队列
		r.roomMatc.Leave(args.Session)
		reply.ReplyError("")
	})
	return nil
}

func (r *Room) InfoRegister(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		if id := args.Session.RoomsUniq("Users"); id == 0 {
			reply.ReplyError("auth.nologin")
			return

		}
		r.roomHall.Join(args.Session, args.Head)
		reply.ReplyError("")
	})
	return nil
}

func (r *Room) GameNewDeck(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		gr := proto.Deck{}
		args.Request.DeJson(&gr)
		// 判断是否登入
		id := uint64(args.Session.RoomsUniq("Users"))
		if id == 0 {
			reply.ReplyError("auth.nologin")
			return
		}
		gr.UserId = id
		gr.UpdatedAt = time.Now()
		NewDeck(id)
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
//			reply.ReplyError("auth.nologin")
//			return
//		}
//		//DelDeck(id, gr.Name)
//		reply.ReplyError("")
//	})
//	return nil
//}

func (r *Room) GameGetDeck(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		// 判断是否登入
		id := args.Session.RoomsUniq("Users")
		if id == 0 {
			reply.ReplyError("auth.nologin")
			return
		}
		// 获取卡组
		odecks := GetDecks(uint64(id))
		reply.Reply(odecks)
	})
	return nil

}

func (r *Room) GameSetDeck(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {

		gr := proto.Deck{}
		args.Request.DeJson(&gr)
		// 判断是否登入
		id := uint64(args.Session.RoomsUniq("Users"))
		if id == 0 {
			reply.ReplyError("auth.nologin")
			return
		}
		gr.UserId = id

		if len(gr.Main) == 0 && len(gr.Extra) == 0 && len(gr.Side) == 0 {
			//DelDeck(id, gr.Id)
		} else {
			SetDeck(id, gr.Id, gr)
		}
		reply.ReplyError("")
	})
	return nil
}

func (r *Room) GameRegister(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		r.GameBC(args, reply, func(game *ygo.YGO, sess *agent.Session) error {
			game.Room.SetHead(sess, args.Head)
			return nil
		})
		reply.ReplyError("")
	})
	return nil
}

func (r *Room) GameCardActionSelectable(args agent.Request, reply *agent.Response) error {
	args.Mutex(reply, func() {
		err := r.GameBC(args, reply, func(game *ygo.YGO, sess *agent.Session) error {
			ar := proto.SelectableRequest{}
			args.Request.DeJson(&ar)
			game.GetPlayer(sess).AddCode(ar.Uniq, ar.Method)
			return nil
		})
		if err != nil {
			reply.ReplyError(err.Error())
		}
	})
	return nil
}

func (r *Room) GameBC(args agent.Request, reply *agent.Response, bc func(*ygo.YGO, *agent.Session) error) error {

	var name string

	args.Session.Data.Get("gameYGO", &name)

	if name == "" {
		return errors.New("auth.nogame1")
	}
	game := r.gameList[name]
	if game == nil {
		return errors.New("auth.nogame2")
	}
	if !game.Room.IsExist(args.Session) {
		return errors.New("auth.nogame3")
	}

	return bc(game, args.Session)
}
