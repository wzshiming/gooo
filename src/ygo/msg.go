package ygo

type MsgCode struct {
	Uniq   uint
	Method uint
}

type MsgChan struct {
	primary   chan MsgCode
	secondary chan MsgCode
}

func NewMsgChan(m func(MsgCode) bool) MsgChan {
	mc := MsgChan{
		primary:   make(chan MsgCode, 128),
		secondary: make(chan MsgCode, 2048),
	}
	go func() {
		for v := range mc.primary {
			if m(v) {
				mc.secondary <- v
			}
		}
	}()
	return mc
}

func (mc *MsgChan) AddCode(uniq, method uint) {
	c := MsgCode{
		Uniq:   uniq,
		Method: method,
	}
	mc.primary <- c
}

func (mc *MsgChan) GetCode() <-chan MsgCode {
	mc.ClearCode()
	return mc.secondary
}

func (mc *MsgChan) ClearCode() {
	for {
		select {
		case <-mc.secondary:
		default:
			return
		}
	}
}

type Call struct {
	Method string      `json:"method"`
	Args   interface{} `json:"args"`
}

func Touch(t *Card, x, y, z int) (string, interface{}) {
	return "touch", map[string]interface{}{
		"uniq": t.ToUint(),
		"x":    x,
		"y":    y,
		"z":    z,
	}
}

func ExprCard(t *Card, le LE_TYPE) (string, interface{}) {
	return "exprCard", map[string]interface{}{
		"uniq": t.ToUint(),
		"expr": le,
	}
}

func SetFront(t *Card) (string, interface{}) {
	return "setFront", map[string]interface{}{
		"desk": t.GetId(),
		"uniq": t.ToUint(),
	}
}

func Message(msg string) (string, interface{}) {
	return "message", map[string]interface{}{
		"message": msg,
	}
}

func MoveCard(t *Card, pos string) (string, interface{}) {
	return "moveCard", map[string]interface{}{
		"uniq":   t.ToUint(),
		"master": t.GetSummoner().Index,
		"pos":    pos,
	}
}

func flashPhases(pl *Player) (string, interface{}) {
	return "flagStep", map[string]interface{}{
		"step": pl.Phases,
		"wait": pl.WaitTime,
	}
}
