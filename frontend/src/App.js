import React, { Component } from 'react';
import logo from './logo.svg';
import './App.css';

class App extends Component {
  componentDidMount() {
    var conn;
    var msg = document.getElementById("msg");
    var log = document.getElementById("log");
    function appendLog(item) {
        var doScroll = log.scrollTop > log.scrollHeight - log.clientHeight - 1;
        log.appendChild(item);
        if (doScroll) {
            log.scrollTop = log.scrollHeight - log.clientHeight;
        }
    }
    document.getElementById("form").onsubmit = function () {
        if (!conn) {
            return false;
        }
        if (!msg.value) {
            return false;
        }
        conn.send(msg.value);
        msg.value = "";
        return false;
    };
    if (window["WebSocket"]) {
        //conn = new WebSocket("ws://" + document.location.host + "/ws");
        conn = new WebSocket("ws://" + document.location.hostname + ":8080/ws");
        conn.onclose = function (evt) {
            var item = document.createElement("div");
            item.innerHTML = "<b>Connection closed.</b>";
            appendLog(item);
        };
        conn.onmessage = function (evt) {
            var messages = evt.data.split('\n');
            for (var i = 0; i < messages.length; i++) {
                var item = document.createElement("div");
                item.innerText = messages[i];
                appendLog(item);
            }
        };
    } else {
        var item = document.createElement("div");
        item.innerHTML = "<b>Your browser does not support WebSockets.</b>";
        appendLog(item);
    }

  }
  render() {
    return (
      <div id="gomud">
      <div id="log"></div>
      <form id="form">
      <input type="text" autoComplete="off" id="msg"/>
      </form>
      </div>
    );
  }
}

export default App;
