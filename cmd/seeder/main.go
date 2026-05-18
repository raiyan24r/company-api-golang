package main

import (
	"company-api/app/api/app"
	"company-api/business/database"
	mysqldb "company-api/foundation/database"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	numCompaniesStr := flag.String("count", "5", "Number of companies to seed")
	flag.Parse()

	numCompanies, err := strconv.Atoi(*numCompaniesStr)
	if err != nil {
		log.Fatalf("Invalid count: %v", err)
	}

	if numCompanies <= 0 {
		log.Fatalf("Count must be greater than 0")
	}

	cfg, err := app.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	mysqlDb, err := mysqldb.Open(cfg.DB)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer mysqlDb.Close()

	dbRepo := database.New(mysqlDb)

	companies := generateCompanies(numCompanies)

	for i, company := range companies {
		id, err := dbRepo.CreateCompany(company)
		if err != nil {
			log.Printf("Failed to insert company %d: %v\n", i+1, err)
			continue
		}
		log.Printf("Inserted company %d/%d (ID: %d)\n", i+1, numCompanies, id)
	}

	log.Printf("Successfully seeded %d companies\n", numCompanies)
}

func generateCompanies(count int) []database.Company {
	// Generate a unique suffix for this run
	seed := time.Now().UnixNano()
	rand.Seed(seed)
	runID := rand.Intn(100000)

	companies := []database.Company{
		{
			Name:              "Tech Innovators Inc.",
			Description:       "Leading technology innovation company",
			AmountOfEmployees: 500,
			Registered:        true,
			Type:              "Technology",
		},
		{
			Name:              "Green Energy Solutions",
			Description:       "Renewable energy and sustainability provider",
			AmountOfEmployees: 250,
			Registered:        true,
			Type:              "Energy",
		},
		{
			Name:              "HealthPlus Corp.",
			Description:       "Healthcare services and medical solutions",
			AmountOfEmployees: 350,
			Registered:        true,
			Type:              "Healthcare",
		},
		{
			Name:              "FinTech Global",
			Description:       "Financial technology and digital banking",
			AmountOfEmployees: 200,
			Registered:        true,
			Type:              "Finance",
		},
		{
			Name:              "CloudSync Systems",
			Description:       "Cloud infrastructure and storage solutions",
			AmountOfEmployees: 180,
			Registered:        true,
			Type:              "Technology",
		},
		{
			Name:              "EcoWorks Ltd.",
			Description:       "Environmental consulting and sustainability",
			AmountOfEmployees: 120,
			Registered:        true,
			Type:              "Environment",
		},
		{
			Name:              "AI Ventures",
			Description:       "Artificial intelligence and machine learning",
			AmountOfEmployees: 300,
			Registered:        true,
			Type:              "Technology",
		},
		{
			Name:              "RetailMax",
			Description:       "Retail management and e-commerce solutions",
			AmountOfEmployees: 400,
			Registered:        true,
			Type:              "Retail",
		},
		{
			Name:              "LogisticsPro",
			Description:       "Supply chain and logistics optimization",
			AmountOfEmployees: 280,
			Registered:        true,
			Type:              "Logistics",
		},
		{
			Name:              "DigitalMarket Co.",
			Description:       "Digital marketing and advertising agency",
			AmountOfEmployees: 150,
			Registered:        true,
			Type:              "Marketing",
		},
	}

	// Add unique suffix to each company name
	for i := range companies {
		companies[i].Name = fmt.Sprintf("%s_%d", companies[i].Name, runID)
	}

	// If count is less than or equal to available companies, return only the requested amount
	if count <= len(companies) {
		return companies[:count]
	}

	// If count is more than available, cycle through and repeat
	result := make([]database.Company, count)
	for i := 0; i < count; i++ {
		baseCompany := companies[i%len(companies)]
		// Add an index suffix to make names unique
		baseCompany.Name = fmt.Sprintf("%s_%d", baseCompany.Name, i/len(companies)+1)
		result[i] = baseCompany
	}

	return result
}
