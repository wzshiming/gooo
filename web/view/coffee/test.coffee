Test = (@scene) ->
  cs = new Player @scene
  console.dir cs
  for i in [1..60]
    cs.Join i + 600, "deck", 10

  for i in [1..15]
    cs.Join i + 500, "extra"

  for i in [1..15]
    cs.Join i + 400, "removed"

  for i in [1..15]
    cs.Join i + 300, "grave"

  for i in [1..7]
    cs.Join i + 200, "hand"

  for i in [1..5]
    cs.Join i + 100, "mzone"

  #c1.SetText()
#  console.dir c1
#  console.dir cs
#  console.dir cs2
  #cs.Clear()
  #cs.push c2
#  console.dir cs
#  console.dir cs.length