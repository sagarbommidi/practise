def bubble_sort(arr)
  for i in (0...arr.length-1) do
    swapped = false
    for j in (0...arr.length-1) do
      if (arr[j] > arr[j+1])
        arr[j], arr[j+1] = arr[j+1], arr[j]
        swapped = true
      end
    end
    break if swapped == false
  end
end

a = [1, 4, 6, 9, 2, 5, 7, 3, 8]
bubble_sort(a)
puts a.join(', ')
