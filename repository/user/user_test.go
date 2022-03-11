package user

// import (
// 	"go-mongodb/config"
// 	"go-mongodb/entities"
// 	"go-mongodb/util"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// var (
// 	name     = "Test1"
// 	password = "Test123"
// 	email    = "Test1@mail.com"
// )

// func TestSetup(t *testing.T) {
// 	//load config if available or set to default
// 	config := config.GetConfigTest()

// 	//initialize database connection based on given config
// 	db := util.MysqlDriver(config)

// 	// cleaning data before testing
// 	db.Migrator().DropTable(&entities.User{})
// 	db.AutoMigrate(&entities.User{})

// 	todoRepository := New(db)

// 	UserCreate(t, todoRepository)
// 	UserGet(t, todoRepository)
// 	UserGetByID(t, todoRepository)
// 	UserUpdate(t, todoRepository)
// 	UserDelete(t, todoRepository)
// }

// func InsertDataUserForGetUsers(repository *UserRepository) error {
// 	user := entities.User{
// 		Name:     "Alta",
// 		Password: "123",
// 		Email:    "alta@gmail.com",
// 	}

// 	var err error
// 	if _, err = repository.Create(user); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func UserCreate(t *testing.T, repository *UserRepository) {
// 	t.Run("UserCreate", func(t *testing.T) {
// 		user := entities.User{
// 			Name:     name,
// 			Email:    email,
// 			Password: password,
// 		}
// 		result, _ := repository.Create(user)
// 		assert.Equal(t, result.Email, user.Email)
// 	})
// }

// func UserGet(t *testing.T, repository *UserRepository) {
// 	t.Run("UserGet", func(t *testing.T) {
// 		result, _ := repository.Get()
// 		assert.Equal(t, len(result), 1)
// 		assert.Equal(t, result[0].Email, email)
// 	})
// }

// func UserGetByID(t *testing.T, repository *UserRepository) {
// 	t.Run("UserGetByID", func(t *testing.T) {
// 		result, _ := repository.GetByID(1)
// 		assert.Equal(t, len(result), 1)
// 		assert.Equal(t, result[0].Email, email)
// 	})
// }

// func UserUpdate(t *testing.T, repository *UserRepository) {
// 	t.Run("UserUpdate", func(t *testing.T) {
// 		user := entities.User{
// 			Name:     "Update1",
// 			Email:    email,
// 			Password: password,
// 		}
// 		_ = repository.Update(user, 1)
// 		result, _ := repository.GetByID(1)
// 		assert.Equal(t, result[0].Name, user.Name)
// 	})
// }

// func UserDelete(t *testing.T, repository *UserRepository) {
// 	t.Run("UserDelete", func(t *testing.T) {
// 		_ = repository.Delete(1)
// 		result, _ := repository.GetByID(1)
// 		assert.Equal(t, len(result), 0)
// 	})
// }
