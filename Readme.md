# Output

```
go run main.go
```

```
Starting Go Cache
Add: BMW
1 - [{BMW}]
Add: Audi
2 - [{Audi}<-->{BMW}]
Add: Land Rover
3 - [{Land Rover}<-->{Audi}<-->{BMW}]
Add: General Motors
4 - [{General Motors}<-->{Land Rover}<-->{Audi}<-->{BMW}]
Add: Porsche
5 - [{Porsche}<-->{General Motors}<-->{Land Rover}<-->{Audi}<-->{BMW}]
Add: Toyota
Remove BMW
5 - [{Toyota}<-->{Porsche}<-->{General Motors}<-->{Land Rover}<-->{Audi}]
Add: 12
Remove Audi
5 - [{12}<-->{Toyota}<-->{Porsche}<-->{General Motors}<-->{Land Rover}]
```