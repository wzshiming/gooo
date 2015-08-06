package ygo

func (ca *Card) Init() {
	ca.Empty()
	ca.original = *ca.baseOriginal
	ca.effects = []*Card{}
	ca.summoner = ca.owner
	ca.isValid = true
	ca.le = LE_None
	ca.registerNormal()
	if ca.IsMonster() {
		ca.RegisterMonster()
	} else if ca.IsMagicAndTrap() {
		ca.registerMagicAndTrap()
	}
	ca.baseOriginal.Initialize.Call(ca)

}

func (ca *Card) registerNormal() {

	// 进入墓地和 移除时 卡牌翻面
	ca.AddEvent(InExtra, ca.FaceUpAttack)
	//ca.AddEvent(InGrave, Disabled)
	ca.AddEvent(InGrave, ca.FaceUpAttack)
	//ca.AddEvent(InRemoved, Disabled)
	ca.AddEvent(InRemoved, ca.FaceUpAttack)
	// 被察觉
	//ca.AddEvent(Realize, ca.FaceUp)

	//	// 发动
	//	ca.AddEvent(Onset, ca.FaceUp)

	// 花费
	ca.AddEvent(Cost, ca.ToGrave)

	// 丢弃
	ca.AddEvent(Discard, func() {
		pl := ca.GetSummoner()
		pl.MsgPub("{self}被丢弃", Arg{"self": ca.GetId()})
		ca.ToGrave()
	})

	// 失效
	ca.AddEvent(Disabled, func() {
		if ca.IsValid() {
			pl := ca.GetSummoner()
			pl.MsgPub("{self}失效", Arg{"self": ca.GetId()})
			ca.EmptyEvent(Offset)
			ca.ToGrave()
			ca.isValid = false
			//ca.CloseEvent(Disabled)
		}
	})

	// 被破坏
	ca.AddEvent(Destroy, func(c *Card) {
		pl := ca.GetSummoner()
		r := Arg{"self": ca.GetId()}
		if c != nil {
			r["rival"] = c.GetId()
		}
		pl.MsgPub("{self}被{rival}破坏", r)
		ca.ToGrave()
	})

	// 被移除
	ca.AddEvent(Removed, func() {
		pl := ca.GetSummoner()
		pl.MsgPub("{self}被移除", Arg{"self": ca.GetId()})
		ca.ToRemoved()
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
		pl.Call(SetFront(ca))
		pl.Call(ExprCard(ca, LE_FaceUpAttack))
		pl.GetTarget().Call(ExprCard(ca, LE_FaceDownAttack))

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
			if ca.GetSummoner().Szone.Len() < 5 {
				ca.ToSzone()
				ca.FaceDownAttack()
				//pl.Msg("覆盖{self}成功!", Arg{"self": ca.GetId()})
			} else {
				pl.MsgPub("覆盖{self}失败，魔陷区场地不足!", Arg{"self": ca.GetId()})
				ca.Dispatch(Destroy)
				ca.StopOnce(s)
			}
		},
	})
}

func (ca *Card) registerMagic() {
	ca.AddEvent(Onset, func() {
		pl := ca.GetSummoner()
		pl.MsgPub("发动魔法卡{self}!", Arg{"self": ca.GetId()})
		ca.FaceUp()
		ca.Dispatch(UseMagic)
	})
}

func (ca *Card) RegisterEffect0(f interface{}) {
	ca.AddEvent(Effect0, f)
}

