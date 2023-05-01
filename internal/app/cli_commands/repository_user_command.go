package cli_commands

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
)

var _ Command = (*UserCommand)(nil)

type UserCommand struct {
	ctx         context.Context
	userService *service.UserService
}

func NewUserCommand(ctx context.Context, userService *service.UserService) *UserCommand {
	return &UserCommand{
		ctx:         ctx,
		userService: userService,
	}
}
func (uc *UserCommand) Exec() {
	fmt.Println("user entity typed. Choose method you want use: create/getAll/getByID/updateUsername/updateEmail/delete")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	method := scanner.Text()

	uc.handleUsersRepositoryCommands(uc.ctx, method)
}

func (uc *UserCommand) Info() string {
	return fmt.Sprintf("user service cli\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n",
		"create - Создать пользователя(зарегистрировать)",
		"getAll - Получить всех пользователей",
		"getUserByID - Получить пользователя по ID",
		"updateUsername - Обновить username пользователя по ID",
		"updateEmail - Обновить email пользователя по ID",
		"delete - Удалить пользователя",
	)
}

func (uc *UserCommand) handleUsersRepositoryCommands(ctx context.Context, method string) {
	switch method {
	case models.CreateUser:
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Please enter username: ")
		scanner.Scan()
		username := scanner.Text()

		fmt.Print("Please enter user email: ")
		scanner.Scan()
		email := scanner.Text()

		id, err := uc.userService.CreateUser(ctx,
			models.Username(username),
			models.UserEmail(email))
		if err != nil {
			log.Fatalf("error occured in usersRepo.CreateUser with err %s", err)
			return
		}
		fmt.Println("New user with id: ", id)

	case models.GetAllUsers:
		fmt.Println("GetAllUsers method called for user")

		users, err := uc.userService.GetAll(ctx)
		if err != nil {
			log.Fatalf("error occured in usersRepo.GetAll with err %s", err)
			return
		}
		usersMarshalled, _ := json.Marshal(users)

		fmt.Printf("%s \n", string(usersMarshalled))

	case models.GetUserByID:
		fmt.Println("GetUserByID method called for user")

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Please enter user ID: ")
		scanner.Scan()
		id := scanner.Text()
		userID, err := models.ParseValueToUserID(id)
		if err != nil {
			log.Fatalf("Invalid ID %s \n", id)
		}

		user, err := uc.userService.GetByID(ctx, userID)
		if err != nil {
			log.Fatalf("error occured in usersRepo.GetByID with err %s", err)
			return
		}

		userMarshalled, _ := json.Marshal(&user)
		fmt.Printf("%s \n", string(userMarshalled))

	case models.UpdateUsername:
		fmt.Println("UpdateUsername method called for user")
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Please enter user ID: ")
		scanner.Scan()
		id := scanner.Text()

		userID, err := models.ParseValueToUserID(id)
		if err != nil {
			log.Fatalf("Invalid ID %s \n", id)
		}

		fmt.Println("Please enter username: ")
		scanner.Scan()
		username := scanner.Text()

		_, err = uc.userService.UpdateUsername(ctx, userID, models.Username(username))
		if err != nil {
			log.Fatalf("error occured in usersRepo.UpdateUsername with err %s", err)
			return
		}
		fmt.Printf("Username of user with ID %d successfully changed\n", userID)

	case models.UpdateEmail:
		fmt.Println("UpdateEmail method called for user")
		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Please enter user ID: ")
		scanner.Scan()
		id := scanner.Text()

		userID, err := models.ParseValueToUserID(id)
		if err != nil {
			log.Fatalf("Invalid ID %s \n", id)
		}

		fmt.Print("Please enter email: ")
		scanner.Scan()
		email := scanner.Text()

		_, err = uc.userService.UpdateEmail(ctx, userID, models.UserEmail(email))
		if err != nil {
			log.Fatalf("error occured in usersRepo.UpdateEmail with err %s", err)
			return
		}
		fmt.Printf("Email of user with ID %d successfully changed\n", userID)

	case models.DeleteUser:
		fmt.Println("DeleteUser method called for user")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Please enter user ID: ")
		scanner.Scan()
		id := scanner.Text()

		userID, err := models.ParseValueToUserID(id)
		if err != nil {
			log.Fatalf("Invalid ID %s \n", id)
		}

		isDeleted, _ := uc.userService.DeleteUser(ctx, userID)
		if !isDeleted {
			log.Fatalf("error occured in usersRepo.Delete with err %s", err)
			return
		}

		fmt.Printf("User with ID %d successfully removed\n", userID)
	}
}
