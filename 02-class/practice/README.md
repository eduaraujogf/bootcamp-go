# Go bases II

## Morning

## Exercise 1 - Wage taxes

A chocolate company needs to calculate the tax on its employees at the time of
deposit the salary, to fulfill its objective it will be necessary to create a function that returns the
salary tax.
>We have information that one of the employees earns a salary of R$50,000 and will be
deducted 17% of salary. Another employee earns a salary of $150,000, and in that
In this case, a further 10% will be deducted.

## Exercise 2 - Calculate average
A school needs to calculate the grade point average (per student). We need to create a function in
which we can pass N number of integers and return the mean.
>Note: If one of the numbers entered is negative, the application should return an error

## Exercise 3 - Calculating salary

A shipping company needs to calculate the salary of its employees based on the number
of hours worked per month and in the employee's category.

>If the category is C, your salary is R$1,000 per hour

>If Category B, your salary is $1,500 per hour plus 20% if over 160 hours
monthly

>If Category A, your salary is $3,000 per hour plus 50% if over 160 hours
monthly

>Calculate the salary of employees according to the information below:
>- Category C employee: 162 hours per month
>- Category B employee: 176 hours per month
>- Category A employee: 172 hours per month


## Exercise 4 - Calculating Statistics

The professors of a university in Colombia need to calculate some statistics of
grades of the students of a course, being necessary to calculate the minimum, maximum and average values
of your notes.
You will need to create a function that indicates what type of calculation should be performed (minimum,
maximum or average) and which returns another function (and a message if the calculation is not
defined) that can be passed an amount N of integers and return the calculation that was
indicated in the previous function
Example:
```go
const (
minimum = "minimum"
average = "average"
maximum = "maximum"
)

...
myFunc, err := operation(minimum)
averageFunc, err := operation(average)
maxFunc, err := operation(maximum)

...
minValue := myFunc(2, 3, 3, 4, 10, 2, 4, 5)
averageValue := averageFunc(2, 3, 3, 4, 1, 2, 4, 5)
maxValue := maxFunc(2, 3, 3, 4, 1, 2, 4, 5)

```

Exercise 5 - Calculation of the amount of food
An animal shelter needs to find out how much food to buy for pets.
At the moment they only have tarantulas, hamsters, dogs and cats, but the forecast
is that there are many more animals to house.
1. Dogs need 10 kg of food
2. Cats 5 kg
3. Hamster 250 grams.
4. Tarantula 150 grams.

We need:
>1. Implement an Animal function that receives as a parameter a value of type text
with the specified animal and that returns a function with a message (if not
the animal exists)
>2. A function for each animal that calculates the amount of food based on the
required amount of the typed animal.
Example:

```go
const (
dog = "dog"
cat = "cat"
)

...
animalDog, msg := Animal(dog)
animalCat, msg := Animal(cat)

...
var amount float64
amount+= animalDog(5)
amount+= animalCat(8)
```