package cli_commands

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	models "gitlab.ozon.dev/daker255/homework-8/internal/app/models"
	service "gitlab.ozon.dev/daker255/homework-8/internal/app/services"
	"log"
	"os"
	"strconv"
)

var _ Command = (*OrderCommand)(nil)

type OrderCommand struct {
	ctx          context.Context
	orderService *service.OrderService
}

func NewOrderCommand(ctx context.Context, orderService *service.OrderService) *OrderCommand {
	return &OrderCommand{
		ctx:          ctx,
		orderService: orderService,
	}
}

func (oc *OrderCommand) Exec() {
	fmt.Println("order entity typed. Choose method you want use: create/getAll/getByID/updateStatus/delete ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	method := scanner.Text()

	oc.handleOrdersRepositoryCommands(oc.ctx, method)
}

func (oc *OrderCommand) Info() string {
	return fmt.Sprintf("order service cli\n\t%s\n\t%s\n\t%s\n\t%s\n\t%s\n",
		"create - Создать заказ",
		"getAll - Получить все заказы",
		"getOrderByID - Получить заказ по ID",
		"updateOrderStatus - Обновить статус заказа по ID",
		"deleteOrder - Удалить заказ",
	)
}

func (oc *OrderCommand) handleOrdersRepositoryCommands(ctx context.Context, method string) {
	switch method {
	case models.CreateOrder:
		fmt.Println("create method called for order")

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Please enter user ID: ")
		scanner.Scan()
		id := scanner.Text()

		userID, err := models.ParseValueToUserID(id)
		if err != nil {
			log.Fatalf("Invalid ID %s \n", id)
		}

		fmt.Print("Please enter product name: ")
		scanner.Scan()
		productName := scanner.Text()

		fmt.Print("Please enter quantity of product(integer): ")
		scanner.Scan()
		quantity := scanner.Text()

		parsedQuantity, err := strconv.Atoi(quantity)
		if err != nil {
			log.Fatalf("Invalid value \n")
		}

		orderID, err := oc.orderService.CreateOrder(ctx, userID, models.ProductName(productName), models.Quantity(parsedQuantity))
		if err != nil {
			log.Fatalf("error occured in ordersRepo.CreateOrder with err %s", err)
			return
		}
		fmt.Print("New order with id: ", orderID)

	case models.GetAllOrders:
		fmt.Println("GetAllOrders method called for order")

		orders, err := oc.orderService.GetAll(ctx)
		if err != nil {
			log.Fatalf("error occured in ordersRepo.GetAll with err %s", err)
			return
		}
		ordersMarshalled, _ := json.Marshal(orders)
		fmt.Printf("%s \n", string(ordersMarshalled))

	case models.GetOrderByID:
		fmt.Println("GetOrderByID method called for order")

		scanner := bufio.NewScanner(os.Stdin)
		fmt.Print("Please enter order ID: ")
		scanner.Scan()
		id := scanner.Text()

		orderID, err := models.ParseValueToOrderID(id)
		if err != nil {
			log.Fatalf("Invalid ID %s \n", id)
		}

		order, err := oc.orderService.GetByID(ctx, orderID)
		if err != nil {
			log.Fatalf("error occured in usersRepo.GetByID with err %s", err)
			return
		}

		orderMarshalled, _ := json.Marshal(&order)
		fmt.Printf("%s \n", string(orderMarshalled))

	case models.UpdateOrderStatus:
		fmt.Println("UpdateOrderStatus method called for order")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Print("Please enter order ID: ")
		scanner.Scan()
		id := scanner.Text()

		orderID, err := models.ParseValueToOrderID(id)
		if err != nil {
			log.Fatalf("Invalid ID %s \n", id)
		}

		fmt.Println("Please enter new order status:")
		scanner.Scan()
		status := scanner.Text()

		_, err = oc.orderService.UpdateOrderStatus(ctx, orderID, models.OrderStatus(status))
		if err != nil {
			log.Fatalf("error occured in ordersRepo.UpdateOrderStatus with err %s", err)
			return
		}

		fmt.Printf("Status of order with ID %d successfully changed\n", orderID)

	case models.DeleteOrder:
		fmt.Println("DeleteOrder method called for order")

		scanner := bufio.NewScanner(os.Stdin)

		fmt.Println("Please enter order ID to delete:")
		scanner.Scan()
		id := scanner.Text()

		orderID, err := models.ParseValueToOrderID(id)
		if err != nil {
			log.Fatalf("Invalid ID %s \n", id)
		}

		_, err = oc.orderService.DeleteOrder(ctx, orderID)
		if err != nil {
			log.Fatalf("error occured in ordersRepo.Delete with err %s", err)
			return
		}
		fmt.Printf("Order with ID %d successfully removed\n", orderID)
	}

}
