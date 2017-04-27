# arr = [-2, -3, 4, -1, -2, 1, 5, -3]
arr = [2,3,-5,2, -1]
len = arr.length
sum, max = arr[0], arr[0]
# for x in 0...len-3 do
#   sum = arr[x] + arr[x+1] + arr[x+2]
#   for y in (x+3)...len do
#     sum += arr[y]
#     max1 = sum if sum > max1
#   end
# end
# puts max1

for x in 1...len do
  sum += arr[x]
  sum = 0 if sum < 0
  max = sum if sum > max
end
puts max