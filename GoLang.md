# GoLang 

[TOC]

## Introduction

### Commandes

```bash
go run main.go // Exécute le fichier main.go
go build // Build et crée un exécutable
go install // Build et place dans le dossier /bin. Le dossier bin étant dans le path, on peut utiliser le programme depuis partout.
```

### Hello World

```go
package main

import "fmt"
func main() {
    fmt.PrintLn("Hello World")
}
```

### Packages

Tout le code doit se situer dans le **$GOPATH/src**. Généralement, un package et le nom du dossier doivent avoir le même nom. Les fichiers et les packages ont des noms en **minuscule**. 

```go
package nomdupackage
```

Pour importer nos propres packages, on doit utiliser le **chemin complet**. Un import par ligne.

```go
import (
	"fmt" // Pour les packages de la lib
    
    "training.go/helloworld/input" // Chemin complet pour les autres
)
```

## Variables

### Les différents types

```go
bool
string
int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32
// represents a Unicode code point
float32 float64
complex64 complex128
```

### Déclaration des variables

Syntaxe de base:

```go
var age int = 20
var age int // age = 0

var name string = "Bob"
var name string // name = ""
```

Syntaxe raccourcie :

```go
var age = 20
var name = "Bob"

age := 20 // Impossible à l'extérieur d'une fonction
name := "Bob" // Pareil
```

### Visibilité des variables

La visibilité d'une variable se fait en fonction de la **case**. Pareil pour les **fonctions**, **structs**...

```go
var name = "Bob" // Visible uniquement par le package
var Name = "Bob" // Visible partout (si un programme importe notre package, il aura accès à cette variable)
```

## Les bases du langage

### Comparaisons

| En français           | En Go |
| --------------------- | ----- |
| Egal                  | ==    |
| Différent             | !=    |
| Strictement inférieur | <     |
| Strictement supérieur | >     |
| Inférieur ou égal     | <=    |
| Supérieur ou égal     | >=    |
| Et                    | &&    |
| Ou                    | \|\|  |
| Non                   | !     |

### Conditions

```go
age :=15
if age > 10 {
    // code
} else if age > 5 {
    // code
} else {
    // code
}
```

### Printf

