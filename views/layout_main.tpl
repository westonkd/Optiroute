<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01//EN" "http://www.w3.org/TR/html4/strict.dtd">
<html lang="en">
<head>
    <meta http-equiv="Content-Type" content="text/html;charset=UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Optiroute</title>
    {{ .HeadScripts }}
    {{ .HeadStyles }}
</head>
<body>

<div class="mdl-layout mdl-js-layout mdl-layout--fixed-drawer
            mdl-layout--fixed-header">
    {{ .Header }}
    <div class="mdl-layout__drawer">
        <span class="mdl-layout-title">
            <img id="logo" src="/static/img/logo.png" alt="Optiroute"/>
        </span>

        <div class="location-form">

            <div id="instructions">
                Use the bellow text box to add locations to your trip.
            </div>

            <div action="#">
                <div class="mdl-textfield mdl-js-textfield mdl-textfield--floating-label">
                    <input class="mdl-textfield__input" type="text" id="loc-to-add"  />
                    <label class="mdl-textfield__label" for="loc-to-add">Location</label>
                </div>
            </div>

            <button id="add-loc-button" class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--colored full-width"
                onclick="addLocation()">
                +
            </button>

            <button id="submit-locs-button" onclick ="sendLocations()" class="mdl-button mdl-js-button mdl-button--raised mdl-js-ripple-effect mdl-button--accent full-width">
                Get Route
            </button>

            <h6 id="loc-header">Locations</h6>

            <ul id="loc-list" class="demo-list-item mdl-list">
                <li id="info-list">
                    <span>
                        Locations added using the above text box will be displayed here.
                    </span>
                </li>
            </ul>

        </div>
    </div>
    <main class="mdl-layout__content">
        <div class="page-content">
            {{ .LayoutContent }}
        </div>
    </main>
</div>

    {{ .Footer }}
</body>
</html>