var ws;
var turn = "Turn : unknown";
var score = "Black : 0 | White : 0"
var me = "You are : unknown"

if(ws != null) {
  ws.close();
  ws = null;
}

ws = new WebSocket("ws://localhost:1112/ws");

ws.onopen = function() {
  console.log("open");
  ws.send("getturn");
  ws.send("getme");
};

ws.onmessage = function(e) {
  console.log("receive :" + e.data);


  if(e.data != "error") {
    var data = e.data.split(',');

    if(data[0] == "win") {
      ws.send("reset");
      $("#game").slideUp(300);
      $("#menu").show(300);

      $("#victory").text(data[1] + " is the winner !!!").show(300);
    } 

    else if(data[2] == "pow") {
      $('.pos' + data[0] + 'y' + data[1]).removeClass("bgblack");
      $('.pos' + data[0] + 'y' + data[1]).removeClass("bgwhite");
    } 

    else if(data[0] == "turn") {
      turn = data[1]
      $("#turn").text("Turn : " + turn);
    } 

    else if(data[0] == "score") {
      $("#score").text(data[1]);
    } 

    else if(data[0] == "me") {
      $("#me").text(data[1]);
    } 

    else {
      $('.pos' + data[0] + 'y' + data[1]).addClass("bg" + data[2]);
      if(data[2] == "black") {
        turn = "white";
      } else {
        turn = "black";
      }
      $("#turn").text("Turn : " + turn);
    }

  }
};

ws.onclose = function(e) {
  console.log("closed");
};

$("#turn").text(turn);
$("#score").text(score);
$("#me").text(me);


$(".selectpvp").click(function() {
  $("#victory").hide();
  $("#menu").slideUp(300);
  $("#game").show(300);
});

$("#reset").click(function() {
  ws.send("reset");
  location.reload();
});

$(".boardclic td").click(function() {
  var stone = $(this).find(".stone");
  var coord = stone.text();
  ws.send(coord);
  ws.send("getscore");
});