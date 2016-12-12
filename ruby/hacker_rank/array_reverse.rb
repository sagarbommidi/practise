#!/bin/ruby

# n = gets.strip.to_i
# arr = gets.strip
# arr = arr.split(' ').map(&:to_i)

arr = [1, 4, 3, 2, 6]
mid_index = (arr.size)/2

for i in (0...mid_index)
  arr[i], arr[arr.size - i - 1] = arr[arr.size - i - 1], arr[i]
end
puts arr.join(' ')