# greedy-algorithm-go

### Goal

To develop the algorithm to find the minimum distance route that can join all capitals of Argentina, using the following heuristics: "From each city go to the nearest city not visited".


### Greedy algorithm

The search strategy follows a heuristic that consists on choosing the optimal option on each local step in the hope of arriving at the optimal overall solution.

The algorithm has the following structure

1. Choose the city of departure.
2. The next city to visit is selected from a list.
3. Validate if the selected city is the one that's the least distance away and hasn't been visited yet, otherwise returns to step 2.
4. Validate if there are cities to visit, in that case, step 2 and 3 is repeated; otherwise, the route is finished.

## Build & Run

For building:
```sh
$ go build *.go
```
For running:
```sh
$ go run *.go
```
