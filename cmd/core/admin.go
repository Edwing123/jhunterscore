package main

import (
	"errors"
	"fmt"
	"os"
	"path"
	"strconv"

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

		offers, err := core.Database.OffersRepository.GetAll()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user
		ViewData.Offers = offers

		core.ClearErrors(c)

		return c.Render("pages/admin/offers/index", ViewData)
	})

	offers.Get("/new", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		companies, err := core.Database.CompaniesRepository.GetAll()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		locations, err := core.Database.LocationsRepository.GetAll()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user
		ViewData.Companies = companies
		ViewData.Locations = locations
		core.ClearErrors(c)

		return c.Render("pages/admin/offers/new", ViewData)
	})

	offers.Post("/__new", core.AdminHandleOffersNew)
	offers.Post("/__delete", core.AdminHandleOffersDelete)

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

	files.Get("/edit/:id<int>", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		errorsBag := forms.Errors{}

		fileId, _ := c.ParamsInt("id")

		file, err := core.Database.FilesRepository.GetById(fileId)
		if err != nil {
			if errors.Is(err, database.ErrNoRows) {
				errorsBag.Add("generic", fmt.Sprintf("El archivo con id %d no existe.", fileId))
				core.SetErrors(errorsBag, c)
				return c.Redirect("/admin/files", fiber.StatusSeeOther)
			}

			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user
		ViewData.File = file
		core.ClearErrors(c)

		return c.Render("pages/admin/files/edit", ViewData)
	})

	files.Post("/__new", core.AdminHandleFilesNew)
	files.Post("/__delete", core.AdminHandleFilesDelete)
	files.Post("/__edit", core.AdminHandleFilesEdit)

	// Companies pages.
	companies := admin.Group("/companies", core.RequireAdmin)

	companies.Get("/", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		companies, err := core.Database.CompaniesRepository.GetAll()
		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user
		ViewData.Companies = companies

		core.ClearErrors(c)

		return c.Render("pages/admin/companies/index", ViewData)
	})

	companies.Get("/new", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user
		core.ClearErrors(c)

		return c.Render("pages/admin/companies/new", ViewData)
	})

	companies.Get("/edit/:id<int>", func(c *fiber.Ctx) error {
		userId := core.GetUserId(c)
		user, err := core.Database.UsersRepository.GetById(userId)

		if err != nil {
			return fiber.ErrInternalServerError
		}

		errorsBag := forms.Errors{}

		companyId, _ := c.ParamsInt("id")

		company, err := core.Database.CompaniesRepository.GetById(companyId)
		if err != nil {
			if errors.Is(err, database.ErrNoRows) {
				errorsBag.Add("generic", fmt.Sprintf("La empresa con id %d no existe.", companyId))
				core.SetErrors(errorsBag, c)
				return c.Redirect("/admin/companies", fiber.StatusSeeOther)
			}

			return fiber.ErrInternalServerError
		}

		ViewData := core.GetCommonViewData(c)
		ViewData.User = user
		ViewData.Company = company
		core.ClearErrors(c)

		return c.Render("pages/admin/companies/edit", ViewData)
	})

	companies.Post("/__new", core.AdminHandleCompaniesNew)
	companies.Post("/__delete", core.AdminHandleCompaniesDelete)
	companies.Post("/__edit", core.AdminHandleCompaniesEdit)
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

	_, err = core.Database.FilesRepository.Create(models.File{
		Name:     filename,
		Path:     filePath,
		Size:     int(fileheader.Size),
		MIMEType: fileheader.Header.Get("Content-Type"),
		UserId:   userId,
	})
	if err != nil {
		if errors.Is(err, database.ErrFileNameTaken) {
			formsErrors.Add("generic", "Ya existe un archivo con este nombre.")
			core.SetErrors(formsErrors, c)
			return c.Redirect("/admin/files/new")
		}

		return fiber.ErrInternalServerError
	}

	formsErrors.Add("success", "Archivo creado con exito.")
	core.SetErrors(formsErrors, c)
	return c.Redirect("/admin/files", fiber.StatusSeeOther)
}

func (core *Core) AdminHandleFilesDelete(c *fiber.Ctx) error {
	errorsBag := forms.Errors{}
	currentUserId := core.GetUserId(c)
	currentUserRole := core.GetUserRole(c)

	redirect := func() error {
		core.SetErrors(errorsBag, c)
		return c.Redirect("/admin/files", fiber.StatusSeeOther)
	}

	fileIdStr := c.FormValue("file_id")
	fileId, err := strconv.Atoi(fileIdStr)
	if err != nil {
		errorsBag.Add("generic", "Hay un problema con el id del archivo.")
		return redirect()
	}

	fileUserId, err := core.Database.FilesRepository.GetFileUserIdByFileId(fileId)
	if err != nil {
		errorsBag.Add("generic", "No se puede eliminar el archivo.")
		return redirect()
	}

	if currentUserId != fileUserId && currentUserRole != "admin" {
		errorsBag.Add("generic", "No puedes eliminar el archivo porque no eres el creador.")
		return redirect()
	}

	file, err := core.Database.FilesRepository.GetById(fileId)
	if err != nil {
		if errors.Is(err, database.ErrNoRows) {
			errorsBag.Add("generic", "No se puede eliminar el archivo.")
			return redirect()
		}

		return fiber.ErrInternalServerError
	}

	err = os.Remove(path.Join(core.FilesDir, file.Path))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	err = core.Database.FilesRepository.Delete(fileId)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	errorsBag.Add("success", "Archivo eliminado.")
	return redirect()
}

