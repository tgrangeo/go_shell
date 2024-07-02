# Custom Shell in Go

This project is a simple custom shell implemented in Go. It supports a few built-in commands and can execute other commands available in the system's `PATH`. The built-in commands include `echo`, `type`, `cd`, and `exit`.

## Features

- **Built-in Commands**:
  - `echo`: Prints the given arguments to the terminal.
  - `type`: Displays whether a command is a shell built-in or the path to the command in the system.
  - `cd`: Changes the current working directory.
  - `exit`: Exits the shell.

- **Execute External Commands**: Executes any external command available in the system's `PATH`.

- **Prompt Display**: Shows the current working directory in the shell prompt.

## Usage

### Built-in Commands

#### echo
Prints the given arguments to the terminal.

```sh
echo Hello, World!
```

#### type
Displays whether a command is a shell built-in or the path to the command in the system.

```sh
type echo
type ls
```

#### cd
Changes the current working directory.

```sh
cd /path/to/directory
cd ~  # Change to home directory
```

#### exit
Exits the shell.

```sh
exit 0
```

### Executing External Commands
Any external command available in the system's `PATH` can be executed.

```sh
ls -l
pwd
```

## How to Run

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/golang-shell.git
    cd golang-shell
    ```

2. Build the project:

    ```sh
    go build -o golang-shell
    ```

3. Run the shell:

    ```sh
    ./golang-shell
    ```

---

Happy Shell Scripting!