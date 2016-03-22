package controller

/*ElevatorController struct
 * Contains:
 *  	elevators: List of Elevators in service
 *  	numFloors: number of floors in the building
 *      waitingPeople: Passengers who are waiting to get on an elevator
 *      unassignedPeople: Passengers who don't an elevator
 *      				  scheduled to pick them up yet
 */
type ElevatorController struct {
    elevators []Elevator
    numFloors int
    waitingPassengers []Passenger
    unassignedPassengers []Passenger
}

/*NewElevatorController (function)
 * ElevatorController Constructor, initalizes values for
 * number of floors and elevators
 * @param {int} floors max number of floors
 * @param {int} elev   number of elevators
 * @returns {*Controller} a pointer to a new ElevatorController
 */
func NewElevatorController(floors, elev int) (*ElevatorController) {
    //new controller on heap
    c := new(ElevatorController)
    c.numFloors = floors
    for i := 0; i < elev; i++ {
        //Construct new elevator
        e := NewElevator(1, c.numFloors)
        c.elevators = append(c.elevators, *e)
    }
    return c
}

////////////////////////////////////////////////////////////////////////////////
// User Accessible API
//    - TimeStep: Step one time unit forward
//    - CallElevator: Request an elevator for a new passenger
//    - UpdateElevator: Change the state of an elevator during runtime
//    - GetSystemStatus: Receive information on the state of all elevators
////////////////////////////////////////////////////////////////////////////////

/*TimeStep (function)
 * Step one time unit forward
 */
func (c *ElevatorController) TimeStep() {
    //assign passengers added by users to elevators
    c.assignPassengers()
    //move the elevators
    c.moveElevators()
}

/*CallElevator (function)
 * Request an elevator for a new passenger
 * @param  {Passenger} p New Passenger
 */
func (c *ElevatorController) CallElevator(p Passenger) {
    //add new passenger to the unassigned queue
    c.unassignedPassengers = append(c.unassignedPassengers, p)
}

/*UpdateElevator (function)
 * Change the state of an elevator during runtime
 * @param  {int} elev which elevator
 * @param  {int} currentFloor currentFloor
 * @param  {int} stop new stop (Optional)
 */
func (c *ElevatorController) UpdateElevator(params... int) {
    //elev int, stop int, currentFloor int) {
    //error checking
    if len(params) < 2 || len(params) > 3 {
        return
    }
    if params[0] >= 0 && params[0] < len(c.elevators) {
        if params[1] <= c.numFloors && params[1] > 0 {
            c.elevators[params[0]].currentFloor = params[1]
            if len(params) == 3 {
                c.elevators[params[0]].NewStop(params[2])
            }
        }
    }
}

/*GetSystemStatus (function)
 * Receive information on the state of all elevators
 * @return  {[]int} elevPos positions of elevators
 * @return  {[]int} elevPas number of passengers in elevators
 * @return  {[]int} elevDir direction of elevators
 * @return  {[][]int} elevDes desinations for all elevators
 */
func (c *ElevatorController) GetSystemStatus() (elevPos []int, elevPas []int, elevDir []Direction, elevDes [][]int) {
    for _, e := range c.elevators {
        elevPos = append(elevPos, e.currentFloor)
        elevPas = append(elevPas, len(e.passengers))
        elevDir = append(elevDir, e.dir)
        elevDes = append(elevDes, e.stops)
    }
    return
}
////////////////////////////////// END OF API //////////////////////////////////


////////////////////////////// PRIVATE FUNCTIONS ///////////////////////////////
/*assignPassengers (function)
 * Assign elevators to pick up passengers added in the time step
 */
func (c *ElevatorController) assignPassengers() {
    //for all unassigned passengers
    for idx := range c.unassignedPassengers {
        //inital min distance is greater than max possible travel
        minDist := c.numFloors + 1
        var elev *Elevator
        for i  := range c.elevators {
            //calculate travel distance per elevator
            var dist int
            if len(c.elevators[i].stops) == 0 || c.elevators[i].stops == nil {
                dist = -1
            } else {
                dist = c.elevators[i].Distance(c.unassignedPassengers[idx])
            }
            //find min (check for -2 since distance returns -1 if elev no moving)
            if dist < minDist && dist > -2 {
                elev = &c.elevators[i]
                minDist = dist
            }
        }
        //if avalible elevator then add passenger start location to requested stops
        if minDist < c.numFloors * 2 {
            elev.NewStop(c.unassignedPassengers[idx].start)
            //add passenger to the list of passengers waiting for elevator
            c.waitingPassengers = append(c.waitingPassengers, c.unassignedPassengers[idx])
        }
    }
    //there should not be a case where a passenger is not assigned, so can nil
    //unassigned people to prevent double assignment
    c.unassignedPassengers = nil
}

/*moveElevators (function)
 * move all elevators in the system and pick up and offload passengers
 */
//TODO: refactor
func (c *ElevatorController) moveElevators() {
    //for all elebators
    for i := range c.elevators {
        //check if elevator has somewhere to go
        if len(c.elevators[i].stops) != 0 {
            //check if elevator has arrived at its next stop in the previous time step
            if c.elevators[i].stops[0] == c.elevators[i].currentFloor {
                //remove the stop from the list of stops
                c.elevators[i].stops = append(c.elevators[i].stops[:0], c.elevators[i].stops[1:]...)
                c.elevators[i].OpenDoors()
                c.elevators[i].PassengerExit();
                //if people waiting on this floor pass slice to passengerEnter
                if p := PeopleOnThisFloor(c.elevators[i].currentFloor, &c.waitingPassengers); p != nil {
                    c.elevators[i].PassengerEnter(&p)
                }
                c.elevators[i].CloseDoors()
            }
        }
        c.elevators[i].Move()
    }
}

///////////////////////////////////////////////////////////////////////////////


////////////////////////////// HELPER FUNCTIONS ///////////////////////////////

/*PeopleOnThisFloor (function)
 * find the people who are waiting on a particular floor
 * Physically removes people on floor from param slice
 * @param  {int} floor floor in question
 * @param  {[]Passenger} slice of Passengers waiting
 * @return  {[]Passenger} Passengers waiting on this floor
 */
func PeopleOnThisFloor(floor int, people *[]Passenger) ([]Passenger) {
    var peopleOnFloor []Passenger
    for i := 0; i < len(*people); i++ {
        if (*people)[i].start == floor {
            peopleOnFloor = append(peopleOnFloor, (*people)[i])
            *people = append((*people)[:i], (*people)[i+1:]...)
        }
    }
    return peopleOnFloor
}

////////////////////////////////////////////////////////////////////////////////
