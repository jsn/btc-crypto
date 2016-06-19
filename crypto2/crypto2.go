package main

import (
    "fmt"
    "flag"
    "os"
    "log"
    "image"
    "image/png"
    "image/color"
    "math"
)

var RED = color.RGBA{255, 0, 0, 255}
var GREEN = color.RGBA{0, 255, 0, 255}
var BLUE = color.RGBA{0, 0, 255, 255}

const cx, cy = 474, 474
const baseR = 360
const stepR = 32

func dieOn(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

type ImageSet interface {
    Set(x, y int, c color.Color)
}

func setSquare(m image.Image, cx int, cy int, r int, c color.Color) {
    for x := -r; x <= r; x ++ {
        for y := -r; y <= r; y ++ {
            m.(ImageSet).Set(cx + x, cy + y, c)
        }
    }
}

func getSquare(m image.Image, cx int, cy int, rad int) (int, int, int) {
    r, g, b, a := 0, 0, 0, 0
    n := (rad * 2 + 1) * (rad * 2 + 1)

    for x := -rad; x <= rad; x ++ {
        for y := -rad; y <= rad; y ++ {
            r1, g1, b1, a1 := m.At(cx + x, cy + y).RGBA()
            r += int(r1 >> 8)
            g += int(g1 >> 8)
            b += int(b1 >> 8)
            a += int(a1 >> 8)
        }
    }
    return r / n, g / n, b / n
}

func getBit(m image.Image, angle float64, rad int, w int) int {
    px := cx + int(math.Sin(angle) * float64(rad))
    py := cy - int(math.Cos(angle) * float64(rad))
    r, g, b := getSquare(m, px, py, w)

    rv := 1
    c := BLUE
    if r >= 250 && g >= 250 && b >= 250 {
        rv = 1
        c = BLUE
    } else {
        rv = 0
        c = RED
    }
    setSquare(m, px, py, w, c)
    fmt.Print(rv)
    return rv
}

func process(m image.Image) {
    step := math.Pi * 2 / (51 * 2)

    for i := step; i < math.Pi * 2; i += 2 * step {
        for j := -3; j <= 3; j ++ {
            if j == 0 {
                getBit(m, i, baseR, 10)
            } else {
                getBit(m, i - step / 2.4, baseR - j * stepR, 1)
                getBit(m, i + step / 2.4, baseR - j * stepR, 1)
            }
        }

        fmt.Println("")
    }
}

func main() {
    flag.Parse()
    args := flag.Args()

    if len(args) != 2 {
        log.Fatal("usage: <input.png> <output.png>")
    }

    file, err := os.Open(flag.Arg(0))
    dieOn(err)
    defer file.Close()

    ofile, err := os.Create(flag.Arg(1))
    dieOn(err)
    defer ofile.Close()

    m, _, err := image.Decode(file)
    dieOn(err)

    process(m)

    png.Encode(ofile, m)
}
