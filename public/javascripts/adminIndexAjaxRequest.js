'use strict';

window.onload = function() {
  getAllPosts();
};

function sendDelete(id) {
  var xhr = new XMLHttpRequest();

  xhr.open("DELETE", "/api/posts/" + id);
  xhr.send(null);

  $('#posts-container').load(document.URL +  ' #posts-container');
  getAllPosts();
}

function getAllPosts() {
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
        posts.reverse();

        for (i = 0; i < posts.length; i++) {
          var card        = document.createElement("div");
          var cardContent = document.createElement("div");
          var editDiv     = document.createElement("div");
          var trashDiv    = document.createElement("div");
          var title       = document.createElement("h5");
          var titleText   = document.createTextNode(posts[i].title);
          var content     = document.createElement("p");
          var contentText = document.createTextNode(posts[i].content.slice(0, 245) + "....");
          var element     = document.getElementById("posts-container");
          var trashIcon   = document.createElement("i");
          var editIcon    = document.createElement("i");
          var trashLink   = document.createElement("a");
          var editLink    = document.createElement("a");

          card.appendChild(cardContent);
          card.className += "card-panel grey lighten-5";

          cardContent.appendChild(title)
          cardContent.className += "card-content black-text";
          cardContent.appendChild(content);

          editLink.appendChild(editDiv);
          trashLink.appendChild(trashDiv);

          editDiv.appendChild(editIcon);
          editDiv.setAttribute("id", "" + posts[i].id);

          trashDiv.appendChild(trashIcon);
          trashDiv.setAttribute("id", "" + posts[i].id);
          trashDiv.setAttribute("onclick", "sendDelete(this.id)");

          trashIcon.className += "right fa fa-trash fa-lg";
          editIcon.className += "right fa fa-pencil-square-o fa-lg";

          title.appendChild(titleText);
          title.className += "card-title";

          content.appendChild(contentText);
          content.appendChild(trashLink);
          content.appendChild(editLink);

          element.appendChild(card);
        }
      }
    }
  };
}
