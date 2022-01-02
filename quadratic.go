package main

import (
    "fmt"
    "log"
    "net/http"
    "strconv"
    "math/cmplx"
)

const (
    pageTop    = `<!DOCTYPE HTML><html><head>
<style>.error{color:#FF0000;}</style></head><title>Quadratic  Solver</title>
<body><h3>Quadratic Solver</h3>
<p>Computes solutions for a quadratic equation of the form ax<sup>2</sup> + bx + c = 0</p>`
    form       = `<form action="/" method="POST">
<input type="text" name="a" size="3">x<sup>2</sup>
<input type="text" name="b" size="3">x
<input type="text" name="c" size="3">
<input type="submit" value="Calculate">
</form>`
    pageBottom = `</body></html>`
    anError    = `<p class="error">%s</p>`
)

type quadratic struct {
    a float64
    b float64
    c float64
    soln1    complex128
    soln2    complex128
}

func main() {
    http.HandleFunc("/", homePage)
    fmt.Printf("Starting server.\n Listening on :9001")
    if err := http.ListenAndServe(":9001", nil); err != nil {
        log.Fatal("failed to start server", err)
    }
}

func homePage(writer http.ResponseWriter, request *http.Request) {
    err := request.ParseForm() // Must be called before writing response
    fmt.Fprint(writer, pageTop, form)
    if err != nil {
        fmt.Fprintf(writer, anError, err)
    } else {
        if quad, message, ok := processRequest(request); ok {
            quad = solve(quad)
            fmt.Fprint(writer, formatAnswer(quad))
        } else if message != "" {
            fmt.Fprintf(writer, anError, message)
        }
    }
    fmt.Fprint(writer, pageBottom)
}

func processRequest(request *http.Request) (quad quadratic, message string, err bool) {
    if field, found := request.Form["a"]; found{
        if x, err := strconv.ParseFloat(field[0], 64); err != nil {
            return quad, "'" + field[0] + "' is invalid", false
        } else {
            quad.a = x
        }
    }
    if field, found := request.Form["b"]; found{
        if x, err := strconv.ParseFloat(field[0], 64); err != nil {
            return quad, "'" + field[0] + "' is invalid", false
        } else {
            quad.b = x
        }
    }
    if field, found := request.Form["c"]; found{
        if x, err := strconv.ParseFloat(field[0], 64); err != nil {
            return quad, "'" + field[0] + "' is invalid", false
        } else {
            quad.c = x
        }
    }
    return quad, "", true
}

func formatAnswer(quad quadratic) string {
    return fmt.Sprintf(`<p>
            %fx<sup>2</sup> + %fx + %f --> x = (%f) or x = (%f)
        </p>`, quad.a, quad.b, quad.c, quad.soln1, quad.soln2)
}

func solve(quad quadratic) quadratic {
    a := complex(quad.a, 0)
    b := complex(quad.b, 0)
    c := complex(quad.c, 0)
    quad.soln1 = (-b  + cmplx.Sqrt(cmplx.Pow(b, 2) - 4.0 * a * c)) / (2*a)
    quad.soln2 = (-b  - cmplx.Sqrt(cmplx.Pow(b, 2) - 4.0 * a * c)) / (2*a)
    return quad
}