// 注册一张通常魔法卡
func (ca *Card) RegisterOrdinaryMagic(f interface{}) {
	ca.registerMagic()
	ca.RegisterEffect0(f)
	ca.AddEvent(Onset, func() {
		if ca.IsInSzone() {
			ca.Dispatch(Effect0)
			ca.Dispatch(Disabled)
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
	ca.registerMagic()
	ca.RegisterEffect0(func() {
		pl := ca.GetSummoner()
		//yg := pl.Game()
		c := pl.Select()
		if c != nil && a.Call(c) {
			// 执行装备 上的效果
			ca.Dispatch(Effect1, ca)

			// 装备卡 离开场地时
			ca.AddEvent(OutSzone, func() {
				// 移除 对目标的监听
				c.RemoveEvent(Change, ca)
				// 卸下装备 执行
				ca.Dispatch(Effect2)
			})
			// 监听目标的改变
			c.AddEvent(Change, ca)

			// 判断目标的改变是否允许
			ca.AddEvent(Change, func() {
				if !c.IsInMzone() || !a.Call(c) {
					ca.Dispatch(Disabled)
				}
			})
		} else {
			ca.Dispatch(Disabled)
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
		pl.MsgPub("发动陷阱卡{self}!", Arg{"self": ca.GetId()})
		ca.FaceUp()
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

func (ca *Card) RegisterMonster() {
	// 进入手牌
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
				pl.MsgPub("{self}需要解放{size}只怪兽作为代价", Arg{"self": ca.GetId(), "size": i})
			}
			for k := 0; k < i; {
				if t := pl.SelectMust(pl.Mzone); t != nil {
					t.Dispatch(Freedom, ca, &k)
				} else {
					pl.MsgPub("{self}召唤失败，祭品不足!", Arg{"self": ca.GetId()})
					ca.Dispatch(Destroy)
					ca.StopOnce(s)
					return
				}
			}
		},
		// 召唤
		Summon: func() {
			pl := ca.GetSummoner()
			if ca.GetSummoner().Mzone.Len() < 5 {
				ca.ToMzone()
				ca.FaceUpAttack()
				ca.SetNotCanChange()
				pl.MsgPub("{self}召唤成功!", Arg{"self": ca.GetId()})
				ca.ShowInfo()
			} else {
				pl.MsgPub("{self}召唤失败，怪兽区场地不足!", Arg{"self": ca.GetId()})
				ca.Dispatch(Destroy)
			}
		},
		// 覆盖
		Cover: func() {
			pl := ca.GetSummoner()
			if ca.GetSummoner().Mzone.Len() < 5 {
				ca.ToMzone()
				ca.FaceDownDefense()
				ca.SetNotCanChange()
				pl.Msg("{self}覆盖成功!", Arg{"self": ca.GetId()})
				ca.OnlyOnce(Flip, func() {
					ca.ShowInfo()
				})
			} else {
				pl.MsgPub("{self}覆盖失败，怪兽区场地不足!", Arg{"self": ca.GetId()})
				ca.Dispatch(Destroy)
			}
		},
	})

	// 进入怪兽区
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
			pl.MsgPub("{self}被解放!", Arg{"self": ca.GetId()})
		},
		Expression: func() {
			pl := ca.GetSummoner()
			if ca.IsCanChange() {
				if ca.IsFaceDownDefense() {
					ca.FaceUpAttack()
					ca.Dispatch(Flip)
					pl.MsgPub("{self}翻转召唤！", Arg{"self": ca.GetId()})
				} else if ca.IsFaceUpDefense() {
					ca.FaceUpAttack()
					pl.MsgPub("{self}变成攻击表示！", Arg{"self": ca.GetId()})
				} else if ca.IsFaceUpAttack() {
					ca.FaceUpDefense()
					pl.MsgPub("{self}变成防御表示！", Arg{"self": ca.GetId()})
				} else {
					pl.Msg("{self}当前状态无法改变表示形式！", Arg{"self": ca.GetId()})
					return
				}
				ca.SetNotCanChange()
			} else {
				pl.Msg("{self}不能改变表示形式！", Arg{"self": ca.GetId()})
			}
		},

		// 发出战斗宣言
		Declaration: func(tar *Card) {
			pl := ca.GetSummoner()

			if tar != nil {
				pl.MsgPub("{self}对{rival}发出战斗宣言！", Arg{"self": ca.GetId(), "rival": tar.GetId()})
			} else {
				pl.MsgPub("{self}对{rival}发出战斗宣言！", Arg{"self": ca.GetId()})
			}

			var b bool
			if tar != nil {
				b = tar.IsFaceDown()
				ca.FaceUp()
				if ca.IsInMzone() && tar.IsInMzone() {
					ca.Dispatch(DamageStep, tar)
				}

			} else {
				if ca.IsInMzone() {
					ca.Dispatch(DamageStep, tar)
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
				pl.MsgPub("{self}攻击了{rival}！", Arg{"self": ca.GetId(), "rival": tar.GetId()})
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
				pl.GetTarget().ChangeHp(-ca.GetAttack())
				pl.MsgPub("{self}直接攻击了{rival}！", Arg{"self": ca.GetId()})
			}

			ca.SetNotCanAttack()
			ca.SetNotCanChange()
		},
	})
}
