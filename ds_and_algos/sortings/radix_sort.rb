def radix_sort(arr)
  max = get_max(arr)
  exp = 1

  begin
    count_sort(arr, exp)
    exp *= 10
  end while (max/exp > 0)
end

def count_sort(arr, exp)
  len = arr.length
  count = Array.new(10, 0)
  output = Array.new(len)

  for i in 0...len do
    count[(arr[i]/exp)%10] += 1
  end

  for i in 1...10 do
    count[i] += count[i-1]
  end

  i = len - 1
  begin
    output[count[(arr[i]/exp)%10] - 1] = arr[i]
    count[(arr[i]/exp)%10] -= 1
    i -= 1
  end while (i >= 0)

  for i in 0...len do
    arr[i] = output[i]
  end
end

def get_max(arr)
  max = arr[0]
  for i in 1...arr.length do
    max = arr[i] if arr[i] > max
  end
  max
end

a = [1, 421, 6, 9, 2, 5, 7, 3, 8, 12, 1]
radix_sort(a)
puts a.join(', ')
