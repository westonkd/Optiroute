//Set of waypoints
var waypoints = [
        "Soda+Springs,Idaho",
        "Sugar+City,Idaho"
];

var locations = [];

function addLocation() {
    var locString = $("#loc-to-add").val();

    // Remove the explanation
    $("#info-list").hide();

    // Don't add an empty string
    if (locString == "") {
        return;
    }

    //Add the location to the array of locations
    locations.push(locString.replaceAll(" ", "+"));
    console.log(locations);

    // Get the list
    var ul = document.getElementById("loc-list");

    // Create the new elements
    var li = document.createElement("li");
    var span = document.createElement("span");

    // Set the proper classes
    li.setAttribute("class", "mdl-list__item");
    span.setAttribute("class", "mdl-list__item-primary-content");

    // Append the elements
    span.appendChild(document.createTextNode(locString));
    li.appendChild(span);
    ul.appendChild(li);

    // Remove the old value from the text box
    $("#loc-to-add").val("");

    // Give focus to the text box
    $("#loc-to-add").focus();
}

String.prototype.replaceAll = function(search, replacement) {
    var target = this;
    return target.split(search).join(replacement);
};

function buildURL(start, end, waypoints) {
        //&origin=Montpelier+Idaho&destination=Paris+Idaho&avoid=tolls&waypoints=Dingle,Idaho|Soda+Srings,Idaho
        // Set up the base
        var url = "https://www.google.com/maps/embed/v1/directions?key=AIzaSyCQykIIlV-s5iZMdGct301A4AC-o8CIPbg";

        // Set up the origin
        url += "&origin=" + start;

        // Set up the destination
        url += "&destination=" + end;

        // Set up the waypoints
        url += "&waypoints=";
        var arrayLength = waypoints.length;
        for (var i = 0; i < arrayLength; i++) {
            url += waypoints[i];

            // Add the pipe if needed
            if (!(i == arrayLength - 1)) {
                url += "|"
            }
        }

        return url
}

$(document).ready(function(){
    //Test
    document.getElementById("google-map").src = buildURL("Montpelier,Idaho", "Rexburg,Idaho", waypoints);

    // So the enter key adds a location
    $("#loc-to-add").keyup(function(event){
        if(event.keyCode == 13){
            $("#add-loc-button").click();
        }
    });
})