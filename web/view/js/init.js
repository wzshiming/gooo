
var site = [
	1,3,
	1,5,
	19,1,
	19,3,
	19,5,
	4,4,
	7,4,
	10,4,
	13,4,
	16,4,
	4,2,
	7,2,
	10,2,
	13,2,
	16,2
]


var camera, scene, renderer;
var controls;

var objects = [];
var targets = { table: [], sphere: [], helix: [], grid: [] };

init();
animate();



function NewCard(name,x,y,z,hand) {
	var element = document.createElement( 'div' );
	element.className = 'card';
	element.style.backgroundColor = 'rgba(0,127,127,' + ( Math.random() * 0.5 + 0.25 ) + ')';
	element.style.backgroundImage = 'url(/pics/' + name + '.jpg)';

	var element2 = document.createElement( 'div' );
	element2.className = 'card';
	element2.style.backgroundImage = 'url(/textures/cover.jpg)';

	if(typeof hand == "function") {
		element.addEventListener( 'click', hand )
		element2.addEventListener( 'click', hand )	
	}
	var object = new THREE.CSS3DDoubleObject( element,element2 );
	object.position.x = x;
	object.position.y = y;
	object.position.z = z;
	scene.add( object );
	return object
}

function NewUnknownCard(x,y,z) {
	var element = document.createElement( 'div' );
	element.className = 'card';
	element.style.backgroundColor = 'rgba(0,127,127,' + ( Math.random() * 0.5 + 0.25 ) + ')';
	element.style.backgroundImage = 'url(/textures/cover.jpg)';
	var object = new THREE.CSS3DObject( element);
	object.position.x = x;
	object.position.y = y;
	object.position.z = z;
	scene.add( object );
	return object
}


