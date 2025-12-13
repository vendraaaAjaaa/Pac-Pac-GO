# 🟡 Pac-Pac-GO

**Pac-Pac-GO** is a simple 2D game inspired by the classic **Pac-Man**, developed using the **Go (Golang) programming language**.

This project was created as a **Final Assignment for the Introduction to Algorithms and Programming course**, with the goal of applying fundamental algorithm concepts, programming logic, and program structure using Go.

---

## 🎯 Project Objectives

The objectives of developing **Pac-Pac-GO** are to:
- Understand **basic algorithmic concepts**
- Apply **control structures** (branching and looping)
- Implement **functions and modularization**
- Manage **game state**
- Practice programming logic through a simple game case study

---

## 🎮 Game Description

Players control the main character to navigate a grid-based map, collect pellets, and avoid enemies.  
The game ends when the player comes into contact with an enemy or when all pellets have been collected.

---

## ✨ Key Features

- 🟡 Simple Pac-Man–style gameplay  
- 🧱 Grid-based map  
- 👻 Enemies with automatic movement  
- 🎯 Scoring system  
- ⌨️ Keyboard controls  
- 🔄 Game logic based on basic algorithms  

---

## 🛠️ Technology Used

- **Programming Language**: Go (Golang)  
- **Paradigm**: Procedural / Modular  
- **Application Type**: Simple 2D Game  

---

## ▶️ How to Run

### Prerequisites

Make sure the following requirements are installed on your system:

- **Go (Golang)** version 1.18 or later  
- A system with OpenGL support (required by Ebiten)

Check the Go installation:
```bash
go version
````

---

### Install Dependencies

This project uses the **Ebiten** game engine.
Install all required dependencies by running:

```bash
go mod tidy
```

If the `go.mod` file does not exist yet, initialize it first:

```bash
go mod init pac-pac-go
go get github.com/hajimehoshi/ebiten/v2
```

---

### Run the Game

Run the game using the following command:

```bash
go run main.go
```

A window titled **"Pacman in Go"** will appear and the game will start automatically.

---

### Build the Executable (Optional)

To build a standalone executable, run:

```bash
go build
```

Run the compiled program:

* **Linux / macOS**

```bash
./pac-pac-go
```

* **Windows**

```bash
pac-pac-go.exe
```

---

### Game Controls

| Key        | Action           |
| ---------- | ---------------- |
| Arrow Keys | Move Pac-Man     |
| R          | Restart the game |
| ESC        | Exit the game    |
