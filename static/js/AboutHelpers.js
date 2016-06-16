
$(document).ready(function() {
    var response;
    var locations = [];

    $('#gen-points').click(function() {
        setRandomPoints();
    })

    $('#optimize-points').click(function() {
        for (var i = 1; i < response.Initial.Locations.length; i++) {
            s.graph.dropEdge('e' + i);
        }

        for (var i = 1; i < response.Final.Locations.length; i++) {
            try {
                s.graph.addEdge({
                    id: 'e' + i,
                    source: response.Final.Locations[i - 1].Name,
                    target: response.Final.Locations[i].Name
                });
            } catch (e) {
                console.log(response.Final.Locations[i].Name)
            }
        }

        console.log(response.Final.Locations);

        // add the final edge
        try {
            s.graph.addEdge({
                id: "eFinal",
                source: response.Final.Locations[0].Name,
                target: response.Final.Locations[19].Name,
            });
        } catch(e){
            console.log("Error adding final edge.");
            // do nothing
        }
        s.refresh();
         $("#res-final").text(response.FinalDistance);
         $("#res-decrease").text(parseInt((response.InitialDistance - response.FinalDistance) / response.InitialDistance * 100) + "%");
    });

    var s = new sigma('graph-container');
    s.settings({labelThreshold: 100});

    function getRandomInt(min, max) {
        return Math.floor(Math.random() * (max - min + 1) + min);
    }

    function showToast(msg) {
        var notification = document.querySelector('.mdl-js-snackbar');
        notification.MaterialSnackbar.showSnackbar({
            message: msg,
            timeout: 2000
        });
    }

    function setRandomPoints() {
        s.graph.clear();
        locations = [];

        $("#optimize-points").prop("disabled", true);
        $("#loading-container").show();

        showToast("Generating random route, please wait.");

        for (var i = 0; i < 20; i++) {
            var randX = getRandomInt(0, 1600);
            var randY = getRandomInt(0, 780);
            var locID = 'n' + i.toString();

            var location = {
                Name: locID,
                Long: randX,
                Lat: randY,
                Id: i + 1
            }

            // Add the location
            locations.push(location);

            s.graph.addNode({
                // Main attributes:
                id: locID,
                label: locID,
                // Display attributes:
                x: randX,
                y: randY,
                size: 1,
                color: '#F44336',
                hover_color: '#F44336'
            })
        }

        $.ajax({
            type: "POST",
            url: "/about",
            data: JSON.stringify(locations),
            success: function(data) {
                $("#res-init").text(data.InitialDistance);

                response = data;
                for (var i = 1; i < data.Initial.Locations.length; i++) {
                    s.graph.addEdge({
                        id: 'e' + i,
                        source: data.Initial.Locations[i - 1].Name,
                        target: data.Initial.Locations[i].Name
                    });
                }

                $("#loading-container").hide();
                $("#optimize-points").prop("disabled", false);
                s.refresh();
                showToast("Complete. Click 'View Optimized Route' to route after optimization was done.")
            },
            fail: function() {
                showToast("There was an error processing your route, please try again.");
            },
            error: function() {
                showToast("There was an error processing your route, please try again.");
            },
            'processData': false,
            'contentType': 'application/json',
        });

        s.refresh();
    }


});