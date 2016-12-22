class Node
  attr_accessor :data, :left, :right, :height

  def initialize(key)
    @left, @right = nil, nil
    @data = key
    @height = 0
  end
end

class AvlTree
  attr_accessor :root

  def initialize
    @root = nil
  end

  def insert(root_node, key)
    if root_node == nil
      node = Node.new(key)
      self.root = node
      return root
    end


  end
end


tree = AvlTree.new()
