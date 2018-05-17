// select buttons and modals
var updateBtn = document.getElementById('updateBtn');
var mapBtn = document.getElementById("mapBtn");
var locationBtn = document.getElementById("locationBtn");

var mapModal = document.getElementById('mapModal');
var closeMap = document.getElementById("closeMap");
var infoModal = document.getElementById('infoModal');
var closeInfo = document.getElementById('closeInfo');
var locationModal = document.getElementById('locationModal');
var closeLocation = document.getElementById('closeLocation');
var errorModal = document.getElementById('errorModal');
var closeError = document.getElementById('closeError');

var currentLocation = $("#locationInput").val();

// handling modals
updateBtn.onclick = function(){
  location.reload();
}

mapBtn.onclick = function() {
    mapModal.style.display = "block";
}

closeMap.onclick = function() {
    mapModal.style.display = "none";
}

closeInfo.onclick = function() {
  infoModal.style.display = "none";
}

locationBtn.onclick = function() {
  locationModal.style.display = "block";
}

closeLocation.onclick = function() {
  locationModal.style.display = "none";
}

closeError.onclick = function() {
    errorModal.style.display = "none";
}

// function that retrieves detailed information about selected planes from the table
function getPlaneInfo(call, model, year, engType, speed, alt){

  var call = JSON.stringify(call);
  call = trimString(call);
  var model = JSON.stringify(model);
  model = trimString(model);
  var year = JSON.stringify(year);
  year = trimString(year);
  var engType = JSON.stringify(engType);
  engType = checkEngineType(engType);
  var speed = JSON.stringify(speed);
  var alt = JSON.stringify(alt);

  $("#infoTable").empty();
  infoModal.style.display = "block";

  $("#infoTable").append("<thead class='thead-dark'>" + "<tr>" +
    "<th>Callsign</th>" +
    "<th>Plane model</th>" +
    "<th>Year built</th>" +
    "<th>Engine type</th>" +
    "<th>Speed</th>" +
    "<th>Altitude</th></tr></thead>");

  // creating an object to hold all data
  var planeData = {
    "reg": call,
    "model": model,
    "year": year,
    "engine": engType,
    "speed": speed,
    "alt": alt
  }

  $("#infoTable").append("<tr></tr>");

  // looping all values in object, append to the table
  Object.values(planeData).forEach(function(value){
    $("#infoTable").append("<td>" + value + "</td>");
  });
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

// function that checks if input location exists in the Beta location array
function checkLocation(){

  var locationExists = Boolean(false);
  var location = $("#locationInput").val();
  $("#currentLocation").append("<h4>Showing result from: " + String(location) + "</h4");
  var locations = [
    "Bergen",
    "Oslo",
    "Kristiansand",
    "Tromsø",
    "Trondheim"
  ]

  for(let i = 0; i < locations.length; i++){
    if(locations[i] == location ){
      locationExists = true
      break;
    }
  }
  if(locationExists == false){
    // not working.... :(
    document.getElementById("searchLocBtn").addEventListener("click", function(event){
      event.preventDefault();
    });
    alert("Sorry, could not locate area. Please try again!");
    //errorModal.style.display = "block";
  }
}
