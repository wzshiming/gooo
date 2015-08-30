package ygo

func (ca *Card) Init() {
	ca.Empty()
	ca.original = *ca.baseOriginal
	ca.effects = []*Card{}
	//ca.summoner = ca.owner
	ca.isValid = true
	//	ca.le = LE
	ca.registerNormal()
	ca.baseOriginal.Initialize.Call(ca)

}

func (ca *Card) registerNormal() {

	//ca.AddEvent(InDeck, ca.SetFaceDownAttack)

	// 进入墓地和 移除时 卡牌翻面
	ca.AddEvent(InExtra, ca.SetFaceUpAttack)

	// 破坏
	ca.AddEvent(Destroy, func(c *Card) {

		ca.ToGrave()
		pl := ca.GetSummoner()
		r := Arg{"self": ca.ToUint()}
		if c != nil {
			r["rival"] = c.ToUint()
		}
		pl.MsgPub("{self}被{rival}破坏", r)
	})

	// 战斗破坏
	ca.AddEvent(DestroyBeBattle, func(c *Card) {
		ca.Dispatch(Destroy, c)
	})

	// 规则破坏
	ca.AddEvent(DestroyBeRule, func(c *Card) {
		pl := ca.GetSummoner()
		ca.ToGrave()
		pl.MsgPub("{self}规则破坏", Arg{"self": ca.ToUint()})
	})

	// 花费
	ca.AddEvent(Cost, func() {
		pl := ca.GetSummoner()
		ca.ToGrave()
		pl.MsgPub("{self}被花费", Arg{"self": ca.ToUint()})
	})

	// 丢弃
	ca.AddEvent(Discard, func() {
		pl := ca.GetSummoner()
		ca.ToGrave()
		pl.MsgPub("{self}被丢弃", Arg{"self": ca.ToUint()})
	})

	// 使用完毕
	ca.AddEvent(Depleted, func() {
		pl := ca.GetSummoner()
		ca.ToGrave()
		pl.MsgPub("{self}使用完毕", Arg{"self": ca.ToUint()})
	})

	// 失效
	ca.AddEvent(Disabled, func() {
		ca.Dispatch(UnegisterGlobalListen)
		if ca.isValid {
			ca.isValid = false
		}
	})

	// 进入墓地和除外
	ca.AddEvent(InGrave, ca.SetFaceUpAttack)
	ca.AddEvent(InRemoved, ca.SetFaceUpAttack)

	// 被移除
	ca.AddEvent(Removed, func() {
		pl := ca.GetSummoner()
		ca.ToRemoved()
		pl.MsgPub("{self}被移除", Arg{"self": ca.ToUint()})
	})

	// 被抽到手牌
	ca.AddEvent(InHand, func() {
		pl := ca.GetSummoner()
		pl.Call(setFront(ca))
		pl.Call(exprCard(ca, LE_FaceUpAttack))
		pl.GetTarget().Call(exprCard(ca, LE_FaceDownAttack))
	})

	// 覆盖
	ca.AddEvent(Use2, func() {
		ca.Dispatch(Cover)
	})
	ca.AddEvent(Use1, func() {
		ca.Dispatch(Onset)
	})

	if ca.IsMonster() {
		ca.registerMonster()
		ca.AddEvent(OutMzone, Disabled)
	} else if ca.IsMagicAndTrap() {
		ca.registerMagicAndTrap()
		ca.AddEvent(OutSzone, Disabled)
	}
}

func (ca *Card) registerMagicAndTrap() {
	ca.Range(InHand, OutHand, Arg{
		// 代价 先覆盖
		Pay: func(s string) {
			if s != Onset && s != Cover {
				return
			}
			pl := ca.GetSummoner()
			ca.ToSzone()
			if ca.IsInSzone() {
				ca.SetFaceDownAttack()
				pl.Msg("覆盖{self}成功!", Arg{"self": ca.ToUint()})
			} else {
				ca.StopOnce(s)
			}
		},
	})
	ca.RegisterPay(func(s string) {
		if s != UseMagic && s != UseTrap {
			return
		}
		ca.SetFaceUp()
		pl := ca.GetSummoner()
		pl.MsgPub("发动{self}!", Arg{"self": ca.ToUint()})
	})
}

func (ca *Card) registerMagic() {
	ca.AddEvent(Onset, func() {
		ca.Dispatch(UseMagic, ca)
	})
}

