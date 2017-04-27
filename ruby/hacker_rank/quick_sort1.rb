# https://www.hackerrank.com/challenges/quicksort1

def quickSort(arr, left, right)
  if(right > left)
    pivot_index = partition(arr, left, right)
    quickSort(arr, left, pivot_index-1)
    quickSort(arr, pivot_index+1, right)
  end
end

def partition(arr, low, high)
  pivot_ele = arr[piv_index]
  left, right = low, high
  while(left < right)
    while(arr[left] <= pivot_ele)
      left += 1
      break if left == high
    end

    while(arr[right] > pivot_ele)
      right -= 1
      break if right == low
    end

    if(left < right)
      arr[left], arr[right] = arr[right], arr[left]
    end
  end
  arr[piv_index], arr[right] = arr[right], pivot_ele
  right
end


ar = [5, 8, 1, 3, 7, 9, 2]
quickSort(ar, 0, ar.length-1)
puts ar.join(', ')
