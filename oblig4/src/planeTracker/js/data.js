var updateBtn = document.getElementById('updateBtn');
var searchBtn = document.getElementById("searchBtn");
var locationBtn = document.getElementById("locationBtn");

var searchModal = document.getElementById('searchModal');
var closeSearch = document.getElementById("closeSearch");
var infoModal = document.getElementById('infoModal');
var closeInfo = document.getElementById('closeInfo');

updateBtn.onclick = function(){
  location.reload();
}

searchBtn.onclick = function() {
    searchModal.style.display = "block";
}

closeSearch.onclick = function() {
    searchModal.style.display = "none";
}

window.onclick = function(event) {
    if (event.target == modal) {
        searchModal.style.display = "none";
    }
}

locationBtn.onclick = function() {
  alert("Coming soon!");
}

function getPlaneInfo(plane){

  $("#infoTable").empty();
  infoModal.style.display = "block";

  $("#infoTable").append("<thead class='thead-dark'>" + "<tr>" +
    "<th>Registration number</th>" +
    "<th>Plane model</th>" +
    "<th>Engine type</th></tr></thead>");

  var testData = {
    "reg": "QWE3",
    "model": "Dash 8 Q400",
    "engine": "Turbo prop"
  }

  $("#infoTable").append("<tr></tr>");

  Object.values(testData).forEach(function(value){
    $("#infoTable").append("<th>" + value + "</th>");
  });

}

closeInfo.onclick = function() {
  infoModal.style.display = "none";
}
