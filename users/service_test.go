package users

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("USER SERVICE TESTS STARTED")
	testResult := m.Run()
	fmt.Println("USER SERVICE TESTS ENDED")
	os.Exit(testResult)
}

func TestService(t *testing.T) {
	t.Run("USERS", func(t *testing.T) {
		t.Parallel()
		t.Run("findAllByNameTest", func(t *testing.T) {
			testReceiverData := TUsers{{
				firstname: "Ark",
				age:       27,
				car: TCar{
					model: "KIA",
					cost:  2000000,
				},
			},
				{
					firstname: "Ivan",
					age:       30,
					car: TCar{
						model: "KIA",
						cost:  2200000,
					},
				},
				{
					firstname: "Alfred",
					age:       32,
					car: TCar{
						model: "BMW",
						cost:  4000000,
					},
				}}
			testParam, expectedResult := "r", TUsers{testReceiverData[0], testReceiverData[2]}

			testResult := testReceiverData.findAllByName(testParam)
			if testResult[0].firstname != expectedResult[0].firstname || testResult[1].firstname != expectedResult[1].firstname {
				t.Error("Incorrect result\n")
			}
		})
	})
}
