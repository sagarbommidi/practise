def insertion_sort(arr)
  for i in (0...arr.length) do
    current = arr[i]
    j = i-1
    while(j > 0 && arr[j] > current)
      arr[j+1] = arr[j]
      j -= 1
    end
    arr[j+1] = current
  end
end

a = [1, 4, 6, 9, 2, 5, 7, 3, 8]
insertion_sort(a)
puts a.join(', ')
