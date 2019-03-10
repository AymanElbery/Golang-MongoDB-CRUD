package main

import (
	"fmt"
	"time"

	"gomongo/config"
	"gomongo/src/modules/user/model"
	"gomongo/src/modules/user/repository"
)

func main() {

	db, err := config.GetMongoDB()

	if err != nil {
		fmt.Println(err)
	}

	userRepository := repository.NewUserRepositoryMongo(db, "users")

	//saveUser(userRepository)
	//updateUser("U1", userRepository)
	//deleteUser("U1", userRepository)
	//getUser("U1", userRepository)
	getAllUser(userRepository)
}

// saveUser ...
func saveUser(userRepository repository.UserRepository) {
	var p model.User
	p.ID = "U2"
	p.FirstName = "Ahmed"
	p.LastName = "Mohamed"
	p.Email = "ahmed.bery@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := userRepository.Save(&p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User Saved")
	}
}

// updateUser ...
func updateUser(id string, userRepository repository.UserRepository) {
	var p model.User
	p.ID = "U1"
	p.FirstName = "Ayman"
	p.LastName = "Elbery"
	p.Email = "ayman.bery@gmail.com"
	p.Password = "123456"
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()

	err := userRepository.Update(id, &p)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User Updated")
	}
}

// deleteUser ...
func deleteUser(id string, userRepository repository.UserRepository) {

	err := userRepository.Delete(id)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("User Deleted")
	}
}

// getUser ...
func getUser(id string, userRepository repository.UserRepository) {

	user, err := userRepository.FindByID(id)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("User:")
	fmt.Println(user.ID)
	fmt.Println(user.FirstName)
	fmt.Println(user.LastName)
	fmt.Println(user.Email)
	fmt.Println(user.CreatedAt)

}

// getAllUser ...
func getAllUser(userRepository repository.UserRepository) {

	users, err := userRepository.FindAll()

	if err != nil {
		fmt.Println(err)
	}

	for _, user := range users {
		fmt.Println("==============")
		fmt.Println(user.ID)
		fmt.Println(user.FirstName)
		fmt.Println(user.LastName)
		fmt.Println(user.Email)
		fmt.Println(user.CreatedAt)
	}

}
