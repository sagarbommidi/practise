def max_seq_count(arr)
  seq_arr = Array.new(arr.length, 1)
  max = 1
  for i in 1...arr.length do
    seq_arr[i] += seq_arr[i-1] if arr[i] > arr[i-1]
    max = seq_arr[i] if seq_arr[i] > max
  end
  max
end


# arr = [10, 15, 1, 3, 5, 8, 25, 6]
# arr = [0, 8, 4, 12, 2, 10, 6, 14, 1, 9, 5, 13, 3, 11, 7, 15]
arr = [62, 92, 96, 43, 28, 37, 92, 5, 3, 54, 93, 83, 22]
puts max_seq_count(arr)


[62, 92, 96, 43, 28, 37, 92, 5, 3, 54, 93, 83, 22]
[01, 02, 03, 01, 01, 02, 03, 1, 1, 03, 04, 04, 01]