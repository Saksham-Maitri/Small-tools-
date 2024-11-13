# Small tools

These are few usefull tools that can save time. 
If you want to add these to your system , just copy the go code into main.go 
then run 

```bash
go mod init <tool name>
```

and then do 

```bash
go mod tidy
```
then make exe file

```bash
go build -o <name of exe file> main.go
```
then make the exe file global 


```bash
sudo mv <name of exe> /usr/local/bin/
```
This way you can make use the commands anywhere 
