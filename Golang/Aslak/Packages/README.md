# Les packages
## Qu'est-ce que c'est ?

Un package n'est rien d'autre qu'un répertoire contenant des fichiers de code go, qui peut être exécuté par d'autres fichiers go.

### Types de packages

il existe deux types de packages :

* **Package exécutable**
* **Package utilitaire**

Dans un package exécutable on retrouve le fameux **package main** qui permet d'informer votre compilateur Go que le paquet doit être compilé en tant que programme exécutable au lieu d'un package utilitaire.

Un package utilitaire n'est pas auto-exécutable, il améliore plutôt les fonctionnalités d'un package exécutable en lui fournissant des fonctionnalités.

## Où créer mon package ?

Pour importer un paquet, il faut utiliser le mot-clé **import** suivi du `nom du paquet`.

Lorsque vous créez un paquet et que vous l'importez sur votre package principal, votre compilateur va d'abord chercher le répertoire du paquet dans le $GOROOT/src/ et si répertoire n'existe pas il va le cherchera dans le dossier `$GOPATH/src/`.

### Informations

``$GOROOT`` et ``$GOPATH`` sont deux variables d'environnement créez lors de l'installation de GO.

### Variables d'environnement Go 

Il est possible d'afficher toutes vos variables d'environnement Go comme suit :

```go
go env
```
### Résultat :

```
set GOCACHE=C:\Users\Admin\AppData\Local\go-build
...
set GOPATH=C:\Users\Admin\go
set GOROOT=C:\Go
...
```        
