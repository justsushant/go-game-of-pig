# Game Of Pig

This excerise deals with a command-line program that simulates the game of Pig. 

It plays with the "hold out until" strategy, with strategy provided in the command line arguments. Results are printed on the standard output, majorly from the POV of Player1. There are three ways to run this:

- if both arguments are numbers, it plays them
- if one of the arguments are range based, it plays the single strategy against the range, skipping the case where both have the same strategy
- if both arguments are range based, then the program prints out the summary instead of the results



This exercise has been solved in a TDD (test driven development) fashion. Please refer to the execise [here](https://one2n.io/go-bootcamp/go-projects/a-game-of-pig).


## Running the code

1. Run the below command to build the binary. It has been saved in the bin directory.
```
make build
```

2. For testing the program, arguments can be passed as follows:
```
./bin/game 21 56
```
OR 
```
./bin/game -p1 21 -p2 56
```