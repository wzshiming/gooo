container = null
face = null
mouse = null

InitGlobal = ->
  container = document.getElementById "ygo"
  face = new Face()
  mouse = new Mouse()

log = (x)->
  console.dir x

gui =  (id, classs) ->
  camera = undefined
  scene = undefined
  renderer = undefined
  controls = undefined

  onWindowResize = ->
    camera.aspect = window.innerWidth / window.innerHeight
    camera.updateProjectionMatrix()
    renderer.setSize window.innerWidth, window.innerHeight
    render()
    return

  animate = ->
    requestAnimationFrame animate
    TWEEN.update()
    controls.update()
    render()
    return

  render = ->
    renderer.render scene, camera
    return

  do ->
    InitGlobal()
    camera = new (THREE.PerspectiveCamera)(40, window.innerWidth / window.innerHeight, 1, 10000)
    camera.position.z = 3000
    #camera.position.y = -2500;
    scene = new (THREE.Scene)
    renderer = new (THREE.CSS3DRenderer)
    renderer.setSize window.innerWidth, window.innerHeight
    renderer.domElement.style.position = 'absolute'
    @game = document.createElement "id"
    container.appendChild @game
    @game.appendChild renderer.domElement
    controls = new (THREE.TrackballControls)(camera, renderer.domElement)
    #    controls.rotateSpeed = 0.5;
    #    controls.minDistance = 500;
    #    controls.maxDistance = 6000;
    controls.addEventListener 'change', render
    controls.noRotate = true
    controls.noZoom = true
    controls.noPan = true
    controls.noRoll = true
    window.addEventListener 'resize', onWindowResize, false
    new classs(scene)
    #new YGO(scene)
    animate()
    return
