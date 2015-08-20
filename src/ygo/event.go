package ygo

func (ca *Card) Init() {
	ca.Empty()
	ca.original = *ca.baseOriginal
	ca.effects = []*Card{}
	//ca.summoner = ca.owner
	ca.isValid = true
	//	ca.le = LE
	ca.registerNormal()
	if ca.IsMonster() {
		ca.registerMonster()
	} else if ca.IsMagicAndTrap() {
		ca.registerMagicAndTrap()
	}
	ca.baseOriginal.Initialize.Call(ca)

}

// 事件各种混乱
// 重新整理
// 精确到  魔陷卡使用过后  魔陷卡被破坏  怪兽卡被解放破坏  等....
func (ca *Card) registerNormal() {

	//ca.AddEvent(InDeck, ca.SetFaceDownAttack)

	// 进入墓地和 移除时 卡牌翻面
	ca.AddEvent(InExtra, ca.SetFaceUpAttack)

	// 花费
	ca.AddEvent(Cost, func() {
		pl := ca.GetSummoner()
		ca.Dispatch(Disabled)
		pl.MsgPub("{self}被花费", Arg{"self": ca.ToUint()})
	})

	// 丢弃
	ca.AddEvent(Discard, func() {
		pl := ca.GetSummoner()
		ca.Dispatch(Disabled)
		pl.MsgPub("{self}被丢弃", Arg{"self": ca.ToUint()})
	})

	// 失效
	ca.AddEvent(Disabled, func() {
		if ca.isValid {
			//pl := ca.GetSummoner()
			ca.isValid = false
			ca.ToGrave()
			ca.CloseEvent(Disabled)
			//pl.MsgPub("{self}失效", Arg{"self": ca.ToUint()})
		}
	})

	// 进入墓地和除外
	ca.AddEvent(InGrave, ca.SetFaceUpAttack)
	ca.AddEvent(InRemoved, ca.SetFaceUpAttack)

	// 离开墓地和除外 初始化
	ca.AddEvent(OutGrave, func() {
		ca.Init()
	})

	ca.AddEvent(OutRemoved, func() {
		ca.Init()
	})

	// 被破坏
	ca.AddEvent(Destroy, func(c *Card) {
		pl := ca.GetSummoner()
		r := Arg{"self": ca.ToUint()}
		if c != nil {
			r["rival"] = c.ToUint()
		}
		ca.Dispatch(Disabled)
		pl.MsgPub("{self}被{rival}破坏", r)
	})

	// 被移除
	ca.AddEvent(Removed, func() {
		pl := ca.GetSummoner()
		ca.ToRemoved()
		pl.MsgPub("{self}被移除", Arg{"self": ca.ToUint()})
	})

	// 覆盖
	ca.AddEvent(Use2, func() {
		ca.Dispatch(Cover)
	})

	// 使用
	ca.AddEvent(Use1, func() {
		if ca.IsMonster() { // 怪兽
			ca.Dispatch(Summon)
		} else {
			ca.Dispatch(Onset)
		}
	})

	// 被抽到手牌
	ca.AddEvent(InHand, func() {
		pl := ca.GetSummoner()
		pl.Call(setFront(ca))
		pl.Call(exprCard(ca, LE_FaceUpAttack))
		pl.GetTarget().Call(exprCard(ca, LE_FaceDownAttack))

	})
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
}

func (ca *Card) registerMagic() {
	ca.RegisterPay(func(s string) {
		if s != UseMagic {
			return
		}
		ca.SetFaceUp()
		pl := ca.GetSummoner()
		pl.MsgPub("发动{self}!", Arg{"self": ca.ToUint()})
	})
	ca.AddEvent(Onset, func() {
		ca.Dispatch(UseMagic)
	})
}

func (ca *Card) RegisterEffect0(f interface{}) {
	ca.AddEvent(Effect0, f)
}

// 注册一张通常魔法卡
func (ca *Card) RegisterOrdinaryMagic(f interface{}) {
	ca.RegisterUnordinaryMagic(f)
	ca.AddEvent(Onset, func() {
		if ca.IsInSzone() {
			ca.Dispatch(Disabled)
		}
	})
}

// 注册一张通常魔法卡手动失效
func (ca *Card) RegisterUnordinaryMagic(f interface{}) {
	ca.registerMagic()
	ca.RegisterEffect0(f)
	ca.AddEvent(Onset, func() {
		if ca.IsInSzone() {
			ca.Dispatch(Effect0)
		}
	})
}

// 推送卡牌能连锁
func (ca *Card) PushChain() {
	yg := ca.GetSummoner().Game()
	yg.AddEvent(Chain, ca)
}

// 注册全局效果监听 直到收到 失效Disabled
func (ca *Card) RegisterGlobalListen(event string, e interface{}) {
	yg := ca.GetSummoner().Game()
	yg.AddEvent(event, e, ca)
	ca.OnlyOnce(Disabled, func() {
		yg.RemoveEvent(event, e, ca)
	})
}

// 注册一个装备魔法卡  装备对象判断  装备上动作 装备下动作
func (ca *Card) RegisterEquipMagic(a Action, f1 interface{}, f2 interface{}) {
	ca.RegisterUnordinaryMagic(func() {
		pl := ca.GetSummoner()
		pl.MsgPub("{self}选择装备的目标!", Arg{"self": ca.ToUint()})
		tar := pl.GetTarget()
		if c := pl.SelectFor(pl.Mzone, tar.Mzone); c != nil && a.Call(c) {

			// 装备卡 离开场地时
			ca.OnlyOnce(Disabled, func() {
				ca.Dispatch(Effect2, c)
			}, c)

			c.OnlyOnce(Disabled, func() {
				ca.Dispatch(Disabled)
			}, ca)

			// 监听目标的改变判断目标的改变是否允许
			c.AddEvent(Change, func() {
				if !c.IsInMzone() || !a.Call(c) {
					ca.Dispatch(Disabled)
				}
			}, ca)

			// 执行装备 上的效果
			ca.Dispatch(Effect1, c)
		} else {
			ca.Dispatch(Disabled)
			pl.MsgPub("选择的目标不合法,{self}被破坏!", Arg{"self": ca.ToUint()})
		}
	})
	ca.AddEvent(Effect1, f1)
	ca.AddEvent(Effect2, f2)
}

