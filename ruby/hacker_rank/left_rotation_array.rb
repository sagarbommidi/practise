#!/bin/ruby

def left_rotate(array, d)
  d = d % array.length
  array = array[d..-1] + array[0...d]
end

arr = [1, 2, 3, 4, 5, 6, 7, 8]
k = 10
arr = left_rotate(arr, k)
puts arr.join(' ')
