package Model

type Meta struct {
	totalRecord, totalPage, page int
}

type Response struct {
	Status  bool
	Data    interface{}
	Message string
	Meta    Meta
}

func (m *Meta) SetPage(page int) {
	m.page = page
}

func (m *Meta) SetTotalPage(totalPage int) {
	m.totalPage = totalPage
}

func (m *Meta) SetTotalRecord(total int) {
	m.totalRecord = total
}

func (m *Meta) GetTotalRecord() int {
	return m.totalRecord
}

func (m *Meta) GetTotalPage() int {
	return m.totalPage
}

func (m *Meta) GetPage() int {
	return m.page
}
