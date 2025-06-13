#!/bin/bash


info () {
  printf "\r  [ \033[00;34m..\033[0m ] $1\n"
}

user () {
  printf "\r  [ \033[0;33m??\033[0m ] $1\n"
}

success () {
  printf "\r\033[2K  [ \033[00;32mOK\033[0m ] $1\n"
}

fail () {
  printf "\r\033[2K  [\033[0;31mFAIL\033[0m] $1\n"
  echo ''
  exit
}
# src: https://github.com/PatentLobster/dotfiles/blob/main/home/dot_dotfiles/scripts/common.sh

set -euo pipefail

APP_URL="http://127.0.0.1:5000"

info "Verifying /health endpoint..."
health_response=$(curl -s -w "\n%{http_code}" "$APP_URL/health")
health_body=$(echo "$health_response" | head -n1)
health_code=$(echo "$health_response" | tail -n1)

if [ "$health_code" -ne 200 ]; then
  fail "/health failed with status $health_code"
fi

success "/health response: $health_body"

info "Verifying /secret endpoint..."
secret_response=$(curl -s -w "\n%{http_code}" "$APP_URL/secret")
secret_body=$(echo "$secret_response" | head -n1)
secret_code=$(echo "$secret_response" | tail -n1)

if [ "$secret_code" -ne 200 ]; then
  fail "/secret failed with status $secret_code"
fi

success "/secret response: $secret_body"

success "All endpoints verified successfully."
