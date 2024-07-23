## Memory managment in go
# happens automatically, not our job to do it

# memory allocation kewyword:
1. new():
a. memory allocation without initialisation
b. you get memory adress 
c. zeroed storage: In Go, the new keyword does allocate storage for a variable, but it does not initialize it beyond zero values. The new keyword returns a pointer to a newly allocated zero value of the specified type.

2. make():
a. memory allocation with initialisation
b. you get memory adress
c. non-zeroed storage: data can be stored initially

# garbage collection : 
the GOGC variable sets a threshold value for collection. once  that collection value is reached gc starts and remove all out of scope and nil values.