# Certains d'entre vous se demandent peut-être l'intérêt de 
# passer des chaînes en minuscules… alors voici un petit exemple.

chaine = str() # Crée une chaîne vide
# On aurait obtenu le même résultat en tapant chaine = ""

while chaine.lower() != "q":
    print("Tapez 'Q' pour quitter...")
    chaine = input()

print("Merci !")

# Vous devez comprendre rapidement ce programme. 
# Dans une boucle, on demande à l'utilisateur de taper la lettre « q » 
# pour quitter. 
# Tant que l'utilisateur saisit une autre lettre, 
# la boucle continue de s'exécuter. 
# Dès que l'utilisateur appuie sur la touche "Q" de son clavier, 
# la boucle s'arrête et le programme affiche « Merci ! ». 

# La petite nouveauté réside dans le test de la boucle : "chaine.lower() != "q"."
# On prend la chaîne saisie par l'utilisateur, on la passe en minuscules et on regarde si elle est différente de « q ». 
# Cela veut dire que l'utilisateur peut taper « q » en majuscule ou en minuscule, dans les deux cas la boucle s'arrêtera.

# Notez que "chaine.lower()" renvoie la chaîne en minuscules mais ne modifie pas la chaîne. 
