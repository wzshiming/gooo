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
  @users = args.users
  @index = args.index
  console.dir args
  if @users instanceof Array
    for v,i in @users
      k = i + @index
      if k >= @users.length
        k = k - @users.length
      if k == 0
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
    MsgInfo '游戏开始'
    console.dir @players
  else
    MsgErr '初始化游戏错误'

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
  else
    MsgErr "exprCard err"
  return

YGO::flagName = (args) ->
  face.SetHTML "回合数", "#{args.round}"
#  t = @users[args.player]
#  @['flag0'].element.innerHTML = '<h1>回合' + args.round + '</h1>'
  return

YGO::message = (args) ->
  face.SetHTML "最新消息", args.message
#  MsgInfo args.message
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

  face.SetHTML "阶段", p
  if typeof args.wait == 'number'
    ti = args.wait / 1000000000
    face.SetHTML "剩余时间", "#{ti}"
    t.stepInterval = window.setInterval((->
      ti--
      face.SetHTML "剩余时间", "#{ti}"
      if ti == 0
        window.clearInterval t.stepInterval
      return
    ), 1000)
  return

YGO::touch = (args)->
  c = Card::Find(args.uniq)
  if c
    c.MoveAdd args

