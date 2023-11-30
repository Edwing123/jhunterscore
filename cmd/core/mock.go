package main

import "edwingarcia.dev/github/jhunterscore/pkg/models"

var (
	mockOffers = []models.Offer{
		{
			Id:          1,
			Title:       "Programador C# con MVC",
			Role:        "programador",
			Company:     "Google",
			Content:     "# Buscamos programador con experiencia en C# y MVC",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1000,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-25",
		},
		{
			Id:          2,
			Title:       "Ingeniero de software web",
			Role:        "programador",
			Company:     "DevTo",
			Content:     "Hola, esta puede ser tu oportunidad de trabajar con nosotros, sigue leyendo.\n\n## Requisitos\n- Experiencia en desarrollo web\n- Experiencia en desarrollo de aplicaciones web\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1300,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          3,
			Title:       "Administrador de base de datos",
			Role:        "dba",
			Company:     "Finanzas SA",
			Content:     "# Buscamos administrador de base de datos con experiencia en MySQL\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1600,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          4,
			Title:       "Contador publico",
			Role:        "contabilidad",
			Company:     "Finanzas SA",
			Content:     "# Requrimos contador publico con experiencia en contabilidad\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Estelí",
			Salary:      950,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-27",
		},
		// Pattern repeated from here.
		{
			Id:          5,
			Title:       "Programador C# con MVC",
			Role:        "programador",
			Company:     "Google",
			Content:     "# Buscamos programador con experiencia en C# y MVC",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1000,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-25",
		},
		{
			Id:          6,
			Title:       "Ingeniero de software web",
			Role:        "programador",
			Company:     "DevTo",
			Content:     "Hola, esta puede ser tu oportunidad de trabajar con nosotros, sigue leyendo.\n\n## Requisitos\n- Experiencia en desarrollo web\n- Experiencia en desarrollo de aplicaciones web\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1300,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          7,
			Title:       "Administrador de base de datos",
			Role:        "dba",
			Company:     "Finanzas SA",
			Content:     "# Buscamos administrador de base de datos con experiencia en MySQL\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1600,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          8,
			Title:       "Contador publico",
			Role:        "contabilidad",
			Company:     "Finanzas SA",
			Content:     "# Requrimos contador publico con experiencia en contabilidad\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Estelí",
			Salary:      950,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-27",
		},
		{
			Id:          9,
			Title:       "Programador C# con MVC",
			Role:        "programador",
			Company:     "Google",
			Content:     "# Buscamos programador con experiencia en C# y MVC",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1000,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-25",
		},
		{
			Id:          10,
			Title:       "Ingeniero de software web",
			Role:        "programador",
			Company:     "DevTo",
			Content:     "Hola, esta puede ser tu oportunidad de trabajar con nosotros, sigue leyendo.\n\n## Requisitos\n- Experiencia en desarrollo web\n- Experiencia en desarrollo de aplicaciones web\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1300,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          11,
			Title:       "Administrador de base de datos",
			Role:        "dba",
			Company:     "Finanzas SA",
			Content:     "# Buscamos administrador de base de datos con experiencia en MySQL\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1600,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          12,
			Title:       "Contador publico",
			Role:        "contabilidad",
			Company:     "Finanzas SA",
			Content:     "# Requrimos contador publico con experiencia en contabilidad\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Estelí",
			Salary:      950,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-27",
		},
		{
			Id:          13,
			Title:       "Programador C# con MVC",
			Role:        "programador",
			Company:     "Google",
			Content:     "# Buscamos programador con experiencia en C# y MVC",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1000,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-25",
		},
		{
			Id:          14,
			Title:       "Ingeniero de software web",
			Role:        "programador",
			Company:     "DevTo",
			Content:     "Hola, esta puede ser tu oportunidad de trabajar con nosotros, sigue leyendo.\n\n## Requisitos\n- Experiencia en desarrollo web\n- Experiencia en desarrollo de aplicaciones web\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1300,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          15,
			Title:       "Administrador de base de datos",
			Role:        "dba",
			Company:     "Finanzas SA",
			Content:     "# Buscamos administrador de base de datos con experiencia en MySQL\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1600,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          16,
			Title:       "Contador publico",
			Role:        "contabilidad",
			Company:     "Finanzas SA",
			Content:     "# Requrimos contador publico con experiencia en contabilidad\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Estelí",
			Salary:      950,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-27",
		},
		{
			Id:          17,
			Title:       "Programador C# con MVC",
			Role:        "programador",
			Company:     "Google",
			Content:     "# Buscamos programador con experiencia en C# y MVC",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1000,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-25",
		},
		{
			Id:          18,
			Title:       "Ingeniero de software web",
			Role:        "programador",
			Company:     "DevTo",
			Content:     "Hola, esta puede ser tu oportunidad de trabajar con nosotros, sigue leyendo.\n\n## Requisitos\n- Experiencia en desarrollo web\n- Experiencia en desarrollo de aplicaciones web\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1300,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          19,
			Title:       "Administrador de base de datos",
			Role:        "dba",
			Company:     "Finanzas SA",
			Content:     "# Buscamos administrador de base de datos con experiencia en MySQL\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Managua",
			Salary:      1600,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-21",
		},
		{
			Id:          20,
			Title:       "Contador publico",
			Role:        "contabilidad",
			Company:     "Finanzas SA",
			Content:     "# Requrimos contador publico con experiencia en contabilidad\n\n## Beneficios\n- Seguro médico\n- Seguro de vida\n- Vacaciones pagadas\n",
			Contract:    "trabajo",
			Location:    "Estelí",
			Salary:      950,
			ContactInfo: "recruit@example.com",
			CreatedAt:   "2023-11-27",
		},
	}
)

func getCompactOffers() []models.CompactOffer {
	var compactOffers []models.CompactOffer

	for _, offer := range mockOffers {
		compactOffers = append(compactOffers, models.CompactOffer{
			Id:        offer.Id,
			Title:     offer.Title,
			Role:      offer.Role,
			Company:   offer.Company,
			Contract:  offer.Contract,
			Salary:    offer.Salary,
			CreatedAt: offer.CreatedAt,
		})
	}
	return compactOffers
}
