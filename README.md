# TaskTracker
simple task tracker

## installation
```sh
git clone https://github.com/ThisIsHyum/TaskTracker/
cd TaskTracker/cli
go build
```

## Usage

- Create task
```sh
./cli add -d "description of task"
 
```

show list
```sh
./cli list
./cli list --status in-progress # show tasks only with status in-progress
```

delete task
```sh
./cli delete --id 1 # delete task with id 0
 
```

update task
```sh
./cli update --id 1 -d "new description" # update description of task with id 1
 
```

mark task
```sh
./cli mark --id 1 --status in-progress # new status of task with id 1
 
```
