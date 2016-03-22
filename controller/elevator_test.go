package controller_test

import (
    "testing"
    "github.com/narendasan/elevator_ctrl_system/controller"
)

func TestDistance(t *testing.T) {
	// test stuff here...
	e := controller.NewElevator(1, 10)
    if e.Distance(*(controller.NewPassenger(3, 5))) != 2 {
        t.Fail()
    }
}
