package transform

import (
	"errors"
	"regexp"
	"strings"

	"github.com/vishkashpvp/cashflow-tracker/server/models"
)

func TransformTransactionType(tt models.TransactionType) (models.TransactionType, error) {
	ttString := string(tt)
	ttLower := strings.ToLower(ttString)
	tt = models.TransactionType(ttLower)
	if !tt.IsValid() {
		return "", errors.New("transaction type must be either 'earning' or 'spending'")
	}
	return tt, nil
}

func TransformCategory(category string) (string, error) {
	regex := regexp.MustCompile("[^a-zA-Z0-9]+")
	category = strings.TrimSpace(category)
	category = regex.ReplaceAllString(category, " ")
	words := strings.Fields(category)
	category = strings.ToUpper(strings.Join(words, "-"))
	if len(category) < 2 {
		return "", errors.New("category must have a minimum length of 2 alphanumeric characters")
	}
	return category, nil
}
