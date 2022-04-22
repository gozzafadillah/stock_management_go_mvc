package test

import (
	"gozzafadillah/config"
	"gozzafadillah/controllers"
	"gozzafadillah/models"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func InsertDataProduct() error {
	product := models.Product{
		ID:              1,
		Code:            "BML-001",
		Product_type_id: 2,
		Name:            "Bimoli 1 L",
		Description:     "Minyak Botol 1L Terbaik",
		Price:           14500,
		Qty:             80,
	}

	var err error
	if err = config.DB.Save(&product).Error; err != nil {
		return err
	}
	return nil
}

func TestGetAllProduct(t *testing.T) {
	defer func() {
		SuccessTest := TestCase{

			Method:         http.MethodGet,
			Name:           "Get All Product",
			Path:           "/product",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Success Get All Product",
		}

		e := InitEchoTestAPI()
		InsertDataProduct()

		req := httptest.NewRequest(SuccessTest.Method, SuccessTest.Path, nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(SuccessTest.Path)
		if assert.NoError(t, controllers.GetAllProductController(c)) {
			assert.Equal(t, SuccessTest.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), SuccessTest.ExpectResponse)
		}
	}()
	FailedTest := TestCase{

		Method:         http.MethodGet,
		Name:           "Get Fail All Product",
		Path:           "/product",
		ExpectStatus:   http.StatusBadGateway,
		ExpectResponse: "Database Empty !!",
	}

	e := InitEchoTestAPI()

	req := httptest.NewRequest(FailedTest.Method, FailedTest.Path, nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(FailedTest.Path)
	if assert.NoError(t, controllers.GetAllProductController(c)) {
		assert.Equal(t, FailedTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), FailedTest.ExpectResponse)
	}
}

func TestCreateProduct(t *testing.T) {
	defer func() {
		SuccessTest_1 := TestCase{

			Method:         http.MethodPost,
			Name:           "Create Product",
			Path:           "/product",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "New Product Success Added",
		}

		e := InitEchoTestAPI()
		reqStr := `{
			"code": "BML-001",
			"product_type": 2,
			"name": "Bimoli",
			"description": "Minyak Elit Eropa",
			"price": 14500,
			"qty": 35
		}`

		req := httptest.NewRequest(SuccessTest_1.Method, SuccessTest_1.Path, strings.NewReader(reqStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(SuccessTest_1.Path)
		if assert.NoError(t, controllers.CreateProduct(c)) {
			assert.Equal(t, SuccessTest_1.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), SuccessTest_1.ExpectResponse)

		}
	}()

	func() {
		SuccessTest_2 := TestCase{

			Method:         http.MethodPost,
			Name:           "Create Product",
			Path:           "/product",
			ExpectStatus:   http.StatusOK,
			ExpectResponse: "Product Updated",
		}

		e := InitEchoTestAPI()
		InsertDataProduct()
		reqStr := `{
			"ID":1,
			"code": "BML-001",
			"product_type": 2,
			"name": "Bimoli",
			"description": "Minyak Elit Eropa",
			"price": 14000,
			"qty": 50
		}`

		req := httptest.NewRequest(SuccessTest_2.Method, SuccessTest_2.Path, strings.NewReader(reqStr))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		c.SetPath(SuccessTest_2.Path)
		if assert.NoError(t, controllers.CreateProduct(c)) {
			assert.Equal(t, SuccessTest_2.ExpectStatus, rec.Code)
			assert.Contains(t, rec.Body.String(), SuccessTest_2.ExpectResponse)

		}
	}()
}

func TestProductIn(t *testing.T) {
	SuccessTest := TestCase{

		Method:         http.MethodPost,
		Name:           "Create Product",
		Path:           "/product/in",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Sukses Menambahkan Product Barang Masuk",
	}

	e := InitEchoTestAPI()
	InsertDataProduct()
	reqStr := `{
		"ID":1,
    	"product_id": 1,
    	"qty": 40
	}`

	req := httptest.NewRequest(SuccessTest.Method, SuccessTest.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(SuccessTest.Path)
	if assert.NoError(t, controllers.ProductInController(c)) {
		assert.Equal(t, SuccessTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), SuccessTest.ExpectResponse)

	}
}

func TestProductOut(t *testing.T) {
	SuccessTest := TestCase{

		Method:         http.MethodPost,
		Name:           "Create Product",
		Path:           "/product/out",
		ExpectStatus:   http.StatusOK,
		ExpectResponse: "Sukses Menambahkan Product Barang Keluar",
	}

	e := InitEchoTestAPI()
	InsertDataProduct()
	reqStr := `{
		"ID":1,
    	"product_id": 1,
    	"qty": 50
	}`

	req := httptest.NewRequest(SuccessTest.Method, SuccessTest.Path, strings.NewReader(reqStr))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	c.SetPath(SuccessTest.Path)
	if assert.NoError(t, controllers.ProductOutController(c)) {
		assert.Equal(t, SuccessTest.ExpectStatus, rec.Code)
		assert.Contains(t, rec.Body.String(), SuccessTest.ExpectResponse)

	}
}
