'use strict';

  function send() {
    var xhr = new XMLHttpRequest();
    var title;
    var content;
    var reqObject;

    title   = document.getElementById("title").value;
    content = document.getElementById("content").value

    reqObject = {
      "title": title,
      "content": content
    }

    xhr.open("POST", "/api/posts");
    xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
    xhr.send(JSON.stringify(reqObject));
    window.location.replace("/admin/posts");
  }
