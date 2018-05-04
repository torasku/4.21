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

function getPlaneInfo(call, model, engType){

  var call = JSON.stringify(call);
  call = trimString(call);
  var model = JSON.stringify(model);
  model = trimString(model);
  var engType = JSON.stringify(engType);
  engType = checkEngineType(trimString(engType));

  $("#infoTable").empty();
  infoModal.style.display = "block";

  $("#infoTable").append("<thead class='thead-dark'>" + "<tr>" +
    "<th>Callsign</th>" +
    "<th>Plane model</th>" +
    "<th>Engine type</th></tr></thead>");

  var planeData = {
    "reg": call,
    "model": model,
    "engine": engType
  }

  $("#infoTable").append("<tr></tr>");

  Object.values(planeData).forEach(function(value){
    $("#infoTable").append("<th>" + value + "</th>");
  });
}

closeInfo.onclick = function() {
  infoModal.style.display = "none";
}

function trimString(s){
  var st = s.replace(/"/g, "");
  return st;
}

function checkEngineType(e){
  var engType;

  if(e == 0){
    engType = "None";
  }
  else if(e == 1){
    engType = "Piston";
  }
  else if(e == 2){
    engType = "Turbo";
  }
  else if(e == 3){
    engType = "Jet";
  }
  else if(e == 4){
    engType = "Electric";
  }
  return engType;
}
