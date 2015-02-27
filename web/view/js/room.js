

(function(){
	var data = [
	{{ range $k, $v := . }}
		"<dir class=\"roon\"><dir class=\"number\">{{ $v.RoomId }}</dir></dir>",
	{{ end }}
		""
	];
	
	for( var i = 0; i != data.length; i++){
		d = data[i];
		if(d != ""){
			var obj = NewElement(Math.random() * 4000 - 2000,Math.random() * 4000 - 2000,-10000,d);
			Change(obj.position,{x:(i%5)*80-150,y:Math.floor(i/5)*-80+50,z:2200},2000);
		}
	}
	
})()