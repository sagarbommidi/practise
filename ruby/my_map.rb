class Array
  def my_map
    return to_enum :my_map unless block_given?
    inject([]) { |result, value| result << yield(value) }
  end
end

a = [1, 2, 3, 4]
b = a.my_map do |x|
  x + 1
end
c = a.my_map(&:to_s)
d = a.my_map # when no block has given

puts b.inspect #=> [2, 3, 4, 5]
puts c.inspect #=> ["1", "2", "3", "4"]
puts d.inspect #=> #<Enumerator: [1, 2, 3, 4]:my_map>
