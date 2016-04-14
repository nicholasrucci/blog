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
          var card        = document.createElement("div");
          var cardContent = document.createElement("div");
          var cardAction  = document.createElement("div");
          var link        = document.createElement("a");
          var title       = document.createElement("span");
          var titleText   = document.createTextNode(posts[i].title);
          var content     = document.createElement("p");
          var contentText = document.createTextNode(posts[i].content);
          var element     = document.getElementById("posts-container");

          card.appendChild(cardContent);
          card.className += "card blue-grey darken-1";

          cardContent.appendChild(title)
          cardContent.className += "card-content white-text";
          cardContent.appendChild(content)
          cardContent.appendChild(cardAction);

          cardAction.className += "card-action";
          cardAction.appendChild(link);

          link.innerHTML = "Click to Read";
          link.setAttribute("href", "#");

          title.appendChild(titleText);
          title.className += "card-title";

          content.appendChild(contentText);

          element.appendChild(card);
        }
      }
    }
  };
})();
