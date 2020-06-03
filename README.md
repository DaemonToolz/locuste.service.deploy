# locuste.service.deploy
LOCUSTE - Service de réception / déploiement et mise à jour

Le project Locuste se divise en 4 grandes sections : 
* Automate (Drone Automata) PYTHON (https://github.com/DaemonToolz/locuste.drone.automata)
* Unité de contrôle (Brain) GOLANG (https://github.com/DaemonToolz/locuste.service.brain)
* Unité de planification de vol / Ordonanceur (Scheduler) GOLANG (https://github.com/DaemonToolz/locuste.service.osm)
* Interface graphique (UI) ANGULAR (https://github.com/DaemonToolz/locuste.dashboard.ui)

![Composants](https://user-images.githubusercontent.com/6602774/83644711-dcc65000-a5b1-11ea-8661-977931bb6a9c.png)

Tout le système est embarqué sur une carte Raspberry PI 4B+, Raspbian BUSTER.
* Golang 1.11.2
* Angular 9
* Python 3.7
* Dépendance forte avec la SDK OLYMPE PARROT : (https://developer.parrot.com/docs/olympe/, https://github.com/Parrot-Developers/olympe)

![Vue globale](https://user-images.githubusercontent.com/6602774/83644783-f10a4d00-a5b1-11ea-8fed-80c3b76f1b00.png)

Détail des choix techniques pour la partie Service de déploiement :

* [Golang] - Rédaction rapide et simple de programmes orientés web, multithreading et multiprocessing intégré au langage

![Updater order](https://user-images.githubusercontent.com/6602774/83646243-aee20b00-a5b3-11ea-9888-d2d07a8c755a.png)

Evolutions à venir : 
* Ajout de nouvelles fonctionnalités (récupération des versions actuelles pour chaque version)
* Démarrager / arrêt de processus
* Partition des versions en versions applicatifs
