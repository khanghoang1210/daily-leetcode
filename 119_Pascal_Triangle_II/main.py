# Time: O(n^2) / Space: O(n)

def getRow(rowIndex: int) :
    arr = [1]
    for i in range(0, rowIndex):
        if i == rowIndex:
            return arr
        else:
            tmp = arr
            arr = []
            for j in range(0, len(tmp) - 1):
                arr.append(tmp[j] + tmp[j+1])
            arr.append(1)
            arr.reverse()
            arr.append(1)
    return arr
            
print(getRow(6)) # [1, 6, 15, 20, 15, 6, 1]