

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
  return

Face::SetButton = (name = null, func = null) ->
  if func
    b = document.createElement('a')
    b.className = "btn btn-primary"
    b.innerText = name
    b.onclick = func
    @SetHTML name, b
  else
    @SetHTML name
  return


Face::SetHTML = (name = null, value = null) ->
  unless name
    for own k of  @infos
      @info.removeChild @infos[k]
      delete @infos[k]
    @infos = {}
    return
  if value
    if typeof value != 'string'
      if @infos[name]
        @info.removeChild @infos[name]
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
    $.get("/cards/json/#{c.id}.json",((data,status)->
      str = ""
      for own k,v of JSON.parse data
        str += "#{k}: #{v}<br>"
      if c.img.src.innerText
        str += "场上状态<br>"
        str += c.img.src.innerText
      t.text.innerHTML = str
    ))
  else
    #@img.src = " "
    @img.style.display = "none"
    #@text.innerText = ""
    @text.style.display = "none"
  return
