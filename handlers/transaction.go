package handlers

import (
	// "fmt"
	// "fmt"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
	dtoresult "waysbook/dto/result"
	dtotransaction "waysbook/dto/transaction"
	"waysbook/models"
	"waysbook/repository"

	"github.com/go-playground/validator/v10"
	jwt "github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
	"github.com/yudapc/go-rupiah"
	"gopkg.in/gomail.v2"
)

type TransactionHandler struct {
	TransactionRepository repository.TransactionRepository
}

func HandlerTransaction(repo repository.TransactionRepository) *TransactionHandler {
	return &TransactionHandler{repo}
}

func (h *TransactionHandler) GetAllTransaction(c echo.Context) error {
	dataResponse, err := h.TransactionRepository.GetAllTransaction()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *TransactionHandler) GetTransactionById(c echo.Context) error {
	var id, _ = strconv.Atoi(c.Param("id"))

	dataResponse, err := h.TransactionRepository.GetTransactionById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: dataResponse,
	})
}

func (h *TransactionHandler) AddTransaction(c echo.Context) error {
	request := new(dtotransaction.TransactionRequest)
	if err := c.Bind(request); err != nil {
		return c.JSON(http.StatusInternalServerError, dtoresult.ErrorResult{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}

	validation := validator.New()
	err := validation.Struct(request)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	userLogin := c.Get("userLogin")
	userId := userLogin.(jwt.MapClaims)["id"].(float64)

	var transactionMatch = false
	var transactionId int

	if !transactionMatch {
		transactionId = int(time.Now().Unix())
		transactionData, _ := h.TransactionRepository.GetTransactionById(transactionId)

		if transactionData.Id == 0 {
			transactionMatch = true
		}
	}

	user, err := h.TransactionRepository.GetUser(int(userId))
	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	var books []models.Book
	var total int = 0
	for _, cartItem := range user.CartItem {
		book, err := h.TransactionRepository.GetBookId(int(cartItem.BookId))
		total += cartItem.Book.Price
		if err != nil {
			return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
		}
		books = append(books, book)
	}

	transaction := models.Transaction{
		Id:           transactionId,
		UserId:       int(userId),
		User:         user,
		Book:         books,
		CounterQty:   len(books),
		TotalPayment: total,
		Status:       request.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	dataResponse, err := h.TransactionRepository.AddTransaction(transaction)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	sendEmail("success", dataResponse)

	var s = snap.Client{}
	s.New(os.Getenv("SERVER_KEY"), midtrans.Sandbox)

	req := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  strconv.Itoa(dataResponse.Id),
			GrossAmt: int64(dataResponse.TotalPayment),
		},
		CreditCard: &snap.CreditCardDetails{
			Secure: true,
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: dataResponse.User.Name,
			Email: dataResponse.User.Email,
		},
	}

	snapResp, _ := s.CreateTransaction(req)

	err = h.TransactionRepository.Delete(int(userId))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	user.CartItem = nil
	_, err = h.TransactionRepository.UpdateUserCart(user)

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: snapResp,
	})
}

func sendEmail(status string, transaction models.Transaction) {
	if status != transaction.Status && status == "success" {
		CONFIG_SMTP_HOST := "smtp.gmail.com"
		CONFIG_SMTP_PORT := 587
		CONFIG_SENDER_NAME := "WaysBook <ibnu.ibnualinsani23@gmail.com>"
		CONFIG_AUTH_EMAIL := os.Getenv("EMAIL_SYSTEM")
		CONFIG_AUTH_PASSWORD := os.Getenv("PASSWORD_SYSTEM")

		var books string
		for _, book := range transaction.Book {
			books += fmt.Sprintf("<li>%s</li>", book.Title)
		}

		price := rupiah.FormatRupiah(float64(transaction.TotalPayment))
		fmt.Println(transaction.User.Email)

		htmlBody := fmt.Sprintf(`
			<!DOCTYPE html>
			<html lang="id">
			<head>
				<meta charset="UTF-8" />
				<meta http-equiv="X-UA-Compatible" content="IE=edge" />
				<meta name="viewport" content="width=device-width, initial-scale=1.0" />
				<title>Dokumen</title>
				<style>
					h1 {
						color: brown;
					}
				</style>
			</head>
			<body>
				<h2>Pembayaran Produk:</h2>
				<ul style="list-style-type:none;">
					%s
					<li>Total pembayaran: %s</li>
					<li>Status: <b>%s</b></li>
					<li>Terima kasih telah melakukan pesanan, harap tunggu jadwal perjalanan. Selamat menikmati perjalanan Anda!</li>
				</ul>
			</body>
			</html>
		`, books, price, status)

		mailer := gomail.NewMessage()
		mailer.SetHeader("From", CONFIG_SENDER_NAME)
		mailer.SetHeader("To", transaction.User.Email)
		mailer.SetHeader("Subject", "Status Transaksi")
		mailer.SetBody("text/html", htmlBody)

		dialer := gomail.NewDialer(
			CONFIG_SMTP_HOST,
			CONFIG_SMTP_PORT,
			CONFIG_AUTH_EMAIL,
			CONFIG_AUTH_PASSWORD,
		)

		err := dialer.DialAndSend(mailer)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Println("Email terkirim ke " + transaction.User.Email)
	}
}


func (h *TransactionHandler) Notification(c echo.Context) error {
	var notificationPayload map[string]interface{}

	if err := c.Bind(&notificationPayload); err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}

	transactionStatus := notificationPayload["transaction_status"].(string)
	fraudStatus := notificationPayload["fraud_status"].(string)
	orderid := notificationPayload["order_id"].(string)

	order_id, _ := strconv.Atoi(orderid)

	transaction, err := h.TransactionRepository.GetTransactionById(int(order_id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, dtoresult.ErrorResult{
			Code:    http.StatusBadGateway,
			Message: err.Error(),
		})
	}

	if transactionStatus == "capture" {
		if fraudStatus == "accept" {
			h.TransactionRepository.UpdateTransaction("success", transaction.Id)
			sendEmail("success", transaction)
		} else if fraudStatus == "deny" {
			h.TransactionRepository.UpdateTransaction("deny", transaction.Id)
		}

	} else if transactionStatus == "settlement" {
		h.TransactionRepository.UpdateTransaction("success", transaction.Id)
		sendEmail("success", transaction)
	} else if transactionStatus == "deny" {
		h.TransactionRepository.UpdateTransaction("failed", transaction.Id)
	} else if transactionStatus == "cancel" || transactionStatus == "expire" {
		h.TransactionRepository.UpdateTransaction("failed", transaction.Id)
	} else if transactionStatus == "pending" {
		h.TransactionRepository.UpdateTransaction("pending", transaction.Id)
	}

	return c.JSON(http.StatusOK, dtoresult.SuccessResult{
		Code: http.StatusOK,
		Data: notificationPayload,
	})
}
