var ws;
var mode;
var turn = "null";
var score = "Black : 0 | White : 0";
var scoreb = 0;
var scorew = 0;
var me = "You are : unknown";
var color = "null";
var host = window.location.hostname;
var portWS;

//init page
$("#turn").text("IDK");
$("#score").text(score);
$("#me").text(me);

function getPortWS() {
  var url = window.location.href;
  var arr = url.split("/");
  var result = arr[2].split(":")[1];
  //portWS = parseInt(result, 0);
  portWS = result;
  //console.log(portWS);
}

function getURLParameter(name) {
  return decodeURI(
  (RegExp(name + '=' + '(.+?)(&|$)').exec(location.search) || [null])[1]);
}

$(document).ready(function() {
  var res = getURLParameter("win");
  if(res == "1") {
    $("#victory").text("YOU WON").show(300);
  } else if(res == "2") {
    $("#victory").text("YOU LOSE").show(300);
  }
   $("#mode").hide(200);
   getPortWS();
});

function sendRules() {
  var dual3 = 0;
  var break5 = 0;
  var timeout = 0;

  if($('#DOUBLE_3').attr('checked')) {
    dual3 = "1";
  }
  if($('#BREAKING_5').attr('checked')) {
    break5 = "1";
  }
  timeout = $('#TIMEOUT').val();
  ws.send("RULES " + dual3 + " " + break5 + " " + timeout);
}

function ConnectWS() {

  ws = new WebSocket("ws://" + host + ":"+portWS+"/ws");

  ws.onopen = function() {
    console.log("Connected to server");
    ws.send("MODE " + mode);
    if(color == "null") {
      ws.send("GETCOLOR");
    }
    ws.send("GETTURN");
    $("#sendrules").removeAttr("disabled");
     $("#reset").removeAttr("disabled");
    $(".infoconnect").css("color", "green");
    $(".infoconnect").html("Connected");
  };

  ws.onmessage = function(e) {
    console.log("receive : '" + e.data + "'");

    if(e.data != "error") {
      var data = e.data.split(' ');

      if(data[0] == "COLOR") {
        color = data[1];
      }

      if(data[0] == "RULES") {
        if (data[1] == "1") {
          $('#DOUBLE_3').attr('checked', 'checked');
        } else if (data[1] == "0") {
          $('#DOUBLE_3').removeAttr('checked');
        }
        if (data[2] == "1") {
          $('#BREAKING_5').attr('checked', 'checked');
        } else if (data[2] == "0") {
          $('#BREAKING_5').removeAttr('checked');
        }
        $('#TIMEOUT').val(data[3]);
      }

      if (data[0] == "HINT") {
        $('.boardclic td div').removeClass("hintstone");
        $('.pos' + data[1] + 'y' + data[2]).addClass("hintstone");
      }

      if(data[0] == "COLOR") {
        color = data[1];
        $("#me").text("You are : " + color);
      }

      if(data[0] == "TURN") {
        turn = data[1];
      }
      //he won
      if(data[0] == "LOSE") {
        ws.send("reset");
        window.location.href = location.protocol + '//' + location.host + location.pathname + "?win=2";
      }

      //you won
      if(data[0] == "WIN") {
        ws.send("reset");
        window.location.href = location.protocol + '//' + location.host + location.pathname + "?win=1";
      }

      //stone captured
      else if(data[0] == "REM") {
        if (turn == "black") {
          scoreb += 2;
        } else if (turn == "white") {
          scorew += 2;
        }
        $('.pos' + data[1] + 'y' + data[2]).removeClass("bgblack");
        $('.pos' + data[1] + 'y' + data[2]).removeClass("bgwhite");
        $('.pos' + data[3] + 'y' + data[4]).removeClass("bgblack");
        $('.pos' + data[3] + 'y' + data[4]).removeClass("bgwhite");
        $("#score").text("Black : "+scoreb+" | White : "+scorew);
      }

      else if(data[0] == "YOURTURN") {
        $("#turn").text("YOUR TURN");
      } else if(data[0] == "ADD") {
        $('.boardclic td div').removeClass("hintstone");
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
    $(".infoconnect").css("color", "red");
    $(".infoconnect").html("Disconnected ;(");
  };

}

//click gamemode
$(".selectpvp").click(function() {
  mode = "pvp";
  ConnectWS();

  $("#victory").hide();
  $("#menu").slideUp();
  $("#game").show();
});

$(".selectpve").click(function() {
  mode = "pve";
  ConnectWS();
  color = "black";
  turn = "black";
  $("#victory").hide();
  $("#menu").slideUp();
  $("#game").show();
  $('#hint').attr('disabled', 'disabled');
});

//togglesetting

$(".settingmenuswitch").click(function() {
  if($('#mode').css('display') == "none" ) {
    $("#mode").show(200);
  } else {
    $("#mode").hide(200);
  }
});

$("#sendrules").click(function() {
   sendRules();
});

//give player an hint
$("#hint").click(function() {
  ws.send("HINT");
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