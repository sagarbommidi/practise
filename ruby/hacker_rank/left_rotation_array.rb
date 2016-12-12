#!/bin/ruby

def left_rotate(array, d)
  temp_ele = array[0]
  (0...array.length).each do |index|
    if index == array.length - 1
      array[index] = temp_ele
    else
      array[index] = array[index + 1]
    end
  end
end

arr = [1, 2, 3, 4, 5]
k = 9
k = k % arr.length
k.times { left_rotate(arr) }
puts arr.join(' ')
