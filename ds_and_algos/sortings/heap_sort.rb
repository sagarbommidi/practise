def heap_sort(arr, len)
  i = (len/2) - 1
  while i >= 0
    heapify(arr, len, i)
    i -= 1
  end

  i = len-1
  while(i >= 0)
    arr[0], arr[i] = arr[i], arr[0]
    heapify(arr, i, 0)
    i -= 1
  end
  puts "sorting done"
end

def heapify(arr, len, root)
  min = root
  left = 2 * root + 1
  right = 2 * root + 2

  if (left < len && arr[left] > arr[min])
    min = left
  end

  if (right < len && arr[right] > arr[min])
    min = right
  end

  if(min != root)
    arr[root], arr[min] = arr[min], arr[root]
    heapify(arr, len, min)
  end
end

a = [1, 4, 6, 9, 2, 5, 7, 3, 8]
heap_sort(a, a.length)
puts a.join(', ')
