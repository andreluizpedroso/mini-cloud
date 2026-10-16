package kv

import (
	"bytes"
	"errors"
	"fmt"
	"sync"
	"testing"
)

func abrir(t *testing.T, dir string) *Store {
	t.Helper()
	s, err := Open(dir)
	if err != nil {
		t.Fatalf("Open: %v", err)
	}
	return s
}

// Etapa 1: API básica.
func TestPutGet(t *testing.T) {
	s := abrir(t, t.TempDir())
	defer s.Close()

	if err := s.Put([]byte("a"), []byte("1")); err != nil {
		t.Fatalf("Put: %v", err)
	}
	got, err := s.Get([]byte("a"))
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if !bytes.Equal(got, []byte("1")) {
		t.Fatalf("Get = %q, quero %q", got, "1")
	}
}

func TestGetInexistente(t *testing.T) {
	s := abrir(t, t.TempDir())
	defer s.Close()

	_, err := s.Get([]byte("nao-existe"))
	if !errors.Is(err, ErrNaoEncontrada) {
		t.Fatalf("erro = %v, quero ErrNaoEncontrada", err)
	}
}

func TestSobrescrever(t *testing.T) {
	s := abrir(t, t.TempDir())
	defer s.Close()

	s.Put([]byte("k"), []byte("v1"))
	s.Put([]byte("k"), []byte("v2"))
	got, _ := s.Get([]byte("k"))
	if !bytes.Equal(got, []byte("v2")) {
		t.Fatalf("Get = %q, quero v2", got)
	}
}

func TestDelete(t *testing.T) {
	s := abrir(t, t.TempDir())
	defer s.Close()

	s.Put([]byte("k"), []byte("v"))
	if err := s.Delete([]byte("k")); err != nil {
		t.Fatalf("Delete: %v", err)
	}
	if _, err := s.Get([]byte("k")); !errors.Is(err, ErrNaoEncontrada) {
		t.Fatalf("após Delete, erro = %v, quero ErrNaoEncontrada", err)
	}
	if err := s.Delete([]byte("nunca-existiu")); err != nil {
		t.Fatalf("Delete de chave inexistente não deve dar erro: %v", err)
	}
}

// Etapa 1: rode com `go test -race`.
func TestConcorrencia(t *testing.T) {
	s := abrir(t, t.TempDir())
	defer s.Close()

	var wg sync.WaitGroup
	for g := 0; g < 8; g++ {
		wg.Add(1)
		go func(g int) {
			defer wg.Done()
			for i := 0; i < 200; i++ {
				k := []byte(fmt.Sprintf("g%d-k%d", g, i))
				if err := s.Put(k, k); err != nil {
					t.Errorf("Put: %v", err)
					return
				}
				if _, err := s.Get(k); err != nil {
					t.Errorf("Get: %v", err)
					return
				}
			}
		}(g)
	}
	wg.Wait()
}

// Etapa 2: durabilidade. Fechar e reabrir não pode perder dados.
func TestPersistencia(t *testing.T) {
	dir := t.TempDir()

	s := abrir(t, dir)
	s.Put([]byte("sobrevive"), []byte("sim"))
	s.Put([]byte("apagada"), []byte("x"))
	s.Delete([]byte("apagada"))
	if err := s.Close(); err != nil {
		t.Fatalf("Close: %v", err)
	}

	s2 := abrir(t, dir)
	defer s2.Close()

	got, err := s2.Get([]byte("sobrevive"))
	if err != nil || !bytes.Equal(got, []byte("sim")) {
		t.Fatalf("após reabrir: Get = %q, %v", got, err)
	}
	if _, err := s2.Get([]byte("apagada")); !errors.Is(err, ErrNaoEncontrada) {
		t.Fatalf("chave apagada voltou após reabrir: %v", err)
	}
}

// Etapa 2 (experimento de quebra): TODO seu.
// Escreva um teste que trunque o arquivo do WAL no meio de um registro e
// prove que Open recupera tudo que estava íntegro e ignora a cauda ruim.
func TestRecuperaDeWALTruncado(t *testing.T) {
	t.Skip("TODO: escreva você mesmo este teste na etapa 2")
}
