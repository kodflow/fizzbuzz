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

THE_ONE="${ENDPOINT}/3/5/100/Fizz/Buzz"
for i in {1..20}; do
    # Construire l'URL
    $(curl -s "$THE_ONE" > /dev/null)
done

for i in {1..100000}; do
    # Générer des valeurs aléatoires
    INT1=$((RANDOM % 10 + 1))
    INT2=$((RANDOM % 10 + 1))
    LIMIT=$((RANDOM % 100 + 1))
    STR1="Fizz"
    STR2="Buzz"

    if ((i % 200 == 0)); then
        $(curl -s "$THE_ONE" > /dev/null)
    fi

    # Construire l'URL
    URL="${ENDPOINT}/${INT1}/${INT2}/${LIMIT}/${STR1}/${STR2}"
    # Envoyer la requête GET
    echo "Test $i: GET $URL"
    $(curl -s "$URL" > /dev/null)
done
