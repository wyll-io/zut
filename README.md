# zut

**Zut** helps you to find the command that you search and how to use it ! Works on *Linux* and *Windows*.

## Why ?

If you ever answered yes to one of this case, *zut* is for you :

 - What the hell is that command that displays the disk usage ?
 - I want to copy something with rsync but I don't remember the syntax
 - What are the options I used to set with netstat ?

## Installation

Dowload a release, extract the executable in `/usr/bin` or `/usr/local/bin`. Then you can :

```
zut version
```

## Usage

### Search a command

`zut search [your search]`

Example : `zut search disk`

### Display commands for a category

The categories are :

 - networking
 - storage
 - file
 - ...

## Favorites

You can define a set of useful command that you often use and put them in a file on your home directory, the file must be nammed `zut.yml`. Like this :

```yml
- rsync -ah --progress source dest
- curl -vvv -I http://
```