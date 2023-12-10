package main

import (
	"errors"
	"fmt"
	"log"
	"path"

	"edwingarcia.dev/github/jhunterscore/pkg/database"
	"edwingarcia.dev/github/jhunterscore/pkg/database/models"
	"edwingarcia.dev/github/jhunterscore/pkg/forms"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/google/uuid"
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

		core.ClearErrors(c)

		return c.Render("pages/admin/files/index", ViewData)
	})

	files.Get("/new", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user
		core.ClearErrors(c)

		return c.Render("pages/admin/files/new", ViewData)
	})

	files.Post("/__new", core.AdminHandleFilesNew)

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

func (core *Core) AdminHandleFilesNew(c *fiber.Ctx) error {
	userId := core.GetUserId(c)

	formsErrors := forms.Errors{}

	filename := c.FormValue("filename")
	if forms.IsEmpty(filename) {
		formsErrors.Add("generic", "El campo de nombre no puede estar vacio.")
		core.SetErrors(formsErrors, c)
		return c.Redirect("/admin/files/new")
	}

	fileheader, err := c.FormFile("file")
	if err != nil {
		formsErrors.Add("generic", "Hay un problema con el archivo, intente de nuevo.")
		core.SetErrors(formsErrors, c)
		return c.Redirect("/admin/files/new")
	}

	if fileheader.Size > MAX_FILE_SIZE {
		formsErrors.Add("generic", "El archivo tiene que pesar <= 5MB.")
		core.SetErrors(formsErrors, c)
		return c.Redirect("/admin/files/new")
	}

	filePath := uuid.NewString()

	err = c.SaveFile(fileheader, path.Join(core.FilesDir, filePath))
	if err != nil {
		formsErrors.Add("generic", "Hay un problema con el archivo, intente de nuevo.")
		core.SetErrors(formsErrors, c)
		return c.Redirect("/admin/files/new")
	}

	fmt.Println(len(filePath))

	file, err := core.Database.FilesRepository.Create(models.File{
		Name:     filename,
		Path:     filePath,
		Size:     int(fileheader.Size),
		MIMEType: fileheader.Header.Get("Content-Type"),
		UserId:   userId,
	})
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, database.ErrFileNameTaken) {
			formsErrors.Add("generic", "Ya existe un archivo con este nombre.")
			core.SetErrors(formsErrors, c)
			return c.Redirect("/admin/files/new")
		}

		return fiber.ErrInternalServerError
	}

	fmt.Println(file)

	formsErrors.Add("generic", "Archivo creado con exito.")
	core.SetErrors(formsErrors, c)
	return c.Redirect("/admin/files", fiber.StatusSeeOther)
}
