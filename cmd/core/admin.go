package main

import (
	"errors"
	"log"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/forms"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
)

func (core *Core) SetupAdmin(app *fiber.App) {
	admin := app.Group("/admin")

	admin.Use(
		helmet.New(),
	)

	admin.Get("/login", func(c *fiber.Ctx) error {
		return c.Render("pages/admin/auth/login", ViewData{
			Path: c.Path(),
		})
	})

	admin.Post("/auth/login", core.AdminHandleLogin)

	admin.Use(
		core.RequireAuth,
	)

	admin.Post("/auth/logout", core.AdminHandleLogout)

	admin.Get("/", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user

		return c.Render("pages/admin/home/index", ViewData)
	})

	// Offers pages.
	offers := admin.Group("/offers")

	offers.Get("/", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user

		return c.Render("pages/admin/offers/index", ViewData)
	})

	// Resources pages.
	resources := admin.Group("/resources")

	resources.Get("/", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user

		return c.Render("pages/admin/resources/index", ViewData)
	})

	// Files pages.
	files := admin.Group("/files")

	files.Get("/", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		files, err := core.Database.FilesRepository.GetAll()
		if err != nil {
			log.Println(err)
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user
		ViewData.Files = files

		return c.Render("pages/admin/files/index", ViewData)
	})

	// Companies pages.
	companies := admin.Group("/companies", core.RequireAdmin)

	companies.Get("/", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user

		return c.Render("pages/admin/companies/index", ViewData)
	})
}

func (core *Core) AdminHandleLogin(c *fiber.Ctx) error {
	formErrors := forms.Errors{}
	username := c.FormValue("username")
	password := c.FormValue("password")

	viewData := core.GetCommonViewData(c)

	if forms.IsEmpty(username) || forms.IsEmpty(password) {
		formErrors.Add("generic", "Campos requeridos")
		viewData.Errors = formErrors

		return c.Render(
			"pages/admin/auth/login",
			viewData,
		)
	}

	id, err := core.Database.AuthRepository.Login(username, password)
	if err != nil {
		if errors.Is(err, database.ErrAuth) {
			formErrors.Add("generic", "Usuario o contraseña invalido")
			viewData.Errors = formErrors

			return c.Render(
				"pages/admin/auth/login",
				viewData,
			)
		}

		return fiber.ErrInternalServerError
	}

	user, err := core.Database.UsersRepository.GetById(id)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	sess := core.GetSession(c)
	sess.Set(USER_ID_KEY, user.Id)
	sess.Set(USER_ROLE_KEY, user.Role)
	sess.Set(IS_LOGGED_IN_KEY, true)

	return c.Redirect("/admin", fiber.StatusSeeOther)
}

func (core *Core) AdminHandleLogout(c *fiber.Ctx) error {
	sess := core.GetSession(c)
	sess.Set(USER_ID_KEY, 0)
	sess.Set(USER_ROLE_KEY, "")
	sess.Set(IS_LOGGED_IN_KEY, false)
	return c.Redirect("/admin/login")
}
