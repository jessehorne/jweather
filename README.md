jweather
===

My weather forecasting service for a code challenge, written in Go using the NWS API.

# Running

```shell
go run main.go
```

Then hit `GET http://127.0.0.1:8080/points?lat=39.7456&long=-97.0892` for example. You should see a response like the following.
`{"short":"Partly Cloudy then Slight Chance Showers And Thunderstorms","character":"moderate"}`

# What I'd do with more time

1. Set up a docker container for this
2. Set up some terraform to run this on cloud run or ec2 or something
3. Add some testing (as you can see there are no testing nor benchmarks)
4. Add better documentation (probably https://github.com/swaggo/swag)
5. The way I'm getting short forecast and what not is probably incorrect as I assume the first period (index 0) is the one I should be using. I don't actually know.
6. Add a way to easily customize User-Agent without changing code. Probably a .env
7. Cacheing to help take the load from NWS servers

What else could I do with more time? What are your recommendations?
