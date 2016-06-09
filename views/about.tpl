<div class="mdl-grid text-container">
    <div class="mdl-cell mdl-cell--1-col">
    </div>
    <div class="mdl-cell mdl-cell--10-col">
        <h1>About Optiroute</h1>
        <p>
            Lorem ipsum dolor sit amet, consectetur adipisicing elit. Blanditiis explicabo neque sit soluta suscipit tempore voluptate voluptatum. Adipisci alias aperiam aspernatur eum modi, molestias, quae, quos ratione unde vero vitae.
        </p>
        <p>Lorem ipsum dolor sit amet, consectetur adipisicing elit. Cumque, dicta, nisi? Esse, inventore ipsum perspiciatis porro quia quod reprehenderit sit tempora. Eius expedita fugit ipsum modi nulla placeat rerum veritatis? Lorem ipsum dolor sit amet, consectetur adipisicing elit. Ad architecto cupiditate ex. Aspernatur cum debitis dolorum eveniet hic, in incidunt itaque maiores maxime molestias, nam natus nihil odit perferendis sequi.</p>

    </div>
    <div class="mdl-cell mdl-cell--1-col">
    </div>
</div>

<div class="mdl-grid">
<div class="mdl-cell mdl-cell--1-col">
</div>
    <div class="mdl-cell mdl-cell--4">
        <p>
            To see the algorithm in action click "Generate Points." This will run the algorithm on a random set of locations, show you the initial route, then allow you to see the optimized route.
        </p>
        <div>
            <button id="gen-points" class="mdl-button mdl-js-button mdl-button--raised">
                Generate Points
            </button>
            <button disabled="disabled" id="optimize-points" class="mdl-button mdl-js-button mdl-button--raised mdl-button--accent">
                View Optimized Route
            </button>
        </div>
        <div id="loading-container">
            <div id="p2" class="mdl-progress mdl-js-progress mdl-progress__indeterminate"></div>
        </div>
        <div id="results">
            <h4>Initial Distance: <span id="res-init"></span></h4>
            <h4>Final Distance: <span id="res-final" ></span></h4>
            <h6>Percent Decrease: <span id="res-decrease"></span></h6>
        </div>
    </div>
    <div class="mdl-cell mdl-cell--6-col">
        <div id="graph-container">

        </div>
    </div>
    <div class="mdl-cell mdl-cell--1-col">
    </div>
</div>



