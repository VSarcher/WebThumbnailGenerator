# Web Thubmail Generator

## ToDo
The goal of this project is to create a Go program that generates thumbnail images of web pages and saves them to a SQLite database.

- Used Fiber Microframework to convey task
- Introduced GORM to integrate with SQLite
  - Image_info table : for storing thumbnail image data
- Used goroutines and channels to implement concurrency
- Containerized using Docker

## How it works.

The program accepts a POST request with a JSON body containing a list of URLs.

| POST | `http://localhost:3000/image`|
| --- | --------------------|

  Request JSON form
 ```sh
  {
      "urls": [
          "google.com",
          "go.dev"
      ]
  }
 ```

For each URL in the list, the program generates a thumbnail image using the [pagepeeker.com](https://pagepeeker.com/) API.

The program saves each URL and its corresponding thumbnail image (encoded as a base64 string) to a SQLite database.

The program processes requests concurrently using goroutines and channels.

The endpoint returns a 200 status code to the user to indicate that the request was successful.

## How to run
Run on your machine
```
go mod tidy
go run cmd/main.go
```
Run with docker container.

```
# create the image
docker build -t web-thumbnail-generator .
```

```
# run a container
docker run --detach --name generator -p 3000:8080 -d web-thumbnail-generator
```

```
# remove container
docker container stop generator
docker container rm generator
```

```
# delete the image
docker rmi web-thumbnail-generator
```