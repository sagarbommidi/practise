def serial_average(str)
   arr = str.split('-')
   avg = "%.2f" % ((arr[1].to_f + arr[2].to_f)/2)
   "#{arr[0]}-#{avg}"
end
puts serial_average('002-10.00-20.00')