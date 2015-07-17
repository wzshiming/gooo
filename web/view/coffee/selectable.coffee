Selectable = (scene) ->
  @scene = scene
  @selec = {}
  @func = null
  @subm = {}
  return

Selectable::register = (func) ->
  @func = func
  return

Selectable::submit = (uniq, name, index) ->
  `var i`
  console.dir name
  if name instanceof Array
    i = 0
    while i != name.length
      @submit uniq, name[i], i
      i++
    for j of @subm
      b = true
      i = 0
      while i != name.length
        if j == name[i]
          b = false
          break
        i++
      if b
        Translate @subm[j].object, {
          x: 0
          y: 0
          z: 4000
        }, 1000
    return
  else if @subm[name] == null
    @subm[name] = new Element(@scene)
  if index == null
    index = 0
  info = undefined
  if name == 4
    info = '召唤'
  else if name == 8
    info = '覆盖'
  else if name == 16
    info = '使用'
  @subm[name].element.innerHTML = '<button type="submit" class="btn btn-primary btn-block" onclick="WsSelectable(' + uniq + ',' + name + ');selectable.unsubmit()"><h2>' + info + '</h2></button>'
  tmp = @scene.AllCard[uniq].object.position
  Translate @subm[name].object, {
    x: tmp.x - 60 + index % 2 * 100
    y: tmp.y + 220 + Math.floor(index / 2) * 100
    z: tmp.z + 30
  }, 1000
  return

Selectable::unsubmit = ->
  for i of @subm
    Translate @subm[i].object, {
      x: 0
      y: 0
      z: 4000
    }, 1000
  return

Selectable::click = (uniq) ->
  if @selec[uniq] == true
    delete @selec[uniq]
    @scene.AllCard[uniq].Unselectable()
  else
    @selec[uniq] = true
    @scene.AllCard[uniq].Selectable()
  if @func
    @func uniq
  return

Selectable::removeall = ->
  for i of @selec
    if @selec[i] == true
      @scene.AllCard[i].Unselectable()
      delete @selec[i]
  return
