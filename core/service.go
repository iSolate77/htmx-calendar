package core

type Service struct {
	Ss SubjectService
	Cs CurriculumService
}

func NewService(sr SubjectRepository, cr CurriculumRepository) *Service {
	return &Service{
		Ss: *NewSubjectService(sr),
		Cs: *NewCurriculumService(cr),
	}
}
