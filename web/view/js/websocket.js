var ws = new WebSocket("ws://127.0.0.1:3006/");
ws.onopen = function(e){
    console.log("onopen");
    console.dir(e);
};
ws.onmessage = function(e){
    console.log("onmessage",e.data);
    console.dir(e);
};
ws.onclose = function(e){
    console.log("onclose");
    console.dir(e);
};
ws.onerror = function(e){
    console.log("onerror");
    console.dir(e);
};
ws.i = 0
ws.sendMsg = function(c1,c2,c3,msg) {
    if (this.readyState == this.OPEN) { 
        s = msg.length + 4;
        this.i =  this.i + 1
        this.send(String.fromCharCode((s & 0xFF00) >> 8,s & 0xFF,this.i & 0xFF,c1 & 0xFF,c2 & 0xFF,c3 & 0xFF) + msg);  
    } 
}