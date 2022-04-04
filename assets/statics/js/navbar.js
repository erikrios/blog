const path = window.location.pathname;

var navItem;

if (path === "/project") {
  navItem = document.getElementById("project");
} else if (path === "/resume") {
  navItem = document.getElementById("resume");
} else if (path === "/about") {
  navItem = document.getElementById("about");
} else {
  navItem = document.getElementById("home");
}

navItem.className = "active";

/* Toggle between adding and removing the "responsive" class to topnav when the user clicks on the icon */
function myFunction() {
  var x = document.getElementById("myTopnav");
  if (x.className === "topnav") {
    x.className += " responsive";
  } else {
    x.className = "topnav";
  }
}
