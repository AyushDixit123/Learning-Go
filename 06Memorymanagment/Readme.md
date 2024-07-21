## Memory managment in go
# happens automatically, not our job to do it

# memory allocation kewyword:
1. new():
a. memory allocation without initialisation
b. you get memory adress 
c. zeroed storage: no data can be stored initially

2. make():
a. memory allocation with initialisation
b. you get memory adress
c. non-zeroed storage: data can be stored initially

# garbage collection : 
the GOGC variable sets a threshold value for collection. once  that collection value is reached gc starts and remove all out of scope and nil values.