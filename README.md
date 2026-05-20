# ezclimon
## Goal
* Make an easy all in one tool for monitoring
* Make it look cool
* Be able to add users
* Look into _Bubble Tea_ for interactive UI in CLI

## Features
### Storage Check
* This function gives us the human readable current storage as `df -h`
* #### Goal
  * Add Inode exhastion with `df -i`

### Network Info
* This function gives us the output of `ip a`
* ##### Goal for network info
  * Change IP address
  * Check if its connected to Gateway_ip

### Memory Check
* This function gives us the output of `free -h`

### Add user
* _work in progress_
* I want to add a gui for adding a user

### Integrity check
* _work in progess_

### Firewall
* _work in progress_

### Critical Service Check
* _work in progress_
* Make a list of crucial services
* Make a loop that `systemd` the list

### Error log
* _work in progress_
* Runs the command `exec.Command("journalctl", "-p", "err", "-n", "10", "--no-pager")`

