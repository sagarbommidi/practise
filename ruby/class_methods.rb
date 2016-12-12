class Developer
  self.instance_eval do
    def backend
      puts "Its a Class method. I am backend developer"
    end
  end

  self.class_eval do
    def frontend
      puts "Its an instance method"
    end
  end

  class << self
    def sr_backend
      puts "Its a Class method. I am senior backend developer"
    end
  end
end

Developer.backend
Developer.sr_backend
Developer.new.frontend
