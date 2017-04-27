# https://www.hackerrank.com/challenges/quicksort2

def quickSort(arr, findex, lindex)
  return if findex >= lindex
  pivot_index = partition(arr, findex, lindex)
  quickSort(arr, findex, pivot_index-1)
  quickSort(arr, pivot_index+1, lindex)
end

def partition(arr, left, right)
  pivot = arr[right]
  pindex = left
  for i in (left...right) do
    if(arr[i] <= pivot)
      arr[pindex], arr[i] = arr[i], arr[pindex]
      pindex += 1
    end
  end
  arr[pindex], arr[right] = arr[right], arr[pindex]
  pindex
end

ar = [5, 8, 1, 3, 7, 9, 2]
quickSort(ar, 0, ar.length-1);
