package ygo

type MsgCode struct {
	Uniq   uint
	Method uint
}

type MsgChan chan MsgCode

func NewMsgChan() MsgChan {
	return make(MsgChan, 128)
}

func (mc *MsgChan) AddCode(uniq, method uint) {
	*mc <- MsgCode{
		Uniq:   uniq,
		Method: method,
	}
}

func (mc *MsgChan) GetCode() <-chan MsgCode {
	return *mc
}

func (mc *MsgChan) ClearCode() {
	for {
		select {
		case <-mc.GetCode():
		default:
			return
		}
	}
}

type Call struct {
	Method string      `json:"method"`
	Args   interface{} `json:"args"`
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
		"master": t.Owner.Index,
		"pos":    pos,
	}
}

func flashPhases(pl *Player) (string, interface{}) {
	return "flagStep", map[string]interface{}{
		"step": pl.Phases,
		"wait": pl.WaitTime,
	}
}