function init() {

	camera = new THREE.PerspectiveCamera( 40, window.innerWidth / window.innerHeight, 1, 10000 );
	camera.position.z = 3000;

	scene = new THREE.Scene();

	// table
	
	console.dir(document);
//	for ( var i = 0; i < site.length; i += 2 ) {

//		//var obj = NewUnknownCard(0,0,0)
//		var obj = NewCard((1000 +( i / 2 )).toString()
//			,Math.random() * 4000 - 2000
//			,Math.random() * 4000 - 2000
//			,Math.random() * 4000 - 2000
//			,function ( event ) {
//				console.log("log");
//			})
//		objects.push( obj );


//		var object = new THREE.Object3D();
//		object.position.x = ( site[ i ] * 100 ) - 1260;
//		object.position.y = - ( site[ i + 1 ] * 150 ) + 600;

//		targets.table.push( object );

//	}
	for ( var i = 0; i < 20; i += 1 ) {

		//var obj = NewUnknownCard(0,0,0)
		var obj = NewCard((1000 + i ).toString()
			,Math.random() * 4000 - 2000
			,Math.random() * 4000 - 2000
			,Math.random() * 4000 - 2000
			,function ( event ) {
				console.log("log");
			})
		objects.push( obj );


//		var object = new THREE.Object3D();
//		object.position.x = ( site[ i ] * 100 ) - 1260;
//		object.position.y = - ( site[ i + 1 ] * 150 ) + 600;

//		targets.table.push( object );

	}
	//scene.remove(objects[0])

	// sphere

	var vector = new THREE.Vector3();

	for ( var i = 0, l = objects.length; i < l; i ++ ) {

		var phi = Math.acos( -1 + ( 2 * i ) / l );
		var theta = Math.sqrt( l * Math.PI ) * phi;

		var object = new THREE.Object3D();

		object.position.x = 1000 * Math.cos( theta ) * Math.sin( phi );
		object.position.y = 1000 * Math.sin( theta ) * Math.sin( phi );
		object.position.z = 1000 * Math.cos( phi );

		vector.copy( object.position ).multiplyScalar( 2 );

		object.lookAt( vector );

		targets.sphere.push( object );

	}

	// helix

	var vector = new THREE.Vector3();

	for ( var i = 0, l = objects.length; i < l; i ++ ) {

		var phi = i * 0.175 + Math.PI;

		var object = new THREE.Object3D();

		object.position.x = 1000 * Math.sin( phi );
		object.position.y = - ( i * 8 ) + 500;
		object.position.z = 1000 * Math.cos( phi );

		vector.x = object.position.x * 2;
		vector.y = object.position.y;
		vector.z = object.position.z * 2;

		object.lookAt( vector );

		targets.helix.push( object );

	}

	// grid

	for ( var i = 0; i < objects.length; i ++ ) {

		var object = new THREE.Object3D();

		object.position.x = ( ( i % 5 ) * 400 ) - 800;
		object.position.y = ( - ( Math.floor( i / 5 ) % 5 ) * 400 ) + 800;
		object.position.z = ( Math.floor( i / 25 ) ) * 1000 - 2000;

		targets.grid.push( object );

	}
	
	
	//

	renderer = new THREE.CSS3DRenderer();
	renderer.setSize( window.innerWidth, window.innerHeight );
	renderer.domElement.style.position = 'absolute';
	document.getElementById( 'container' ).appendChild( renderer.domElement );
	controls = new THREE.TrackballControls( camera, renderer.domElement );
	//controls.rotateSpeed = 0;
	//controls.minDistance = 3000;
	//controls.maxDistance = 3000;
	controls.addEventListener( 'change', render );
	controls.noRotate = true;
	controls.noZoom = true;
	controls.noPan = true;
	controls.noRoll = true;
	//

	//controls = new THREE.TrackballControls( camera, renderer.domElement );
	//controls.rotateSpeed = 0.5;
	//controls.minDistance = 500;
	//controls.maxDistance = 6000;
	//controls.addEventListener( 'change', render );

	var button = document.getElementById( 'table' );
	button.addEventListener( 'click', function ( event ) {

		Moves(objects, targets.table, 2000 );

	}, false );

	var button = document.getElementById( 'sphere' );
	button.addEventListener( 'click', function ( event ) {

		Moves(objects, targets.sphere, 1000 );

	}, false );

	var button = document.getElementById( 'helix' );
	button.addEventListener( 'click', function ( event ) {

		Moves(objects, targets.helix, 2000 );

	}, false );

	var button = document.getElementById( 'grid' );
	button.addEventListener( 'click', function ( event ) {
		Moves(objects, targets.grid, 2000 );
	}, false );
	
	var button = document.getElementById( 'test' );
	button.addEventListener( 'click', function ( event ) {
		Eval("Moves(objects, targets.grid, 2000 );")
	}, false );


	//Moves(objects, targets.table, 200 );

	var object = new THREE.Object3D();

	object.rotation.x = 1
	object.rotation.y = 2
	object.rotation.z = 3

	Move(objects[0],object,3000)
	
	var fun = function(){
		Moves(objects, targets.grid, 2000 )
		window.setTimeout("Moves(objects, targets.helix, 2000 )",3000)
		window.setTimeout("Moves(objects, targets.sphere, 2000 )",6000)
	}
	
	fun()
	var timer =  window.setInterval(fun,9000)
	//clearInterval(timer)

	window.addEventListener( 'resize', onWindowResize, false );

}

function Change(object, target, duration) {
	
	new TWEEN.Tween( object )
		.to( target, duration )
		.easing( TWEEN.Easing.Exponential.InOut )
		.start();
		
}

function Move( object, target, duration ) {

	Change(object.position,{ x: target.position.x, y: target.position.y, z: target.position.z },(Math.random() + 0.5 ) *  duration)
	
	Change(object.rotation,{ x: target.rotation.x, y: target.rotation.y, z: target.rotation.z },(Math.random() + 0.5 ) *  duration)

}

