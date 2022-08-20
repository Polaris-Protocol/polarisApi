package controllers

import (
	"net/http"
	"polarisApi/configs"
	"polarisApi/models"
	"polarisApi/responses"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var projectsCollection *mongo.Collection = configs.GetCollection(configs.DB, "projects")
var ordersCollection *mongo.Collection = configs.GetCollection(configs.DB, "orders")
var validate = validator.New()

func UpdatedOrder(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var order models.Orders
	order_id := c.Param("orderId")
	validator_id := c.Param("validatorId")

	defer cancel()

	order_id_uint, err := strconv.ParseUint(order_id, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	validator_id_uint, err := strconv.ParseUint(validator_id, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	update := bson.M{}
	if validator_id_uint == 1 {
		update = bson.M{"valid_1": true}

	} else if validator_id_uint == 2 {

		update = bson.M{"valid_2": true}

	} else if validator_id_uint == 3 {

		update = bson.M{"valid_3": true}

	} else {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": "Invalid Request!"}})
	}

	result, err := ordersCollection.UpdateOne(ctx, bson.M{"order_id": order_id_uint}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})

}

func GetOrders(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var orders []models.Orders
	chain := c.Param("chain")

	defer cancel()

	results, err := ordersCollection.Find(ctx, bson.M{"chain": chain})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)

	for results.Next(ctx) {
		var singleOrder models.Orders

		if err = results.Decode(&singleOrder); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		orders = append(orders, singleOrder)

	}

	return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": orders}})

}

func LinkWalletToUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Users
	discord_id := c.Param("discordId")
	chain := c.Param("chain")

	defer cancel()

	discord_id_uint, err := strconv.ParseUint(discord_id, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	update := bson.M{}

	if chain == "wax" {
		update = bson.M{"wax_wallet": user.Wax_wallet}

	} else if chain == "eth" {

		update = bson.M{"eth_wallet": user.Eth_wallet}

	} else if chain == "bnb" {

		update = bson.M{"bnb_wallet": user.Bnb_wallet}

	} else if chain == "polygon" {

		update = bson.M{"polygon_wallet": user.Polygon_wallet}
	} else {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": "Chain not supported"}})
	}

	result, err := userCollection.UpdateOne(ctx, bson.M{"disc_id": discord_id_uint}, bson.M{"$set": update})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": result}})

}

func GetProjects(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var projects []models.Projects
	defer cancel()

	results, err := projectsCollection.Find(ctx, bson.M{})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	defer results.Close(ctx)
	for results.Next(ctx) {
		var singleProject models.Projects
		if err = results.Decode(&singleProject); err != nil {
			return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})
		}

		projects = append(projects, singleProject)
	}

	return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": projects}})
}

func GetProjectByServerId(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var project models.Projects
	serverId := c.Param("serverId")

	defer cancel()

	serverIdInteger, err := strconv.ParseInt(serverId, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := c.Bind(&project); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	error := projectsCollection.FindOne(ctx, bson.M{"server_id": serverIdInteger}).Decode(&project)

	if error != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": error.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": project}})

}

func GetUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Users
	userId := c.Param("discordId")

	defer cancel()

	userIdInteger, err := strconv.ParseInt(userId, 10, 64)

	if err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	error := userCollection.FindOne(ctx, bson.M{"disc_id": userIdInteger}).Decode(&user)

	if error != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": error.Error()}})
	}

	return c.JSON(http.StatusOK, responses.UserResponse{Status: http.StatusOK, Message: "success", Data: &echo.Map{"data": user}})

}

func RegisterUser(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var user models.Users
	disc_id := c.Param("discordId")

	defer cancel()
	discord_id, err := strconv.ParseUint(disc_id, 10, 64)
	//validate the request body
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})
	}

	newUser := models.Users{

		Disc_id: discord_id,
	}

	result, err := userCollection.InsertOne(ctx, newUser)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})

	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})
}

func RegisterProject(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	proj := new(models.Projects)

	server_id := c.Param("serverId")
	discord_server_id, err := strconv.ParseUint(server_id, 10, 64)

	defer cancel()

	if err := c.Bind(&proj); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(proj); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})

	}

	newProject := models.Projects{
		Remaining_giveaways: proj.Remaining_giveaways,
		Owner_disc_id:       proj.Owner_disc_id,
		Tier:                proj.Tier,
		Balance:             proj.Balance,
		Server_id:           discord_server_id,
		Chain:               proj.Chain,
	}

	result, err := projectsCollection.InsertOne(ctx, newProject)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})

	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})

}

func NewOrder(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	order := new(models.Orders)

	defer cancel()

	if err := c.Bind(&order); err != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": err.Error()}})
	}

	if validationErr := validate.Struct(order); validationErr != nil {
		return c.JSON(http.StatusBadRequest, responses.UserResponse{Status: http.StatusBadRequest, Message: "error", Data: &echo.Map{"data": validationErr.Error()}})

	}

	result, err := ordersCollection.InsertOne(ctx, order)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, responses.UserResponse{Status: http.StatusInternalServerError, Message: "error", Data: &echo.Map{"data": err.Error()}})

	}

	return c.JSON(http.StatusCreated, responses.UserResponse{Status: http.StatusCreated, Message: "success", Data: &echo.Map{"data": result}})

}

func StandardResponse(c echo.Context) error {
	return c.JSON(200, "Polaris Api")
}
