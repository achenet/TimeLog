# TimeLog
A simple program to log time spent on tasks.

## Building from source
This program requires Go version 1.15 or later.

To build this program from source, first clone the github repo.
```
git clone github.com/achenet/TimeLog
```

Then get all necessary dependencies and build the binary, using the `-o` flag to specify a name.
```
go mod vendor
go build -o timelog
```

## Usage
There are currently 3 commands.

The `timelog` will list all tasks and the time spent working on them, as well as whether or not they are in progress.

The `timelog start TASK_NAME` will start logging time for the task with name `TASK_NAME`. If one does not exist, it will be created.

The `timelog stop TASK_NAME` will stop logging time for the task with name `TASK_NAME`, if it exists and is currently in progress. Otherwise, it will do nothing.
