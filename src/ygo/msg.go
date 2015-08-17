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

func remind(t uint) (string, interface{}) {
	return "remind", map[string]interface{}{
		"uniq": t,
	}
}

func touch(t uint, x, y, z int) (string, interface{}) {
	return "touch", map[string]interface{}{
		"uniq": t,
		"x":    x,
		"y":    y,
		"z":    z,
	}
}

func exprCard(t *Card, le le_type) (string, interface{}) {
	return "exprCard", map[string]interface{}{
		"uniq": t.ToUint(),
		"expr": le,
	}
}

func setFront(t *Card) (string, interface{}) {
	return "setFront", map[string]interface{}{
		"desk": t.GetId(),
		"uniq": t.ToUint(),
	}
}

func message(format string, a Arg) (string, interface{}) {
	return "message", map[string]interface{}{
		"message": format,
		"params":  a,
	}
}

func setPick(cs *Cards, pl *Player) (string, interface{}) {
	if cs != nil {
		return "setPick", map[string]interface{}{
			"master": pl.Index,
			"uniqs":  cs.Uniqs(),
		}
	}
	return "setPick", map[string]interface{}{
		"master": pl.Index,
		"uniqs":  []int{},
	}

}

func setCardFace(t *Card, a Arg) (string, interface{}) {
	return "setCardFace", map[string]interface{}{
		"uniq":   t.ToUint(),
		"params": a,
	}
}

func moveCard(t *Card, pos ll_type) (string, interface{}) {
	return "moveCard", map[string]interface{}{
		"uniq":   t.ToUint(),
		"master": t.GetSummoner().Index,
		"pos":    pos,
	}
}

func flashPhases(pl *Player) (string, interface{}) {
	return "flagStep", map[string]interface{}{
		"step": pl.Phases,
		"wait": pl.PassTime,
	}
}

func setFace(a Arg) (string, interface{}) {
	return "setFace", a
}

func flagName(pl *Player) (string, interface{}) {
	return "flagName", map[string]interface{}{
		"round":  pl.GetRound(),
		"player": pl.Index,
	}
}
