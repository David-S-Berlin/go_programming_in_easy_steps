package cube

type Dims struct {
width, length, height int
}


func (d *Dims) area() int {
return d.width * d.length
}

func (d *Dims) SetSize( w, l, h int) {
d.width = w
  d.length = l
  d.height = h
}

func (d *Dims) GetVolume() int {
return d.width * d.length * d.height
}

func (d *Dims) GetArea() int {
  return d.area()
}
