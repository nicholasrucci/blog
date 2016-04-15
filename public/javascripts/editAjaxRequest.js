'use strict';

var xhr = new XMLHttpRequest();
var postID = window.location.href.split('/');

postID = postID[postID.length - 2];

xhr.open('GET', '/api/posts/' + postID);
xhr.send(null);

xhr.onreadystatechange = function() {
  var DONE = 4; // readyState 4 means the request is done.
  var OK = 200; // status 200 is a successful return.
  var post, posts, i;

  if (xhr.readyState === DONE) {
    if (xhr.status === OK) {

      post = JSON.parse(xhr.responseText);

      var title   = document.getElementById("title");
      var content = document.getElementById("content");

      title.setAttribute("value", post.title);
      content.innerHTML = post.content;

    }
  }
};

function send() {
  var xhr = new XMLHttpRequest();
  var title;
  var content;
  var reqObject;
  var postID;

  var postID = window.location.href.split('/');
  postID = postID[postID.length - 2];

  title   = document.getElementById("title").value;
  content = document.getElementById("content").value

  reqObject = {
    "title": title,
    "content": content
  }

  xhr.open("PUT", "/api/posts/" + postID);
  xhr.setRequestHeader("Content-Type", "application/json; charset=UTF-8");
  xhr.send(JSON.stringify(reqObject));
  window.location.replace("/posts/" + postID);
}
