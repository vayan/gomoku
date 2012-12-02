var ws;
var turn = "Turn : unknown";
var score = "Black : 0 | White : 0"
var me = "You are : unknown"

if(ws != null) {
  ws.close();
  ws = null;
}

var host = window.location.hostname;
ws = new WebSocket("ws://" + host + ":1112/ws");

//init page
$("#turn").text(turn);
$("#score").text(score);
$("#me").text(me);

ws.onopen = function() {
  console.log("open");
  ws.send("getturn");
  ws.send("getme");
};

ws.onmessage = function(e) {
  console.log("receive :" + e.data);


  if(e.data != "error") {
    var data = e.data.split(',');

    //someone won
    if(data[0] == "win") { 
      ws.send("reset");
      $("#game").slideUp();
      $("#menu").show();
      $("#victory").text(data[1] + " is the winner !!!").show(300);
    } 
    //stone captured
    else if(data[2] == "pow") {
      $('.pos' + data[0] + 'y' + data[1]).removeClass("bgblack");
      $('.pos' + data[0] + 'y' + data[1]).removeClass("bgwhite");
    } 
    //update turn
    else if(data[0] == "turn") {
      turn = data[1]
      $("#turn").text("Turn : " + turn);
    } 
    // update score
    else if(data[0] == "score") {
      $("#score").text(data[1]);
    } 
    //who I am WHO ?!
    else if(data[0] == "me") {
      if (data[1] == " You are OBS") {
        $("#reset").attr("disabled", "disabled");
      }
      $("#me").text(data[1]);
    } 
    //new stone
    else {
      $('.boardclic td div').removeClass("newstone");
      $('.pos' + data[0] + 'y' + data[1]).addClass("bg" + data[2]);
      $('.pos' + data[0] + 'y' + data[1]).addClass("newstone");
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

//click gamemode
$(".selectpvp").click(function() {
  $("#victory").hide();
  $("#menu").slideUp();
  $("#game").show();
});


//reset all game
$("#reset").click(function() {
  ws.send("reset");
  location.reload();
});

//new move
$(".boardclic td").click(function() {
  var stone = $(this).find(".stone");
  var coord = stone.text();
  ws.send(coord);
  ws.send("getscore");
});