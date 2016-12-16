def quick_sort(arr, findex, lindex)
  return if findex >= lindex
  pivot_index = partition1(arr, findex, lindex)
  quick_sort(arr, findex, pivot_index-1)
  quick_sort(arr, pivot_index+1, lindex)
end

# The below partition is based on the PIVOT element as last element in sub array
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

# The below partition is based on the PIVOT element as first element in sub array
def partition1(arr, left, right)
  pele = arr[left]
  pindex = right
  i = right
  while(i > left)
    if(arr[i] > pele)
      arr[i], arr[pindex] = arr[pindex], arr[i]
      pindex -= 1
    end
    i -= 1
  end
  arr[left], arr[pindex] = arr[pindex], arr[left]
  pindex
end

a = [1, 4, 6, 9, 2, 5, 7, 3, 8]
quick_sort(a, 0, a.length-1)
puts a.join(', ')
