# Deep Neural Network from scratch

This deep neural network does the simple task of predicting the value of a written digit.
Written from scratch with no third library.

## Going over the basics

### Perceptrons

A perceptron is an artificial neuron that takes some inputs $x_1, x_2, x_3$ and procudes a single binary output. It uses weights $w_1, w_2, w_3$ to detemine which input matters more in predicting the output. Thus the output will related to whether the sum $\sum_j w_j x_j$ is bigger than a threshold.

![](data/1.png)

### Network of perceptrons

Generall, a perceptron won't be able to do much, so to accomplish a task such as predicting hand-written numbers we need a network of perceptrons that is basically layers of perceptrons put in front of each other. The layers between the inputs and the output are called hidden layers. An example of a two hidden layered network looks like the following:

![](data/2.png)

### Bias

A bias is basically the threshold moved to the other side of the inequality and defines the measure of how easy to get the perceptron to output the correct value.

### Sigmoid

A sigmoid functtion also called logistic function, is a method by which we map whatever number into a small range spanning between 0 and 1.
$$\sigma(z) \equiv \frac{1}{1+e^{-z}}$$
For our case, written such as:
$$ \frac{1}{1+\exp(-\sum_j w_j x_j-b)}$$

### Cost function

A cost function (also called loss function) lets us define how well our prediction is doing, and is denoted with the following formula, with $n$ being the size of training inputs, $a$ as the vector of outputs:

$$C(w,b) \equiv \frac{1}{2n} \sum_x \| y(x) - a\|^2$$

$C(w,b) \approx0$ when $y(x)$ is equal to the input $a$. This will enable us to determine the best set of weights and biases so that our model behaves in a good way.

### Gradient Descent Algorithm

We said earlier that we need to find parameters $w$ and $b$ such as the C function leans toward 0, one way for doing it is to use Gradient descent which is an iterative optimization algorithm used to find the minimum of a function. It works by adjusting the parameters of a model in small steps, guided by the negative gradient (opposite direction of the steepest ascent) of the function, until it reaches a local or global minimum.

For the gradient descent to work properly we need to choose a good learning rate $\eta$ that is small enough so that $\Delta C > 0$ doesn't happen.

Gradient descent update equations become:

$$w_k \rightarrow w_k' = w_k-\eta \frac{\partial C}{\partial w_k}$$

$$b_l \rightarrow b_l' = b_l-\eta \frac{\partial C}{\partial b_l}$$

### Stochastic Gradient Descent

To compute the gradient $\nabla C$ we need to compute the gradients $\nabla C_x$ for each training input $x$, and calculate the average $\nabla C = \frac{1}{n}\sum_x \nabla C_x$, which proves to be a slow process for large number of inputs. We will use Stochastic GD which takes a number $m$ of randomly chosen inputs and use them to ge the average $\nabla C_x$.

The update equations thus become:

$$
w_k  \rightarrow  w_k' = w_k-\frac{\eta}{m}
  \sum_j \frac{\partial C_{X_j}}{\partial w_k}\\
$$

$$
b_l  \rightarrow  b_l' = b_l-\frac{\eta}{m}
\sum_j \frac{\partial C_{X_j}}{\partial b_l},
$$

TODOs

- Write tests for Matrice operations
- Document code

#### Refs:

Training/Test data used: https://www.kaggle.com/competitions/digit-recognizer/data
