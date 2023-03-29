![image](https://user-images.githubusercontent.com/42554663/228620879-5b7ba8ef-9ffc-4c82-ae48-4e21f8b49354.png)

### Random wallpaper from unsplash

Download random image from unsplash in HD resolution using golang. 

### Installation

Clone the repository. 

```bash
git clone https://github.com/Awan/random_wallpaper.git
```

Go to the directory where the project is cloned. 

and build it.

```bash
go build wallpaper.go
```

If you want to build for android:

```bash
GOARCH=arm go build -ldflags="-s -w" wallpaper.go
```

### Usage

```bash
wallpaper <some keyword>
```
For example, if you want to get some random mountains image:

```bash
wallpaper mountains
```

