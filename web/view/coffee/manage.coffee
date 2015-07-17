
getCardsSize = (deck) ->
  main = {}
  for v,k in deck
    if !v.id
      continue
    main[v.id] = main[v.id] or 0
    main[v.id] += 1
  mainr = []
  for k,v of main
    mainr.push
      id: parseInt(k)
      size: v
  mainr

Manage = (@scene) ->
  @hold = null
  @layout = {}
  @layout.side = new Hand 0 , -6 , 0, 10
  @layout.side.seat = "side"
  @layout.extra = new Hand 0 , -4 , 0, 10
  @layout.extra.seat = "extra"
  @layout.main = new Vast 0 , -2 , 0, 10, 15
  @layout.main.seat = "main"
  @layout.result = new Vast 13 , -6 , 0, 6, 10
  @layout.result.seat = "result"
  for k,v of @layout
    mouse.Hint v
    mouse.Alone v


  t = this

  f = (c,rio) ->
    if rio == 3
      t.layout.side.Push c
      c.seat = t.layout.side.seat
    else if rio == 1
      c.Remove()

  mouse.Drag @layout.main, f

  mouse.Drag @layout.extra, f

  mouse.Drag @layout.side, (c,rio) ->
    if rio == 7
      $.get("/cards/json/#{c.id}.json",((data,status)->
        d = JSON.parse data
        ty = d["卡片种类"]
        if ty == '融合怪兽' or ty == '仪式怪兽' or ty == '同调怪兽' or ty == 'XYZ怪兽'
          t.layout.extra.Push c
          c.seat = t.layout.extra.seat
        else
          t.layout.main.Push c
          c.seat = t.layout.main.seat
      ))
    else if rio == 1
      c.Remove()
  mouse.Drag @layout.result, (c,rio) ->
    if rio == 5
      $.get("/cards/json/#{c.id}.json",((data,status)->
        d = JSON.parse data
        ty = d["卡片种类"]
        if ty == '融合怪兽' or ty == '仪式怪兽' or ty == '同调怪兽' or ty == 'XYZ怪兽'
          t.AddCard t.layout.extra,c.id ,1,t.layout.extra.seat
        else
          t.AddCard t.layout.main,c.id ,1,t.layout.main.seat
      ))
  @decks = {}
  @deck = {}
  t = this
  WsGameGetDeck (d) ->
    log d
    for k, v of  d
      log v
      t.AddCards k,v
    return
  return

Manage::AddCards = (k,v) ->
  t = this
  @deck[v.name] = v
  @decks[v.name] = new Pile 18 ,  -4 + k *2, 0
  @decks[v.name] .addEventListener "click",((event) ->
    c = event.toElement.hold
    log c
    if t.hold and !t.use
      t.layout.side.MoveTo t.hold
      t.layout.extra.MoveTo t.hold
      t.layout.main.MoveTo t.hold
      t.hold = null
    if t.hold == null
      t.hold = c.hold
      t.use = true
      query = document.createElement('input')
      query.id = "query"
      query.type = "text"
      query.addEventListener "keyup", ((event) ->
        WsCardFind { query: event.srcElement.value }, (data) ->
          t.layout.result.Clear()
          if data
            for v in data
              t.AddCard t.layout.result,v,1,"result"
      ),false
      face.SetHTML "搜索",query
      face.SetHTML "卡组",v.name
      face.SetButton "保存", (event) ->
        WsGameSetDeck
          main: getCardsSize t.layout.main
          extra: getCardsSize t.layout.extra
          side: getCardsSize t.layout.side
          name: v.name

        t.layout.side.MoveTo t.hold
        t.layout.extra.MoveTo t.hold
        t.layout.main.MoveTo t.hold
        t.hold = null
        face.SetButton "保存"
        face.SetHTML "卡组"
        face.SetHTML "搜索"
        t.layout.result.Clear()
      while t.hold.Length() != 0
        c = t.hold.Pop()
        t.layout[c.seat].Push c

    return
  ),false
  if v.side
    for x in v.side
      @AddCard @decks[v.name],x.id,x.size,"side"
  if v.extra
    for x in v.extra
      @AddCard @decks[v.name],x.id,x.size,"extra"
  if v.main
    for x in v.main
      @AddCard @decks[v.name],x.id,x.size,"main"
  return


Manage::AddCard = (deck,id,size = 1,seat) ->
  if deck
    t = this
    for x in [0...size]
      c = new Card(@scene)
      c.FaceUp()
      c.Attack()
      c.SetFront id
      c.seat = seat
      deck.Push c
  return