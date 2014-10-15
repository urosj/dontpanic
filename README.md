A simple fuzzing test golang util.

The input is any interface {} object. Using reflection, the values
of the interface are replaced by random values of proper type. In case 
we are dealing with a map, the keys remain the same, but the values change.

For some types, initial zero value is replaced with random value without 
additional initialization.

In case of maps and slices, some values need to be initialized in the
struct which are then replaced with random values. See example.