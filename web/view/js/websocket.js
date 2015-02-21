var ws = new WebSocket("ws://127.0.0.1:3006/");
ws.onopen = function(e){
};
ws.onmessage = function(e){
    d = e.data
    this.event( d.charCodeAt(2), d.charCodeAt(3), d.charCodeAt(4), d.charCodeAt(5), $.parseJSON(d.slice(6, d.length)) )
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
        if(!isString(msg)) {
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

function isString(object){
    return object && typeof object==='string'
}