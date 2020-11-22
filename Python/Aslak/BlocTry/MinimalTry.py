# On va parler ici de bloc "try". 
# Nous allons en effet mettre les instructions que nous souhaitons tester dans un premier bloc 
# et les instructions à exécuter en cas d'erreur dans un autre bloc. Sans plus attendre, voici la syntaxe :

try:
    # Bloc à essayer
except:
    # Bloc qui sera exécuté en cas d'erreur

# Dans l'ordre, nous trouvons :

# - Le mot-clé "try" suivi des deux points « : » (try signifie « essayer » en anglais).
# - Le bloc d'instructions à essayer.
# - Le mot-clé "except" suivi, une fois encore, des deux points « : ». Il se trouve au même niveau d'indentation que le "try".
# - Le bloc d'instructions qui sera exécuté si une erreur est trouvée dans le premier bloc.

# Reprenons notre test de conversion en enfermant dans un bloc "try" l'instruction susceptible de lever une exception.

annee = input()
try: # On essaye de convertir l'année en entier
    annee = int(annee)
except:
    print("Erreur lors de la conversion de l'année.")

# Vous pouvez tester ce code en précisant plusieurs valeurs différentes pour la variable "annee", comme « 2010 » ou « annee2010 ».