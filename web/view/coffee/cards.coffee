# ok


# ������
Cards = ->
  return

Cards:: = []


#ѹ��
Cards::Push = (c) ->
  if c
    c.Placed()
    @push c
    c.SetHold this
  return

#����
Cards::Pop = ->
  c = @pop()
  if c
    c.SetHold()
  c

Cards::Insert = (i,c) ->
  @splice(i, 0, c)
  c.SetHold this
  return
#�Ƴ�
Cards::Remove = (i) ->
  c = @splice(i, 1)[0]
  c.SetHold()
  c

#����
Cards::Placed = (u) ->
  i = @Index(u)
  if i != -1
    @Remove i
  i

#��ȡ����
Cards::Index = (u) ->
  if typeof u == "object"
    for v, k in this
      if v == u
        return k
  else
    for v, k in this
      if v.uniq == u
        return k
  return -1

#���
Cards::Clear = ->
  while @Length() != 0
    @Pop().Remove()
  return

#ȫ���ƶ���
Cards::MoveTo = (to) ->
  while @Length() != 0
    to.Push @Pop()
  return

Cards::Length = ->
  @length