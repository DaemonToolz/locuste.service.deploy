# locuste.service.deploy
LOCUSTE - Service de réception / déploiement et mise à jour

[![Codacy Badge](https://app.codacy.com/project/badge/Grade/a9b0ddc3726644e8a27249c395f0ed48)](https://www.codacy.com/manual/axel.maciejewski/locuste.service.deploy?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=DaemonToolz/locuste.service.deploy&amp;utm_campaign=Badge_Grade)

[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=DaemonToolz_locuste.service.deploy&metric=alert_status)](https://sonarcloud.io/dashboard?id=DaemonToolz_locuste.service.deploy)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=DaemonToolz_locuste.service.deploy&metric=sqale_rating)](https://sonarcloud.io/dashboard?id=DaemonToolz_locuste.service.deploy)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=DaemonToolz_locuste.service.deploy&metric=reliability_rating)](https://sonarcloud.io/dashboard?id=DaemonToolz_locuste.service.deploy)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=DaemonToolz_locuste.service.deploy&metric=security_rating)](https://sonarcloud.io/dashboard?id=DaemonToolz_locuste.service.deploy)
[![Vulnerabilities](https://sonarcloud.io/api/project_badges/measure?project=DaemonToolz_locuste.service.deploy&metric=vulnerabilities)](https://sonarcloud.io/dashboard?id=DaemonToolz_locuste.service.deploy)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=DaemonToolz_locuste.service.deploy&metric=bugs)](https://sonarcloud.io/dashboard?id=DaemonToolz_locuste.service.deploy)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=DaemonToolz_locuste.service.deploy&metric=coverage)](https://sonarcloud.io/dashboard?id=DaemonToolz_locuste.service.deploy)

<img width="2610" alt="locuste-uploader-banner" src="https://user-images.githubusercontent.com/6602774/84285953-5aec9e80-ab3e-11ea-8a18-60a0f613ef0b.png">

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
