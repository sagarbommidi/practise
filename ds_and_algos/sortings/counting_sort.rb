def counting_sort(arr)
  min = arr.min
  max = arr.max
  k = max - min + 1
  count = Array.new(k, 0)
  len = arr.length
  output = Array.new(len)

  for i in 0...len do
    count[arr[i]-min] += 1
  end

  for i in 1...k do
    count[i] += count[i-1]
  end

  for i in 0...len do
    output[count[arr[i] - min] - 1] = arr[i]
    count[arr[i] - min] -= 1
  end
  output
end

a = [1, 4, 6, 9, 2, 5, 7, 3, 8, 122, 2]
a = counting_sort(a)
puts a.join(', ')
