package main

import (
    "fmt";
    "bufio";
    "strings";
    "os";
    "strconv";
    "./controller";
)

func main() {
    if len(os.Args) == 2 &&  os.Args[1] == "cli"{
         cli();
    } else {
        testFile(os.Args[1]);
    }
}

func testFile(path string){
    inFile, err := os.Open(path)
    if err {
        fmt.Errorf("INVALID PATH")
        return 
    }
    defer inFile.Close()
    scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)
    i := 0
    var c *controller.ElevatorController
    for scanner.Scan() {
        str := scanner.Text()
        stringSlice := strings.Split(str, " ")
        if i == 0 {
            f,_ := strconv.Atoi(stringSlice[0])
            e,_ := strconv.Atoi(stringSlice[1])
            c = controller.NewElevatorController(f, e)
        } else if (stringSlice[0] == "pass") {
            if len(stringSlice) == 3 {
                s,_ := strconv.Atoi(stringSlice[1])
                f,_ := strconv.Atoi(stringSlice[2])
                c.CallElevator(*(controller.NewPassenger(s, f)))
            } else {
                fmt.Errorf("invalid command: %s", str)
            }
        } else if (stringSlice[0] == "step") {
            c.TimeStep()
        } else if (stringSlice[0] == "status") {
            printStatus(c.GetSystemStatus())
        } else if (stringSlice[0] == "update") {
            if len(stringSlice) == 4 {
                e,_ := strconv.Atoi(stringSlice[1])
                f,_ := strconv.Atoi(stringSlice[2])
                s,_ := strconv.Atoi(stringSlice[1])
                c.UpdateElevator(e, f, s)
            } else if len(stringSlice) == 3 {
                e,_ := strconv.Atoi(stringSlice[1])
                f,_ := strconv.Atoi(stringSlice[2])
                c.UpdateElevator(e, f)
            } else {
                fmt.Errorf("invalid command: %s", str)
            }
        }
        i++
    }
}

func cli() {
    fmt.Printf("Number of Floors:")
    var f int
    fmt.Scanf("%d\n", &f)
    var e int
    fmt.Printf("Number of Elevators:")
    fmt.Scanf("%d\n", &e)
    c := controller.NewElevatorController(f, e)
    t := 0
    fmt.Printf("\n")
    for {
        fmt.Printf("========= Time %d =========\n", t)
        fmt.Printf("Number of New Passengers (or type enter for no new):")
        var p int
        fmt.Scanf("%d\n", &p)
        for i := 0; i < p; i++ {
            fmt.Printf("Start Floor for Passanger %d: ", i)
            var f int
            fmt.Scanf("%d\n", &f)
            fmt.Printf("End Floor for Passanger %d: ", i)
            var d int
            fmt.Scanf("%d\n", &d)
            c.CallElevator(*(controller.NewPassenger(f,d)))
            fmt.Printf("\n")
        }
        fmt.Printf("\n")
        c.TimeStep()
        printStatus(c.GetSystemStatus())
        t++
    }
}

/*printStatus(function)
 * Formats output of GetSystemStatus
 * @param  {[]int} elevPos positions of elevators
 * @param  {[]int} elevPas number of passengers in elevators
 * @param  {[]int} elevDir direction of elevators
 * @param  {[][]int} elevDes desinations for all elevators
 */
func printStatus(elevPos []int, elevPas []int, elevDir []controller.Direction, elevDest [][]int) {
    for idx := range elevPos {
        fmt.Printf("------- Elevator %d --------\n", idx)
        fmt.Printf("Current Floor: %d\n", elevPos[idx])
        fmt.Printf("Number of Passengers: %d\n", elevPas[idx])
        fmt.Printf("Direction: %d\n", elevDir[idx])
        fmt.Printf("Requested Stops: ")
        fmt.Println(elevDest[idx])
        fmt.Printf("\n")
    }
}