Voir documentation [ici](https://golang.org/pkg/fmt/).

### Switch

Normal:

```go
switch <variable> {
    case 1: // code
    case 2, 3: // code
    default: // Default code
}
```

Autre:

```go
heure := 10;
switch {
    case heure < 6:
    	// c'est la nuit
	case heure < 18 && heure > 14: 
    	// c'est l'aprem
	default : 
    	//autre
}
```

### Les fonctions

Logique de visibilité avec la première lettre(maj ou non).

```go
func <nomdelafnct> (<param1> type, <param2> type) <TypeRetour> {
    // code
}
```

Le type de retour est **absent** lorsqu'une fonction ne renvoie rien.

On peut **déclarer** la variable de retour dans la signature de la fonction. Exemple :

```go
func sumNamedReturn(x, y, z int) (sum int) {
	sum = x + y + z
    return sum // le sum est facultatif (on peut juste mettre return) mais déconseillé
}
```

Possibilité de return **plusieurs** variables :

```go
func MultReturn() (int,string) {
    return 4, "hello"
}
num, name := MultReturn() // exemple d'appel
_, name := MultReturn() // ignore la premiere variable
```

### Les tableaux

La taille d'un tableau est **définitive**. Le contenu est toujours initialisé à la valeur par **défaut** (0 pour les integers, "" pour les strings...)

```go
var t[5] int; //Déclaration
odds := [6]int{1,3,5,7} // Déclation avec valeurs
t[3] = 12; //Affectation
fmt.Println("names=%v", t) // Affiche le tableau en entier, peu importe son type 
len(t) // Longeur d'un tableau
```

### Les slices

Ce sont des tableaux de taille **dynamique**. 

```go
s := make([]int, 3) // Crée un tableau avec len(3) (et cap de 3)
letters := []string{"g", "o", "l", "a", "n", "g"} // Crée automatiquement un slice
s = append(s,12) 

sub1 := letters[:2] // l a n g (pointeur)
sub1 :=letters[1:3] // g o l (pointeur)

subCopy = make([]string, len(sub1))
copy(subCopy, sub1) // Copie donc plus de pointeurs

```

Lorsque on ajoute un élément à un slice :

- Si on dépasse la taille du tableau
- Un nouveau tableau est alloué, de capacité doublée

Plus d'information [ici](https://blog.golang.org/go-slices-usage-and-internals)

### Boucles

```go
for i:=0; i<5 ; i++ { // version longue
    fmt.Println(i)
}

i:=1
for i <= 3 { // version while
    fmt.Println(i)
    i = i+1
}

for { // boucle infinie
    // code
    continue // permet de sauter au tour suivant
    break // sort de la boucle
}
```

### Range

Permet **d'itérer** sur des tableaux

```go
nums := []int{15,-2}
for i, num := range nums {
    fmt.Printf("nums[%d]=%d\n", i, num)
}

for _, c := range "golang" {
    fmt.Printf("%v", string(c)) // Obliger de cast en string sinon on a la valeur du byte
}
```

### Gestion d'erreurs

Contrairement au Java ou le C#, Go n'utilise pas les exceptions. Il utilise des **codes d'erreurs** (comme C) :

```go
func MyFunc() (int, error) {...}

v, err := MyFunc()
if err!=nil {
    // il y a eu une erreur
}
```

Exemple plus complet : 

```go
func readFile(filename string) (string, error) {
    dat, err := ioutil.ReadFile(filename)
    if err != nil {
        return "", err
    }
    if len(dat) == 0 {
        return "", fmt.Errorf("Empty content (filename=%v)", filename)
    }
    return dat, nil
}

func main() {
    dat, err := readFile("test.txt")
    if err != nil {
        fmt.Printf("Error while reading file: %v\n", err)
    }
}
```



### Defer

`defer` permet d'exécuter du code à la sortie d'une fonction. Cela peut -être très pratique pour fermer un fichier a la sortie d'une fonction par exemple. **Attention** `defer` fonctionne en mode **LIFO**(Last In First Out).

````go
func myNameIs() {
    defer fmt.Println("Slim Shady")
    fmt.Println("My name is")
}
func main() {
    defer fmt.Println("Hello")
    defer fmt.Println("Goodbye")
    myNameIs()
}
// On obtient dans l'ordre : My name is / Slim Shady / Goodbye / Hello
````



## Structs

### Définition

```go
type Address struct {
    street, city string
}
type User struct {
    Name string
    Email string
    Age int
    Addr Address
}
func main() {
    var p Person
    p.Name = "Bob"
    p.Addr.city = "Lyon"
}
```

Quelques points sur les structs 

- Un `struct` ne peut contenir que des variables (pas de méthodes).
- La règle de visibilité de package s'applique pour :
  - le `struct` lui même.
  - les variables à l'intérieur.

### Déclaration

 ```go
// Un à un
var p1 player.Player;
p1.Name = "Bob"
p1.Age = 10

// Clé valeur
p2 := player.Player {
    Name: "Bob",
    Age: 10,
    Avatar: player.Avatar{
        Url : "http://url.jpg"
    }
}

// Par une méthode (conseillé)
// Dans le package player
func New(name string) Player {
    return Player {
        Name : name
    }
}
// Dans le main
p3 := player.New("Bob")
 ```

### Embedded Struct

L'héritage en go n'existe pas. En effet, Go a choisi une approche totalement différente. On préfère la composition avec l'Embedded Struct. Exemple : 

```go
type Address struct {
    City string
}
type User struct {
    Name string
    Address // pas de nom de variable!
}
```

Ici, un User est une adresse. On peut donc directement écrire :

```go
var u = User
u.City = "Londres" // City est directement accessible
```

Exemple plus concret [ici](02-exemple-struct/main.go).

### Struct et Receiver function

Permet d'écrire des méthodes pour des structs.

```go
func (u User) SayHello() {
    fmt.Printf("Hello %v!\n", u.Name)
}
```

Utilisation : 

```go
u := User{"Bob"}
u.SayHello() 
```

**Attention**, la méthode SayHello() reçoit une copie de u. On parle de méthode avec un **value receiver**, SayHello() ne peut donc pas **modifier** le struct original. En créant une méthode nommée `String() string` on redéfinie la manière d'afficher notre struct (équivalent à toString en Java).

### Pointeurs

```go
x := -42
s := "Bob"
c := &x // Création d'un pointer vers la variable x
i := *p // Recupère la valeur de ce sur quoi pointe p (-42) 
```

Exemple plus complet [ici](03-exemple-pointeurs/main.go])

### Pointer Receiver

Permet de modifier les structs grâce aux receivers.

```go
type User struct {
    Name string
}
func (u *User) SayHello() {
    fmt.Printf("Hello %v", u.Name)
}

func (u *User) UpdateName(name string) {
    u.Name := name
}
```

Exemple plus complet [ici](04-exemple-pointer-receiver/main.go)

## Maps

