// based on https://gist.github.com/erikcorry/5bcd95b916527969a4210a3aa32bd7a0

package main

func main() {
    // test starting out with two values (123, 123)
    a := []int{}
    a = append(a, 123)
    a = append(a, 123)
    b := a[0:]
    c := a[0:]
    b = append(b, 456)
    c = append(c, 999)
    println("the last element of b is:", b[len(b)-1], " (expected 456)")
    println("the last element of c is:", c[len(c)-1], " (expected 999)")


    // test starting out with three values (123, 123, 123)
    x := []int{}
    x = append(x, 123)
    x = append(x, 123)
    x = append(x, 123)
    y := x[0:]
    z := x[0:]
    y = append(y, 456)
    z = append(z, 999)
    println("the last element of y is:", y[len(y)-1], " (expected 456)  wat?!")
    println("the last element of z is:", z[len(z)-1], " (expected 999)")
}
