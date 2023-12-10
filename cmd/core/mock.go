package main

import "edwingarcia.dev/github/jhunterscore/pkg/database/models"

var (
	fakeLogo = "https://cdn-icons-png.flaticon.com/512/2702/2702602.png"

	mockOffers = []models.Offer{
		{
			Id:             1,
			Title:          "Programador C# con MVC",
			Role:           "programador",
			Company:        "Google",
			Content:        "# Buscamos programador con experiencia en C# y MVC",
			Contract:       "trabajo",
			Location:       "Managua",
			Salary:         1000,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-25",
			CompanyLogoURL: fakeLogo,
		},
		{
			Id:             2,
			Title:          "Ingeniero de software web",
			Role:           "programador",
			Company:        "DevTo",
			Content:        "Hola, esta puede ser tu oportunidad de trabajar con nosotros, sigue leyendo.\n\n## Requisitos\n- Experiencia en desarrollo web\n- Experiencia en desarrollo de aplicaciones web\n",
			Contract:       "trabajo",
			Location:       "Managua",
			Salary:         1300,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-21",
			CompanyLogoURL: fakeLogo,
		},
		{
			Id:             3,
			Title:          "Administrador de base de datos",
			Role:           "dba",
			Company:        "Finanzas SA",
			Content:        "# Buscamos administrador de base de datos con experiencia en MySQL\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:       "trabajo",
			Location:       "Managua",
			Salary:         1600,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-21",
			CompanyLogoURL: fakeLogo,
		},
		{
			Id:             4,
			Title:          "Contador publico",
			Role:           "contabilidad",
			Company:        "Finanzas SA",
			Content:        "# Requrimos contador publico con experiencia en contabilidad\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:       "trabajo",
			Location:       "Estelí",
			Salary:         950,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-27",
			CompanyLogoURL: fakeLogo,
		},
		// Pattern repeated from here.
		{
			Id:             5,
			Title:          "Administrador para SuperExpress",
			Role:           "administrador",
			Company:        "CNCC",
			Content:        "# Buscamos encargado de administracion para SuperExpress",
			Contract:       "trabajo",
			Location:       "Esteli",
			Salary:         1000,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-25",
			CompanyLogoURL: fakeLogo,
		},
		{
			Id:             6,
			Title:          "Programador Ruby on Rails",
			Role:           "programador",
			Company:        "DevTo",
			Content:        "Hola, esta puede ser tu oportunidad de trabajar con nosotros, sigue leyendo.\n\n## Requisitos\n- Experiencia en desarrollo web\n- Experiencia en desarrollo de aplicaciones web con Ruby On Rails\n",
			Contract:       "trabajo",
			Location:       "Managua",
			Salary:         1300,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-21",
			CompanyLogoURL: fakeLogo,
		},
		{
			Id:             7,
			Title:          "Administrador de base de datos MySQL",
			Role:           "dba",
			Company:        "UdeM",
			Content:        "# Buscamos administrador de base de datos con experiencia en MySQL\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:       "trabajo",
			Location:       "Managua",
			Salary:         1100,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-21",
			CompanyLogoURL: fakeLogo,
		},
		{
			Id:             8,
			Title:          "Gerente de proyectos informaticos",
			Role:           "manager",
			Company:        "Finanzas SA",
			Content:        "# Se busca gerente de proyectos de software\n\nEres una persona con experiencia en manejando proyectos de software? Esta puede ser tu oportunidad.",
			Contract:       "trabajo",
			Location:       "Managua",
			Salary:         1550,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-27",
			CompanyLogoURL: fakeLogo,
		},
		{
			Id:             9,
			Title:          "Desarrollador de software Senior con Golang",
			Role:           "programador",
			Company:        "Google Nic",
			Content:        "Buscamos a un desarrollador de software senior (5+ anios de experiencia) con Go, y base de datos MySQL.",
			Contract:       "trabajo",
			Location:       "Managua",
			Salary:         2400,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-25",
			CompanyLogoURL: fakeLogo,
		},
		{
			Id:             10,
			Title:          "Ingeniero de datos",
			Role:           "data-engineer",
			Company:        "DevTo",
			Content:        "Hola, esta puede ser tu oportunidad de trabajar con nosotros, sigue leyendo.\n\n## Requisitos\n- Experiencia en desarrollo web\n- Experiencia en extraccion y analisis de datos\n",
			Contract:       "trabajo",
			Location:       "Managua",
			Salary:         1800,
			ContactInfo:    "recruit@example.com",
			CreatedAt:      "2023-11-21",
			CompanyLogoURL: fakeLogo,
		},
	}

	mockResources = []models.Resource{
		{
			Id:        1,
			Title:     "Aprender sobre los CV",
			Author:    "Edwin Garcia",
			Content:   "Hola, esta es una guia para aprender a crear un CV.\n\n## Que es un CV?\n\nUn CV es un documento que se utiliza para presentar tus habilidades y experiencia profesional a un empleador.\n\n## Que debe contener un CV?\n\n- Información personal\n- Experiencia laboral\n- Educación\n- Habilidades\n- Referencias\n",
			Summary:   "Guia para aprender a crear un CV",
			CreatedAt: "2023-11-21",
		},
		{
			Id:        2,
			Title:     "Terminos comunes en el entorno laboral",
			Author:    "Kenroy Norori",
			Content:   "Hola, esta es una guia para aprender terminos comunes en el entorno laboral.\n\n## Que es un CV?\n\nUn CV es un documento que se utiliza para presentar tus habilidades y experiencia profesional a un empleador.\n\n## Que debe contener un CV?\n\n- Información personal\n- Experiencia laboral\n- Educación\n- Habilidades\n- Referencias\n",
			Summary:   "Terminos comunes en el entorno laboral",
			CreatedAt: "2023-11-18",
		},
		{
			Id:        3,
			Title:     "Consejos para afrontar una entrevista de trabajo",
			Author:    "Enmanuel Olivas",
			Summary:   "Consejos para afrontar una entrevista de trabajo",
			Content:   "# Que es una entrevista?\n\nUna entrevista es un proceso de selección que se realiza a los candidatos para un puesto de trabajo.\n\n## Consejos\n\n- Prepararse para la entrevista\n- Vestirse adecuadamente\n- Llegar a tiempo\n- Ser amable\n- Ser honesto\n- Ser positivo\n- Ser agradecido\n",
			CreatedAt: "2023-11-25",
		},
		{
			Id:        4,
			Title:     "Mi experiencia trabajando para el gobierno",
			Author:    "Carlos Lopez",
			Summary:   "Mi experiencia trabajando para el gobierno",
			Content:   "# Mi experiencia trabajando para el gobierno\n\nHola, mi nombre es Carlos Lopez y esta es mi experiencia trabajando para el gobierno.\n\n## Como fue mi experiencia?\n\nMi experiencia fue muy buena, aprendi mucho y conoci a muchas personas.\n\n## Que aprendi?\n\nAprendi a trabajar en equipo, a ser responsable y a ser puntual.\n\n## Que recomiendo?\n\nRecomiendo trabajar para el gobierno porque es una experiencia que te ayuda a crecer como persona y profesionalmente.\n",
			CreatedAt: "2023-11-29",
		},
	}

	mockCompanies = []models.Company{
		{
			Id:      1,
			Name:    "Google",
			LogoURL: fakeLogo,
		},
		{
			Id:      2,
			Name:    "DevTo",
			LogoURL: fakeLogo,
		},
		{
			Id:      3,
			Name:    "Finanzas SA",
			LogoURL: fakeLogo,
		},
	}
)

func getCompactOffers(offers []models.Offer) []models.CompactOffer {
	var compactOffers []models.CompactOffer

	for _, offer := range mockOffers {
		compactOffers = append(compactOffers, models.CompactOffer{
			Id:             offer.Id,
			Title:          offer.Title,
			Role:           offer.Role,
			Company:        offer.Company,
			Contract:       offer.Contract,
			Salary:         offer.Salary,
			CreatedAt:      offer.CreatedAt,
			CompanyLogoURL: offer.CompanyLogoURL,
		})
	}

	return compactOffers
}

func getCompactResources(resources []models.Resource) []models.CompactResource {
	var compactResources []models.CompactResource

	for _, resource := range mockResources {
		compactResources = append(compactResources, models.CompactResource{
			Id:        resource.Id,
			Title:     resource.Title,
			Author:    resource.Author,
			CreatedAt: resource.CreatedAt,
		})
	}
	return compactResources
}
