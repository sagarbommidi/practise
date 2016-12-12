def selection_sort(arr)
  for i in (0...arr.length) do
    min = i
    for j in (i+1...arr.length) do
      min = j if arr[j] < arr[min]
    end
    arr[i], arr[min] = arr[min], arr[i]
  end
end

a = [1, 4, 6, 9, 2, 5, 7, 3, 8]
selection_sort(a)
puts a.join(', ')
