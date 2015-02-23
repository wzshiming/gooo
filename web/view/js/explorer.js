
(function() {
	var explorer = window.navigator.userAgent ;
	//ie 
	if (explorer.indexOf("MSIE") >= 0) {
		alert("ie");
	}
	//firefox 
	else if (explorer.indexOf("Firefox") >= 0) {
		alert("Firefox");
	}
	//Chrome
	else if(explorer.indexOf("Chrome") >= 0){
		alert("Chrome");
	}
	//Opera
	else if(explorer.indexOf("Opera") >= 0){
		alert("Opera");
	}
	//Safari
	else if(explorer.indexOf("Safari") >= 0){
		alert("Safari");
	}
})()