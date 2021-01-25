package main

import (
	"fmt"

	"github.com/alvinarthas/grpc-test-init/model"
)

func main() {
	var (
		user1 = &model.User{
			Id:       "u001",
			Name:     "Sylvana Windrunner",
			Password: "f0r Th3 H0rD3",
			Gender:   model.UserGender_FEMALE,
		}

		userList = &model.UserList{
			List: []*model.User{
				user1,
			},
		}

		garage1 = &model.Garage{
			Id:   "g001",
			Name: "Kalimdor",
			Coordinate: &model.GarageCoordinate{
				Latitude:  23.2212847,
				Longitude: 53.22033123,
			},
		}

		garageList = &model.GarageList{
			List: []*model.Garage{
				garage1,
			},
		}

		garageListByUser = &model.GarageListByUser{
			List: map[string]*model.GarageList{
				user1.Id: garageList,
			},
		}
	)

	// =========== original
	fmt.Printf("# ==== Original\n       %#v \n", user1)

	// =========== as string
	fmt.Printf("# ==== As String\n       %v \n", user1.String())

	fmt.Println(userList)
	fmt.Println(garageListByUser)
}
