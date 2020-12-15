# Lire et écrire dans un fichier en Go

Vous pouvez efficacement lire et écrire sur des fichiers à l'aide du langage de programmation go.

### Ce que j'ai fais dans le fichier __EcrireSurFichier.go__

Je reviens ici sur ces lignes de code :

```go
file, err := os.OpenFile("test.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
defer file.Close() // on ferme automatiquement le fichier après l'avoir manipulé
```

La fonction ``os.OpenFile()``, propose vraiment beaucoup d'options, que je vous explique ci-dessous :

* Premier paramètre : correspond au nom du fichier à ouvrir
* Deuxième paramètre : ici ce sont des options spécialement dédiées au fichier que vous allez manipuler :
  * ``os.O_CREATE`` : Permet de créer le fichier si il n'existe pas.
  * ``os.O_WRONLY`` : Permet de rendre le fichier (dans votre programme) accessible en écriture seulement.
  * ``os.O_APPEND`` : Permet de ne pas écraser le fichier quand vous écrivez dessus (supprimez cette option si vous souhaitez écraser le fichier).
* Troisième paramètre : les droits d'accès de votre fichier.

**Il est absolument important de fermer le fichier à la fin de votre programme**. D'où l'utilisation de la fonction ``close()``, j'ai rajouté le mot clé ``defer``, ce mot-clé permet d'exécuter la ligne de code en question jusqu'à la fin d'exécution d'une fonction.
