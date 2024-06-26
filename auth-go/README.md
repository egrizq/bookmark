## Ubuntu Cheat Sheet

### Navigation

| Command    | Purporse              |
| ---------- | --------------------- |
| cd ..      | back                  |
| cd dirname | open directory        |
| pwd        | show current location |
| clear      | remove history        |

### Folder

| Command                    | Purporse                     |
| -------------------------- | ---------------------------- |
| ls                         | list all directory           |
| mkdir dirname              | create a directory           |
| rmdir dirname              | delete a directory           |
| rm -r dirname              | delete a directory with file |
| mv dirname new_location    | move directory               |
| cp -r dirname new_location | copy a directory             |
| mv dirname new_dirname     | rename directory             |

### File

| Command                  | Purporse                |
| ------------------------ | ----------------------- |
| touch filename           | create a file           |
| rm filename              | delete a file           |
| mv filename new_location | move to a new directory |
| mv filename new_filename | rename filename         |

### Setup

1.  Update lastest Software
    - sudo apt update
    - sudo apt upgrade
    - sudo apt dist-upgrade
    - sudo apt autoremove
2.  Additional Driver
    - sudo apt install git curl unzip wget
3.  Install Programming Language
    1. Node.js & typescript
       - sudo apt install npm
       - npm -v
       - curl -fsSL https://fnm.vercel.app/install | bash
       - node -v
       - sudo npm install -g pnpm
       - pnpm -v
       - npm install -g typescript
       - tsc -v
    2. Golang
       - sudo apt-get update && sudo apt-get -y install golang-go
       - golang version
    3. Python
       - python3 --version
       - python already installed
4.  empty