func (core *Core) AdminHandleFilesEdit(c *fiber.Ctx) error {
	errorsBag := forms.Errors{}
	currentUserId := core.GetUserId(c)
	currentUserRole := core.GetUserRole(c)

	redirectIndex := func() error {
		core.SetErrors(errorsBag, c)
		return c.Redirect("/admin/files", fiber.StatusSeeOther)
	}

	redirectEdit := func(id int) error {
		core.SetErrors(errorsBag, c)
		return c.Redirect(fmt.Sprintf("/admin/files/edit/%d", id), fiber.StatusSeeOther)
	}

	fileIdStr := c.FormValue("file_id")
	fileId, err := strconv.Atoi(fileIdStr)
	if err != nil {
		errorsBag.Add("generic", "Hay un problema con el id del archivo.")
		return redirectIndex()
	}

	filename := c.FormValue("filename")
	if forms.IsEmpty(filename) {
		errorsBag.Add("generic", "El campo nombre no puede estar vacio.")
		return redirectEdit(fileId)
	}

	fileUserId, err := core.Database.FilesRepository.GetFileUserIdByFileId(fileId)
	if err != nil {
		errorsBag.Add("generic", "No se puede editar el archivo.")
		return redirectIndex()
	}

	if currentUserId != fileUserId && currentUserRole != "admin" {
		errorsBag.Add("generic", "No puedes editar el archivo porque no eres el creador.")
		return redirectIndex()
	}

	_, err = core.Database.FilesRepository.Update(models.File{
		Id:   fileId,
		Name: filename,
	})
	if err != nil {
		if errors.Is(err, database.ErrFileNameTaken) {
			errorsBag.Add("generic", "Ya existe un archivo con este nombre.")
			return redirectEdit(fileId)
		}

		return fiber.ErrInternalServerError
	}

	errorsBag.Add("success", "Archivo editado.")
	return redirectIndex()
}

func (core *Core) AdminHandleCompaniesNew(c *fiber.Ctx) error {
	formsErrors := forms.Errors{}

	redirect := func() error {
		core.SetErrors(formsErrors, c)
		return c.Redirect("/admin/companies/new", fiber.StatusSeeOther)
	}

	companyName := c.FormValue("name")
	if forms.IsEmpty(companyName) {
		formsErrors.Add("generic", "El campo de nombre no puede estar vacio.")
		return redirect()
	}

	logoURL := c.FormValue("logo_url")
	if forms.IsEmpty(logoURL) {
		formsErrors.Add("generic", "El campo de logo URL no puede estar vacio.")
		return redirect()
	}

	_, err := core.Database.CompaniesRepository.Create(models.Company{
		Name:    companyName,
		LogoURL: logoURL,
	})
	if err != nil {
		if errors.Is(err, database.ErrCompanyNameTaken) {
			formsErrors.Add("generic", "Ya existe una empresa con este nombre.")
			core.SetErrors(formsErrors, c)
			return redirect()
		}

		return fiber.ErrInternalServerError
	}

	formsErrors.Add("success", "Empresa registrada con exito.")
	core.SetErrors(formsErrors, c)
	return c.Redirect("/admin/companies", fiber.StatusSeeOther)
}

func (core *Core) AdminHandleCompaniesDelete(c *fiber.Ctx) error {
	errorsBag := forms.Errors{}

	redirect := func() error {
		core.SetErrors(errorsBag, c)
		return c.Redirect("/admin/companies", fiber.StatusSeeOther)
	}

	companyIdStr := c.FormValue("company_id")
	companyId, err := strconv.Atoi(companyIdStr)
	if err != nil {
		errorsBag.Add("generic", "Hay un problema con el id de la empresa.")
		return redirect()
	}

	err = core.Database.CompaniesRepository.Delete(companyId)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	errorsBag.Add("success", "Empresa eliminada.")
	return redirect()
}

