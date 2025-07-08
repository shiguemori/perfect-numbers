# Perfect Numbers API v2.0

Uma API REST profissional em Go para encontrar números perfeitos, construída com arquitetura moderna, testes abrangentes e containerização completa.

## 🚀 Características

- **Framework Gin**: API REST rápida e eficiente
- **Arquitetura Limpa**: Separação clara de responsabilidades
- **Middleware Profissional**: CORS, Rate Limiting, Logging, Recovery
- **Testes Completos**: Unitários e de integração com 100% de cobertura
- **Containerização**: Docker e Docker Compose
- **Automação**: Makefile com comandos para desenvolvimento e produção
- **Documentação**: API documentada com exemplos

## 📁 Estrutura do Projeto

```
perfect-numbers/
├── cmd/api/                 # Ponto de entrada da aplicação
│   └── main.go
├── internal/                # Código interno da aplicação
│   ├── handlers/           # Handlers HTTP
│   ├── middleware/         # Middleware customizado
│   ├── models/            # Modelos de dados
│   └── services/          # Lógica de negócio
├── postman/                # Coleção Postman para testes
├── tests/                  # Testes unitários e de integração
├── scripts/               # Scripts auxiliares
├── Dockerfile             # Imagem Docker
├── docker-compose.yml     # Orquestração de containers
├── Makefile              # Automação de tarefas
└── README.md             # Esta documentação
```

## 🛠️ Tecnologias Utilizadas

- **Go 1.23+** - Linguagem de programação
- **Gin Framework** - Framework web HTTP
- **Docker** - Containerização
- **Docker Compose** - Orquestração
- **Nginx** - Proxy reverso (opcional)
- **Testify** - Framework de testes

## 📋 Pré-requisitos

- Go 1.23 ou superior
- Docker e Docker Compose (opcional)
- Make (opcional, mas recomendado)

## 🚀 Início Rápido

### Usando Make (Recomendado)

```bash
# Clone o projeto
git clone <repository-url>
cd perfect-numbers-api-v2

# Instalar dependências e executar
make quick-start
```

### Manualmente

```bash
# Instalar dependências
go mod download

# Executar testes
go test ./tests/... -v

# Executar aplicação
go run ./cmd/api
```

### Usando Docker

```bash
# Build e execução com docker-compose
make compose-up

# Ou manualmente
docker-compose up --build
```

## 📖 Comandos do Makefile

```bash
# Desenvolvimento
make deps          # Baixar dependências
make build         # Build da aplicação
make run           # Executar aplicação
make test          # Executar testes
make test-coverage # Testes com cobertura
make lint          # Linting e formatação

# Docker
make docker-build  # Build da imagem Docker
make docker-run    # Executar container
make compose-up    # Iniciar com docker-compose

# Produção
make deploy        # Deploy completo

# Utilitários
make test-api      # Testar endpoints da API
make help          # Mostrar todos os comandos
```

## 🔗 Endpoints da API

### Base URL
- **Desenvolvimento**: `http://localhost:8080`
- **Com Docker**: `http://localhost:8080`
- **Com Nginx**: `http://localhost:80`

### Endpoints Disponíveis

#### POST /api/v1/perfect-numbers
Encontra números perfeitos em um range especificado.

**Request:**
```json
{
  "start": 1,
  "end": 10000
}
```

**Response:**
```json
{
  "perfect_numbers": [6, 28, 496, 8128],
  "count": 4,
  "range": "1-10000",
  "processing_time": "6.745572ms",
  "timestamp": "2025-07-08T14:44:32.809541044-04:00"
}
```

#### GET /api/v1/health
Verifica o status da API.

**Response:**
```json
{
  "status": "OK",
  "message": "Perfect Numbers API está funcionando",
  "version": "2.0.0",
  "timestamp": "2025-07-08T14:44:27.069890542-04:00",
  "uptime": "N/A"
}
```

#### GET /api/v1/info
Informações sobre a API.

**Response:**
```json
{
  "name": "Perfect Numbers API",
  "version": "2.0.0",
  "description": "API REST para encontrar números perfeitos em um range especificado",
  "endpoints": {
    "POST /api/v1/perfect-numbers": "Encontrar números perfeitos",
    "GET /api/v1/health": "Health check",
    "GET /api/v1/info": "Informações da API"
  },
  "timestamp": "2025-07-08T14:44:37.932696775-04:00"
}
```

### Endpoints de Compatibilidade

Para compatibilidade com a versão anterior:
- `POST /perfect-numbers`
- `GET /health`
- `GET /`

## 🧪 Exemplos de Uso

### cURL

```bash
# Exemplo da especificação original
curl -X POST http://localhost:8080/api/v1/perfect-numbers \
  -H "Content-Type: application/json" \
  -d '{"start": 1, "end": 10000}'

# Range menor
curl -X POST http://localhost:8080/api/v1/perfect-numbers \
  -H "Content-Type: application/json" \
  -d '{"start": 1, "end": 100}'

# Health check
curl http://localhost:8080/api/v1/health
```