// 注册一张通常魔法卡手动失效
func (ca *Card) RegisterUnordinaryMagic(f interface{}) {
	ca.registerMagic()
	ca.AddEvent(Effect0, f)
	ca.AddEvent(UseMagic, func() {
		pl := ca.GetSummoner()
		if ca.IsValid() {
			ca.Dispatch(Effect0)
			pl.MsgPub("发动{self}成功!", Arg{"self": ca.ToUint()})
		} else {
			pl.MsgPub("发动{self}失败!", Arg{"self": ca.ToUint()})
		}
	})
}

// 注册一张通常魔法卡
func (ca *Card) RegisterOrdinaryMagic(f interface{}) {
	ca.registerMagic()
	ca.AddEvent(Effect0, f)
	ca.AddEvent(UseMagic, func() {
		pl := ca.GetSummoner()
		if ca.IsValid() {
			ca.Dispatch(Effect0)
			pl.MsgPub("发动{self}成功!", Arg{"self": ca.ToUint()})
			ca.Dispatch(Depleted)
		} else {
			pl.MsgPub("发动{self}失败!", Arg{"self": ca.ToUint()})
		}
	})
}

// 注册一个通常的陷阱卡下个回合才能发动的效果
func (ca *Card) registerTrap(event string, e interface{}) {
	ca.AddEvent(InSzone, func() {
		pl := ca.GetSummoner()
		pl.OnlyOnce(RoundEnd, func() {
			ca.RegisterGlobalListen(event, e)
		}, ca, event, e)
	}, event, e)
	ca.AddEvent(Trigger, func() {
		ca.Dispatch(UnegisterGlobalListen)
		ca.Dispatch(UseTrap)
	})
}

// 注册一张通常的陷阱卡 效果
func (ca *Card) RegisterUnordinaryTrap(event string, e interface{}) {
	ca.registerTrap(event, e)

	ca.AddEvent(UseTrap, func() {
		pl := ca.GetSummoner()
		if ca.IsValid() {
			ca.Dispatch(Chain)
			pl.MsgPub("发动{self}成功!", Arg{"self": ca.ToUint()})
		} else {
			pl.MsgPub("发动{self}失败!", Arg{"self": ca.ToUint()})
		}
	})
}

func (ca *Card) RegisterOrdinaryTrap(event string, e interface{}) {
	ca.registerTrap(event, e)
	ca.AddEvent(UseTrap, func() {
		pl := ca.GetSummoner()
		if ca.IsValid() {
			ca.Dispatch(Chain)
			pl.MsgPub("发动{self}成功!", Arg{"self": ca.ToUint()})
			ca.Dispatch(Depleted)
		} else {
			pl.MsgPub("发动{self}失败!", Arg{"self": ca.ToUint()})
		}
	})
}

// 推送卡牌能连锁
func (ca *Card) PushChain(f interface{}) {
	yg := ca.GetSummoner().Game()
	yg.AddEvent(Chain, ca)
	ca.EmptyEvent(Chain)
	ca.AddEvent(Chain, f)
}

// 注册全局效果监听 直到收到 失效Disabled
func (ca *Card) RegisterGlobalListen(event string, e interface{}) {
	yg := ca.GetSummoner().Game()
	yg.AddEvent(event, e, ca)
	ca.OnlyOnce(UnegisterGlobalListen, func() {
		yg.RemoveEvent(event, e, ca)
	}, e, event)
}

// 注册全局效果监听 直到收到 失效Disabled
func (ca *Card) UnegisterGlobalListen() {
	ca.Dispatch(UnegisterGlobalListen)
}