func (core *Core) AdminHandleCompaniesEdit(c *fiber.Ctx) error {
	errorsBag := forms.Errors{}

	redirectIndex := func() error {
		core.SetErrors(errorsBag, c)
		return c.Redirect("/admin/companies", fiber.StatusSeeOther)
	}

	redirectEdit := func(id int) error {
		core.SetErrors(errorsBag, c)
		return c.Redirect(fmt.Sprintf("/admin/companies/edit/%d", id), fiber.StatusSeeOther)
	}

	companyIdStr := c.FormValue("company_id")
	companyId, err := strconv.Atoi(companyIdStr)
	if err != nil {
		errorsBag.Add("generic", "Hay un problema con el id de la empresa.")
		return redirectIndex()
	}

	companyName := c.FormValue("name")
	if forms.IsEmpty(companyName) {
		errorsBag.Add("generic", "El campo nombre no puede estar vacio.")
		return redirectEdit(companyId)
	}

	logoURL := c.FormValue("logo_url")
	if forms.IsEmpty(logoURL) {
		errorsBag.Add("generic", "El campo logo URL no puede estar vacio.")
		return redirectEdit(companyId)
	}

	_, err = core.Database.CompaniesRepository.Update(models.Company{
		Id:      companyId,
		Name:    companyName,
		LogoURL: logoURL,
	})
	if err != nil {
		if errors.Is(err, database.ErrCompanyNameTaken) {
			errorsBag.Add("generic", "Ya existe una empresa con este nombre.")
			return redirectEdit(companyId)
		}

		return fiber.ErrInternalServerError
	}

	errorsBag.Add("success", "Empresa guardada.")
	return redirectIndex()
}

func (core *Core) AdminHandleOffersNew(c *fiber.Ctx) error {
	formsErrors := forms.Errors{}
	userId := core.GetUserId(c)

	redirect := func() error {
		core.SetErrors(formsErrors, c)
		return c.Redirect("/admin/offers/new", fiber.StatusSeeOther)
	}

	title := c.FormValue("title")
	if forms.IsEmpty(title) {
		formsErrors.Add("generic", "El campo titulo no puede estar vacio.")
		return redirect()
	}

	role := c.FormValue("role")
	if forms.IsEmpty(role) {
		formsErrors.Add("generic", "El campo rol no puede estar vacio.")
		return redirect()
	}

	companyIdStr := c.FormValue("company_id")
	if forms.IsEmpty(companyIdStr) {
		formsErrors.Add("generic", "Tienes que seleccionar una empresa.")
		return redirect()
	}

	contract := c.FormValue("contract")
	if forms.IsEmpty(contract) {
		formsErrors.Add("generic", "Tienes que seleccionar un tipo de contrato.")
		return redirect()
	}

	locationIdStr := c.FormValue("location_id")
	if forms.IsEmpty(locationIdStr) {
		formsErrors.Add("generic", "Tienes que seleccionar una ubicacion.")
		return redirect()
	}

	salaryStr := c.FormValue("salary")
	if forms.IsEmpty(salaryStr) {
		formsErrors.Add("generic", "El campo salario no puede estar vacio.")
		return redirect()
	}

	contactInfo := c.FormValue("contact_info")
	if forms.IsEmpty(contactInfo) {
		formsErrors.Add("generic", "El campo informacion de contacto no puede estar vacio.")
		return redirect()
	}

	content := c.FormValue("content")
	if forms.IsEmpty(content) {
		formsErrors.Add("generic", "El campo contenido no puede estar vacio.")
		return redirect()
	}

	companyId, _ := strconv.Atoi(companyIdStr)
	locationId, _ := strconv.Atoi(locationIdStr)
	salary, _ := strconv.Atoi(salaryStr)

	_, err := core.Database.OffersRepository.Create(models.Offer{
		Title:       title,
		Role:        role,
		CompanyId:   companyId,
		LocationId:  locationId,
		Salary:      float64(salary),
		Content:     content,
		Contract:    contract,
		ContactInfo: contactInfo,
		UserId:      userId,
	})
	if err != nil {
		if errors.Is(err, database.ErrOfferTitleTaken) {
			formsErrors.Add("generic", "Ya existe una ofeta con este titulo.")
			core.SetErrors(formsErrors, c)
			return redirect()
		}

		return fiber.ErrInternalServerError
	}

	formsErrors.Add("success", "Oferta creada con exito.")
	core.SetErrors(formsErrors, c)
	return c.Redirect("/admin/offers", fiber.StatusSeeOther)
}

func (core *Core) AdminHandleOffersDelete(c *fiber.Ctx) error {
	errorsBag := forms.Errors{}
	currentUserId := core.GetUserId(c)
	currentUserRole := core.GetUserRole(c)

	redirect := func() error {
		core.SetErrors(errorsBag, c)
		return c.Redirect("/admin/offers", fiber.StatusSeeOther)
	}

	offerIdStr := c.FormValue("offer_id")
	offerId, err := strconv.Atoi(offerIdStr)
	if err != nil {
		errorsBag.Add("generic", "Hay un problema con el id de la oferta.")
		return redirect()
	}

	offerUserId, err := core.Database.OffersRepository.GetUserIdByOfferId(offerId)
	if err != nil {
		errorsBag.Add("generic", "No se puede eliminar el archivo.")
		return redirect()
	}

	if currentUserId != offerUserId && currentUserRole != "admin" {
		errorsBag.Add("generic", "No puedes eliminar la oferta porque no eres el creador.")
		return redirect()
	}

	err = core.Database.OffersRepository.Delete(offerId)
	if err != nil {
		return fiber.ErrInternalServerError
	}

	errorsBag.Add("success", "Oferta eliminada.")
	return redirect()
}