func (ca *Card) RegisterPlaceMagic(f interface{}) {
	ca.registerMagic()
}

// 注册一个通常的陷阱卡   下个回合才能发动  以及 连锁的事件
func (ca *Card) registerTrap(event string, e interface{}, f interface{}) {
	ca.AddEvent(InSzone, func() {
		pl := ca.GetSummoner()
		pl.OnlyOnce(RoundEnd, func() {
			ca.RegisterGlobalListen(event, e)
		})
	})

	ca.AddEvent(Trigger+event, func() {
		pl := ca.GetSummoner()
		ca.SetFaceUp()
		pl.MsgPub("发动{self}!", Arg{"self": ca.ToUint()})
		ca.Dispatch(UseTrap)
		if ca.IsInSzone() {
			ca.Dispatch(Effect0)
			ca.Dispatch(Disabled)
		}
	})
	ca.RegisterEffect0(f)
}

func (ca *Card) RegisterOrdinaryTrap(event string, e interface{}, f interface{}) {
	ca.registerTrap(event, e, f)
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
						tm := pl.SelectForCards(is)
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
		SummonFusion: SummonSpecial,
	})

}

func (ca *Card) registerMonster() {

	ca.AddEvent(SummonSpecial, func() {
		pl := ca.GetSummoner()
		ca.isValid = true
		ca.ToMzone()
		ca.SetFaceUpAttack()
		ca.ShowInfo()
		pl.MsgPub("{self}特殊召唤！", Arg{"self": ca.ToUint()})
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
				if t := pl.SelectFor(pl.Mzone); t != nil {
					t.Dispatch(Freedom, ca, &k)
				} else {
					ca.Dispatch(Destroy)
					ca.StopOnce(s)
					pl.MsgPub("{self}召唤失败，祭品不足!", Arg{"self": ca.ToUint()})
					return
				}
			}

		},
		// 召唤
		Summon: func() {
			pl := ca.GetSummoner()
			ca.ToMzone()
			if ca.IsInMzone() {
				ca.SetFaceUpAttack()
				ca.SetNotCanChange()
				pl.MsgPub("{self}召唤成功!", Arg{"self": ca.ToUint()})
				ca.ShowInfo()
			}
		},
		// 覆盖
		Cover: func() {
			pl := ca.GetSummoner()
			ca.ToMzone()
			if ca.IsInMzone() {
				ca.SetFaceDownDefense()
				ca.SetNotCanChange()
				pl.Msg("{self}覆盖成功!", Arg{"self": ca.ToUint()})
				ca.OnlyOnce(Flip, func() {
					ca.ShowInfo()
				})
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
					ca.Dispatch(SummonFlip)
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
			ca.Dispatch(Flip)
			pl.MsgPub("{self}翻转召唤！", Arg{"self": ca.ToUint()})
		},

		// 翻转
		Flip: func() {
			ca.SetFaceUpAttack()
		},

		// 发出战斗宣言
		Declaration: func(tar *Card) {
			pl := ca.GetSummoner()

			if tar != nil {
				pl.MsgPub("{self}对{rival}发出战斗宣言！", Arg{"self": ca.ToUint(), "rival": tar.ToUint()})
			} else {
				pl.MsgPub("{self}对{rival}发出战斗宣言！", Arg{"self": ca.ToUint()})
			}

			var b bool
			if tar != nil {
				b = tar.IsFaceDown()
				ca.SetFaceUp()
				if ca.IsInMzone() && tar.IsInMzone() {
					ca.Dispatch(DamageStep, tar)
				}

			} else {
				if ca.IsInMzone() {
					ca.Dispatch(DamageStep)
				}
			}
			if b {
				tar.Dispatch(Flip)
			}
		},
		// 战斗判定
		DamageStep: func(tar *Card) {
			pl := ca.GetSummoner()
			if tar != nil {
				plt := tar.GetSummoner()
				pl.MsgPub("{self}攻击了{rival}！", Arg{"self": ca.ToUint(), "rival": tar.ToUint()})
				if tar.IsAttack() {
					t := ca.GetAttack() - tar.GetAttack()
					if t > 0 {
						tar.Dispatch(Destroy, ca)
						plt.ChangeHp(-t)
					} else if t < 0 {
						ca.Dispatch(Destroy, tar)
						pl.ChangeHp(t)
					} else {
						tar.Dispatch(Destroy, ca)
						ca.Dispatch(Destroy, tar)
					}
				} else if tar.IsDefense() {
					t := ca.GetAttack() - tar.GetDefense()
					if t > 0 {
						tar.Dispatch(Destroy, ca)
					} else if t < 0 {
						pl.ChangeHp(t)
					}
				}
				ca.Dispatch(Fought, tar)
				tar.Dispatch(Fought, ca)
			} else {
				ca.Dispatch(Deduct)
				pl.GetTarget().ChangeHp(-ca.GetAttack())
				pl.MsgPub("{self}直接攻击了{rival}！", Arg{"self": ca.ToUint()})
			}

			ca.SetNotCanAttack()
			ca.SetNotCanChange()
		},
	})
}
