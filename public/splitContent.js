'use strict';

(function() {
  window.onload = function () {
    var x = document.getElementById("blog-content");
    var content = x.innerHTML;
    var splitContent = content.trim().split("\n");

    x.innerHTML = "";

    for (var i = 0; i < splitContent.length; i++) {
        var p = document.createElement("p");
        p.innerHTML = splitContent[i];
        x.appendChild(p);
    }

  }
})();
