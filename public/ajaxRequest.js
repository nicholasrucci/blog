/* eslint no-empty-label: 0 */

'use strict';

(function() {
  var xhr = new XMLHttpRequest();

  xhr.open('GET', '/api/posts');
  xhr.send(null);

  xhr.onreadystatechange = function() {
    var DONE = 4; // readyState 4 means the request is done.
    var OK = 200; // status 200 is a successful return.
    var post, posts, i;

    if (xhr.readyState === DONE) {
      if (xhr.status === OK) {

        posts = JSON.parse(xhr.responseText);
        console.log(posts);

        for (i = 0; i < posts.length; i++) {
          var div = document.createElement("div");
          var title = document.createElement("h2");
          div.appendChild(title)
          var titleText = document.createTextNode(posts[i].title);
          title.appendChild(titleText);

          var element = document.getElementById("posts-container");
          element.appendChild(div);
        }
      }
    }
  };
})();
