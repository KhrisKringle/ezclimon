# ezclimon
## Goal
* Make an easy all in one tool for monitoring
* Make it look cool
* Be able to add users
* Look into _Bubble Tea_ for interactive UI in CLI

## Features
### Storage Check
* This function gives us the human readable current storage as `df -h`
* Add Inode exhastion with `df -i`

### Network Info
* This function gives us the output of `ip a`
#### Goal for network info
* Change IP address

### Memory Check
* This function gives us the output of `free -h`

### Add user
* I want to add a gui for adding a user

### Integrity check
* work in progess

### Firewall
* work in progress

### Critical Service Check
* Make a list of crucial services
* Make a loop that `systemd` the list

### Error log
* Runs the command `exec.Command("journalctl", "-p", "err", "-n", "10", "--no-pager")`

