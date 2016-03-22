package controller

import (
    "sort"
)

/*Elevator struct
 * Contains:
 *  	passengers: List of passengers in elevator
 *  	stops: list of stops in service order
 *      currentFloor: Current position of elevator
 *      dir: Direction of movement
 *      max: Max floor value
 *      min: Min floor value
 *      isOpen: Serves as a lock preventing editing of elevator contents
 */
type Elevator struct {
    passengers []Passenger
    stops []int
    currentFloor int
    dir Direction
    max int
    min int
    isOpen bool //lock
    /*
       TODO:
       weight (type semaphore)
       limit concurrent use of resource
    */
}

/*NewElevator (function)
 * Elevator Constructor
 * @param {int} min min number of floors
 * @param {int} max max number of floors
 * @returns {*Elevator} a pointer to a new elevator
 */
func NewElevator(min, max int) (*Elevator) {
    //create new elevator on heap
    e := new(Elevator)
    //intialize values
    e.max = max
    e.min = min
    e.dir = up
    e.currentFloor = 1
    e.isOpen = false
    return e
}

/*OpenDoors (function)
 * Open elevator doors (allow change of contents)
 */
func (e *Elevator) OpenDoors() {
    e.isOpen = true
}

/*CloseDoors (function)
 * Close elevator doors (prevent change of contents)
 */
func (e *Elevator) CloseDoors() {
    e.isOpen = false
}

/*GetStatus (function)
 * Get status of the elevator
 * @return {int} passengers   number of passengers
 * @return {int} currentFloor position
 * @return {int} isOpen       state of lock
 */
func (e *Elevator) GetStatus() (passangers int, currentFloor int, isOpen bool) {
    return len(e.passengers), e.currentFloor, e.isOpen
}

/*Move (function)
 * Move elevator in set direction or turn arround
 */
func (e *Elevator) Move() {
    //safety :)
    if e.isOpen {
        return
    }
    //If there is no where to go don't move
    if e.stops == nil || len(e.stops) == 0  {
        return
    }
    //If elevators new stop is above current position and elevator is going down change direction
    //Else If elevators new stop is below current position and elevator is going up change direction
    if e.stops[0] - e.currentFloor > 0  && e.dir < 0 {
        e.dir = up
    } else if e.stops[0] - e.currentFloor < 0  && e.dir > 0 {
        e.dir = down
    }

    next := e.currentFloor + int(e.dir)
    //If next not out of range move 1 in direction of elevator
    //Otherwise change direction and move
    if next <= e.max && next >= e.min {
        e.currentFloor = next
    } else {
        if e.dir == up {
            e.dir = down
        } else {
            e.dir = up
        }
        e.currentFloor += int(e.dir)
    }
}

/*PassengerExit (function)
 * Allow passengers who requested for this floor to exit
 */
func (e *Elevator) PassengerExit() {
    //Doors must be open
    if !e.isOpen {
        return
    }
    //TODO: Refactor peopleOnThisFloor for use here,
    //find way to remove elements from one array from the other

    //Search passenger list for people who need to get off here and remove from list
    for i,p := range e.passengers {
        if p.dest == e.currentFloor {
            e.passengers = append(e.passengers[:i], e.passengers[i+1:]...)
        }
    }
}

/*PassengerEnter (function)
 * Allow passengers who requested for this floor to exit
 * @param {*[]Passenger} newPassengers List of passengers getting on on this floor
 */
func (e *Elevator) PassengerEnter(newPassengers *[]Passenger) {
    for _,person := range *newPassengers {
        //Add new stop to queue
        e.NewStop(person.dest)
        //add passenger to the list
        e.passengers = append(e.passengers, person)
    }
}

/*Distance (function)
 * @param {Passenger} p Passenger to be picked up
 * @return {int} distance how far away the elevator is
 */
func (e *Elevator) Distance(p Passenger) int {
    // If an elevator is not moving then it should pick up the new passenger
    if e.currentFloor > p.start {
        // If current floor is above passenger and elevator going down
        // Then the distance is the number of floors between passenger and elevator
        if e.dir == down {
            return e.currentFloor - p.start
        }
        // Else the distance is number of floors to the top from the elevators
        // position and the distance from the top to the passenger
        return e.furthestStop() - e.currentFloor + e.furthestStop() - p.start
    } else {
        // If current floor is below passenger and elevator going up
        // Then the distance is the number of floors between passenger and elevator
        if e.dir == up {
            return p.start - e.currentFloor
        }
        // Else the distance is number of floors to the bottom from the elevators
        // position and the distance from the bottom to the passenger
        return e.currentFloor - e.furthestStop() + p.start - e.furthestStop()
    }
}

/*furthestStop(function)
 * @return {int} which stop is furthest away
 */
func (e *Elevator) furthestStop() int {
    if len(e.stops) == 0 || e.stops == nil {
        return e.currentFloor
    }
    small, large := FindMaxMin(e.stops)
    if e.dir == up {
        return large;
    }
    return small
}

/*NewStop (function)
 * @param {int} floor New floor to stop at
 */
func (e *Elevator) NewStop(floor int) {
    //error checking
    if floor >= e.min && floor <= e.max {
        //append stop to the list
        e.stops = append(e.stops, floor)
        //remove duplicate stops in the list
        e.stops = RemoveDuplicates(e.stops)
        //order stops in most effiecent order
        e.stops = orderStops(e.stops, e.currentFloor, e.dir)
    }
}

/*OrderStops (function)
 * @param {[]int} s stops
 * @param {int} pivot current position
 * @param {Direction} dir
 * @param {[]int} ordered list
 */
func orderStops(s []int, pivot int, dir Direction) []int {
    var greater []int
    var less []int
    //sort for those greater than and those less than current location
    for _,e := range s {
        if e >= pivot {
            greater = append(greater, e)
        } else {
            less = append(less, e)
        }
    }

    sort.Ints(greater)
    sort.Sort(sort.Reverse(sort.IntSlice(less)))
    //if going up create a list that is the stops greater or equal to current
    //location and then reverse sorted those stops less than the pivot
    if dir == up {
        return append(greater, less...)
    }
    if len(greater) != 0 && greater != nil {
        if greater[0] == pivot {
            greater = greater[1:]
            less = append([]int{pivot}, less...)
        }
    }
    //if going up create a list that is the stops reverse sorted those stops
    // less or equal and then those that are greater sorted
    return append(less, greater...)

}
