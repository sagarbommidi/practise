class Array
  def my_reverse
    i = 0
    j = length - 1
    while i < j do
      self[i], self[j] = self[j], self[i]
      i += 1
      j -= 1
    end
    self
  end
end

a = [1,2,3,4,5,6,7,8]
puts a.my_reverse.inspect
