def merge_sort(arr)
  return arr if arr.length < 2
  mid = arr.length/2
  left_arr = arr[0...mid]
  right_arr = arr[mid...arr.length]
  merge_sort(left_arr)
  merge_sort(right_arr)
  merge_arrays(arr, left_arr, right_arr)
end

def merge_arrays(arr, left, right)
  lindex = left.length
  rindex = right.length
  i, j, k = 0, 0, 0
  while(i < lindex && j < rindex) do
    if left[i] <= right[j]
      arr[k] = left[i]
      i += 1
    else
      arr[k] = right[j]
      j += 1
    end
    k += 1
  end

  while(i < lindex) do
    arr[k] = left[i]
    k += 1
    i += 1
  end

  while(j < rindex) do
    arr[k] = right[j]
    k += 1
    j += 1
  end
end

a = [1, 4, 6, 9, 2, 5, 7, 3, 8]
merge_sort(a)
puts a.join(', ')
