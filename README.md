# Pokédex CLI (Go)

A simple **Pokédex-like REPL** built in **Go** that fetches data from the [PokeAPI](https://pokeapi.co/) and caches results for faster repeated access.  
This project was built as part of the [Boot.dev Build a Pokedex CLI](https://www.boot.dev/courses/build-pokedex-cli-golang) guided course.

---

## 🚀 Features

- **Interactive REPL**: Type commands to explore Pokémon data right in your terminal.
- **Pokédex Management**: Catch, inspect, and view Pokémon you’ve collected.
- **HTTP Networking**: Learn how to make API calls in Go.
- **Caching Layer**: Reduce redundant API requests with an in-memory cache.
- **JSON Parsing**: Deserialize API responses into Go structs.

---

## 🧑‍💻 Commands

Here are some of the commands available inside the REPL:

| Command             | Description                                                               |
| ------------------- | ------------------------------------------------------------------------- |
| `map`               | Lists available Pokémon location areas.                                   |
| `explore <area>`    | Shows Pokémon encounters in a given area.                                 |
| `catch <pokemon>`   | Attempts to catch the given Pokémon (with some chance of failure).        |
| `inspect <pokemon>` | Displays detailed information about a caught Pokémon (stats, types, etc). |
| `pokedex`           | Lists all Pokémon you have successfully caught.                           |
| `help`              | Shows all available commands.                                             |
| `exit`              | Quits the program.                                                        |

---

## 📚 Chapters Learned

This project covers three main areas:

1. **REPL**  
   Build a basic Read-Eval-Print Loop from scratch in Go.

2. **Cache**  
   Implement an in-memory caching system to optimize repeated API calls.

3. **Pokédex**  
   Combine the REPL and caching system into a fully functional Pokédex CLI.

---

## 🛠️ Tech Stack

- **Language**: Go
- **API**: [PokeAPI](https://pokeapi.co/)
- **Concepts**: REPL, HTTP requests, JSON parsing, Caching, Concurrency

---

## ⚡ Getting Started

### Prerequisites

- [Go](https://go.dev/) (1.21+ recommended)

### Installation

```bash
git clone https://github.com/<your-username>/pokedex-cli.git
cd pokedex-cli
go run main.go
```


pokedex > help
Available commands:
  map
  explore <area>
  catch <pokemon>
  inspect <pokemon>
  pokedex
  help
  exit

pokedex > map
canalave-city-area
eterna-city-area
...

pokedex > explore eterna-city-area
Found pokemon:
  - bidoof
  - zubat
  - gastly

pokedex > catch bidoof
You caught bidoof!

pokedex > inspect bidoof
Name: bidoof
Height: 5
Weight: 200
Stats:
  - hp: 59
  - attack: 45
  - defense: 40
  - special-attack: 35
  - special-defense: 40
  - speed: 31
Types:
  - normal


 What I Learned

    How to implement a REPL in Go.

    Making and handling HTTP requests.

    Parsing and working with JSON APIs.

    Writing a caching system with concurrency safety.

    Structuring and testing Go CLI applications.
