##
## STEP 1 - BUILD
##

# specify the base image to  be used for the application, alpine or ubuntu
FROM golang:1.20.2-alpine as builder

# create a working directory inside the image
WORKDIR /app

# copy Go modules and dependencies to image
COPY go.mod .

# copy directory files i.e all files ending with .go
COPY . .

# download Go modules and dependencies
RUN go mod download

# compile application
RUN go build -o /restaurantApi ./api/cmd/restaurantSrv

##
## STEP 2 - DEPLOY
##
FROM scratch
WORKDIR /
COPY --from=builder /restaurantApi /restaurantApi

ENTRYPOINT ["/restaurantApi"]
