import fonctions
from donnees import tries_number
import os

print("""Bienvenue, on va jouer au jeu du pendu. 
        Vous sélectionnez une lettre et cherchez le mot""")



mystery_word = fonctions.mystery_word()

starred = fonctions.starred_word(mystery_word)

starred_word = starred[0]

list_word = starred[1]

# Demande votre nom
playername = input('rentrez votre nom : ')

fonctions.print_player_score(playername)
print('Devinez le mot : ',starred_word)


i = 0
# 10 chances
while i < tries_number:
    user_char = input('Choisissez votre caractère : ').lower()

    test_user_char = fonctions.test_char(user_char, list_word, starred_word)

    starred_word = test_user_char
    print(starred_word)

    user_answer = input("Essayer de déviner le mot ou appuyer sur start (reste " + str(10-i) + " chances) : ").lower()  

    test_user_answer = fonctions.test_answer(user_answer, mystery_word, playername)
    if test_user_answer == True:
        break
    else:
        print(starred_word)
    if i == 9:
        print("""Vous avez perdu ! Ooh le mauvais ! 
        tu peux tricher en regardant le fichier donnees.py !
        solution : """, mystery_word)
        fonctions.print_score()
    i += 1


# pour windows
os.system("pause")