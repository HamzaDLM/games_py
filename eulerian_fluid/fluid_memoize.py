import math


class Memoize:
    def __init__(self, f):
        self.f = f
        self.memo = {}

    def __call__(self, *args):
        if not args in self.memo:
            self.memo[args] = self.f(*args)
        # Warning: You may wish to do a deepcopy here if returning objects
        return self.memo[args]


class Fluid:
    def __init__(
        self, N: int, dt: float, diffussion: float, viscosity: float, iter: int
    ) -> None:
        self.size = N
        self.dt = dt
        self.diff = diffussion
        self.visc = viscosity
        self.s: list[float] = [0] * (N * N)
        self.density: list[float] = [0] * (N * N)
        self.Vx: list[float] = [0] * (N * N)
        self.Vy: list[float] = [0] * (N * N)
        self.Vx0: list[float] = [0] * (N * N)
        self.Vy0: list[float] = [0] * (N * N)
        self.iter = iter

    def add_density(self, x: int, y: int, amount: int) -> None:
        """Add density to the element that will be added to the water (e.g. soy sauce)"""
        index = IX(x, y)
        self.density[index] += amount

    def add_velocity(self, x: int, y: int, amount_x: float, amount_y: float) -> None:
        index = IX(x, y)
        self.Vx[index] += amount_x
        self.Vy[index] += amount_y

    def step(self) -> None:
        """Go through time with the fluid"""

        diffuse(1, self.Vx0, self.Vx, self.visc, self.dt, self.iter, self.size)
        diffuse(2, self.Vy0, self.Vy, self.visc, self.dt, self.iter, self.size)

        project(self.Vx0, self.Vy0, self.Vx, self.Vy, self.iter, self.size)

        advect(1, self.Vx, self.Vx0, self.Vx0, self.Vy0, self.dt, self.size)
        advect(2, self.Vy, self.Vy0, self.Vx0, self.Vy0, self.dt, self.size)

        project(self.Vx, self.Vy, self.Vx0, self.Vy0, self.iter, self.size)

        diffuse(0, self.s, self.density, self.diff, self.dt, self.iter, self.size)
        advect(0, self.density, self.s, self.Vx, self.Vy, self.dt, self.size)

    def fade_density(self, amount):
        for i in range(0, len(self.density)):
            if self.density[i] > 0:
                self.density[i] -= amount


def diffuse(
    b: int, x: list[float], x0: list[float], diff: float, dt: float, iter: int, N: int
) -> None:
    a: float = dt * diff * (N - 2) * (N - 2)
    lin_solve(b, x, x0, a, 1 + 6 * a, iter, N)


@Memoize
def constrain(value, minimum, maximum):
    if value < minimum:
        return minimum
    elif value > maximum:
        return maximum
    else:
        return value


@Memoize
def IX(
    x: int, y: int, N: int = 128
) -> int:  # TODO: Fix value for N (should be dynamic)
    """Return 2D location as a 1D index"""
    x = constrain(x, 0, N - 1)
    y = constrain(y, 0, N - 1)
    return int(x + y * N)


def lin_solve(
    b: int, x: list[float], x0: list[float], a: float, c: float, iter: int, N: int
) -> None:
    cRecip: float = 1.0 / c
    for k in range(0, iter):
        for j in range(1, N - 1):
            for i in range(1, N - 1):
                x[IX(i, j)] = (
                    x0[IX(i, j)]
                    + a
                    * (
                        x[IX(i + 1, j)]
                        + x[IX(i - 1, j)]
                        + x[IX(i, j + 1)]
                        + x[IX(i, j - 1)]
                    )
                    * cRecip
                )
        set_bnd(b, x, N)


