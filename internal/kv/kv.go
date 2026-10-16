// Package kv é a fase 1 do mini-cloud: um banco chave-valor persistente.
//
// ESQUELETO. A implementação é sua (veja ROADMAP.md, etapas 1 a 4).
// Hoje todas as funções devolvem ErrNaoImplementado, então os testes falham.
package kv

import "errors"

var (
	// ErrNaoImplementado é devolvido enquanto você não implementa a função.
	ErrNaoImplementado = errors.New("kv: não implementado")

	// ErrNaoEncontrada é devolvido por Get quando a chave não existe.
	ErrNaoEncontrada = errors.New("kv: chave não encontrada")
)

// Store é o banco chave-valor.
//
// TODO (etapa 1): decidir os campos. Pergunta-guia: o que precisa ser
// protegido de acesso concorrente?
type Store struct{}

// Open abre (ou cria) um banco no diretório dir.
//
// TODO (etapa 2): ao abrir, reler o WAL e reconstruir o estado.
func Open(dir string) (*Store, error) {
	return nil, ErrNaoImplementado
}

// Put grava o valor da chave.
func (s *Store) Put(key, value []byte) error {
	return ErrNaoImplementado
}

// Get devolve o valor da chave, ou ErrNaoEncontrada.
func (s *Store) Get(key []byte) ([]byte, error) {
	return nil, ErrNaoImplementado
}

// Delete remove a chave. Remover uma chave inexistente não é erro.
//
// TODO (etapa 3): em disco, deleção vira tombstone. Por quê?
func (s *Store) Delete(key []byte) error {
	return ErrNaoImplementado
}

// Close libera os recursos e garante que tudo foi para o disco.
func (s *Store) Close() error {
	return ErrNaoImplementado
}
