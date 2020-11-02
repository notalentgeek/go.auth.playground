FROM   golang:1.15.3-buster
COPY   main /main
EXPOSE 8000
CMD    ["/main"]
