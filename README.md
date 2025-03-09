# Serveur de Chat TCP en Go

## Description
Ce projet implémente un serveur de chat TCP en Go permettant aux clients de se connecter via Netcat (`nc`) ou tout autre client TCP. Chaque client doit fournir un nom à la connexion, et les messages sont échangés entre les utilisateurs en temps réel.

## Fonctionnalités
- Connexion multiple de clients via TCP
- Attribution d'un nom à chaque utilisateur
- Envoi de messages à tous les clients connectés
- Déconnexion propre des clients
- Utilisation de séquences ANSI pour améliorer l'affichage

## Prérequis
- Go (version 1.18 ou supérieure)
- Un terminal compatible avec Netcat (`nc`)

## Installation
1. Clonez le dépôt :
   ```sh
   git clone https://github.com/votre-utilisateur/chat-tcp-go.git
   cd chat-tcp-go
   ```
2. Compilez et lancez le serveur :
   ```sh
   go run main.go
   ```
3. Compilez et lancez le serveur avec une prot spécifique par exempl 8001:
   ```sh
   go run main.go 8001
   ```
## Utilisation
1. Démarrez le serveur :
   ```sh
   go run main.go
   ```
2. Connectez un client avec Netcat :
   ```sh
   nc localhost 8080
   ```
3. Saisissez votre nom lorsque demandé.
4. Envoyez des messages aux autres clients connectés.
5. Pour quitter, fermez simplement la connexion (`CTRL+C`).


## Auteur
Développé par [Mohammed aidi] maidi
Développé par [Mohammed aidi] maidi

