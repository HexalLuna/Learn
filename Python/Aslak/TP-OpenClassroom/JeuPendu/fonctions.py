# -*-coding:Utf-8 -*
import donnees
import random
import os
import pickle

def mystery_word():

    random_number = random.randrange(0,len(donnees.words)-1)
    return donnees.words[random_number]

def starred_word(word):

    i = 0
    my_list = list()
    starred_word = ""
    while i < len(word):
        my_list.append(word[i])
        starred_word = starred_word + '*'
        i += 1
    return starred_word, my_list

def test_char(char, list_word, starred_word):


    while len(char) != 1:
        char = input('Vous ne pouvez choisir qu\'une seule lettre : ').lower()
    starred_w_list = list(starred_word)
    i = 0

    while i < len(list_word):
        if char == list_word[i]:
            print('vous avez trouvé un caractère')
            starred_w_list[i] = char
        i += 1
    starred_w = ''.join(starred_w_list)
    return starred_w

def test_answer(answer, word, playername):

    if answer == word:
        print('Bravo vous avez gagné')
        save_score(playername)
        return True
    else:
        print('ca n\'est pas ça')
        return False

def save_score(playername):

    # i read file if exists
    if os.path.exists('score'):
        with open('score','rb') as file:
            my_unpickler =  pickle.Unpickler(file)
            if os.path.getsize('score') > 0:
                score_recup = my_unpickler.load()
            else:
                score_recup = {}

        with open('score','wb') as file:
            my_pickler = pickle.Pickler(file)
            if playername in score_recup:
                score_recup[playername] += 1
                my_pickler.dump(score_recup)
                print('Votre score a été sauvegardé')
            else:
                score_recup[playername] = 1
                print(score_recup)
                my_pickler.dump(score_recup)
            print('Votre score a été sauvegardé')
        print_score()
    else:

        with open('score','wb') as file:
            my_pickler = pickle.Pickler(file)
            score = {playername : 1}
            my_pickler.dump(score)
        print_score()

def print_score():

    if os.path.exists('score'):
        with open('score','rb') as file:
            my_unpickler =  pickle.Unpickler(file)
            score = my_unpickler.load()
            print("Rappel des scores : ", score)

def print_player_score(playername):
 
    if os.path.exists('score'):
        with open('score','rb') as file:
            if os.path.getsize('score') > 0:
                my_unpickler =  pickle.Unpickler(file)
                score = my_unpickler.load()
                if playername in score:
                    print("{} a {} points".format(playername,score[playername]))
