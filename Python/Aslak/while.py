# Ici nous allons étudier la boucle while

# On rajoute pour compliqué la chose une condition
# Si n est supérieur à 5, afficher reponse correcte et terminer le programme
# Sinon, afficher reponse incorecte et redemander de donner un nombre n supérieur ou égale à 5

while True:
    n = int(input("donnez un nombre n > ou = à 5 : "))
    print("vous avez répondu", n)
    if n >= 5:
        print("reponse correcte")
        break
    else:
        print("reponse incorrecte")
