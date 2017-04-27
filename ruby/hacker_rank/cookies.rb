# Enter your code here. Read input from STDIN. Print output to STDOUT
#!/bin/ruby

def get_count(array, len, sweetness)
  heap_sort(array, len)
  return -1 if array[0] < 1 && array[1] < 1
  count = 0
  while (len > 1 && array[0] < sweetness)
    new_sweet = array[0] + 2 * array[1]
    array[0] = new_sweet
    array[1], array[len-1] = array[len-1], nil
    len -= 1
    place_new_sweet(array, sweet)
    # new_arr.concat(array[2..-1])
    return -1 if array[0] < sweetness && len < 2
    heap_sort(array, len)
    count += 1
  end
  count
end

def place_new_sweet(arr, sweet)
  arr[0], arr[1] = -1
  
end

def heap_sort(arr, len)
  @sortcount += 1
  for i in 0...len/2 do
    max_heapify(arr, len, i)
  end

  i = len-1
  while(i >= 0)
    arr[0], arr[i] = arr[i], arr[0]
    max_heapify(arr, i, 0)
    i -= 1
  end
end

def max_heapify(arr, len, root)
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
    max_heapify(arr, len, min)
  end
end

@sortcount = 0
str = (0...100000).map { |x| 1 }.join(' ')
array = str.strip.split(' ').map(&:to_i)
puts "prepared array"
puts get_count(array, array.length, 105800)
# puts get_count(arr, arr.length, k)
