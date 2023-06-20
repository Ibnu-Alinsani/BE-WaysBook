package middleware

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"

	"github.com/labstack/echo/v4"
)

func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("file")

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		src, err := file.Open()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		defer src.Close()

		var ctx = context.Background()
		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
		var API_KEY = os.Getenv("API_KEY")
		var API_SECRET = os.Getenv("API_SECRET")

		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
		resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "waysbook/book"})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		image, err := c.FormFile("image")

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		srcImage, err := image.Open()

		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}

		defer srcImage.Close()

		dataImage, err := cld.Upload.Upload(ctx, srcImage, uploader.UploadParams{Folder: "waysbook/thumbnail"})

		if err != nil {
			fmt.Println(err.Error())
		}

		c.Set("dataFile", resp.SecureURL)
		c.Set("filePublicId", resp.PublicID)
		c.Set("dataImage", dataImage.SecureURL)
		c.Set("imagePublicId", dataImage.PublicID)
		return next(c)
	}
}

// image user
// func UploadFileUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		image, err := c.FormFile("avatar")

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, "tidak ada image")
// 		}

// 		src, err := image.Open()

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, "gagal membuka gambar")
// 		}

// 		defer src.Close()

// 		var ctx = context.Background()
// 		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
// 		var API_KEY = os.Getenv("API_KEY")
// 		var API_SECRET = os.Getenv("API_SECRET")

// 		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)
// 		resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "waysbook/user"})

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		c.Set("imageUser", resp.SecureURL)
// 		c.Set("imageUserPublicId", resp.PublicID)
// 		return next(c)
// 	}
// }

// func UpdateUser(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		id, _ := strconv.Atoi(c.Param("id"))

// 		user, _ := repository.UserRepository.GetUserById(repository.RepositoryUser(postgresql.DB) ,id)

// 		image, err := c.FormFile("avatar")

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, "tidak ada image")
// 		}

// 		src, err := image.Open()

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, "gagal membuka gambar")
// 		}

// 		defer src.Close()

// 		var ctx = context.Background()
// 		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
// 		var API_KEY = os.Getenv("API_KEY")
// 		var API_SECRET = os.Getenv("API_SECRET")

// 		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

// 		if user.PublicIDAvatar != "" {
// 			_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
// 				PublicID: user.PublicIDAvatar})

// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		}

// 		resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "waysbook/user"})

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		c.Set("imageUser", resp.SecureURL)
// 		c.Set("imageUserPublicId", resp.PublicID)
// 		return next(c)
// 	}
// }
// func UpdateBook(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {

// 		id, _ := strconv.Atoi(c.Param("id"))

// 		book, _ := repository.BookRepository.GetBookById(repository.RepositoryBook(postgresql.DB) ,id)

// 		var ctx = context.Background()
// 		var CLOUD_NAME = os.Getenv("CLOUD_NAME")
// 		var API_KEY = os.Getenv("API_KEY")
// 		var API_SECRET = os.Getenv("API_SECRET")

// 		cld, _ := cloudinary.NewFromParams(CLOUD_NAME, API_KEY, API_SECRET)

// 		if book.PublicIdBook != "" && book.PublicIdThumbnail != "" {
// 			_, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
// 				PublicID: book.PublicIdBook})

// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}

// 			_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
// 				PublicID: book.PublicIdThumbnail})

// 			if err != nil {
// 				fmt.Println(err.Error())
// 			}
// 		}

// 		file, err := c.FormFile("file")

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err)
// 		}

// 		src, err := file.Open()

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err)
// 		}

// 		defer src.Close()

// 		resp, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{Folder: "waysbook/book"})

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		image, err := c.FormFile("image")

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err)
// 		}

// 		srcImage, err := image.Open()

// 		if err != nil {
// 			return c.JSON(http.StatusBadRequest, err)
// 		}

// 		defer srcImage.Close()

// 		dataImage, err := cld.Upload.Upload(ctx, srcImage, uploader.UploadParams{Folder: "waysbook/thumbnail"})

// 		if err != nil {
// 			fmt.Println(err.Error())
// 		}

// 		c.Set("dataFile", resp.SecureURL)
// 		c.Set("filePublicId", resp.PublicID)
// 		c.Set("dataImage", dataImage.SecureURL)
// 		c.Set("imagePublicId", dataImage.PublicID)
// 		return next(c)
// 	}
// }
