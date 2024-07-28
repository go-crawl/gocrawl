package scheduler

type Scheduler struct {
	visited map[string]bool
	queue   []string
}

func NewScheduler() *Scheduler {
	return &Scheduler{
		visited: make(map[string]bool),
		queue:   []string{},
	}
}

func (s *Scheduler) Add(url string) {
	if !s.visited[url] {
		s.visited[url] = true
		s.queue = append(s.queue, url)
	}
}

func (s *Scheduler) NextRequest() *Request {
	if len(s.queue) == 0 {
		return nil
	}
	url := s.queue[0]
	s.queue = s.queue[1:]
	return &Request{URL: url}
}

type Request struct {
	URL string
}
