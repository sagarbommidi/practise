class Foo
  def pub_method
    puts "Its a public method"
  end

  private
  def pri_method(val=nil)
    puts "Its a Private method #{val}"
  end
end

f = Foo.new
# f.pri_method #=> private method `pri_method' called for #<Foo:0x007fc8d10c1470> (NoMethodError)

f.send(:pri_method) #=> Its a Private method
f.send(:pri_method, "in Ruby") #=> Its a Private method in Ruby

f.method(:pri_method).call #=> Its a Private method
f.method(:pri_method).call("in Ruby") #=> Its a Private method in Ruby

eval("f.pub_method") #=> Its a public method
# eval("f.pri_method") #=> `eval': private method `pri_method' called for #<Foo:0x007fabbe0acb40> (NoMethodError)

f.instance_eval {pri_method} #=> Its a Private method
