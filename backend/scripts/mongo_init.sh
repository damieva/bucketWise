#!/bin/bash
set -e

DB_NAME="${MONGO_INITDB_DATABASE:-bucketWise}"

echo "🏗️ Inicializando base de datos '$DB_NAME'..."

mongosh <<EOF
use $DB_NAME

if (!db.getCollectionNames().includes("categories")) {
  db.createCollection("categories");
  print("✅ Colección 'categories' creada");
} else {
  print("ℹ️ Colección 'categories' ya existe");
}

if (!db.getCollectionNames().includes("transactions")) {
  db.createCollection("transactions");
  print("✅ Colección 'transactions' creada");
} else {
  print("ℹ️ Colección 'transactions' ya existe");
}
EOF
