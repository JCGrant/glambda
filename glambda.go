/*

variables:   x, y, z...
functions:   λ x • x + 1
application: (λ x • x + 1) 2
               = 2 + 1
               = 3

multi variable functions:
\ x . \ y . x + y
\ x y . x + y


integers:
0 = \ f x . x
1 = \ f x . f x
2 = \ f x . f (f x)
3 = \ f x . f (f (f x))
...
n = \ f x . f^n x

integer ops:
succ = \ n f x . f (n f x)
add  = \ n m f x . n f (m f x)
add  = \ n m f x . n succ m


bools:
true  = \ x y . x
false = \ x y . y

boolean ops:
not = \ x . x false true
and = \ x y . x y false
or  = \ x y . x true y

*/

package glambda

import (
	"fmt"
)

func Run() {
	code := `
	
	test = \ a b . a b


	test2 = \ x y . x (y y)
	
	
	`
	nodes := parse(code)
	for _, node := range nodes {
		fmt.Printf("%s\n", node)
	}
}
