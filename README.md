# mini-cloud

Um "mini-AWS" construído do zero, em Go, como **projeto impossível**: um laboratório de engenharia de sistemas que nunca fica pronto.

## Por que existe

Sair da zona de conforto de dados/ML e aprender o que fica por baixo: disco, concorrência, rede, consenso e sistemas distribuídos. Cada peça é construída na mão e usa a anterior.

## Componentes (visão de longo prazo)

| # | Componente | Inspiração | Estado |
|---|-----------|------------|--------|
| 1 | Storage engine chave-valor (WAL, LSM-tree) | RocksDB / LevelDB | **Fase 1 — em andamento** |
| 2 | Object storage com replicação e checksums | S3 | — |
| 3 | Fila / log particionado, consumer groups | Kafka | — |
| 4 | Consenso (eleição de líder, replicação de log) | Raft | — |
| 5 | Orquestrador de containers | Kubernetes | — |
| 6 | Gateway, load balancer, service discovery | Envoy / Consul | — |
| 7 | Observabilidade (métricas, tracing, logs) | Prometheus / OpenTelemetry | — |
| 8 | Inferência de modelos na própria plataforma | SageMaker | — |

## Como trabalhar aqui

- Leia o `CLAUDE.md`: ele define como o Claude Code / Codex deve atuar como professor.
- O plano da fase atual está em `ROADMAP.md`.
- Decisões de design vão em `docs/adr/` (copie o `0000-template.md`).
- Ao fim de cada sessão, escreva no `docs/diario/` com as suas próprias palavras.

## Começando

```bash
go version            # use a versão estável mais recente
go test -race ./...   # no início, os testes vão falhar: é o ponto de partida
```

Referências de apoio: *Designing Data-Intensive Applications* (Kleppmann), curso MIT 6.5840 (Raft), *Build Your Own Database* (James Smith).
