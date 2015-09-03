YGO = (@scene) ->
  @players = []
  t = this
  WsGameRegister (e) ->
    if e.method
      if typeof t[e.method] == 'function'
        t[e.method] e.args
      else
        MsgErr e.method
  return

YGO::init = (args) ->
  face.SetButton "BP", ->
    WsSelectable 0, 4

  face.SetButton "MP2", ->
    WsSelectable 0, 5

  face.SetButton "EP", ->
    WsSelectable 0, 6

  face.SetButton "Yes", ->
    WsSelectable 0, 10

  face.SetButton "No", ->
    WsSelectable 0, 11

  face.SetButton "Defeat", ->
    WsSelectable 0, 666

  @users = args.users
  @index = args.index
  if @users instanceof Array
    for v,i in @users
      k = i + @index
      if k >= @users.length
        k = k - @users.length
      if k == 0
        @playerindex = i
        angle = 0
        x = 0
        y = 0
      else if k == 1
        angle = PI
        x = 0
        y = 0
      else if k == 2
        angle = PI * 0.5
        x = 0
        y = 0
      else if k == 3
        angle = PI * 1.5
        x = 0
        y = 0
      @players[i] = new Player(@scene, i, x, y, angle, v)
#MsgInfo '游戏开始了'

  else
    MsgErr '初始化游戏错误'

YGO::remind = (args) ->
  c = Card::Find(args.uniq)
  if c
    c.Remind()
  else
    MsgErr "remind err"

YGO::setPick = (args) ->
  pl = @players[args.master]
  pick = pl.decks.pick
  while pick.Length() != 0
    c = pick.Pop()
    if c.hold
      c.Update()
    else
      pl.Join c.uniq, "deck"
  for v in args.uniqs
    pl.Join v, "pick"


YGO::setCardFace = (args) ->
  c = Card::Find(args.uniq)
  if c
    for own k,v of args.params
      c.SetHTML k, v
  else
    MsgErr "setCardFace err"
  return

YGO::moveCard = (args) ->
  @players[args.master].Join args.uniq, args.pos
  return

YGO::setFront = (args) ->
  c = Card::Find(args.uniq)
  if c
    c.SetFront(args.desk)
  else
    MsgErr "setFront err"

  return

YGO::exprCard = (args) ->
  c = Card::Find(args.uniq)
  if c
    if (args.expr & 1 << 30) != 0
      c.FaceUp()
    else if (args.expr & 1 << 29) != 0
      c.FaceDown()
    if (args.expr & 1 << 28) != 0
      c.Attack()
    else if (args.expr & 1 << 27) != 0
      c.Defense()
  else
#MsgErr "exprCard err"
  return

YGO::flagName = (args) ->
  face.SetHTML "回合数", args.round
  return

YGO::setFace = (args) ->
  for own k,v of args
    face.SetHTML k, v

YGO::message = (args) ->
  for own k,v of args.params
    if v == @users[@playerindex].name
      args.params[k] = "您"
    c = Card::Find(v)
    if c
      CardInfo c, (data)->
        args.params[k] = " #{data["type"]} 「#{data["name"]}」 "

  m = args.message.format args.params
  face.Msg m
  return

YGO::flagStep = (args) ->
  t = this
  if t.stepInterval
    window.clearInterval t.stepInterval

  p = if args.step == 1
    "DP"
  else if args.step == 2
    "SP"
  else if args.step == 3
    "MP1"
  else if args.step == 4
    "BP"
  else if args.step == 5
    "MP2"
  else if args.step == 6
    "EP"
  else if args.step == 7
    "Chain"
  else
    "现在什么阶段?"
  face.SetHTML "阶段", p
  if typeof args.wait == 'number'
    ti = args.wait / 1000000000
    face.SetHTML "剩余时间", "#{ti}"
    t.stepInterval = window.setInterval((->
      ti--
      face.SetHTML "剩余时间", "#{ti}"
      if ti <= 0
        ti = 0
        window.clearInterval t.stepInterval
      return
    ), 1000)
  return
YGO::changeHp = (args)->
  c = Card::Find(args.uniq)
  if c
    c.SetHTML "HP", "#{args.hp}"

YGO::setPortrait = (args)->
  c = Card::Find(args.uniq)
  if c
    c.SetPortrait args.desk

YGO::touch = (args)->
  c = Card::Find(args.uniq)
  if c
    c.MoveAdd args

YGO::over = ()->
  ExitPage()
  MsgInfo "游戏结束"

YGO::trigg = (args) ->
  if @trigger
    for v in @trigger
      c = Card::Find v
      if c
        c.SetClass "card"
  @trigger = args.uniqs
  if @trigger
    for v in @trigger
      c = Card::Find v
      if c
        c.SetClass "card0"
