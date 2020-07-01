package pageloop

type Pager struct {
	cur     int
	perPage int
	length  int
}

func NewPager(perPage, length int) *Pager {
	if perPage <= 0 {
		perPage = length
	}

	return &Pager{
		perPage: perPage,
		length:  length,
	}
}

func (p *Pager) Next() bool {
	return p.cur < p.length
}

func (p *Pager) Index() (int, int) {
	begin := p.cur
	end := p.cur + p.perPage

	p.cur = end
	if p.length < end {
		return begin, p.length
	}

	return begin, end
}
