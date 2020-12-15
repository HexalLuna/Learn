# Les goroutines
## Qu'est-ce que c'est ?

Elles permettent de créer des programmes multi-threads simplement.

## Concurrence et parallélisme

Avant d'aborder des goroutines il est essentiel de comprendre la différence entre concurrence et parallélisme.

### Concurrence

La concurrence est la capacité de **traiter plusieurs de choses à la fois**, par exemple :

Un humain normal doit d'abord finir sa bouchée avant de pouvoir parler, une fois qu'il aura fini sa bouchée il pourra parler ensuite une fois qu'il aura fini de parler il pourra encore une fois reprendre une autre bouchée et reparler juste après.

Dans cet exemple la personne est capable capable de gérer plusieurs de choses (manger et parler) dans un intervalle de temps différent.

### Parallélisme

Le parallélisme permet de **traiter beaucoup de choses en même temps**, je m'explique :

L'humain normal de l'exemple précèdent devient un humain mutant avec un bras et une bouche en plus. Cette fois-ci il est capable de manger et de parler en même temps.

## Pourquoi les goroutines ?

L'un des aspects les plus intéressants dans Go est son modèle de concurrence, il rend la création de programmes multi-threads simples.

Go est capable d'effectuer plusieurs opérations simultanément. C'est particulièrement important sur les processeurs multicœurs actuels. Les programmes n'utilisant qu'un seul cœur laisse une grande partie de la puissance de traitement perdue, coup de chance car Go nous permet d'utiliser pleinement les cœurs de notre processeur grâce aux **goroutines**.
