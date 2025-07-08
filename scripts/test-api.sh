#!/bin/bash

# Script para testar a API Perfect Numbers
# Uso: ./scripts/test-api.sh [HOST] [PORT]

HOST=${1:-localhost}
PORT=${2:-8080}
BASE_URL="http://${HOST}:${PORT}"

echo "=== Testando Perfect Numbers API ==="
echo "URL Base: $BASE_URL"
echo

# Cores para output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

test_endpoint() {
    local method=$1
    local endpoint=$2
    local data=$3
    local expected_status=$4
    local description=$5
    
    echo -e "${BLUE}Testando: $description${NC}"
    echo "Endpoint: $method $endpoint"
    
    if [ "$method" = "POST" ]; then
        response=$(curl -s -w "\n%{http_code}" -X POST "$BASE_URL$endpoint" \
            -H "Content-Type: application/json" \
            -d "$data")
    else
        response=$(curl -s -w "\n%{http_code}" "$BASE_URL$endpoint")
    fi
    
    # Separar body e status code
    body=$(echo "$response" | head -n -1)
    status_code=$(echo "$response" | tail -n 1)
    
    if [ "$status_code" = "$expected_status" ]; then
        echo -e "${GREEN}✅ PASSOU${NC} (Status: $status_code)"
        if command -v jq &> /dev/null; then
            echo "$body" | jq .
        else
            echo "$body"
        fi
    else
        echo -e "${RED}❌ FALHOU${NC} (Status: $status_code, Esperado: $expected_status)"
        echo "$body"
    fi
    echo
}

# Verificar se o servidor está rodando
echo -e "${YELLOW}Verificando se o servidor está rodando...${NC}"
if ! curl -s "$BASE_URL/health" > /dev/null; then
    echo -e "${RED}❌ Servidor não está respondendo em $BASE_URL${NC}"
    echo "Certifique-se de que a API está rodando:"
    echo "  make run"
    echo "  ou"
    echo "  make docker-run"
    exit 1
fi
echo -e "${GREEN}✅ Servidor está rodando${NC}"
echo

# Testes
test_endpoint "GET" "/health" "" "200" "Health Check"

test_endpoint "GET" "/api/v1/info" "" "200" "API Info"

test_endpoint "POST" "/api/v1/perfect-numbers" '{"start": 1, "end": 10}' "200" "Range pequeno (1-10)"

test_endpoint "POST" "/api/v1/perfect-numbers" '{"start": 1, "end": 100}' "200" "Range médio (1-100)"

test_endpoint "POST" "/api/v1/perfect-numbers" '{"start": 1, "end": 10000}' "200" "Range grande (1-10000) - Exemplo da especificação"

test_endpoint "POST" "/api/v1/perfect-numbers" '{"start": 28, "end": 28}' "200" "Range específico (28-28)"

test_endpoint "POST" "/api/v1/perfect-numbers" '{"start": 10, "end": 5}' "400" "Validação: Start > End"

test_endpoint "POST" "/api/v1/perfect-numbers" '{"start": -1, "end": 10}' "400" "Validação: Start negativo"

test_endpoint "POST" "/api/v1/perfect-numbers" '{"start": 1, "end": 2000000}' "400" "Validação: End muito grande"

test_endpoint "POST" "/api/v1/perfect-numbers" '{"start": "invalid", "end": 10}' "400" "Validação: JSON inválido"

test_endpoint "POST" "/perfect-numbers" '{"start": 1, "end": 100}' "200" "Compatibilidade: Endpoint sem versão"

echo -e "${BLUE}=== Testes concluídos ===${NC}"

