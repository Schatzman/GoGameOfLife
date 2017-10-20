# Conway's Game of Life

## The Rules
1. Any live cell with fewer than two live neighbours dies, as if caused by underpopulation.
2. Any live cell with two or three live neighbours lives on to the next generation.
3. Any live cell with more than three live neighbours dies, as if by overpopulation.
4. Any dead cell with exactly three live neighbours becomes a live cell, as if by reproduction.

This version in Go by Dan Schatz-Miller by strict TDD, October 2017.

## To Build
1. Move the `GameOfLife/` directory to your `$GOPATH/src`
2. `cd` into the `GameOfLife/` directory
3. `go build -o GameOfLife` (or `-o GameOfLife.exe` on Windows)
4. `GameOfLife`
### Alternately, if you've...
- created `src/` and `bin/` at your `$GOPATH`
- set your `$GOSRCPATH`
- added your `$GOBINPATH` to your `$PATH`
2. `cd $GOSRCPATH`
3. `go install GameOfLife`
4. `GameOfLife`