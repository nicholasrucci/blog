'use strict';

(function() {
  var xhr = new XMLHttpRequest();
  var postID = window.location.href.split('/');

  postID = postID[postID.length - 1];

  xhr.open('GET', '/api/posts/' + postID);
  xhr.send(null);

  xhr.onreadystatechange = function() {
    var DONE = 4; // readyState 4 means the request is done.
    var OK = 200; // status 200 is a successful return.
    var post, posts, i;

    if (xhr.readyState === DONE) {
      if (xhr.status === OK) {

        post = JSON.parse(xhr.responseText);
        splitContent(post.content);

      }
    }
  };

  function splitContent(content) {
    var x = document.getElementById("blog-content");
    var splitContent = content.trim().split("\n");

    x.innerHTML = "";

    for (var i = 0; i < splitContent.length; i++) {
        var p = document.createElement("p");
        p.innerHTML = splitContent[i];
        x.appendChild(p);
    }

  }
})();
