package controller

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"project-tdd/anggunpermata/config"
	"project-tdd/anggunpermata/database"
	"project-tdd/anggunpermata/model"
	"testing"

	"github.com/labstack/echo"
)

func testGetBookController(t *testing.T, bookController echo.HandlerFunc) {
	//request
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)
	//test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: expecting %d", statusCode)
	}
	body := rec.Body.Bytes()
	var books []model.Book
	if err := json.Unmarshal(body, &books); err != nil {
		t.Error(err)
	}
	if len(books) != 1 {
		t.Errorf("Expecting One Book, got %#v instead", books)
		return
	}
	if books[0].Title != "Story Book" {
		t.Errorf("expected Story Book, got %#v instead", books[0].Title)
	}
}

func testPostBookController(t *testing.T, bookController echo.HandlerFunc) {
	// coba request
	reqBody, err := json.Marshal(map[string]string{
		"title": "Abc",
	})
	if err != nil {
		t.Error(err)
		return
	}
	req, err := http.NewRequest("POST", "/books", bytes.NewBuffer(reqBody))
	if err != nil {
		t.Error(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	e := echo.New()
	c := e.NewContext(req, rec)
	bookController(c)
	// test
	statusCode := rec.Result().StatusCode
	if statusCode != 200 {
		t.Errorf("Response is not 200: %d", statusCode)
	}
	body := rec.Body.Bytes()
	var book model.Book
	if err := json.Unmarshal(body, &book); err != nil {
		t.Error(err)
	}
	if book.Title != "Abc" {
		t.Errorf("Expected Abc, got %#v instead", book.Title)
	}
}

func TestDBGetBookController(t *testing.T) {
	//bikin db
	db, err := database.CreateDB(config.TEST_DB_CONNECTION_STRING)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Book{})
	//MUST ASK
	db.Delete(&model.Book{}, "1=1")
	m := model.NewGormBookModel(db)
	//bikin controller
	bookController := CreateGetBookController(m)
	if err != nil {
		t.Error(err)
	}
	// insert data baru
	m.Insert(model.Book{Title: "Story Book"})
	testGetBookController(t, bookController)
}

func TestDBPostBookController(t *testing.T) {
	//bikin db
	db, err := database.CreateDB(config.TEST_DB_CONNECTION_STRING)
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&model.Book{})
	db.Delete(&model.Book{}, "1=1")
	m := model.NewGormBookModel(db)
	//bikin controller
	bookController := CreatePostBookController(m)
	if err != nil {
		t.Error(err)
	}
	//test controller
	testPostBookController(t, bookController)
	db.Delete(&model.Book{}, "1=1")
}