function Moves( objects, targets, duration ) {

	//TWEEN.removeAll();
	
	for ( var i = 0; i < objects.length; i ++ ) {

		Move(objects[ i ],targets[ i ],duration)

	}
	
//	new TWEEN.Tween( this )
//		.to( {}, duration * 1.5 )
//		.onUpdate( render )
//		.start();

}

function onWindowResize() {

	camera.aspect = window.innerWidth / window.innerHeight;

	camera.updateProjectionMatrix();

	renderer.setSize( window.innerWidth, window.innerHeight );

	render();

}

function animate() {

	requestAnimationFrame( animate );
	
	TWEEN.update();

	controls.update();

	render()

}

function render() {

	renderer.render( scene, camera );

}

function Pos(x,y,z) {
	var target = new THREE.Object3D();
	target.position.x = x
	target.position.y = y
	target.position.z = z
	return target
}

function Rand(){
	return Math.random() * 4000 - 2000
}

function RandPos() {
	var target = new THREE.Object3D();
	target.position.x = Rand()
	target.position.y = Rand()
	target.position.z = Rand()
	return target
}


function NewElement(pos,data) {

	var element = document.createElement( 'div' );
	if(typeof data == "string"){
		element.innerHTML = data;
	}else if(typeof data.documentElement == "object"){ 
		element.innerHTML = data.documentElement.innerHTML;
	}else {
		console.log( typeof data.documentElement )
		element = data;
	}
	element.addEventListener( 'click', function ( event ) {
		event.target.focus();
	})
	var object = new THREE.CSS3DObject(element);
	
	object.position.x = pos.position.x;
	object.position.y = pos.position.y;
	object.position.z = pos.position.z;
	object.rotation.x = pos.rotation.x;
	object.rotation.y = pos.rotation.y;
	object.rotation.z = pos.rotation.z;
	scene.add( object );
	return object;
}

var Funcs = {};

var FuncsEvent = function(path,funcname,eventfunc){
	CurrentFunc = function(){
		$.get(path,function(data,status){
			var obj = NewElement(Pos(Rand(),Rand(),-10000),data);
			Change(obj.position,{x:0,y:0,z:2600},2000);
			Funcs[funcname] = function(u){
				eventfunc(u)
				Change(obj.position,{x:0,y:0,z:4000},2000);
				window.setTimeout(function(){
					scene.remove(obj);
				},2000);
				return false;
			};
		});
	}
	CurrentFunc();
}

var CurrentFunc = function(){}


var FuncsInit = {};
FuncsInit['{{ key "Chan.Info.Rooms" }}'] = function(d){
	for(var i = 0; i != d.length; i++){
		var element = document.createElement( 'div' );
		element.className = 'control-group';
		//element.style.backgroundColor = 'rgba(0,127,127,0.5)';
		var controls = document.createElement( 'div' );
		controls.className = 'controls';
		
		element.appendChild( controls );
	
		var symbol = document.createElement( 'button' );
		symbol.className = 'btn btn-lg btn-primary btn-block';
		symbol.innerText = d[i].r + ". " + d[i].n;
		//symbol.onclick = "Funcs.chan( " + i.i + " )";
		controls.appendChild( symbol );
		var obj = NewElement(RandPos(),element)
		Change(obj.position,{x:(i%4)*80-150,y:Math.floor(i/4)*-80+50,z:2200},2000);
	}
}

FuncsInit['{{ key "Auth.Auth.LogIn" }}'] = function(d){
	FuncsEvent("/chan.html","chan",function(u){
		ws.sendMsg({{ key "Chan.Use.Chan" }},{u:u});
		ws.sendMsg({{ key "Chan.Info.Rooms" }},{});
	})	
}
FuncsInit["login"] = function(d){
	FuncsEvent("/login.html","login",function(){
		var b = document.formlogin;
		ws.sendMsg({{ key "Auth.Auth.LogIn" }},{n:b.username.value,p:b.password.value});
	})	
}
FuncsInit['login']()

