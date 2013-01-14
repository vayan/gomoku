var ws;
var turn = "black";
var score = "Black : 0 | White : 0";
var me = "You are : unknown";
var canplay = false;
var color = "black";

if(ws != null) {
  ws.close();
  ws = null;
}

var host = window.location.hostname;
ws = new WebSocket("ws://" + host + ":1112/ws");

//init page
$("#turn").text("IDK");
$("#score").text(score);
$("#me").text(me);

function change_turn() {

}

ws.onopen = function() {
  console.log("open");
  ws.send("getturn");
  ws.send("getme");
};

ws.onmessage = function(e) {
  console.log("receive : '" + e.data + "'");

  if(e.data != "error") {
    var data = e.data.split(' ');

    //he won
    if(data[0] == "LOSE") { 
      $("#game").slideUp();
      $("#menu").show();
      $("#victory").text("YOU LOSE").show(300);
    } 

    //you won
    if(data[0] == "WIN") { 
      ws.send("reset");
      $("#game").slideUp();
      $("#menu").show();
      $("#victory").text("YOU WIN").show(300);
    } 
    //stone captured
    else if(data[0] == "REM") {
      $('.pos' + data[1] + 'y' + data[2]).removeClass("bgblack");
      $('.pos' + data[1] + 'y' + data[2]).removeClass("bgwhite");
      $('.pos' + data[3] + 'y' + data[4]).removeClass("bgblack");
      $('.pos' + data[3] + 'y' + data[4]).removeClass("bgwhite");
    } 
    // update score
    else if(data[0] == "score") {
      $("#score").text(data[1]);
    } 
    //who I am WHO ?!
    // else if(data[0] == "me") {
    //   if (data[1] == " You are OBS") {
    //     $("#reset").attr("disabled", "disabled");
    //   }
    //   $("#me").text(data[1]);
    // } 
    //new stone
    else if (data[0] == "YOURTURN") {
      $("#turn").text("YOUR TURN");
    }
    else if(data[0] == "ADD") {
      $('.boardclic td div').removeClass("newstone");
      console.log("ADD pierre in X "+data[1]+" Y "+data[2]);
      if (turn == color) {
        $("#turn").text("NOT YOU");
      }
      $('.pos' + data[1] + 'y' + data[2]).addClass("bg" + turn);
      $('.pos' + data[1] + 'y' + data[2]).addClass("newstone");
      if(turn == "black") {
        turn = "white";
      } else {
        turn = "black";
      }
      //$("#turn").text("Turn : " + turn);
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
  ws.send("PLAY "+coord);
  ws.send("getscore");
});