# greedy-algorithm-go

### Goal

To develop the algorithm to find the minimum distance route that can join all capitals of Argentina, using the following heuristics: "From each city go to the nearest city not visited".


### Greedy algorithm

The search strategy follows a heuristic that consists of choosing the optimal option at each local step in the hope of arriving at an optimal overall solution.

The algorithm has the following structure

- Choose the city of departure.
- The next city to visit is selected from a list.
- Validate that the selected city is the one that is the least distance away and has not been visited, otherwise returns to step 2.
- It is validated that there are cities to visit, in this case, step 2 and 3 is repeated; otherwise, the route is finished.

![N|Solid](https://github.com/manerajona/greedy-algorithm-go/blob/master/greedy.gif)


## Build & Run

For building:
```sh
$ go build *.go
```
For running:
```sh
$ go run *.go
```
