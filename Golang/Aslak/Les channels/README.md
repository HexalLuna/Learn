# Les channels
## Qu'est-ce que c'est ?

Les channels permettent de créer, connecter, synchroniser et communiquer les différentes goroutines.
> Les channels sont utilisés avec les goroutines pour envoyer des données (int, string, struct…) d'une goroutine et les recevoir dans une autre goroutine. C'est un moyen de connecter les différentes goroutines, c'est un moyen de communication et de synchronisation entre les goroutines. **La transmission des channels se fait qu'avec des goroutines.**

### Déclarer un channel

Pour déclarer votre channel, vous utiliserez le mot-clé ``make`` avec le mot-clé ``chan`` suivit du type de donnée que vous souhaitez transiter.

```go
ch := make(chan typeDeValeur)
```
