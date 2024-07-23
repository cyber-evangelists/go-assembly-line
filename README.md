# Car Assembly Line Scheduling 

## Overview
This project demonstrates a solution to the car assembly line scheduling problem using dynamic programming. The goal is to determine the minimum time required to assemble a car given two assembly lines, each with a series of stations, and additional constraints such as extra time required for switching between lines and the time taken to enter and exit each line

### Files and Directories

- **main.go**: The main Go file containing the implementation of the car assembly line scheduling algorithm.

## Algorithm Description
The algorithm calculates the minimum time required to assemble a car on two assembly lines. It uses dynamic programming to keep track of the minimum time needed to reach each station on both lines.

The approach considers the following factors:

-**a:** Time taken to process the car at each station.

-**ct:** Extra time required to switch from one line to another.

-**e:** Time taken to enter the assembly lines.

-**x:** Time taken to exit the assembly lines.


## Installation


### Steps
1. Clone the repository:
    ```sh
    git clone https://github.com/your-repo/car-assembly-line-scheduling.git
    cd car-assembly-line-scheduling

    ```
2. Build and run the Go application:
     ```sh
    go run main.go

    ```

## Conclusion

This project provides a practical implementation of the car assembly line scheduling problem, demonstrating how dynamic programming can be applied to optimize production processes. By considering factors such as processing times, switching costs, and entry/exit times, the algorithm effectively calculates the minimum time required to complete the assembly of a car.
## Acknowledgements

```bash
This app was made with 💖 by Hamza under the guidance of Sir Husnain.
```
