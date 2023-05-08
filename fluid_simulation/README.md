# Fluid simulation - PyGame

Ported using as reference the paper of **Jos Stam** on **Real-Time Fluid Dynamics for Games**: https://www.dgp.toronto.edu/public_user/stam/reality/Research/pdf/GDC03.pdf

TODO:

- Re-write fluid.py in Cython inshaalah to enhance performance as the Python implementation is painfully slow.
- Add better density coloring.
- Add visual arrows indicating velocity vectors.
- Enable changing parameters (density, velocity...) on GUI.

Also, isn't this a gorgeous equation :O

$$
\begin{equation}
{\displaystyle \rho {\frac {\mathrm {D} \mathbf {u} }{\mathrm {D} t}}=\rho \left ({\frac {\partial \mathbf {u} }{\partial t}}+(\mathbf {u} \cdot \nabla )\mathbf {u} \right)=-\nabla p+\nabla \cdot \left\\{\mu \left[\nabla \mathbf {u} +(\nabla \mathbf {u} )^{\mathrm {T} }-{\tfrac {2}{3}}(\nabla \cdot \mathbf {u} )\mathbf {I} \right]+\zeta (\nabla \cdot \mathbf {u} )\mathbf {I} \right\\}+\rho \mathbf {g}.}
\end{equation}
$$

<p align="center">
Navier-Stokes equation (general form)
</p>

PS: Not an academic work, just a for fun thing.
