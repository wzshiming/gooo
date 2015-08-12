

Face =  ->
  @face = document.createElement "Face"
  container.appendChild @face
  @face.className = "face"

  @img = document.createElement('img')
  @face.appendChild @img

  @text = document.createElement('p')
  @face.appendChild @text

#  @phases = document.createElement "phases"
#  container.appendChild @phases
#  @phases.className = "phases"

  @infos = {}
  @info = document.createElement('div')
  @info.className = "phases"
  container.appendChild @info

  @chat = document.createElement('div')
  @chat.className = "chat"
  container.appendChild @chat
  return

Face::Msg = (msg) ->
  b = document.createElement('p')
  b.innerText = CurentTime() + "\n" + msg + "\n"
  @chat.appendChild b
  @chat.scrollTop = @chat.scrollHeight
  return

Face::SetInput = (name = null, u = null) ->
  if u
    b = document.createElement('input')
    b.id = name
    b.type = "text"
    b.className = "form-control"
    b.name = name
    b.placeholder = name
    @SetHTML name,b
  else
    @SetHTML name
  return b


Face::SetButton = (name = null, func = null) ->
  if func
    b = document.createElement('a')
    b.className = "btn btn-primary"
    b.innerText = name
    b.onclick = func
    @SetHTML name, b
  else
    @SetHTML name
  return b


Face::SetHTML = (name = null, value = null) ->
  unless name
    #log @infos
    for own k,v of @infos
      @info.removeChild @infos[k]
      delete @infos[k]
    @infos = {}
    return
  if value
    if typeof value == 'object'
      if @infos[name]
        @info.removeChild @infos[name]
        delete @infos[name]
      @info.appendChild value
      @infos[name] = value
    else
      unless @infos[name]
        @infos[name] = document.createElement('p')
        @info.appendChild @infos[name]
      @infos[name].innerHTML = "#{name}: #{value}"
  else if @infos[name]
    @info.removeChild @infos[name]
    delete @infos[name]
  return

Face::ShowCard = (c) ->
  if c
    @img.src = c.img.src
    @img.style.display = ""
    @text.style.display = ""
    t = this
    #@text.innerText = c.img.src.innerText
    #console.dir c
    CardInfo c,(data)->
      str = ""
      for own k,v of JSON.parse data
        str += "#{k}: #{v}<br>"
#      if c.img.src.innerText
#        str += "场上状态<br>"
#        str += c.img.src.innerText
      t.text.innerHTML = str

  else
    #@img.src = " "
    @img.style.display = "none"
    #@text.innerText = ""
    @text.style.display = "none"
  return
