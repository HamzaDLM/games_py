package main

import (
	"math"
)

const (
	N    int = 124
	iter int = 15
)

type Fluid struct {
	size int64
	dt   float64
	diff float64
	visc float64

	s       []float64
	density []float64

	Vx []float64
	Vy []float64
	Vz []float64

	Vx0 []float64
	Vy0 []float64
	Vz0 []float64
}

func step(f Fluid) {
	diffuse(1, f.Vx0, f.Vx, f.visc, f.dt)
	diffuse(2, f.Vy0, f.Vy, f.visc, f.dt)

	project(f.Vx0, f.Vy0, f.Vx, f.Vy)

	advect(1, f.Vx, f.Vx0, f.Vx0, f.Vy0, f.dt)
	advect(2, f.Vy, f.Vy0, f.Vx0, f.Vy0, f.dt)

	project(f.Vx, f.Vy, f.Vx0, f.Vy0)

	diffuse(0, f.s, f.density, f.diff, f.dt)
	advect(0, f.density, f.s, f.Vx, f.Vy, f.dt)

}

// Fluid methods
func addDensity(f Fluid, x int, y int, amount float64) {
	index := IX(x, y)
	f.density[index] += amount
}

func addVelocity(f Fluid, x int, y int, amountX float64, amountY float64) {
	index := IX(x, y)
	f.Vx[index] += amountX
	f.Vy[index] += amountY
}

// Fluid functions
func setBnd(b int, x []float64) {
	for i := 1; i < N-1; i++ {
		if b == 2 {
			x[IX(i, 0)] = -x[IX(i, 1)]
			x[IX(i, N-1)] = -x[IX(i, N-2)]
		} else {
			x[IX(i, 0)] = x[IX(i, 1)]
			x[IX(i, N-1)] = x[IX(i, N-2)]
		}
	}
	for j := 1; j < N-1; j++ {
		if b == 1 {
			x[IX(0, j)] = -x[IX(1, j)]
			x[IX(N-1, j)] = -x[IX(N-2, j)]
		} else {
			x[IX(0, j)] = x[IX(1, j)]
			x[IX(N-1, j)] = x[IX(N-2, j)]
		}
	}

	x[IX(0, 0)] = 0.5 * (x[IX(1, 0)] + x[IX(0, 1)])
	x[IX(0, N-1)] = 0.5 * (x[IX(1, N-1)] + x[IX(0, N-2)])
	x[IX(N-1, 0)] = 0.5 * (x[IX(N-2, 0)] + x[IX(N-1, 1)])
	x[IX(N-1, N-1)] = 0.5 * (x[IX(N-2, N-1)] + x[IX(N-1, N-2)])
}

func advect(b int, d []float64, d0 []float64, velocX []float64, velocY []float64, dt float64) {
	var i0, i1, j0, j1 float64

	var dtx float64 = dt * (float64(N) - 2)
	var dty float64 = dt * (float64(N) - 2)

	var s0, s1, t0, t1 float64
	var tmp1, tmp2, x, y float64

	for j := 1; j < N-1; j++ {
		for i := 1; i < N-1; i++ {

			tmp1 = dtx * velocX[IX(i, j)]
			tmp2 = dty * velocY[IX(i, j)]
			x = float64(i) - tmp1
			y = float64(j) - tmp2

			if x < 0.5 {
				x = 0.5
			}
			if x > float64(N)+0.5 {
				x = float64(N) + 0.5
			}
			i0 = math.Floor(x)
			i1 = i0 + 1.0

			if y < 0.5 {
				y = 0.5
			}
			if y > float64(N)+0.5 {
				y = float64(N) + 0.5
			}
			j0 = math.Floor(y)
			j1 = j0 + 1.0

			s1 = x - i0
			s0 = 1.0 - s1
			t1 = y - j0
			t0 = 1.0 - t1

			i0i := int(i0)
			i1i := int(i1)
			j0i := int(j0)
			j1i := int(j1)

			d[IX(i, j)] = s0*(t0*d0[IX(i0i, j0i)]) + (t1 * d0[IX(i0i, j1i)]) + s1*(t0*d0[IX(i1i, j0i)]) + (t1 * d0[IX(i1i, j1i)])

		}
	}
	setBnd(b, d)
}

func project(velocX []float64, velocY []float64, p []float64, div []float64) {
	for i := 1; i < N-1; i++ {
		for j := 1; j < N-1; j++ {
			div[IX(i, j)] = -0.5 * (velocX[IX(i+1, j)] - velocX[IX(i-1, j)] + velocY[IX(i, j+1)] - velocY[IX(i, j-1)]) / float64(N)
			p[IX(i, j)] = 0
		}
	}

	setBnd(0, div)
	setBnd(0, p)
	linSolve(0, p, div, 1, 6)
	for i := 1; i < N-1; i++ {
		for j := 1; j < N-1; j++ {
			velocX[IX(i, j)] -= 0.5 * (p[IX(i+1, j)] - p[IX(i-1, j)]) * float64(N)
			velocY[IX(i, j)] -= 0.5 * (p[IX(i, j+1)] - p[IX(i, j-1)]) * float64(N)
		}
	}
	setBnd(1, velocX)
	setBnd(2, velocY)

}

func linSolve(b int, x []float64, x0 []float64, a float64, c float64) {
	var cRecip float64 = 1.0 / c
	for k := 0; k < iter; k++ {
		for j := 1; j < N-1; j++ {
			for i := 1; i < N-1; i++ {
				x[IX(i, j)] = x0[IX(i, j)] + a*(x[IX(i+1, j)]+x[IX(i-1, j)]+x[IX(i, j+1)]+x[IX(i, j-1)])*cRecip
			}
		}
		setBnd(b, x)
	}

}

func diffuse(b int, x []float64, x0 []float64, diff float64, dt float64) {
	var a float64 = dt * diff * (float64(N) - 2) * (float64(N) - 2)
	linSolve(b, x, x0, a, 1+6*a)
}

// Helper functions
func IX(x int, y int) int {
	x = constrain(x, 0, N-1)
	y = constrain(x, 0, N-1)
	return x + y*N
}

func constrain[V int | float64](val V, min V, max V) V {
	if val > max {
		return max
	} else if val < min {
		return min
	} else {
		return val
	}
}

// Create an instance of fluid
func newFluid(size int64, dt float64, diff float64, visc float64) Fluid {
	fluid := Fluid{
		size: size,
		dt:   dt,
		diff: diff,
		visc: visc,

		s:       make([]float64, N*N),
		density: make([]float64, N*N),

		Vx: make([]float64, N*N),
		Vy: make([]float64, N*N),
		Vz: make([]float64, N*N),

		Vx0: make([]float64, N*N),
		Vy0: make([]float64, N*N),
		Vz0: make([]float64, N*N),
	}

	return fluid
}
