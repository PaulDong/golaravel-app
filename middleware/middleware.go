package middleware

import (
	"bnlogic/data"

	"github.com/PaulDong/golaravel"
)

type Middleware struct{
	App *golaravel.Golaravel
	Models data.Models
}