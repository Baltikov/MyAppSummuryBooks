package model

import (
	"fmt"
	"log"
	"testapi/bd"
)

type Faq struct {
	ID          int64  `json:"faq_id"`
	Description string `json:"faq_description"`
	Category    string `json:"category"`
}

func GETAll(limit, page int) ([]Faq, error) {
	query := "SELECT * FROM faq LIMIT ? OFFSET ?"
	offset := (page - 1) * limit

	log.Printf("Executing query: %s with params: %d, %d", query, limit, offset)
	rowsFaq, err := bd.DB.Query(query, limit, offset)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return nil, fmt.Errorf("error querying FAQs: %v", err)
	}
	defer rowsFaq.Close()

	var faqs []Faq
	for rowsFaq.Next() {
		var faq Faq
		err = rowsFaq.Scan(&faq.ID, &faq.Description, &faq.Category)
		if err != nil {
			log.Printf("Error scanning FAQ: %v", err)
			return nil, fmt.Errorf("error scanning FAQ: %v", err)
		}
		faqs = append(faqs, faq)
	}

	if err = rowsFaq.Err(); err != nil {
		log.Printf("Error iterating over FAQs: %v", err)
		return nil, fmt.Errorf("error iterating over FAQs: %v", err)
	}

	return faqs, nil
}
