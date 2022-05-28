# Go Bases III

## Exercise 1 - Save file
A company that sells cleaning products needs:
>1. Implement a functionality to save a text file, with the information
of purchased products, separated by semicolons (csv).
>2. Must have product ID, price and quantity.
>3. These values ​​can be hardcoded or manually written to a variable.

## Exercise 2 - Read file
The same company needs to read the stored file, for this it requires that:
>- The tabulated values ​​are printed on the screen, with title (on the left for the ID and on the
right for Price and Quantity), the price, the quantity and below the price the total
should be displayed (adding price per quantity)

Example:
|  ID | Price  | Quantity  | 
| --------- | -------- | ------- |
|  111223 | 30012.00  | 1 |
|444321 |1000000.00 |4|
|434321 |50.50| 1|
||4030062.50| |

## Day 2

## Exercise 1 - Social Network

A social media company needs to implement a user structure with roles
that add information to the structure. To optimize and save memory, they need
that the user structure occupies the same place in memory for the main program and
for the functions:
>- The structure must have the following fields: First name, Last name, age, email and
password
  
And it should implement the following features:
>- change name: allows me to change first and last name
>- change age: allows me to change age
>- change email: allows me to change email
>- change a password: allows me to change a password

## Exercise 2 - E-commerce (Part II)
A large web sales company needs to add functionality to add
products to users. To do this, they ask users and products to have the
same memory address in program main and functions.

Common Structures:
>- User: First name, Last name, E-mail, Products (array of products).
>- Product: Name, price, quantity.
 
Some characteristic functions:
>- New product: receives name and price, and returns a product.
>- Add product: receives user, product and quantity, does not return anything, adds
the user's product.
>- Delete products: receive a user, delete the user's products.

## Exercise 3 - Calculate Price (Part II)
A national company is responsible for selling products, services and maintenance.
For this, they need to run a program that is responsible for calculating the total price
of Products, Services and Maintenance. Due to strong demand and to optimize speed,
they execute the sum calculation to be performed in parallel through **3 go routines**.

Structure of 3 structures:
>- Products: name, price, quantity.
>- Services: name, price, minutes worked.
>- Maintenance: name, price.
Required 3 functions:
>- Add Products: receives a product array and returns the total price (price *
the amount).
>- Sum Services: receives a service matrix and returns the total price (price * average
hour worked, if he does not come to work for 30 minutes, he will be billed as if
had worked half an hour).
>- Add Maintenance: receives a maintenance matrix and delivers the total price.

3 must be fulfilled concomitantly and at the end the final value must be shown in the
screen (adding the total of 3).

## Afternoon


## Exercise 1 - Social Network

A social media company needs to implement a user structure with roles
that add information to the structure. To optimize and save memory, they require
that the user structure occupies the same place in memory for the main program and
for the functions:
>- The structure must have the following fields: First name, Last name, age, email and
password
And they must implement the following features:
>- change name: allows me to change first and last name
>- change age: allows me to change age
>- change email: allows me to change email
>- change password: allows me to change the password


## Exercise 2 - E-commerce (Part II)
A large web sales company needs to add functionality to add
products to users. To do this, they require users and products to have the
same memory address in program main and functions.

Required structures:
>- User: First name, Last name, E-mail, Products (array of products).
>- Product: Name, price, quantity.
Some necessary functions:
>- New product: receives name and price, and returns a product.
>- Add product: receives user, product and quantity, does not return anything, adds
the product to the user.
>- Delete products: receive a user, delete the user's products.

## Exercise 3 - Calculate Price (Part II)
A national company is responsible for selling products, services and maintenance.
For this, they need to run a program that is responsible for calculating the total price
of Products, Services and Maintenance. Due to strong demand and to optimize speed,
they require the sum calculation to be performed in parallel through 3 go routines.

We need 3 structures:
>- Products: name, price, quantity.
>- Services: name, price, minutes worked.
>- Maintenance: name, price.
We need 3 functions:
>- Add Products: receives a product array and returns the total price (price *
the amount).
>- Sum Services: receives a service array and returns the total price (price * average
hour worked, if he does not come to work for 30 minutes, he will be billed as if
had worked half an hour).
>- Add Maintenance: receives a maintenance array and returns the total price.

The 3 must be executed concurrently and at the end the final value must be shown in the
screen (adding the total of 3).


## Exercise 4 - Ordering
A systems company needs to analyze which ranking algorithms to use for its
services.
For them, it is necessary to instantiate 3 arrays with unordered random values
>- an array of integers with 100 values
>- an array of integers with 1000 values
>- an array of integers with 10,000 values

To instantiate the variables, use rand:
```go
package main
import (
"math/rand"
)

func main() {
variable := rand.Perm(100)
variable2 := rand.Perm(1000)
variable3 := rand.Perm(10000)
}
```
Each must be sorted by:
>- Insertion
>- group
>- selection

One routine for each sort run
I have to wait to finish sorting 100 numbers to continue sorting by
1000 and then sorting 10000.
Finally, I must measure the time of each one and show the result on the screen, to know which
sorting got better for each arrangement.



