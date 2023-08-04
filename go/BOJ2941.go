package main

import (
	"fmt"
    "strings"
)


func main() {
	var n string
	fmt.Scan(&n)
    var ans int = len(n)
	a := strings.Split(n, "")
    var nList [105]string
    for i:=0; i < len(a); i++ {
        nList[i] = a[i]
    }
    for i := 0; i < len(n); i++ {
        // d
        if nList[i] == "d" {
            //dz=
            if nList[i+1] == "z" && nList[i+2] == "=" {
                ans--; ans--;
                i++; i++;
            } else if i < len(nList) -1 && nList[i+1] == "-" { // d-
                ans--;
                i++;
            }
        } else if nList[i] == "c" || nList[i] == "s" || nList[i] == "z" { //c= s= z=
            if nList[i+1] == "=" || nList[i+1] == "-" {
                ans--;
                i++;
            }
        } else if (nList[i] == "l" && nList[i+1] == "j") || (nList[i] == "n" && nList[i+1] == "j"){
            ans--;
            i++;
        }
    }
    fmt.Print(ans)
}
