module Abc
  def abc
    puts 'Its an instance method in Module'
  end

  def self.xyz
    puts "Its a class method in module"
  end
end

class Inc
  include Abc
end

class Exc
  extend Abc
end

Inc.new.abc
# Inc.new.xyz #=> Raise error
# Inc.xyz #=> Raise error, undefined method xyz for Inc object

Exc.abc
# Exc.new.xyz #=> Raise error, undefined method xyz for Exc object
