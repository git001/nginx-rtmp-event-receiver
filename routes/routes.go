package routes

import (
	"fmt"
	"net/url"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type MyAuth struct {
	CameraID string `form:"camid"`
	Token    string `form:"token"`
	Tcurl    string `form:"tcurl"`
	Call     string `form:"call"`
	App      string `form:"app"`
	ClientID string `form:"clientid"`
}

var resp string = "okay"

func CheckAccess(inputBody MyAuth) bool {
	myUrl, error := url.ParseRequestURI(inputBody.Tcurl)

	if error != nil {
		fmt.Printf("CheckAccess() error %#v", error)
		return false
	}

	//	fmt.Printf("CheckAccess() true\n")
	//	fmt.Printf("CheckAccess() myUrl: %s\n", inputUrl)
	if viper.GetBool("config.debug") == true {
		fmt.Printf("CheckAccess() myUrl: %#v\n", inputBody)
	}

	if viper.GetBool("streams."+inputBody.App) == true {
		if viper.GetBool("config.verbose") == true {
			fmt.Printf("CheckAccess() accesslist true for streams.%s\n", inputBody.App)
		}
		return true
	}

	if (myUrl.Query().Get("camid") == "") || (myUrl.Query().Get("token") == "") {
		if viper.GetBool("config.verbose") == true {
			fmt.Printf("CheckAccess() false for streams.%s\n", inputBody.App)
		}
		return false
	}

	return true
}
func OnConnet(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnPlay(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnPlayList(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnPublish(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnDone(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnPlayDone(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnPublishDone(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnRecordStarted(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnRecordDone(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func OnUpdate(c *fiber.Ctx) error {
	// check for the incoming request body
	b := new(MyAuth)

	if err := c.BodyParser(b); err != nil {
		return err
	}

	if viper.GetBool("config.verbose") == true {
		fmt.Printf("body: %#v\n", b)
		c.Context().Logger().Printf("body: %#v\n", b)
	}

	if CheckAccess(*b) == false {
		return c.Status(fiber.StatusForbidden).JSON("Access Denied")
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