// 注册一个装备魔法卡  装备对象判断  装备上动作 装备下动作
func (ca *Card) RegisterEquipMagic(a Action, f1 interface{}, f2 interface{}) {
	ca.RegisterUnordinaryMagic(func() {
		pl := ca.GetSummoner()
		pl.MsgPub("{self}选择装备的目标!", Arg{"self": ca.ToUint()})
		tar := pl.GetTarget()
		if c := pl.SelectForWarn(pl.Mzone, tar.Mzone, a); c != nil {

			// 装备卡 离开场地时
			ca.OnlyOnce(Disabled, func() {
				ca.Dispatch(Effect2, c)
			}, c)

			c.OnlyOnce(Disabled, func() {
				ca.Dispatch(Depleted)
			}, ca)

			// 监听目标的改变判断目标的改变是否允许
			c.AddEvent(Change, func() {
				if !c.IsInMzone() || !a.Call(c) {
					ca.Dispatch(Depleted)
				}
			}, ca)

			// 执行装备 上的效果
			ca.Dispatch(Effect1, c)
			pl.MsgPub("{self}装备成功!", Arg{"self": ca.ToUint()})
		} else {
			ca.Dispatch(DestroyBeRule)
			pl.MsgPub("选择的目标不合法,{self}被破坏!", Arg{"self": ca.ToUint()})
		}
	})
	ca.AddEvent(Effect1, f1)
	ca.AddEvent(Effect2, f2)
}

func (ca *Card) RegisterPlaceMagic(f interface{}) {
	ca.registerMagic()
}

func (ca *Card) RegisterPay(f interface{}) {
	ca.AddEvent(Pay, f)
}

func (ca *Card) RegisterFlip(f interface{}) {
	ca.AddEvent(Flip, f)
}

func (ca *Card) RegisterFusionMonster(names ...string) {
	h := map[string]int{}
	for _, v := range names {
		h[v]++
	}
	ca.Range(InExtra, OutExtra, Arg{
		Pay: func(s string) {
			if s != SummonFusion {
				return
			}
			pl := ca.GetSummoner()
			se := NewCards()
			cs := NewCards(&pl.Hand.Cards, &pl.Mzone.Cards)
			for k, v := range h {
				is := cs.Find(func(c *Card) bool {
					return c.GetName() == k
				})
				if is.Len() == v {
					is.ForEach(func(c *Card) bool {
						se.EndPush(c)
						return true
					})
				} else if is.Len() > v {
					for i := 0; i != v; i++ {
						tm := pl.SelectForPopup(is)
						if tm == nil {
							pl.MsgPub("选择非法目标 {self}, 召唤失败!", Arg{"self": ca.ToUint()})
							ca.StopOnce(s)
							return
						}
						is.PickedFor(tm)
						se.EndPush(tm)
					}
				} else {
					pl.MsgPub("材料不足 {self}, 召唤失败!", Arg{"self": ca.ToUint()})
					ca.StopOnce(s)
					return
				}
			}
			se.ForEach(func(c *Card) bool {
				c.Dispatch(Cost)
				return true
			})
		},
		SummonFusion: func() {
			ca.Dispatch(SummonSpecial, ca)
		},
	})

}

