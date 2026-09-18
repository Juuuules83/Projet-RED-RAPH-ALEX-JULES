# Projet RED — GOLANG VS B1

Projet RPG en ligne de commande réalisé en Go.

## Lancer le projet

Depuis le dossier racine :

```bash
go run ./src
```

Ne lance pas uniquement `main.go` : Go compile les fichiers `.go` du dossier `src` ensemble.

## Architecture

```text
Projet-RED-GOLANG-VS-B1/
├── src/
│   ├── main.go
│   ├── character.go
│   ├── menu.go
│   ├── inventory.go
│   ├── potion.go
│   └── utils.go
├── docs/
├── go.mod
└── README.md
```

## Organisation

- `main.go` : point d'entrée.
- `character.go` : personnage et initialisation.
- `menu.go` : menus et affichage des informations.
- `inventory.go` : menu inventaire.
- `potion.go` : potions.
- `utils.go` : fonctions génériques, nettoyage du terminal, attente et effet d'écriture.

## Affichage

Chaque menu nettoie le terminal avant de s'afficher. Le joueur ne doit voir que le menu dans lequel il se trouve.

Les couleurs ANSI et `typeWriter` sont préparés pour les futurs dialogues.
