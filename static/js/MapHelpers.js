// A messy collection of UI helper functions

var locations = [];

function validateLocations() {
    var validPattern = /^[0-9a-zA-Z,\s-+]+$/;

    //Check for at least 4 locations
    if (locations.length < 4) {
        showToast("Please include at least 4 locations in the route.");
        return false;
    }

    for (var i = 0; i < locations.length; i++) {
        if(!validPattern.test(locations[i])) {
            showToast('"' + locations[i] + '"' + " contains invalid characters.");
            return false;
        }
    }

    return true;
}

function removeLocation(loc) {
    loc = loc.replaceAll(" ", "+");

    console.log("removing " + loc);

    for (var i=locations.length-1; i>=0; i--) {
        if (locations[i] === loc) {
            locations.splice(i, 1);
            break;
        }
    }
}

function showToast(msg) {
    var notification = document.querySelector('.mdl-js-snackbar');
    notification.MaterialSnackbar.showSnackbar(
      {
        message: msg,
        timeout: 5000
      }
    );
}

function sendLocations() {
    console.log("sending:");
    console.log(locations)

    // Check for errors
    if (!validateLocations()) {
        return;
    }

    $.ajax({
      type: "POST",
      url: "/route",
      data: JSON.stringify(locations),
      success: function(locations){
        console.log(toType(locations));

        locations = locations.route;

        origin = locations[0];
        destination = locations[0];
        waypoints = locations.slice(1);

        console.log("origin");
        console.log(origin);

        // Print the results
        console.log("Final Route:");
        console.log(locations);

        // Set the map
       document.getElementById("google-map").src = buildURL(origin, destination, waypoints);
      },
      fail: function(){
        showToast("There was an error processing your route, please try again.");
      },
      error: function(){
          showToast("There was an error processing your route, please try again.");
      },
      'processData': false,
      'contentType': 'application/json',
    });
}

var toType = function(obj) {
  return ({}).toString.call(obj).match(/\s([a-zA-Z]+)/)[1].toLowerCase()
}

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
    li.setAttribute("id", locString + "li");
    span.setAttribute("class", "mdl-list__item-primary-content");

    var icon = document.createElement("span");
    icon.setAttribute("class","delete-icon");
    icon.appendChild(document.createTextNode("✖"));

    // Append the elements
    span.appendChild(document.createTextNode(locString));
    li.appendChild(span);
    li.appendChild(icon)
    ul.appendChild(li);

    // Remove the old value from the text box
    $("#loc-to-add").val("");

    // Give focus to the text box
    $("#loc-to-add").focus();

    $(".delete-icon").click(function(e){
        $(this).parent().hide();

        removeLocation($(this).parent().children().eq(0).text());

        console.log(locations);
    });
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
    // So the enter key adds a location
    $("#loc-to-add").keyup(function(event){
        event.preventDefault();
        if(event.keyCode == 13){
            $("#add-loc-button").click();
        }
    });
})