package main

import (
	"github.com/rs/zerolog/log"
	"github.com/samber/lo"
)

func main() {

	a := []int{1, 2, 3}
	res := lo.Map(a, func(item int, index int) int {
		return item * 3
	})
	log.Info().Any("values", res).Send()
}
