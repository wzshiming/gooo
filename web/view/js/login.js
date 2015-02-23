
var loginObj

var loginFunc


(function(){
	$.get("login.html",function(data,status){
	    function NewLogin(x,y,z) {
			var element = document.createElement( 'div' );
			element.innerHTML=(data);
			element.addEventListener( 'click', function ( event ) {
				event.target.focus()
				console.dir(event);
			})
			var object = new THREE.CSS3DObject(element);
			object.position.x = x;
			object.position.y = y;
			object.position.z = z;
			scene.add( object );
			return object
		}
		var target = new THREE.Object3D();
		target.position.x = 0
		target.position.y = 0
		target.position.z = 2600
		loginObj = NewLogin(0,0,-50000)
		Move(loginObj,target,2000)
		loginFunc = function(){
			var b = document.formlogin
			ws.sendMsg({{ key "Auth.Auth.LogIn" }},{n:b.username.value,p:b.password.value})
			return false;
		}
		
	});
})()
