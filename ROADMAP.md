# ROADMAP — Fase 1: storage engine chave-valor

Meta: um banco chave-valor persistente, durável e concorrente, escrito em Go sem dependências. Duração estimada: 2 a 3 semanas, no seu ritmo.

Regra do jogo: para cada etapa, **você escreve primeiro**; o professor (Claude Code / Codex) pergunta, revisa e quebra o seu código de propósito.

## Etapa 0 — Preparação (1 sessão)

- [ ] Instalar Go e rodar `go test ./...` (os testes iniciais devem falhar).
- [ ] Ler `internal/kv/kv_test.go` e explicar com suas palavras o que cada teste exige.
- [ ] Registrar o ADR 0001: formato de chave/valor (só `[]byte`? strings?).

**Perguntas-guia:** o que é uma goroutine? o que um `sync.Mutex` protege e o que ele não protege?

## Etapa 1 — Memória + API (1 a 2 sessões)

- [ ] `Put`, `Get`, `Delete` em memória (um `map` protegido por `sync.RWMutex`).
- [ ] Teste de concorrência com `go test -race`.

**Experimento de quebra:** remova o lock e rode com `-race`. Leia o relatório de data race.

## Etapa 2 — Durabilidade com WAL (3 a 4 sessões)

- [ ] Escrever cada operação num arquivo append-only (write-ahead log) antes de aplicar na memória.
- [ ] Formato de registro com tamanho e checksum (CRC32).
- [ ] Recuperação: ao abrir, reler o WAL e reconstruir o estado.
- [ ] Entender `fsync`: quando chamar, quanto custa.

**Perguntas-guia:** e se a máquina cair no meio de um registro? e se o arquivo ficar com a cauda corrompida?
**Experimento de quebra:** truncar o WAL no meio de um registro e provar que a recuperação ignora a cauda ruim sem perder o que estava íntegro.
**ADR sugerido:** `fsync` a cada escrita vs em lote (trade-off durabilidade x throughput).

## Etapa 3 — Memtable + SSTables (4 a 5 sessões)

- [ ] Quando a memória passar de um limite, gravar os dados ordenados num arquivo imutável (SSTable).
- [ ] Leitura procura na memtable e depois nas SSTables, da mais nova para a mais antiga.
- [ ] Deleção por tombstone.
- [ ] Índice esparso para não varrer o arquivo inteiro.

**Perguntas-guia:** por que ordenar? por que arquivos imutáveis? o que acontece com uma chave sobrescrita 10 vezes?

## Etapa 4 — Compactação e Bloom filter (3 a 4 sessões)

- [ ] Compactação: juntar SSTables, descartando versões antigas e tombstones.
- [ ] Bloom filter para evitar leituras de disco de chaves que não existem.
- [ ] Benchmark antes/depois de cada otimização, com número, não com sensação.

**Experimento de quebra:** 10 milhões de escritas e medir amplificação de escrita e leitura.

## Etapa 5 — Servidor (2 sessões)

- [ ] Expor o KV via TCP ou HTTP simples (`cmd/kvserver`).
- [ ] Cliente de linha de comando.
- [ ] Teste de carga com vários clientes simultâneos.

## Critério para fechar a fase 1

Você consegue explicar, sem olhar o código, o caminho completo de um `Put` até o disco e de um `Get` de volta, e o que acontece se a energia cair em cada ponto desse caminho.

## Próximas fases (não adiantar)

Fase 2: object storage · Fase 3: fila/log particionado · Fase 4: Raft · Fase 5: orquestrador · e assim por diante (ver `README.md`).
