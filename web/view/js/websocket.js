//var ws = new WebSocket("ws://127.0.0.1:3006/");
var ws = new WebSocket("ws://127.0.0.1:3006/conn");
ws.onopen = function(e){
	console.log("onopen");
};
ws.onmessage = function(e){
    d = e.data
	msg = JSON.parse(d.slice(4, d.length))

    this.event( d.charCodeAt(0), d.charCodeAt(1), d.charCodeAt(2), d.charCodeAt(3), msg )
};
ws.onclose = function(e){
    console.log("onclose");
    console.dir(e);
};
ws.onerror = function(e){
};
ws.i = 0
ws.sendMsg = function(c1,c2,c3,msg) {
    if (this.readyState == this.OPEN) { 
        if(typeof msg.exec != "string") {
            msg = JSON.stringify(msg)
        }
        this.i += 1;
        this.send(String.fromCharCode(this.i & 0xFF, c1 & 0xFF, c2 & 0xFF, c3 & 0xFF) + msg);  
    } 
}

ws.event = function(i, c1, c2, c3, msg){
	var n = c1 + "," + c2 + "," +c3
	var f = FuncsInit[n]
	if((typeof msg.e == "string") && msg.e != ""){
		CurrentFunc();
	}else if(typeof f == "function"){
		f(msg)
	}
    console.log(i,msg);
};


var Eval = function(code){ 
	if(!!(window.attachEvent && !window.opera)){ 
		execScript(code); 
	}else{ 
		window.eval(code); 
	} 
} 