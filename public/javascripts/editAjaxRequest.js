'use strict';

(function() {
  var xhr = new XMLHttpRequest();
  var postID = window.location.href.split('/');

  console.log(postID);

  postID = postID[postID.length - 2];

  console.log(postID);

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
})();
