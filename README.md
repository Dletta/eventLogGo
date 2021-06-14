# eventLogGo
Command-line event, task, note logger


### What it does:

- Manually log events, tasks, notes or others via the command line.
- Print a list of events (etc) via ```list``` command
- Save event list to Disk

### Implementation Details:
- Small String Parser that assumes
  ```Event Type: Event description or anything you want to write here```
- Saves to Append-Only File
  ```Event Type: Description etc // Timestamp in Unix Epoch```
- Currently no dependencies

### Future Plans

- [ ] Edit Lines

- [ ] Search Entries from Command Line
