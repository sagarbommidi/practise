def quick_sort(arr, findex, lindex)
  return if findex >= lindex
  pivot_index = partition(arr, findex, lindex)
  quick_sort(arr, findex, pivot_index-1)
  quick_sort(arr, pivot_index+1, lindex)
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


a = [1, 4, 6, 9, 2, 5, 7, 3, 8]
quick_sort(a, 0, a.length-1)
puts a.join(', ')
