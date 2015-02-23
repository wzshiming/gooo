//var ws = new WebSocket("ws://127.0.0.1:3006/");
var ws = new WebSocket("ws://127.0.0.1:3006/conn/");
ws.onopen = function(e){
};
ws.onmessage = function(e){
    d = e.data
	msg = JSON.parse(d.slice(6, d.length))
	if((typeof msg.exec == "string") && msg.exec != ""){
		console.log(d);
		Eval(msg.exec)
		delete msg.exec 
	}
    this.event( d.charCodeAt(2), d.charCodeAt(3), d.charCodeAt(4), d.charCodeAt(5), msg )
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
        s = msg.length + 4;
        this.i += 1;
        this.send(String.fromCharCode((s & 0xFF00) >> 8, s & 0xFF, this.i & 0xFF, c1 & 0xFF, c2 & 0xFF, c3 & 0xFF) + msg);  
    } 
}

ws.event = function(i, c1, c2, c3, msg){
    console.log(i,msg);
};


var Eval = function(code){ 
	if(!!(window.attachEvent && !window.opera)){ 
		execScript(code); 
	}else{ 
		window.eval(code); 
	} 
} 