# http-server-demo

This demonstrates use of the Go package httprouter
at <https://github.com/julienschmidt/httprouter>.

To run this, cd to `$GOPATH` and enter
`go run src/github.com/mvolkmann/http-server-demo/main.go`.

## URLs to try

GET localhost:8080 outputs "Welcome!".
GET localhost:8080/hello/mark outputs "hello, mark!".

GET localhost:8080/coupon/5 outputs a coupon image URL.

GET localhost:8080/public/foo.html outputs the contents of `$GOPATH/public/foo.html`.

POST localhost:8080/activity-level outputs
POST localhost:8080/body-condition outputs
POST localhost:8080/breed outputs
POST localhost:8080/food-recommendations outputs
POST localhost:8080/weight-range outputs
