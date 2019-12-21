embyoracle
==========

You find it hard to find a movie to watch while browsing your Mediaserver library?
embyoracle accesses your list of `unwatched` movies and suggests a random entry to watch

## build 
```
go build
```

## setup env 
```
export EMBY_USER=username
export EMBY_URL=http://127.0.0.1:8096
export EMBY_TOKEN=YoUrApItOkEn
```

## Docker 

```
docker build -t embyoracle .
docker run -ti -e EMBY_URL='http://127.0.0.1:8096' -e EMBY_TOKEN='YoUrApItOkEn' -e EMBY_USER='username'
```

# ToDo 
- create detailed Movie view
- Add graphical interfacem