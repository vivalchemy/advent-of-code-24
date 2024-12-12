# Problem Statement 1

## Description
Compare the list elements one by one

---

## Breakdown

1. Sort the both the lists
2. Pair the numbers based on the indexes
3. Calculate the distance between the two numbers
4. Add all the distances
5. Print the total distance

---

## Steps

1. Read from the stdin piped to the program/ Open a file and read it
2. Seperate the lists by alternating between the received numbers
3. Sort each list individually(some sorting optimization can be done)
4. Based on the length of either of the list loop over both the lists and calculate the total sum
> [!NOTE]
> Assumption that the length of both the list will be same
5. Print the result



# Problem Statement 2

## Description
Find the similarities between the two lists

---

## Breakdown

1. Get the number of time a number on left list repeats on the right list
2. Multiple the left number by the count from the right list
3. Add the numbers to obtain the final result

---

## Steps

1. Read from the stdin piped to the program/ Open a file and read it
2. Seperate the lists by alternating between the received numbers
3. Create a map with number relating to count for the right array
4. iterate the left list and multiple with count using the map
5. return the sum of all the multiplications
