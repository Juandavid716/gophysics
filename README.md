<p align="center"><img src="https://github.com/Juandavid716/gophysics/blob/master/logo.png"></p>

# Gophysics - A Physics library for the Go programming language.

```
This is a Physical library for Go Language.
Which includes multiple Geometrical calculations like Circle, Lines etc in different forms.
```
## Installation

Installation is done using `go get`.
```
go get -u github.com/Juandavid716/gophysics
```

#### These are following equations calculation which is supported by gophysics.
- [x]  **Constant acceleration equations**



# Constant acceleration equations

# Velocity

### Equations



* <a href="https://www.codecogs.com/eqnedit.php?latex=V&space;=&space;V_{o}&space;&plus;&space;at" target="_blank"><img src="https://latex.codecogs.com/gif.latex?V&space;=&space;V_{o}&space;&plus;&space;at" title="V = V_{o} + at" /></a>
* <a href="https://www.codecogs.com/eqnedit.php?latex=V^{2}&space;=&space;Vo^{2}&space;&plus;&space;2a(x_{f}-x_{o})" target="_blank"><img src="https://latex.codecogs.com/gif.latex?V^{2}&space;=&space;Vo^{2}&space;&plus;&space;2a(x_{f}-x_{o})" title="V^{2} = Vo^{2} + 2a(x_{f}-x_{o})" /></a>

### Example

...


...


## Functions
| Equations                                   | Method                            | Arguments       |  Results | 
| :-------------                         |:-------------                           | :-----            |  :------  | 
| <a href="https://www.codecogs.com/eqnedit.php?latex=V&space;=&space;V_{o}&space;&plus;&space;at" target="_blank"><img src="https://latex.codecogs.com/gif.latex?V&space;=&space;V_{o}&space;&plus;&space;at" title="V = V_{o} + at" /></a>|     |           |     | 
|  <a href="https://www.codecogs.com/eqnedit.php?latex=V^{2}&space;=&space;Vo^{2}&space;&plus;&space;2a(x_{f}-x_{o})" target="_blank"><img src="https://latex.codecogs.com/gif.latex?V^{2}&space;=&space;Vo^{2}&space;&plus;&space;2a(x_{f}-x_{o})" title="V^{2} = Vo^{2} + 2a(x_{f}-x_{o})" /></a>|   |            |   | 


# Displacement 

### Equations
  
* <a href="https://www.codecogs.com/eqnedit.php?latex=x=&space;\frac{1}{2}(V&space;&plus;&space;V_{o})t&space;&plus;&space;x_{0}" target="_blank"><img src="https://latex.codecogs.com/gif.latex?x=&space;\frac{1}{2}(V&space;&plus;&space;V_{o})t&space;&plus;&space;x_{0}" title="x= \frac{1}{2}(V + V_{o})t + x_{0}" /></a>
 
 * <a href="https://www.codecogs.com/eqnedit.php?latex=x&space;=&space;V_{0}t&plus;&space;\frac{1}{2}at^{2}&space;&plus;&space;x_{o}" target="_blank"><img src="https://latex.codecogs.com/gif.latex?x&space;=&space;V_{0}t&plus;&space;\frac{1}{2}at^{2}&space;&plus;&space;x_{o}" title="x = V_{0}t+ \frac{1}{2}at^{2} + x_{o}" /></a>


### Example

```
package main

import (
	"fmt"

	ellipse "github.com/Juandavid716/gophysics"
)


func main() {

	// Init the function where parameters are (initial velocity, acceleration, initial displacement)
	fx := disp.DisplaceAcc(5, 6, 4)

	// After t seconds, the displacement is..
	fmt.Println(fx(0))

	// Init the function where parameters are (initial velocity, final velocity, initial displacement)
	fv := disp.DisplaceVec(5, 5, 4)

	// After t seconds, the displacement is..
	fmt.Println(fv(3))
}

Output

4
19
```


## Functions
| Equations                                   | Method                            | Arguments       |  Results | 
| :-------------                         |:-------------                           | :-----            |  :------  | 
| <a href="https://www.codecogs.com/eqnedit.php?latex=x=&space;\frac{1}{2}(V&space;&plus;&space;V_{o})t&space;&plus;&space;x_{0}" target="_blank"><img src="https://latex.codecogs.com/gif.latex?x=&space;\frac{1}{2}(V&space;&plus;&space;V_{o})t&space;&plus;&space;x_{0}" title="x= \frac{1}{2}(V + V_{o})t + x_{0}" /></a>|  DisplaceAcc()   |     Vo, a, Xo   float64   |   float64  | 
| <a href="https://www.codecogs.com/eqnedit.php?latex=x&space;=&space;V_{0}t&plus;&space;\frac{1}{2}at^{2}&space;&plus;&space;x_{o}" target="_blank"><img src="https://latex.codecogs.com/gif.latex?x&space;=&space;V_{0}t&plus;&space;\frac{1}{2}at^{2}&space;&plus;&space;x_{o}" title="x = V_{0}t+ \frac{1}{2}at^{2} + x_{o}" /></a>|  DisplaceVec() |    Vo, Vf, Xo   float64    | float64  | 

