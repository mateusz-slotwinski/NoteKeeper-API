package services

import (
	context "context"
	time "time"

	jwt "github.com/golang-jwt/jwt"
	bson "go.mongodb.org/mongo-driver/bson"
	mongo "go.mongodb.org/mongo-driver/mongo"
	bcrypt "golang.org/x/crypto/bcrypt"

	Config "NoteKeeperAPI/src/config"
	Database "NoteKeeperAPI/src/database"
	Models "NoteKeeperAPI/src/database/models"
	Helpers "NoteKeeperAPI/src/helpers"
	Requests "NoteKeeperAPI/src/types/requests"
)

var AccountCollection *mongo.Collection = Database.GetCollection(Database.DB, "accounts")

type AuthService struct{}

func (v AuthService) Register(userData Requests.Register) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := AccountCollection.InsertOne(ctx, bson.D{
		{Key: "Name", Value: userData.Name},
		{Key: "Email", Value: userData.Email},
		{Key: "Login", Value: userData.Login},
		{Key: "Password", Value: userData.Password},
	})
	Helpers.PrintError(err)
}

func (v AuthService) Login(userData Requests.Login) *Models.Account {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var User Models.Account
	err := AccountCollection.FindOne(ctx, bson.D{
		{Key: "Login", Value: userData.Login},
	}).Decode(&User)

	if err == mongo.ErrNoDocuments {
		return nil
	}
	Helpers.PrintError(err)

	err = bcrypt.CompareHashAndPassword([]byte(User.Password), []byte(userData.Password))
	if err != nil {
		return nil
	}
	Helpers.PrintError(err)

	return &User
}

func (v AuthService) ValidateRegister(userData Requests.Register) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	checkEmail, err := AccountCollection.CountDocuments(ctx, bson.M{"Email": userData.Email})
	Helpers.PrintError(err)

	checkLogin, err := AccountCollection.CountDocuments(ctx, bson.M{"Login": userData.Login})
	Helpers.PrintError(err)

	return checkEmail == 0 && checkLogin == 0
}

func (v AuthService) CreateToken(data interface{}) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": data,
	})

	tokenString, err := token.SignedString([]byte(Config.JWT_Secret))
	Helpers.PrintError(err)

	return tokenString
}

func (v AuthService) HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	Helpers.PrintError(err)
	return string(hashedPassword)
}
