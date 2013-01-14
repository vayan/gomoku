var ws;
var mode;
var turn = "null";
var score = "Black : 0 | White : 0";
var me = "You are : unknown";
var color = "null";

if(ws != null) {
  ws.close();
  ws = null;
}

var host = window.location.hostname;

//init page
$("#turn").text("IDK");
$("#score").text(score);
$("#me").text(me);

function ConnectWS() {

  ws = new WebSocket("ws://" + host + ":1112/ws");

  ws.onopen = function() {
    console.log("Connected to server");
    ws.send("MODE " + mode);
    if(color == "null") {
      ws.send("GETCOLOR");
    }
    ws.send("GETTURN");
  };

  ws.onmessage = function(e) {
    console.log("receive : '" + e.data + "'");

    if(e.data != "error") {
      var data = e.data.split(' ');

      if(data[0] == "COLOR") {
        color = data[1];
      }

      if(data[0] == "COLOR") {
        color = data[1];
      }

      if(data[0] == "TURN") {
        turn = data[1];
      }
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
      // // update score
      // else if(data[0] == "score") {
      //   $("#score").text(data[1]);
      // } 
      else if(data[0] == "YOURTURN") {
        $("#turn").text("YOUR TURN");
      } else if(data[0] == "ADD") {
        $('.boardclic td div').removeClass("newstone");
        if(turn == color) {
          $("#turn").text("NOT YOU");
        }
        $('.pos' + data[1] + 'y' + data[2]).addClass("bg" + turn);
        $('.pos' + data[1] + 'y' + data[2]).addClass("newstone");
        if(turn == "black") {
          turn = "white";
        } else {
          turn = "black";
        }
      }

    }
  };


  ws.onclose = function(e) {
    console.log("Disconnected from server");
  };

}

//click gamemode
$(".selectpvp").click(function() {
  ConnectWS();
  mode = "pvp"
  $("#victory").hide();
  $("#menu").slideUp();
  $("#game").show();
});

$(".selectpve").click(function() {
  ConnectWS();
  color = "black";
  turn = "black"
  mode = "pve";
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
  ws.send("PLAY " + coord);
  //ws.send("getscore");
});