var map;
var geocoder;
var circle;

function initMap() {
  geocoder = new google.maps.Geocoder();
  var mapProp = {
      center:new google.maps.LatLng(59.188161, 9.612769),
      zoom:6,
  };
  map = new google.maps.Map(document.getElementById("googleMap"), mapProp);
  setMarkers();
}

// get input value and convert to coordinates for map
function getAddress() {
  if(circle != null){
      circle.setMap(null);
  }
  var address = document.getElementById("searchInput").value;
  geocoder.geocode({'address': address}, function(result, status) {
    var coordinates = result[0].geometry.location;
    map.setCenter(coordinates);
    circle = new google.maps.Circle({
            strokeColor: '#FF0000',
            strokeOpacity: 0.8,
            strokeWeight: 2,
            fillColor: '#FF0000',
            fillOpacity: 0.35,
            map: map,
            center: coordinates,
            radius: 100000
    });
  });
}

// append coordinates to hidden form and then send to server
function submit() {
  var address = document.getElementById("searchInput").value;
  geocoder.geocode({'address': address}, function(result, status) {
    var coordinates = result[0].geometry.location;
    document.getElementById("coordinates").setAttribute("value", coordinates);
    document.getElementById("coordinates-form").submit();
  });
}

// reset input field and remove circle from map for each time "search location" button is pushed
function reset() {
  $("#searchInput").empty();
  circle.setMap(null);
}
