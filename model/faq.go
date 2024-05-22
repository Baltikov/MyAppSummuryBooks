package model

import (
	"fmt"
	"testapi/bd"
)

type Faq struct {
	ID          int64  `json:"faq_id"`
	Description string `json:"faq_description"`
	Category    string `json:"category"`
}

func GETAll() ([]Faq, error) {
	query := "SELECT * FROM faq"
	rowsFaq, err := bd.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rowsFaq.Close()
	var faqs []Faq
	for rowsFaq.Next() {
		var faq Faq
		err = rowsFaq.Scan(&faq.ID, &faq.Description, &faq.Category)
		if err = rowsFaq.Err(); err != nil {
			return nil, fmt.Errorf("error iterating over FAQs: %v", err)
		}
		faqs = append(faqs, faq)
	}
	return faqs, nil
}
