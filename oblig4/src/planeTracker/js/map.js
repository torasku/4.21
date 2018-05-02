var map;

function initMap() {
  map = new google.maps.Map(document.getElementById('map'), {
    zoom: 6,
    center: {lat: 59.188161, lng: 9.612769}
  });
}

/*function createMarker(float lat, float long){
  const marker = new google.maps.Marker({
    map: map,
    draggable: true,
    animation: google.maps.Animation.DROP,
    position: {lat: lat, lng: long}
  });
}

function addListener(marker){
  marker.addListener('click', function(){
    displayData(getPlaneId(marker));
  });
}

function displayData() {

}

function getPlaneId() {

}*/
