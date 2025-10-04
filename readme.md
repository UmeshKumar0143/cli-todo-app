
---

```markdown
# 🧭 CLI BASED TODO APP  
### 📝 Todo CLI – Simple Command-Line Todo Manager

A lightweight command-line **Todo Manager** written in Go.  
Manage your daily tasks right from your terminal — add, list, edit, mark done, and delete todos, all stored locally in a JSON file.

---

## 🚀 Features

- ✅ Add new todo items  
- 📜 List all current tasks  
- ✏️ Edit existing tasks  
- ☑️ Mark tasks as completed  
- 🗑️ Delete individual tasks  
- 💣 Delete all tasks at once  
- 💾 Persistent storage using JSON  
- 🆘 Built-in help command (`todo -h`)



## ⚙️ Installation

### 1️⃣ Clone the repository

```bash
git clone https://github.com/umeshkumar0143/todo-cli.git
cd todo-cli
````

### 2️⃣ Build the binary

```bash
go build -o todo
```

### 3️⃣ (Optional) Install globally

To use `todo` as a global command:

```bash
sudo mv todo /usr/local/bin/
```

Now you can run:

```bash
todo -h
```

from anywhere in your terminal 🎉

---

## 🧠 Usage

### ➕ Add a new task

```bash
todo add "Learn Go basics"
```

### 📋 List all tasks

```bash
todo list
```

### ✅ Mark a task as done

```bash
todo done 2
```

### 🖊️ Edit a task

```bash
todo edit 3 "Read Go documentation"
```

### ❌ Delete a task

```bash
todo delete 1
```

### 💣 Delete all tasks

```bash
todo clear
```

### 🆘 Show help

```bash
todo -h
```

---

## 🗂️ Data Storage

All tasks are stored locally in:

```
/todo-cli/internal/storage/todo.json
```

You can change this path inside your Go file:

```go
const File_name = "/your/custom/path/todo.json"
```

---

## 🧰 Tech Stack

* **Language:** Go (Golang)
* **Storage:** JSON file (local persistence)

---

## 💬 Example Output

```bash
$ todo list
1. [  ] Learn Go basics
2. [✓] Build my first CLI tool
3. [  ] Explore Go concurrency
```



