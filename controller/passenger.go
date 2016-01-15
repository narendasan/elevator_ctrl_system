package controller

/*Passenger struct
 * Contains:
 *  	dir: direction the passenger wants to go
 *  	start: where the passenger is to be picked up
 *      dest: where the passenger wants to go
 */
type Passenger struct {
    dir Direction
    start int
    dest int
}

/*NewPassenger (function)
 * Elevator Constructor
 * @param {int} startFloor starting floor
 * @param {int} destFloor destination floor
 * @returns {*Passenger} a pointer to a new passenger
 */
func NewPassenger(startFloor int, destFloor int) (*Passenger) {
    p := new(Passenger)
    if startFloor > destFloor {
        p.dir = down;
    } else {
        p.dir = up;
    }
    p.start = startFloor
    p.dest = destFloor
    return p
}
