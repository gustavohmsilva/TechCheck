// Package tech seria o dominio da app, as partes logicas do que a app vai fazer
// ela nao entende sobre sql, ou http, mas pode entender de permissoes e validaçoes
package tech

type QueryOptions struct {
	Find   string
	Amount int
	Offset int
}

func (q *QueryOptions) Validate() error {
	// coisas para verificar os valores e retornar erro

	return nil
}
