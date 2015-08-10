YGO = (@scene) ->
  @players = []
  t = this
  WsGameRegister (e) ->
    if typeof t[e.method] == 'function'
      t[e.method] e.args
    else
      MsgErr e.method
  return

YGO::init = (args) ->
  face.SetButton "BP", ->
    WsSelectable 1,4

  face.SetButton "MP2", ->
    WsSelectable 1,5

  face.SetButton "EP", ->
    WsSelectable 1,6

  face.SetButton "Yes", ->
    WsSelectable 1,10

  face.SetButton "No", ->
    WsSelectable 1,11
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
      @players[i] = new Player(@scene, i, x, y, angle)
      face.SetHTML v.name, "#{v.hp}"
    #MsgInfo '游戏开始了'

  else
    MsgErr '初始化游戏错误'

YGO::remind = (args) ->
  c = Card::Find(args.uniq)
  if c
    c.Remind()
  else
    MsgErr "remind err"

YGO::setCardFace = (args) ->
  c = Card::Find(args.uniq)
  if c
    for own k,v of args.params
      c.SetHTML k,v


  else
    MsgErr "setCardFace err"
  #  if @scene.AllCard[args.uniq] == null
  #    new Card(@scene, args.uniq, @players[args.master])
  #  @scene.AllCard[args.uniq].master.moveCard args.uniq, args.pos
  return

YGO::moveCard = (args) ->
  @players[args.master].Join args.uniq, args.pos
  #  if @scene.AllCard[args.uniq] == null
  #    new Card(@scene, args.uniq, @players[args.master])
  #  @scene.AllCard[args.uniq].master.moveCard args.uniq, args.pos
  return

YGO::setFront = (args) ->
  c = Card::Find(args.uniq)
  if c
    c.SetFront(args.desk)
  else
    MsgErr "setFront err"
  #@scene.AllCard[args.uniq].SetFront args.desk
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
  #else
    #MsgErr "exprCard err"
  return

YGO::flagName = (args) ->
  face.SetHTML "回合数", args.round
#  t = @users[args.player]
#  @['flag0'].element.innerHTML = '<h1>回合' + args.round + '</h1>'
  return

YGO::setFace = (args) ->
  for own k,v of args
    face.SetHTML k,v

YGO::message = (args) ->
#  for own k,v of args.params
#    c = Card::Find(v)
#    if c
#      CardInfo c,(data)->
#        args.params[k] = data["中文名"] # 不支持中文

  m = args.message.format args.params

  face.Msg  m
#  @msg.element.innerHTML = '<h1>' + args.message + '</h1>'
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

YGO::touch = (args)->
  c = Card::Find(args.uniq)
  if c
    c.MoveAdd args

