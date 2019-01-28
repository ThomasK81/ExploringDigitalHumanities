# R is like a fast calculator

232047294879347 * 242424 / 12313 +2

# in R <- is used to assign a value to a variable. You can also use =

test_variable <- 232047294879347 * 242424 / 12313 +2
test_variable

#you can use rm() to remove a variable, object
rm(test_variable)
test_variable = 2
test_variable

# you do not have to remove to update or change the variable
# and you can change the variable using it's own value
test_variable <- test_variable * test_variable
test_variable

#### The datatype of each variable is inferred

#### R's datatypes 

# Logical is either TRUE or FALSE
variable <- FALSE 
class(variable)

# Numeric is a number that can include fractual parts, i.e. 2.34 
variable <- 2.34 
class(variable)

# Integer is a whole number
variable <- 2L 
variable
class(variable)

##### Sidetrack

# Integers are cheaper than Numerics
object.size(c(1L,2L,3L,4L,5L)) < object.size(c(1,2,3,4,5))
# Unless they are empty
object.size(numeric()) == object.size(integer())
# Than they are both 40bytes
object.size(numeric())

##### Back to data types

# You can also express complex numbers
variable <- 3 + 2i
variable
class(variable)
# And plot them
plot(variable)

# Strings are called characters in R
variable <- "hello world"
variable
class(variable)

# And there is also raw data
variable <- charToRaw(variable)
variable
class(variable)

# raw objects are small. empty raw objects are also 40bytes though

object.size(raw())

### The types you will use most in DH are characters, numeric, and logical

### Technically everything that you have produced so far are vectors
### Vectors can be combined using the c() function

VectorA <- 2
VectorB <- c(1L,2L)
c(VectorA,VectorB)
class(c(VectorA,VectorB))
class(c(VectorB))

### A list is an R-object which can contain many different types
### like vectors, functions and even another list inside it.

list(VectorA, VectorB, list(VectorA, VectorB), mean)

### A matrix is a two-dimensional rectangular data set of ONE data type. 
### It can be created using a vector input to the matrix function.
### Create a matrix.
M <- matrix( c('a','a','b','c','b','a'), nrow = 2, ncol = 3, byrow = TRUE)
M

# Hold on, how does the matrix() function work
?matrix

# Compare
M <- matrix(c('a','a','b','c','b','a'), nrow = 2, ncol = 3, byrow = FALSE)
M

# Got it?

### A factor stores the vector and all its unique values
# Create a vector.
apples <- c('iOS','iPhone','MBP','iPhone','MBP')
# Create a factor object.
factor_apples <- factor(apples)

# Have a look
apples
factor_apples
nlevels(factor_apples)
factor_apples[4]
as.character(factor_apples[4])

### Data Frames are mixed data presented in tabular form
# Create the data frame.
data_types <- 	data.frame(
  types = c("logical", "numerical","integer", "complex", "character", "raw"), 
  usage_probability = c(1, 1, 1, 0.1, 1, 0.5), 
  got_it = c(rep("yes", 5), "no"),
  answer = rep(42L, 6)
)
plot(data_types$types,data_types$usage_probability)


# Data Structures often used in programming

# Strings and arrays

## string in R

newString <- "test"
length(newString)
nchar(newString)
gregexpr('e', newString)[[1]][1]
substr(newString, 2, 2)

##one dimensional arrays of simple types in R are best expressed as vectors. 
# there is the datatype array in R which can deal with multidimensionional array

intArray <- c(7L,8L,9L,10L)
length(intArray)
intArray 

## sidetrack
intRealArray <- array(c(7L,8L,9L,10L),dim = c(4))
intRealArray
intRealArray <- array(c(7L,8L,9L,10L),dim = c(3,3,3))

### A two dimensional array is a matrix
M <- matrix(c('dr','who'), nrow = 3, ncol = 2)
A <- array(c('dr','who'),dim = c(3,2))
A == M
is.matrix(A)

## Back to class

intArray[0]
intArray[1]

which(intArray == 9)

charArray <- c("one","two")
charArray[2]
boolArray <- c(T,F)
boolArray[2]
numArray <- c(7.8,9.1)
numArray[2]

typeof(intArray)
typeof(numArray)
typeof(charArray)
typeof(boolArray)

typeof(newString)

## hash tables
## R uses something like it that is called environment and lists

mylist <- list()
for (i in 1:length(LETTERS)) {
  mylist[LETTERS[i]] <- list(rnorm(5, mean = 2))
}

names(emptylist)
mylist[c("A","B")]

hashTableIsh <- new.env()
list2env(mylist, envir = hashTableIsh)

hashTableIsh[["A"]]
# now it is hashed, but I cannot extract a vector anymore
hashTableIsh[[c("A", "B")]]
hashTableIsh[c("A", "B")]

get("A", envir=hashTableIsh)

## there is a hashmap library though

# Linked Lists
## linked list are directly implemented in R using list
## R does the memory management for you... ... 

lst <- list() # creates an empty (length zero) list
lst
lst[[1]] <- 1 # automagically extends the lst
lst[[2]] <- 2 # ditto
lst

lst <- vector("list", 5)
lst
lst <- list(1, 2, 3, 4, 5)
lst
lst <- lst[-2]
lst

# Queues and stacks 
## they are implemented usually via third party libraries
## you can write your own though

queue <- function(){ 
  e <- new.env() 
  q <- list() 
  assign("q",q,envir=e) 
  class(e) <- c("queue","environment") 
  e 
} 

push.queue <- function(e,v){ 
  q <- c(v,get("q",envir=e)) 
  assign("q",q,envir=e) 
  v 
} 

pop.queue <- function(e){ 
  q <- get("q",envir=e) 
  v <- q[[length(q)]] 
  if(length(q)==1){ 
    assign("q",list(),e) 
  }else{ 
    assign("q",q[1:(length(q)-1)],e) 
  } 
  return(v) 
} 

print.queue <- function(x,...){ 
  print(get("q",envir=x)) 
} 

q <- queue()
push.queue(q,"hello") 
push.queue(q,"world")
push.queue(q,42)

pop.queue(q) 
pop.queue(q) 
pop.queue(q) 
pop.queue(q) 


# for Stack change either pop.queue or push.queue

pop.stack <- function(e){ 
  q <- get("q",envir=e) 
  v <- q[[1]] 
  if(length(q)==1){ 
    assign("q",list(),e) 
  }else{ 
    assign("q",q[2:(length(q))],e) 
  } 
  return(v) 
}

push.queue(q,"hello") 
push.queue(q,"world")
push.queue(q,42)

pop.stack(q)

# Tree  
## Tree is just a list of lists
## (R likes copying (a lot), so if you want to store the data as a more complex graph, you need to use a graph db)

tree <- list(list(1, 2), list(3, list(4, 5)))

# left child: list(1, 2)
tree[[1]]

# right child
tree[[2]]

# left child of right child:list(4, 5)
tree[[2]][[1]]



