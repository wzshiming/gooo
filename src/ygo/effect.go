package ygo

func (ca *Card) EffectplMonsterManipulate(c *Card) {
	if !ca.IsMonster() {
		return
	}
	pl := c.GetSummoner()
	ca.isValid = true
	ca.SetSummoner(pl)
	if ca.IsInMzone() {
		ca.ToMzone()
	} else {
		ca.Dispatch(SummonSpecial)
	}

	c.AddEvent(Disabled, func() {
		ca.SetSummoner(ca.GetOwner())
		if ca.IsInMzone() {
			ca.ToMzone()
		}
	})
}
