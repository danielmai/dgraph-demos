cd "$(dirname "$0")"
openssl rand -hex 8 | tr -d '\n' > ./enc-secret
openssl rand -hex 16 | tr -d '\n' > ./acl-secret
