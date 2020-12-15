# Les channels
## Qu'est-ce que c'est ?

Les channels permettent de créer, connecter, synchroniser et communiquer les différentes goroutines.
> Les channels sont utilisés avec les goroutines pour envoyer des données (int, string, struct…) d'une goroutine et les recevoir dans une autre goroutine. C'est un moyen de connecter les différentes goroutines, c'est un moyen de communication et de synchronisation entre les goroutines. **La transmission des channels se fait qu'avec des goroutines.**

### Déclarer un channel

Pour déclarer votre channel, vous utiliserez le mot-clé ``make`` avec le mot-clé ``chan`` suivit du type de donnée que vous souhaitez transiter.

```go
ch := make(chan typeDeValeur)
```

### Les channels sont bloquants

Les envois et les réceptions sur un channel est bloquant par défaut. Qu'est-ce que ça veut dire ? Lorsqu'une donnée est envoyée à un channel, le contrôle est bloqué dans l'instruction d'envoi jusqu'à ce qu'une autre goroutine lise depuis ce channel. De la même manière, lorsque des données sont lues sur un channel, la lecture est bloquée jusqu'à ce qu'une certaine goroutine écrit des données sur ce channel.

C'est cette propriété des channels qui permet aux goroutines de communiquer efficacement sans l'utilisation de verrous explicites ou de variables conditionnelles.

### Deadlock

Par défaut, les channels sont dit **unbuffered**, ce qui signifie qu'ils n'accepteront 
pas de **récepteur** ( chan<-) que s'il existe un **expéditeur** ( <- chan) correspondant prêt à recevoir la valeur envoyée, l'inverse est aussi vrai.

Voila ce que j'entend par expéditeur et récepteur :

* **récepteur** : c'est le moment où on entre une valeur dans notre channel
* **expéditeur** : c'est le moment où on lit une valeur depuis notre channel

### Buffered Channels

il est possible de rendre nos channels buffered, c'est-à-dire de posséder autant de récepteurs que la taille du buffer de notre channel, par exemple un channel avec une taille de buffer de 30 peut posséder 30 récepteurs.

Pour rendre notre channel buffered, il suffit d'indiquer la longueur de notre buffer comme second argument de notre canal.

```ch := make(chan type-de-valeur, taille_du_buffer)```
