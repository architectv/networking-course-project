// +build bdd_e2e

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/architectv/networking-course-project/backend/pkg/handlers"
	"github.com/architectv/networking-course-project/backend/pkg/models"
	"github.com/architectv/networking-course-project/backend/pkg/repositories"
	"github.com/architectv/networking-course-project/backend/pkg/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/jmoiron/sqlx"
	. "github.com/smartystreets/goconvey/convey"
)

const (
	UsernameTestDB = "postgres"
	PasswordTestDB = "1234"
	HostTestDB     = "localhost"
	PortTestDB     = "5432"
	DBnameTestDB   = "yak_test_db"
	SslmodeTestDB  = "disable"
)

func openTestDatabase() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		HostTestDB, PortTestDB, UsernameTestDB, DBnameTestDB, PasswordTestDB, SslmodeTestDB))
	return db, err
}

func prepareTestDatabase() (*sqlx.DB, error) {
	db, err := openTestDatabase()
	schema, err := ioutil.ReadFile("scripts/init.sql")
	if err != nil {
		log.Fatal(err)
	}
	db.MustExec(string(schema))
	return db, err
}

var LogDisable = false

func Test_E2E_App(t *testing.T) {
	db, err := prepareTestDatabase()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}
	defer db.Close()

	repos := repositories.NewRepository(db)
	services := services.NewService(repos)
	handlers := handlers.NewHandler(services)

	app := fiber.New()
	if LogDisable {
		app.Use(logger.New(logger.Config{Output: ioutil.Discard}))
	} else {
		app.Use(logger.New())
	}

	handlers.RegisterHandlers(app)

	// Example
	Convey("Given params", t, func() {
		// Given params
		Convey("When to do", func() {
			// When to do
			Convey("Then should be...", func() {
				// Then should be...
			})
		})
	})

	// Register
	const (
		ExpectedUserId        = "5"
		ExpectedProjectId     = "4"
		ExpectedBoardId       = "4"
		ExpectedListId        = "4"
		ExpectedTaskId1       = "7"
		ExpectedTaskId2       = "8"
		InputNickname         = "e2eUser"
		InputEmail            = "e2eUser@test.com"
		InputPassword         = "qwerty"
		InputTitle            = "E2E Test Project"
		InputDescription      = "LW2 testing"
		InputBoardTitle       = "E2E Test Board"
		InputListTitle        = "todo list"
		InputTaskTitle1       = "Task 1"
		InputTaskDescription1 = "some description"
		InputTaskTitle2       = "Task 2"
		InputTaskDescription2 = "some description"
		Uid                   = 5
		Lid                   = 4
	)
	Convey("Given input params for new user", t, func() {
		expectedStatus := fiber.StatusOK
		expectedUserId := ExpectedUserId
		expectedBody := fmt.Sprintf(`{"code":200,"message":"OK","data":{"uid":%s}}`, expectedUserId)
		inputNickname := InputNickname
		inputPassword := InputPassword
		inputEmail := InputEmail
		inputBody := fmt.Sprintf(`{"nickname": "%s", "email": "%s", "password": "%s"}`, inputNickname, inputEmail, inputPassword)

		Convey("When post signup request", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/v1/users/signup",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})

	token := ""
	// Login
	Convey("Given input params for login", t, func() {
		expectedStatus := fiber.StatusOK
		inputNickname := InputNickname
		inputPassword := InputPassword
		inputBody := fmt.Sprintf(`{"nickname": "%s", "password": "%s"}`, inputNickname, inputPassword)

		Convey("When post signin request", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/v1/users/signin",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				apiResp := &models.ApiResponse{}
				err = json.Unmarshal(body, apiResp)
				So(err, ShouldBeNil)
				So(apiResp.Code, ShouldEqual, expectedStatus)
				token = apiResp.Data.(map[string]interface{})["token"].(string)
			})
		})
	})

	// Get user info
	Convey("Given nothing", t, func() {
		expectedStatus := fiber.StatusOK

		Convey("When get user request", func() {
			req := httptest.NewRequest(
				"GET",
				"/api/v1/users/",
				nil,
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				apiResp := &models.ApiResponse{}
				err = json.Unmarshal(body, apiResp)
				So(err, ShouldBeNil)
				So(apiResp.Code, ShouldEqual, expectedStatus)
				userMap := apiResp.Data.(map[string]interface{})["user"].(map[string]interface{})
				userJSON, err := json.Marshal(userMap)
				So(err, ShouldBeNil)
				user := &models.User{}
				err = json.Unmarshal(userJSON, user)
				So(err, ShouldBeNil)
				uid, err := strconv.Atoi(ExpectedUserId)
				So(err, ShouldBeNil)
				So(uid, ShouldEqual, user.Id)
			})
		})
	})

	// Create project
	Convey("Given input params for project", t, func() {
		expectedStatus := fiber.StatusOK
		expectedProjectId := ExpectedProjectId
		expectedBody := fmt.Sprintf(`{"code":200,"message":"OK","data":{"projectId":%s}}`, expectedProjectId)
		inputTitle := InputTitle
		inputDescription := InputDescription
		inputBody := fmt.Sprintf(`{"title":"%s", "description":"%s"}`, inputTitle, inputDescription)

		Convey("When post project request", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/v1/projects",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})

	// Get project
	Convey("Given nothing", t, func() {
		expectedStatus := fiber.StatusOK

		Convey("When get project request", func() {
			req := httptest.NewRequest(
				"GET",
				"/api/v1/projects/"+ExpectedProjectId,
				nil,
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				apiResp := &models.ApiResponse{}
				err = json.Unmarshal(body, apiResp)
				So(err, ShouldBeNil)
				So(apiResp.Code, ShouldEqual, expectedStatus)
				projectMap := apiResp.Data.(map[string]interface{})["project"].(map[string]interface{})
				projectJSON, err := json.Marshal(projectMap)
				So(err, ShouldBeNil)
				project := &models.Project{}
				err = json.Unmarshal(projectJSON, project)
				So(err, ShouldBeNil)
				So(project.Title, ShouldEqual, InputTitle)
				So(project.Description, ShouldEqual, InputDescription)
			})
		})
	})

	// Create board
	Convey("Given input params for board", t, func() {
		expectedStatus := fiber.StatusOK
		expectedBoardId := ExpectedBoardId
		expectedBody := fmt.Sprintf(`{"code":200,"message":"OK","data":{"boardId":%s}}`, expectedBoardId)
		inputBoardTitle := InputBoardTitle
		inputBody := fmt.Sprintf(`{"title":"%s"}`, inputBoardTitle)

		Convey("When post board request", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/v1/projects/"+ExpectedProjectId+"/boards",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})

	// Get board
	Convey("Given nothing", t, func() {
		expectedStatus := fiber.StatusOK

		Convey("When get board request", func() {
			req := httptest.NewRequest(
				"GET",
				"/api/v1/projects/"+ExpectedProjectId+"/boards/"+ExpectedBoardId,
				nil,
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				apiResp := &models.ApiResponse{}
				err = json.Unmarshal(body, apiResp)
				So(err, ShouldBeNil)
				So(apiResp.Code, ShouldEqual, expectedStatus)
				boardMap := apiResp.Data.(map[string]interface{})["board"].(map[string]interface{})
				boardJSON, err := json.Marshal(boardMap)
				So(err, ShouldBeNil)
				board := &models.Board{}
				err = json.Unmarshal(boardJSON, board)
				So(err, ShouldBeNil)
				So(board.Title, ShouldEqual, InputBoardTitle)
				So(board.OwnerId, ShouldEqual, Uid)
				pid, err := strconv.Atoi(ExpectedProjectId)
				So(err, ShouldBeNil)
				So(board.ProjectId, ShouldEqual, pid)
			})
		})
	})

	// Create list
	Convey("Given input params for list", t, func() {
		expectedStatus := fiber.StatusOK
		expectedListId := ExpectedListId
		expectedBody := fmt.Sprintf(`{"code":200,"message":"OK","data":{"listId":%s}}`, expectedListId)
		inputListTitle := InputListTitle
		inputBody := fmt.Sprintf(`{"title":"%s"}`, inputListTitle)

		Convey("When post list request", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/v1/projects/"+ExpectedProjectId+"/boards/"+ExpectedBoardId+"/lists",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})

	// Create task 1
	Convey("Given input params for task 1", t, func() {
		expectedStatus := fiber.StatusOK
		expectedTaskId1 := ExpectedTaskId1
		expectedBody := fmt.Sprintf(`{"code":200,"message":"OK","data":{"taskId":%s}}`, expectedTaskId1)
		inputTaskTitle1 := InputTaskTitle1
		inputTaskDescription1 := "some description"
		inputBody := fmt.Sprintf(`{"title":"%s", "description":"%s"}`, inputTaskTitle1, inputTaskDescription1)

		Convey("When post task 1 request", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/v1/projects/"+ExpectedProjectId+"/boards/"+ExpectedBoardId+"/lists/"+ExpectedListId+"/tasks",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})

	// Create task 2
	Convey("Given input params for task 2", t, func() {
		expectedStatus := fiber.StatusOK
		expectedTaskId2 := ExpectedTaskId2
		expectedBody := fmt.Sprintf(`{"code":200,"message":"OK","data":{"taskId":%s}}`, expectedTaskId2)
		inputTaskTitle2 := InputTaskTitle2
		inputTaskDescription2 := "some description"
		inputBody := fmt.Sprintf(`{"title":"%s", "description":"%s"}`, inputTaskTitle2, inputTaskDescription2)

		Convey("When post task 2 request", func() {
			req := httptest.NewRequest(
				"POST",
				"/api/v1/projects/"+ExpectedProjectId+"/boards/"+ExpectedBoardId+"/lists/"+ExpectedListId+"/tasks",
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})

	// Update task 2
	Convey("Given input params for upd task 2", t, func() {
		expectedStatus := fiber.StatusOK
		expectedBody := fmt.Sprintf(`{"code":200,"message":"OK","data":{}}`)
		inputBody := fmt.Sprintf(`{"position":0}`)

		Convey("When put task 2 request", func() {
			req := httptest.NewRequest(
				"PUT",
				"/api/v1/projects/"+ExpectedProjectId+"/boards/"+ExpectedBoardId+"/lists/"+ExpectedListId+"/tasks/"+ExpectedTaskId2,
				bytes.NewBufferString(inputBody),
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})

	// Get tasks
	Convey("Given nothing", t, func() {
		expectedStatus := fiber.StatusOK
		inputTaskTitle := []string{InputTaskTitle2, InputTaskTitle1}
		inputTaskDescription := []string{InputTaskDescription2, InputTaskDescription1}
		taskCount := 2

		Convey("When get tasks request", func() {
			req := httptest.NewRequest(
				"GET",
				"/api/v1/projects/"+ExpectedProjectId+"/boards/"+ExpectedBoardId+"/lists/"+ExpectedListId+"/tasks",
				nil,
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				apiResp := &models.ApiResponse{}
				err = json.Unmarshal(body, apiResp)
				So(err, ShouldBeNil)
				So(apiResp.Code, ShouldEqual, expectedStatus)
				tasksInterface := apiResp.Data.(map[string]interface{})["tasks"].([]interface{})
				var tasksSlice []map[string]interface{}
				for _, item := range tasksInterface {
					tasksSlice = append(tasksSlice, item.(map[string]interface{}))
				}
				So(len(tasksSlice), ShouldEqual, taskCount)
				var tasks []*models.Task
				for _, item := range tasksSlice {
					taskJSON, err := json.Marshal(item)
					So(err, ShouldBeNil)
					task := &models.Task{}
					err = json.Unmarshal(taskJSON, task)
					So(err, ShouldBeNil)
					tasks = append(tasks, task)
				}
				lid, err := strconv.Atoi(ExpectedListId)
				So(err, ShouldBeNil)
				for i, item := range tasks {
					So(item.Position, ShouldEqual, i)
					So(item.ListId, ShouldEqual, lid)
					So(item.Title, ShouldEqual, inputTaskTitle[i])
					So(item.Description, ShouldEqual, inputTaskDescription[i])
				}
			})
		})
	})

	// Delete task 2
	Convey("Given nothing", t, func() {
		expectedStatus := fiber.StatusOK
		expectedBody := fmt.Sprintf(`{"code":200,"message":"OK","data":{}}`)

		Convey("When delete task 2 request", func() {
			req := httptest.NewRequest(
				"DELETE",
				"/api/v1/projects/"+ExpectedProjectId+"/boards/"+ExpectedBoardId+"/lists/"+ExpectedListId+"/tasks/"+ExpectedTaskId2,
				nil,
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				So(string(body), ShouldEqual, expectedBody)
			})
		})
	})

	// Get tasks after delete task 2
	Convey("Given nothing", t, func() {
		expectedStatus := fiber.StatusOK
		inputTaskTitle := []string{InputTaskTitle1}
		inputTaskDescription := []string{InputTaskDescription1}
		taskCount := 1

		Convey("When get tasks request", func() {
			req := httptest.NewRequest(
				"GET",
				"/api/v1/projects/"+ExpectedProjectId+"/boards/"+ExpectedBoardId+"/lists/"+ExpectedListId+"/tasks",
				nil,
			)
			req.Header.Set("Authorization", "Bearer "+token)
			req.Header.Set("Content-type", "application/json")

			resp, err := app.Test(req, -1)
			body, bodyErr := ioutil.ReadAll(resp.Body)

			Convey("Then status should be OK", func() {
				So(err, ShouldBeNil)
				So(bodyErr, ShouldBeNil)
				So(resp.StatusCode, ShouldEqual, expectedStatus)
				apiResp := &models.ApiResponse{}
				err = json.Unmarshal(body, apiResp)
				So(err, ShouldBeNil)
				So(apiResp.Code, ShouldEqual, expectedStatus)
				tasksInterface := apiResp.Data.(map[string]interface{})["tasks"].([]interface{})
				var tasksSlice []map[string]interface{}
				for _, item := range tasksInterface {
					tasksSlice = append(tasksSlice, item.(map[string]interface{}))
				}
				So(len(tasksSlice), ShouldEqual, taskCount)
				var tasks []*models.Task
				for _, item := range tasksSlice {
					taskJSON, err := json.Marshal(item)
					So(err, ShouldBeNil)
					task := &models.Task{}
					err = json.Unmarshal(taskJSON, task)
					So(err, ShouldBeNil)
					tasks = append(tasks, task)
				}
				lid, err := strconv.Atoi(ExpectedListId)
				So(err, ShouldBeNil)
				for i, item := range tasks {
					So(item.Position, ShouldEqual, i)
					So(item.ListId, ShouldEqual, lid)
					So(item.Title, ShouldEqual, inputTaskTitle[i])
					So(item.Description, ShouldEqual, inputTaskDescription[i])
				}
			})
		})
	})
}