def project(
    velocX: list[float],
    velocY: list[float],
    p: list[float],
    div: list[float],
    iter: int,
    N: int,
) -> None:
    for j in range(1, N - 1):
        for i in range(1, N - 1):
            div[IX(i, j)] = (
                -0.5
                * (
                    velocX[IX(i + 1, j)]
                    - velocX[IX(i - 1, j)]
                    + velocY[IX(i, j + 1)]
                    - velocY[IX(i, j - 1)]
                )
                / N
            )
            p[IX(i, j)] = 0

    set_bnd(0, div, N)
    set_bnd(0, p, N)
    lin_solve(0, p, div, 1, 6, iter, N)

    for j in range(1, N - 1):
        for i in range(1, N - 1):
            velocX[IX(i, j)] -= 0.5 * (p[IX(i + 1, j)] - p[IX(i - 1, j)]) * N
            velocY[IX(i, j)] -= 0.5 * (p[IX(i, j + 1)] - p[IX(i, j - 1)]) * N
    set_bnd(1, velocX, N)
    set_bnd(2, velocY, N)


def advect(
    b: int,
    d: list[float],
    d0: list[float],
    velocX: list[float],
    velocY: list[float],
    dt: float,
    N: int,
) -> None:
    dtx: float = dt * (N - 2)
    dty: float = dt * (N - 2)

    for j in range(1, N - 1):
        for i in range(1, N - 1):
            tmp1 = dtx * velocX[IX(i, j)]
            tmp2 = dty * velocY[IX(i, j)]
            x = i - tmp1
            y = j - tmp2

            if x < 0.5:
                x = 0.5
            if x > N + 0.5:
                x = N + 0.5
            i0 = math.floor(x)
            i1 = i0 + 1.0

            if y < 0.5:
                y = 0.5
            if y > N + 0.5:
                y = N + 0.5
            j0 = math.floor(y)
            j1 = j0 + 1.0

            s1 = x - i0
            s0 = 1.0 - s1
            t1 = y - j0
            t0 = 1.0 - t1

            i0i = int(i0)
            i1i = int(i1)
            j0i = int(j0)
            j1i = int(j1)

            d[IX(i, j)] = (
                s0 * (t0 * d0[IX(i0i, j0i)])
                + (t1 * d0[IX(i0i, j1i)])
                + s1 * (t0 * d0[IX(i1i, j0i)])
                + (t1 * d0[IX(i1i, j1i)])
            )

    set_bnd(b, d, N)


def set_bnd(b: int, x: list[float], N: int = 128) -> None:
    """a way to keep fluid from leaking out of your box"""

    for i in range(1, N - 1):
        if b == 2:
            x[IX(i, 0)] = -x[IX(i, 1)]
            x[IX(i, N - 1)] = -x[IX(i, N - 2)]
        else:
            x[IX(i, 0)] = x[IX(i, 1)]
            x[IX(i, N - 1)] = x[IX(i, N - 2)]

    for j in range(1, N - 1):
        if b == 1:
            x[IX(0, j)] = -x[IX(1, j)]
            x[IX(N - 1, j)] = -x[IX(N - 2, j)]
        else:
            x[IX(0, j)] = x[IX(1, j)]
            x[IX(N - 1, j)] = x[IX(N - 2, j)]

    x[IX(0, 0, 0)] = 0.33 * (x[IX(1, 0, 0)] + x[IX(0, 1, 0)] + x[IX(0, 0, 1)])
    x[IX(0, N - 1, 0)] = 0.33 * (
        x[IX(1, N - 1, 0)] + x[IX(0, N - 2, 0)] + x[IX(0, N - 1, 1)]
    )
    x[IX(N - 1, 0, 0)] = 0.33 * (
        x[IX(N - 2, 0, 0)] + x[IX(N - 1, 1, 0)] + x[IX(N - 1, 0, 1)]
    )
    x[IX(N - 1, N - 1, 0)] = 0.33 * (
        x[IX(N - 2, N - 1, 0)] + x[IX(N - 1, N - 2, 0)] + x[IX(N - 1, N - 1, 1)]
    )
