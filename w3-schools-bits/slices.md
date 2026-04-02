## The Summary of GO Slices:

Slices are more powerful and flexible than arrays.

Slices are used to store multiple values of the same type in a single variable and the length of a slice can grow and shrink as you see fit.

Ways to create a slice includes:

Using the []datatype{values} format.
Creating a slice from an array.
Using the make() function.

*We can access a specific slice element by referring to the index number.
*We can append elements to the end of a slice using the append()function.
*The copy() function creates a new underlying array with only the required elements for the slice. This will reduce the memory used for the program. 