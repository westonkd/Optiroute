# Optiroute

[Optiroute](http://optiroute.herokuapp.com) is a web application that uses a genetic algorithm to find an optimal tour of a set of locations. Genetic algorithms are inspired by natural evolutionary processes and can be summarized in five steps:

## Initialization
The algorithm begins by creating a set of “random guesses” that might be a good route. This set is called the population. These guesses are usually very poor but are used to create better guesses. The total distance of each route is then calculated and stored.

## Selection
Next we need to select some of the best routes from the initial population for “breeding”. There are many ways to accomplish this, but Optiroute uses a process called a tournament select.

## Crossover
After this selection occurs, two of the selected routes (called parents) are combined in a process called crossover. During crossover the best aspects of each parent route are combined to form a single child route which is (usually) more efficient than its parents. Crossover is used to create an entire new generation of routes.

## Mutation
Just like in nature, the algorithm occasionally makes small changes to the child routes. These changes are random and help ensure that the algorithm is homing in on the global maximum rather than a local maximum.

## Repeat
By completing crossover and mutation, we now have an entirely new population of routes that is better than our initial population. Now we repeat the above steps on our new population until we find a route that is sufficiently short.


