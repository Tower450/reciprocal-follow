# Reciprocal Follow

```bash
⠀⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⣾⠀
⠀⣿⣿⣷⣄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣴⣿⣿⣿⠀
⠀⣿⣿⣿⣿⣷⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⣿⣿⣿⣿⠀
⠀⣿⣿⣿⣿⡟⢀⣾⣿⣷⡖⠒⣀⣀⣤⣶⣾⣿⣷⣶⣤⣤⣴⣾⣆⠘⣿⣿⣿⠀
⠀⣿⣿⣿⡿⠁⣾⣿⣿⣿⣧⣀⠙⠛⠛⠛⠋⣈⠻⢿⣿⣿⣿⣿⣿⣧⠈⢿⣿⠀
⠀⣿⣿⣿⠁⣼⣿⠟⠻⠿⣿⣿⣿⣷⣶⣾⣿⠿⣷⣄⡉⠻⣿⣿⣿⣿⣧⠈⢿⠀
⠀⠛⠛⠃⠀⠻⠋⣠⣶⠄⠙⠛⢻⣿⣿⣿⣷⣦⣈⠛⢿⣦⣄⠙⠻⠛⠁⠀⠀⠀
⠀⠀⠀⠀⠀⢠⣾⠟⠁⣠⣿⠇⠈⠛⣿⣿⣈⠙⠿⣷⣦⡈⠛⠛⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠁⢠⣾⡿⠁⣠⣾⠆⠸⠉⠻⣷⣤⡈⠙⠿⠃⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠉⠀⣾⠟⢁⣴⡶⠀⣤⡀⠙⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠰⣿⠟⠁⠀⠛⠟⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
```

## Dependencies

Go

```
go get golang.org/x/net/html
```


## How to use


Go HTML and JSON

1. Download your `followers_*.html` and `following.html` from (instagram)[https://www.instagram.com/download/request/]
2. Once downloaded they are located in `instagram-<username>-2025-01-01-<hash>/connections/followers_and_following`
3. Move them inside the current `html/` folder

4. Since instagram official download gives you HTML use this command
   ```bash
   go run mutual_follow.go
   ```
  ***if you have json for a reason***
 
  ```bash
  go run mutual_follow_json.go
  ```

Python

```bash
pip install instaloader beautifulsoup4
```

```bash
python3 mutual_follow.py
```
