# EverPhone Assignment API 

### API endpoint:
http://localhost:8088/assign-gift/:emplayeeName

### Assignment

1. Fetched the employee and gifts from file and based on employee's interest try to find similar interst in gift if there is a match then return the gift
otherwise try to find similar interest gift.<br>

2. In case of employee rush to the website then multiple threads will get created and try to find similar interst gift. In this case chances are we can assign same gift to two different person. To prevent such thing put lock(mutex) before start finding similar interest to make sure no other thead is accessing same gift array.<br>

3. In case we do not have bought appropriate gifts, create the category of similar kind of interest. In case employee do not match have exact match for gift then we can fetch the similar interest from category and again try to match.<br>

4. Once an employee has been assigned a gift that’s it, no change can be enforce by having another structure which contain employee and gift. Before assigning gift to employee, search the employee name in this array of structure. If name present then return otherwise try to find gift to the employee.<br>

5. If in future employees to be able to return a gift once then just go through the assigned strucure and delete the employee and try to find the new gift for the employee.<br>


### Compliation

I have used windows for the development as my windows WSL Ubuntu is having some network issue.<br>

#### To run the source code:
```
 go run .\main.go
 ```
