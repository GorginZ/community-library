# community-library

small unserious project to practice go, intended to be backend for a frontend 'library' website. 

### **required**: 

- [Go](https://golang.org/)
- [Docker](https://www.docker.com/get-started/)
--------
### Endpoints

##### /books
Methods: GET

gets all books 

example: may return ```[{"isbn":"1","title":"Brave New World","author":"Aldous Huxley"},{"isbn":"2","title":"Das Kapital Volume One","author":"Karl Marx"},{"isbn":"3","title":"If This Is A Man","author":"Primo Levi"},{"isbn":"3","title":"Sexus","author":"Henry Miller"}]``` response
_______

build:

```docker build -t community-library<tag> .```

>NOTE: if on macOS with newer architecture you may need to explicitely set platform:

``` --platform linux/amd64 ```

run:

```docker run --rm -p 8080:8080 community-library:<tag>```


## publish to GCR

```bash
#for GCR
# docker build --build-arg version="$version" -t "asia.gcr.io/gorg-364804/community-library:$tag" .

# docker push "asia.gcr.io/gorg-364804/community-library:$tag"
```
