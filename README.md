
#Elevator Control System

##Requirements
Design and implement an elevator control system. What data structures, interfaces and algorithms will you need? Your elevator control system should be able to handle a few elevators — up to 16.

You can use the language of your choice to implement an elevator control system. In the end, your control system should provide an interface for:

    Querying the state of the elevators (what floor are they on and where they are going),
    receiving an update about the status of an elevator,
    receiving a pickup request,
    time-stepping the simulation.

For example, we could imagine in Scala an interface like this:
```scala
trait ElevatorControlSystem {
  def status(): Seq[(Int, Int, Int)]
  def update(Int, Int, Int)
  def pickup(Int, Int)
  def step()
}
```
Here we have chosen to represent elevator state as 3 integers:

Elevator ID, Floor Number, Goal Floor Number

An update alters these numbers for one elevator. A pickup request is two integers:

Pickup Floor, Direction (negative for down, positive for up)

This is not a particularly nice interface, and leaves some questions open. For example, the elevator state only has one goal floor; but it is conceivable that an elevator holds more than one person, and each person wants to go to a different floor, so there could be a few goal floors queued up. Please feel free to improve upon this interface!

The most interesting part of this challenge is the scheduling problem. The simplest implementation would be to serve requests in FCFS (first-come, first-served) order. This is clearly bad; imagine riding such an elevator! Please discuss how your algorithm improves on FCFS in your write-up.

Please provide a source tarball (or link to a GitHub repository) containing code in the language of your choice, as well as a README discussing your solution (and providing build instructions). The accompanying documentation is an important part of your submission. It counts to show your work.

Good luck!

##Implementation

Golang was selected as the language in this submission due to its modern array implementation (slices) and modern interpretation of object oriented programming (featuring interfaces instead of classes)
I also wanted to learn a new language while I do this assignment

This challenge is very similar to assigning processes to multicore processors and disk reading with multiple arms in operating systems.

As such I defined two major components to this problem:
  1. Deciding which elevator will most efficiently serve the passenger
  2. Ordering the servicing of passengers to minimize overall travel time.

  _A tertiary point is load balancing - making sure that all elevators are being used_

I addressed these components thusly:
  1. In order to to decide which elevator will most efficiently service the passenger I intially considered two components the distance to pickup and the travel time.
      1. Pickup time is calculated by the distance that between the elevator and the passenger if the elevator is moving towards the passenger and if not the distance between the elevator and its furthest stop plus the distance between that stop and the passenger
      1. Travel time is constant for all elevators that are moving in the same direction that the passenger wants to (for those going in the oppisite direction it is the distance to the furthest floor in the queue in the direction the elevator is moving from the elevators current position plus the distance between that floor and the passenger). I don't do this calucation and instead enforce a rule that elevators only pick up people who move in the same direction as the elevator (like real life).
  2. To minimize travel time I prevent duplicate stop for the same elevator as a first step. I then sort the order of stops around the current location and direction of the elevator
      * Ex.
          * if the elevator is at floor 5 and going down and had to go to floor 5,3,4,6,8,10 then the ordered list would be 5,4,3,6,8,10
          * if up then 5,6,8,10,4,3
   3. To handle load balancing I automatically assign new passengers to elevators that arent moving (currently), It is possible to use a semaphore for each elevator to further prevent overloading one elevator.


##API
Functions are provided to do the following actions from a user level:

* Query the state of all elevators including:
    * Current Position
    * Number of Passengers
    * Queue of stops requested by passengers
    * Direction of movement

```golang
   func (c *Controller) GetSystemStatus() (elevPos []int, elevPas []int, elevDir []Direction, elevDes [][]int);

```

* Update the state of an elevator in runtime:
```golang
func (c *Controller) UpdateElevator(elev int,  currentFloor int, (optional) stop int)
```

* Request Pickup for a new passenger:
```golang
func (c *Controller) CallElevator(p Passenger)
```
* Time step the simulation:
```golang
func (c *Controller) TimeStep()
```


##Execution
In a standard Golang env just running ```go run main.go cli``` should start the system and allow to see the system run

*THIS DOES NOT WORK YET (Ran out of time)*
Otherwise  ```go run main.go [path to file]``` will run a test file