func (ca *Card) registerMonster() {

	ca.AddEvent(Onset, func() {
		ca.Dispatch(Summon, ca)
	})
	ca.RegisterPay(func(s string) {
		if s != SummonSpecial && s != SummonFlip && s != Flip && s != Summon {
			return
		}
		ca.SetFaceUpAttack()
		ca.ShowInfo()
	})

	ca.AddEvent(SummonSpecial, func() {
		pl := ca.GetSummoner()
		ca.ToMzone()
		ca.SetNotCanChange()
		pl.MsgPub("{self}特殊召唤成功!", Arg{"self": ca.ToUint()})
	})

	// 手牌
	ca.Range(InHand, OutHand, Arg{
		// 代价
		Pay: func(s string) {
			if s != Summon && s != Cover {
				return
			}
			pl := ca.GetSummoner()
			pl.ResetReplyTime()
			i := 0
			if ca.GetLevel() > 6 {
				i += 2
			} else if ca.GetLevel() > 4 {
				i += 1
			}
			if i != 0 {
				pl.MsgPub("{self}需要解放{size}只怪兽作为代价", Arg{"self": ca.ToUint(), "size": i})
			}
			for k := 0; k < i; {
				if t := pl.SelectForWarn(pl.Mzone); t != nil {
					t.Dispatch(Freedom, ca, &k)
				} else {
					ca.StopOnce(s)
					pl.MsgPub("{self}召唤失败，祭品不足!", Arg{"self": ca.ToUint()})
					return
				}
			}
		},
		// 召唤
		Summon: func() {
			if ca.IsValid() {
				pl := ca.GetSummoner()
				ca.ToMzone()
				ca.SetNotCanChange()
				pl.MsgPub("{self}召唤成功!", Arg{"self": ca.ToUint()})
			}
		},
		// 覆盖
		Cover: func() {
			pl := ca.GetSummoner()
			if ca.IsInHand() {
				ca.ToMzone()
				ca.SetFaceDownDefense()
				ca.SetNotCanChange()
				pl.Msg("{self}覆盖成功!", Arg{"self": ca.ToUint()})
			}
		},
	})

	// 怪兽区
	ca.AddEvent(OutMzone, func() {
		ca.HideInfo()
	})
	ca.Range(InMzone, OutMzone, Arg{
		// 被解放
		Freedom: func(c *Card, i *int) {
			pl := ca.GetSummoner()
			if i != nil {
				*i++
			}
			ca.ToGrave()
			pl.MsgPub("{self}被解放!", Arg{"self": ca.ToUint()})
		},
		Expression: func() {
			pl := ca.GetSummoner()
			if ca.IsCanChange() {
				if ca.IsFaceDownDefense() {
					ca.Dispatch(SummonFlip, ca)
				} else if ca.IsFaceUpDefense() {
					ca.SetFaceUpAttack()
					pl.MsgPub("{self}变成攻击表示！", Arg{"self": ca.ToUint()})
				} else if ca.IsFaceUpAttack() {
					ca.SetFaceUpDefense()
					pl.MsgPub("{self}变成防御表示！", Arg{"self": ca.ToUint()})
				} else {
					pl.Msg("{self}当前状态无法改变表示形式！", Arg{"self": ca.ToUint()})
					return
				}
				ca.SetNotCanChange()
			} else {
				pl.Msg("{self}不能改变表示形式！", Arg{"self": ca.ToUint()})
			}
		},

		// 翻转召唤
		SummonFlip: func() {
			pl := ca.GetSummoner()
			pl.MsgPub("{self}翻转召唤！", Arg{"self": ca.ToUint()})
		},

		// 翻转
		Flip: func() {
			pl := ca.GetSummoner()
			pl.MsgPub("{self}翻转！", Arg{"self": ca.ToUint()})
		},

		// 发出战斗宣言
		Declaration: func(c *Card) {
			pl := ca.GetSummoner()

			if c != nil {
				pl.MsgPub("{self}对{rival}发出战斗宣言！", Arg{"self": ca.ToUint(), "rival": c.ToUint()})
			} else {
				pl.MsgPub("{self}对{rival}发出战斗宣言！", Arg{"self": ca.ToUint()})
			}

			var b bool
			if c != nil {
				b = c.IsFaceDown()
				ca.SetFaceUp()
				if ca.IsInMzone() && c.IsInMzone() {
					ca.Dispatch(DamageStep, c)
				}

			} else {
				if ca.IsInMzone() {
					ca.Dispatch(DamageStep)
				}
			}
			if b {
				c.Dispatch(Flip)
			}
		},
		// 战斗判定
		DamageStep: func(c *Card) {
			pl := ca.GetSummoner()
			if c != nil {
				tar := c.GetSummoner()
				pl.MsgPub("{self}攻击了{rival}！", Arg{"self": ca.ToUint(), "rival": c.ToUint()})
				if c.IsAttack() {
					t := ca.GetAttack() - c.GetAttack()
					if t > 0 {
						c.Dispatch(DestroyBeBattle, ca)
						tar.ChangeHp(-t)
						c.Dispatch(Deduct, c, tar)
					} else if t < 0 {
						ca.Dispatch(DestroyBeBattle, c)
						pl.ChangeHp(t)
						ca.Dispatch(Deduct, ca, pl)
					} else {
						c.Dispatch(DestroyBeBattle, ca)
						ca.Dispatch(DestroyBeBattle, c)
					}
				} else if c.IsDefense() {
					t := ca.GetAttack() - c.GetDefense()
					if t > 0 {
						c.Dispatch(DestroyBeBattle, ca)
					} else if t < 0 {
						pl.ChangeHp(t)
						ca.Dispatch(Deduct, ca, pl)
					}
				}
				ca.Dispatch(Fought, c)
				c.Dispatch(Fought, ca)
			} else {
				tar := pl.GetTarget()
				tar.ChangeHp(-ca.GetAttack())
				ca.Dispatch(Deduct, ca, tar)
				pl.MsgPub("{self}直接攻击了{rival}！", Arg{"self": ca.ToUint()})
			}

			ca.SetNotCanAttack()
			ca.SetNotCanChange()
		},
	})
}
