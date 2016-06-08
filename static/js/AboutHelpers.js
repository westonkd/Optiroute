$(document).ready(function(){
    $('#gen-points').click(function(){
        setRandomPoints();
    })

    $('#optimize-points').click(function(){
            $.ajax({
              type: "POST",
              url: "/about",
              data: JSON.stringify(locations),
              success: function(locations){

                console.log(locations);
                showToast("Success.");

                // Set the map
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
     });

    var locations = [];

    var s = new sigma('graph-container');
    s.defaultNodeColor = "#1390F5";

    function getRandomInt(min, max) {
      return Math.floor(Math.random() * (max - min + 1) + min);
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

    function setRandomPoints(){
        s.graph.clear();
        locations = [];

        for (var i = 0; i < 50; i++) {
            var randX = getRandomInt(0,800);
            var randY = getRandomInt(0,390);
            var locID = 'n' + i;

            var location = {
                Name: locID,
                Long: randX,
                Lat: randY,
                Id: i
            }

            // Add the location
            locations.push(location);

            s.graph.addNode({
                  // Main attributes:
                  id: locID,
                  label: i,
                  // Display attributes:
                  x: randX,
                  y: randY,
                  size: 1,
                  color: '#F44336'
                })
        }

        s.refresh();
        console.log(locations);
    }


});



