#!/bin/ruby

def get_cookies(array, sweetness)
  sort_min(array)
  count = 0
  return -1 if array[0] < 1 && array[1] < 1
  while (array.length > 1 && array[0] < sweetness)
    new_sweet = array[0] + 2 * array[1]
    puts new_sweet
    new_arr = [new_sweet]
    new_arr.concat(array[2..-1])
    return -1 if new_arr[0] < sweetness && new_arr.length < 2
    sort_min(new_arr)
    count += 1
    array = new_arr
  end
  count
end

def sort_min(arr)
  for i in [0, 1] do
    min = i
    for j in (i+1...arr.length) do
      min = j if arr[j] < arr[min]
    end
    arr[i], arr[min] = arr[min], arr[i]
  end
end


# n, k = gets.strip.split(' ').map(&:to_i)
# arr = []
# for i in 0...n do
#   arr[i] = gets.strip.to_i
# end
str = (0...100000).map { |x| 1 }.join(' ')
array = str.strip.split(' ').map(&:to_i)
puts "prepared array"
# puts str
# sort_min(array)
# puts array[0...5].join(', ')
# puts get_cookies(array, 105823341)

puts get_cookies(array, 105800)



# all ones 100000 105823341
