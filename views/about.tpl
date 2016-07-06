<div class="mdl-grid text-container">
    <div class="mdl-cell mdl-cell--1-col">
    </div>
    <div class="mdl-cell mdl-cell--10-col">
        <h1>About Optiroute</h1>
        <p>
            Optiroute uses a genetic algorithm to find an optimal tour of a set of locations. Genetic algorithms are inspired by natural evolutionary processes and can be summarized in five steps:
        </p>
        <h2>
            Initialization
        </h2>
        <p>
            The algorithm begins by creating a set of “random guesses” that might be a good route. This set is called the population. These guesses are usually very poor but are used to create better guesses. The total distance of each route is then calculated and stored.
        </p>
        <h2>
            Selection
        </h2>
        <p>
            Next we need to select some of the best routes from the initial population for “breeding”. There are many ways to accomplish this, but Optiroute uses a process called a
            <a href="https://en.wikipedia.org/wiki/Tournament_selection">tournament select</a>.
        </p>
        <h2>
            Crossover
        </h2>
        <p>
            After this selection occurs, two of the selected routes (called parents) are combined in a process called crossover. During crossover the best aspects of each parent route are combined to form a single child route which is (usually) more efficient than its parents. Crossover is used to create an entire new generation of routes.
        </p>
        <h2>
            Mutation
        </h2>
        <p>
            Just like in nature, the algorithm occasionally makes small changes to the child routes. These changes are random and help ensure that the algorithm is homing in on the global maximum rather than a local maximum.
        </p>
        <h2>Repeat</h2>
        <p>
            By completing crossover and mutation, we now have an entirely new population of routes that is better than our initial population. Now we repeat the above steps on our new population until we find a route that is sufficiently short.
        </p>

    </div>
    <div class="mdl-cell mdl-cell--1-col">
    </div>
</div>

<div class="mdl-grid">
<div class="mdl-cell mdl-cell--1-col">
</div>
    <div class="mdl-cell mdl-cell--4">
        <p>
            Steps to see the algorithm in action:
        </p>
        <ol>
            <li>Click "Generate Points" to generate a random set of points and show one of the <em>initial</em> routes the algorithm will use.</li>
            <li>Click "View Optimized Route" to show the final optimized route.</li>
        </ol>
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



