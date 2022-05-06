# Project Description

## English

For the concurrent programming paradigm, you will only implement the brute force solution of the
knapsack problem. A Go version of this solution is given to you in its regular recursive form (with a
single thread). Your mission is to use concurrency in order to accelerate the resolution of this problem.

Since the brute force approach consists in exploring the binary solution tree, a first naïve approach
would be to subdivide the problem into two parts, the left sub-tree and the right sub-tree, and to create
two threads in order to concurrently solve these subproblems. Once these two threads completed, only
the best solution is kept.

Following this idea, instead of solwing the problem this way:

#### KnapSack(W, weights , values)
  
You should proceed as follow:

#### last := len(weights)-1
//best solution with the last item
#### go KnapSack(W - weights [last], weights [:last], values [:last])
// best solution without the last item
#### go KnapSack(W, weights [:last], values [:last])
//code here to synchronize and then determine which one is the best solution

You can even go one step further and solve the two sub-sub-problems of each two sub-problems above
using 4 concurrent threads. You can also proceed similarly for the next level (8 sub-problems) and so
on. This sounds like a good approach except that at some point the sub-problems will become so simple
(i.e. having so few items) that there won’t be any value in splitting them into concurrent threads.
Moreover, we know that the number of nodes in the solution tree is O(2n); consequently, the number of
threads to create in order to have all sub-problems concurrently solved would be prohibitive. Even if the
go threads are very lightweight they still require some resources such that the total amount of memory
that this exhaustive approach would require could crash your machine.

Consequently if our objective is to make our program to run faster, there should be an optimal choice in
terms of number of threads to create. That is to say that that a sub-problem should be solved
concurrently only if the number of items in this sub-problem is greater than a given threshold.

You are then asked to solve the knapsack problem following the brute force strategy and using
concurrent programming. You have to experimentally determine the optimal number of threads that
should be created in order to obtain a solution as quickly as possible.
Your program must read the input data from a file et display the name of the chosen items, their total
value and the execution runtime. 

## Français

Pour la programmation suivant le paradigme concurrent, vous aurez seulement à programmer la solution
force brute de notre problème du sac à dos. Une version Go de cette solution vous est donnée dans sa
forme récursive régulière (avec un seul fil d’exécution). Votre mission est d’utiliser la concurrence afin
d’accélérer la résolution de ce problème.

Puisque l’approche force brute consiste à explorer l’arbre binaire des solutions possible, une première
approche naïve serait de subdiviser le problème en deux parties, le sous-arbre de gauche et le sous arbre
de droite, et de lancer deux fils concurrents afin résoudre ces deux sous-problèmes. Une fois ces deux
fils terminés, on ne garde que la meilleure des deux solutions obtenues.

Suivant cette idée, au lieu de résoudre le problème comme suit :
#### KnapSack(W, weights , values)
  
Il faudrait faire :

#### last := len(weights)-1
//best solution with the last item
#### go KnapSack(W - weights [last], weights [:last], values [:last])
// best solution without the last item
#### go KnapSack(W, weights [:last], values [:last])
//code here to synchronize and then determine which one is the best solution
  
Poursuivant cette idée, vous pourriez aussi résoudre les deux sous-sous problèmes de chacun des sous-
problème ci-dessus en utilisant 4 fils concurrents. Vous pourriez aussi procéder de façon similaire pour
le niveau suivant de l’arbre (8 sous-problèmes) et ainsi de suite. Ceci parait être une bonne approche
sauf qu’à un certain point, les sous-problèmes à résoudre seront si simple (i.e. auront si peu d’items)
qu’il n’y aura plus aucun gain à les résoudre de façon concurrente. Par surcroit, nous savons que le
nombre de nœuds dans l’arbres des solutions est O(2n)ce qui implique que le nombre de fils
d’exécution à créer afin de résoudre tous les sous-problèmes serait prohibitifs. En effet, même si les go
routines sont très légères, elles mobilisent tout de même une certaine quantité de ressources, en
conséquence la quantité totale de mémoire exigée par cette approche exhaustive sera telle que votre
machine pourrait bien ne pas le supporter.

En conséquence, puisque notre objectif est de permettre à notre programme de s’exécuter aussi
rapidement que possible, il devrait exister un choix optimal pour ce qui est du nombre de fils
d’exécution à créer. C’est-à-dire qu’un sous-problème devrait être résolu de façon concurrente
seulement si le nombre d’items est supérieur à un certain seuil.

On vous demande donc de résoudre le problème du sac à dos en utilisant la méthode force brute
programmée de façon concurrente. Vous devez déterminer, de façon expérimentale, le nombre de fils
d’exécution qui devrait être utilisé afin d’obtenir une exécution aussi rapide que possible.
Votre programme doit aussi lire les données du problème dans un fichier et retourner le nom des items
sélectionnés, la valeur totale de ceux-ci et le temps d’exécution de la solution. 
