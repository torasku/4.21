var map;
var geocoder;

function initMap() {
  geocoder = new google.maps.Geocoder();
  var mapProp = {
      center:new google.maps.LatLng(59.188161, 9.612769),
      zoom:6,
  };
  map = new google.maps.Map(document.getElementById("googleMap"), mapProp);
  setMarkers();
}

function getAddress() {
  var address = document.getElementById("searchInput").value;
  //$("#currentLocation").append(address);
  geocoder.geocode({'address': address}, function(result, status) {
    var coordinates = result[0].geometry.location;
    map.setCenter(coordinates);
    var marker = new google.maps.Marker ({
       map: map,
       position: result[0].geometry.location
     });
  });
}

function submit() {
  var address = document.getElementById("searchInput").value;
  geocoder.geocode({'address': address}, function(result, status) {
    var coordinates = result[0].geometry.location;
    document.getElementById("coordinates").setAttribute("value", coordinates);
    document.getElementById("coordinates-form").submit();
  });
}

/*function setMarkers() {
    var latlng = {lat: {{.Lat}}, lng: {{.Long}}};
    var marker = new google.maps.Marker ({
      map: map,
      position: latlng
    });
}*/
