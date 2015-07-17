# ok


Player = (@scene, @index = 0, @x = 0, @y = 0, @a = 0) ->
  @decks = {}

  @inDeckSize = 0
  # 手牌
  @decks.hand = new Hand @x + 6, @y + 6, @a

  # 卡组
  @decks.deck = new Pile @x + 16, @y + 5, @a

  # 额外卡组
  @decks.extra = new Dig @x + 4, @y + 5, @a

  # 排除卡组
  @decks.removed = new Dig @x + 16, @y + 1, @a

  # 墓地
  @decks.grave = new Dig @x + 16, @y + 3, @a

  # 场地卡
  @decks.field = new Pile @x + 4, @y + 3, @a

  #怪物卡区
  @decks.mzone = new Rows @x + 6, @y + 2, @a

  #魔陷区
  @decks.szone = new Rows @x + 6, @y + 4, @a

  for k,v of @decks
    if k != "deck"
      mouse.Hint v
      mouse.Drag v, (c,rio)->
        WsSelectable c.uniq, rio
      mouse.Lay v
  return


Player::Join = (u, lay, s = null) ->
  if @decks[lay]
    c = Card::Find u
    unless c
      if lay != "deck" and @decks.deck.Length() != 0
        c = @decks.deck.Pop()
        c.SetUniq u
      else
        c = new Card(@scene)
        c.FaceDown()
        c.Attack()
        timer = null
        if lay == "deck"
          c.SetUniq  "#{ @index }_{ @inDeckSize++ }"
          c.FaceDown()
        else
          c.SetUniq u
    @decks[lay].Push c, s
  return