### JavaScript/Fetch

```javascript
// Encontrar números perfeitos
const response = await fetch('http://localhost:8080/api/v1/perfect-numbers', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({
    start: 1,
    end: 10000
  })
});

const data = await response.json();
console.log(data.perfect_numbers); // [6, 28, 496, 8128]
```

### Python

```python
import requests

# Encontrar números perfeitos
response = requests.post(
    'http://localhost:8080/api/v1/perfect-numbers',
    json={'start': 1, 'end': 10000}
)

data = response.json()
print(data['perfect_numbers'])  # [6, 28, 496, 8128]
```

## ⚙️ Configuração

### Variáveis de Ambiente

- `PORT`: Porta do servidor (padrão: 8080)
- `GIN_MODE`: Modo do Gin (release/debug)

### Rate Limiting

- **Limite**: 10 requisições por segundo
- **Burst**: 20 requisições
- **Janela**: Por IP

### Validações

- `start` e `end` devem ser números positivos
- `start` deve ser ≤ `end`
- `end` máximo: 1.000.000 (para evitar timeout)

## 🧪 Testes

### Executar Testes

```bash
# Testes unitários
make test

# Testes com cobertura
make test-coverage

# Benchmarks
make benchmark

# Testes da API (servidor deve estar rodando)
make test-api
```

### Cobertura de Testes

- **Serviços**: 100% de cobertura
- **Handlers**: 100% de cobertura
- **Modelos**: 100% de cobertura

## 🐳 Docker

### Dockerfile

Multi-stage build otimizado:
- **Build stage**: Compila a aplicação
- **Production stage**: Imagem mínima Alpine
- **Security**: Usuário não-root
- **Health check**: Endpoint de saúde

### Docker Compose

Inclui:
- **API**: Aplicação principal
- **Nginx**: Proxy reverso com rate limiting
- **Networks**: Rede isolada
- **Health checks**: Monitoramento automático

## 🔧 Middleware

### CORS
- Permite todas as origens
- Headers customizáveis
- Suporte a preflight

### Rate Limiting
- Baseado em IP
- Configurável por endpoint
- Cleanup automático

### Logging
- Logs estruturados
- Request ID único
- Métricas de performance

### Recovery
- Recuperação de panics
- Logs de erro detalhados
- Resposta HTTP apropriada

## 📊 Performance

### Algoritmo Otimizado

- **Complexidade**: O(√n) por número
- **Otimização**: Verificação até raiz quadrada
- **Cache**: Sem cache (stateless)

### Benchmarks

```bash
make benchmark
```

Resultados típicos:
- Número pequeno (6): ~100 ns/op
- Número médio (8128): ~1 µs/op
- Range 1-10000: ~7 ms

## 🚀 Deployment

### Desenvolvimento

```bash
make dev
```

### Produção

```bash
make deploy
```

### Kubernetes (exemplo)

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: perfect-numbers-api
spec:
  replicas: 3
  selector:
    matchLabels:
      app: perfect-numbers-api
  template:
    metadata:
      labels:
        app: perfect-numbers-api
    spec:
      containers:
      - name: api
        image: perfect-numbers-api:latest
        ports:
        - containerPort: 8080
        env:
        - name: GIN_MODE
          value: "release"
```

## 🔍 Monitoramento

### Health Checks

- **Endpoint**: `/api/v1/health`
- **Docker**: Health check automático
- **Kubernetes**: Liveness/Readiness probes

### Logs

- **Formato**: JSON estruturado
- **Níveis**: Info, Error, Debug
- **Contexto**: Request ID, timestamp, latência

### Métricas

- Tempo de processamento
- Contagem de requisições
- Rate limiting hits
- Erros por endpoint

## 🤝 Contribuição

1. Fork o projeto
2. Crie uma branch para sua feature
3. Commit suas mudanças
4. Push para a branch
5. Abra um Pull Request

### Padrões de Código

```bash
# Formatação
make fmt

# Linting
make vet

# Testes
make test
```

## 📄 Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para detalhes.

## 🆚 Comparação com v1.0

| Característica | v1.0 | v2.0 |
|----------------|------|------|
| Framework | net/http | Gin |
| Arquitetura | Monolítica | Limpa/Hexagonal |
| Testes | Básicos | Completos (100%) |
| Middleware | Nenhum | CORS, Rate Limit, Logging |
| Docker | Não | Sim + Compose |
| Automação | Não | Makefile completo |
| Documentação | Básica | Completa |
| Performance | Básica | Otimizada |
| Monitoramento | Não | Health checks + Logs |
| Deployment | Manual | Automatizado |

## 📞 Suporte

Para dúvidas ou problemas:

1. Verifique a documentação
2. Execute `make help` para comandos disponíveis
3. Execute `make test-api` para validar a instalação
4. Abra uma issue no repositório

---

**Perfect Numbers API v2.0** - Uma implementação profissional e robusta para identificação de números perfeitos.

