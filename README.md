# Job Tracker CLI

A simple **Go command-line tool** to track your jobs and tasks. Start and end jobs directly from the terminal, and keep a **record in SQL** of when each job was started and completed.  

The tool is lightweight, extensible, and with the goal of being integrated with APIs for reporting.  

---

## Prerequisites

- Go 1.25 or later installed on your system.  
  You can download Go from [https://golang.org/dl/](https://golang.org/dl/).

## Current features

- Add tasks
- Store tasks in SQL database (SQLite)
- Start and end a task, tracking the duration

## In progress

- View all tasks
- Mark tasks as completed / deleted / archived
- Format start / end times in 15min intervals
- Arguments for specifying start / end time
- Store record of working time in SQL
- Provide interface for API to be used from stored data
- Visually improve the cli

---

## Installation
### Windows

1. Clone the repository:  
```bash
git clone https://github.com/SamoThreeSixty/job-tracking-cli
cd job-tracker-cli
```
2. Build the CLI tool:
```bash
go build -o job_tracker.exe ./cmd
```

3. Run the CLI:
```
./job_tracker.exe
```

4. (optional if windows) Add this to your bin
``` powershell 
go build -o C:\Users\SamoT\go\bin\job_tracker.exe ./cmd
```