arr = []  #initiate a list
arr = [0]*12  #initiate sized list
arr.append(20) #add element to list
arr[0]=444 #change element to list
len(arr) #lenght of the list

#list extending (list can contain different types(int, strings, bool)
arr.extend(['a', 5, True, 'hello'])

#get element by index
arr[16] #'hello' 
arr[14] #5

#results
print(arr[0]) #444
print(arr[16]) #'hello'
print(arr[15]) #True
print(arr[14]) #5
print((arr))