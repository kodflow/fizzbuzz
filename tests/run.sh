#!/bin/sh
# test_fizzbuzz.sh
# Script pour tester le endpoint /fizzbuzz de manière aléatoire
#
# Ce script génère des valeurs aléatoires pour int1, int2, limit, str1 et str2,
# puis il utilise curl pour envoyer une requête GET au endpoint /fizzbuzz.
#
# Parameters:
# - Aucun
#
# Returns:
# - Affiche la réponse du serveur

ENDPOINT="http://localhost/fizzbuzz" 

for i in {1..100000}; do
    # Générer des valeurs aléatoires
    INT1=$((RANDOM % 10 + 1))
    INT2=$((RANDOM % 10 + 1))
    LIMIT=$((RANDOM % 100 + 1))
    STR1="Fizz"
    STR2="Buzz"

    # Construire l'URL
    URL="${ENDPOINT}/${INT1}/${INT2}/${LIMIT}/${STR1}/${STR2}"

    # Envoyer la requête GET
    echo "Test $i: GET $URL"
    RESPONSE=$(curl -s "$URL")
done
