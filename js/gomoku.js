var ws;
var turn = "black";

if (ws != null) {
	ws.close();
	ws = null;
}

ws = new WebSocket("ws://localhost:1112/ws");

ws.onopen = function () {
	console.log("open");
};

ws.onmessage = function (e) {
	console.log("receive :"+e.data);

	if (e.data != "error") {
		var data = e.data.split(',');

		if (data[2] == "remove"){
			$('.pos'+data[0]+'y'+data[1]).removeClass("bgblack");
			$('.pos'+data[0]+'y'+data[1]).removeClass("bgwhite");
		} else {
			$('.pos'+data[0]+'y'+data[1]).addClass("bg"+data[2]);
			if (data[2] == "black") {
				turn = "white";
			} else {
				turn = "black";
			}
			$("#turn").text(turn+" Turn");
		}
	}
};

ws.onclose = function (e) {
	console.log("closed");
};

$("#turn").text(turn+" Turn");

$("#reset").click(function() {	
	ws.send("reset");
	location.reload();
});

$(".boardclic td").click(function() {
	var stone = $(this).find(".stone");
	var coord = stone.text();
	
	ws.send(coord);
});
