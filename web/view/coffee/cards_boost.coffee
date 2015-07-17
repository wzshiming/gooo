# ok


gridX = (x) ->
  x * 130 - 1300

gridY = (y) ->
  -(y * 130)

cards_boost = (@pos, @s = 100)->
  return

cards_boost:: = new Cards()

cards_boost::Update = (s = @s) ->
  if @queue.length != 0
    @queue.unshift s
    @queue.pop()
    return

  t = this
  setTimeout (->
    t.Flash t.queue.pop()
    e = t.queue.pop()
    if e
      t.Update e
  ), s

  return

cards_boost::Flash = (s = @s) ->
  for v, k in this
      v.MoveTo @pos(k), s
  return

cards_boost::Placed = (u, s = null) ->
  i = Cards::Placed.call this, u
  @Update s
  i

cards_boost::Insert = (i,c, s = null) ->
  if c
    Cards::Insert.call this, i, c
    for v,k in @event
      c.addEventListener v...
    @Update s
  return

cards_boost::Remove = (i, s = null) ->
  c = Cards::Remove.call this, i
  if c
    for v,k in @event
      c.removeEventListener v...
  @Update s
  c

cards_boost::Push = (c, s = null)->
  if c
    Cards::Push.call this, c
    for v,k in @event
      c.addEventListener v...
    @Update s
  return

cards_boost::Pop = (s = null) ->
  c = Cards::Pop.call this
  if c
    for v,k in @event
      c.removeEventListener v...
  @Update s
  c

cards_boost::addEventListener = (s,f,b) ->
  l = [s,f,b]
  @event.push l
  for v,k in this
    v.addEventListener l...
  return



# ��ʾ��ѡ��
Pick =  (@x = 0, @y = 0, @a = 0)->
  @queue = []
  @event = []
  @rotation = (new (THREE.Matrix4)).makeRotationZ(@a)
  return

Pick:: = new cards_boost((i)->
  object = new (THREE.Object3D)
  object.matrix.makeTranslation gridX(@x + i), gridY(@y), 1000 + i
  object.applyMatrix @rotation
  object.position
)

Pick::Push = (c, s = null)->
  if c
    @push c
    @Update s
  return

Pick::Pop = ->
  r = @pop this
  @Update()
  r

# ����һ���
Pile =  (@x = 0, @y = 0, @a = 0)->
  @queue = []
  @event = []
  @rotation = (new (THREE.Matrix4)).makeRotationZ(@a)
  return

Pile:: = new cards_boost((i)->
  object = new (THREE.Object3D)
  object.matrix.makeTranslation gridX(@x), gridY(@y), i + 1
  object.applyMatrix @rotation
  object.position
)


Dig =  (@x = 0, @y = 0, @a = 0, @s = 100)->
  @queue = []
  @event = []
  @rotation = (new (THREE.Matrix4)).makeRotationZ(@a)
  @p = 0
  t = this
  @addEventListener "mouseover", (->
    if t.p == 1
      return
    t.p = 1
    t.Update @s
    return
  ), false
  @addEventListener "mouseout", (->
    if t.p == 0
      return
    t.p = 0
    t.Update @s
    return
  ), false
  return

Dig:: = new cards_boost((i)->
  object = new (THREE.Object3D)
  if @p == 0
    object.matrix.makeTranslation gridX(@x), gridY(@y), i + 1
    object.applyMatrix @rotation
    object.position
  else
    object.matrix.makeTranslation gridX(@x)  + ( i - @Length()/2) * 5 , gridY(@y), 50 + i * 2
    z = 1
    if @Length() > 40
      z = 1 - (@Length() - 40) / @Length()
    object.applyMatrix (new (THREE.Matrix4)).makeRotationZ(@a - (i - @Length() / 2) / 180 * PI * 9 * z)
    object.position#.multiplyScalar 0.9
)


# λ�� ����
flex = (x,y,a,r,i,l)->
  object = new (THREE.Object3D)
  z = 1
  if l > r
    z = 1 - (l - r) / l
  object.matrix.makeTranslation gridX(x + r / 2 + 0.5 +  (i - l / 2) * 1.2 * z ), gridY(y), i + 1
  object.applyMatrix a
  object.position
#���ϵ�
Hand =  (@x = 0, @y = 0, @a = 0, @b = 8, @c = 10)->
  @queue = []
  @event = []
  @rotation = (new (THREE.Matrix4)).makeRotationZ(@a)
  return

Hand:: = new cards_boost((i)->
  flex(@x,@y,@rotation,@b,i,@Length())
)

Vast =  (@x = 0, @y = 0, @a = 0,@b = 8,@c=10)->
  @queue = []
  @event = []
  @rotation = (new (THREE.Matrix4)).makeRotationZ(@a)
  return

Vast:: = new cards_boost((i)->
  l = parseInt(i / @c)
  m = i % @c

  if l == parseInt(@Length() / @c)
    z = @Length() % @c
    flex(@x,@y + l*1.5 ,@rotation,@b,m,z)
  else
    flex(@x,@y + l*1.5 ,@rotation,@b,m,@c)
)

#�����̿���
Rows =  (@x = 0, @y = 0, @a = 0)->
  @queue = []
  @event = []
  @rotation = (new (THREE.Matrix4)).makeRotationZ(@a)
  return

Rows:: = new cards_boost((i)->
  object = new (THREE.Object3D)
  object.matrix.makeTranslation gridX(@x + i * 2), gridY(@y), i + 1
  object.applyMatrix @rotation
  object.position
)